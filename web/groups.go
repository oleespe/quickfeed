package web

import (
	"context"
	"net/http"

	"github.com/autograde/aguis/database"
	"github.com/autograde/aguis/models"
	"github.com/autograde/aguis/scm"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
)

// PatchGroup updates status of a group
func PatchGroup(logger logrus.FieldLogger, db database.Database) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := parseUint(c.Param("gid"))
		if err != nil {
			return err
		}
		var ngrp UpdateGroupRequest
		if err := c.Bind(&ngrp); err != nil {
			return err
		}
		if ngrp.Status > models.Teacher {
			return echo.NewHTTPError(http.StatusBadRequest, "invalid payload")
		}

		user := c.Get("user").(*models.User)
		// TODO: This check should be performed in AccessControl.
		if user.IsAdmin == nil || !*user.IsAdmin {
			// Only admin / teacher can update status of a group
			return c.NoContent(http.StatusForbidden)
		}

		// we need the remote identities of the group's users
		oldgrp, err := db.GetGroup(true, id)
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				return echo.NewHTTPError(http.StatusNotFound, "group not found")
			}
			return err
		}
		users := oldgrp.Users

		course, err := db.GetCourse(oldgrp.CourseID)
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				return echo.NewHTTPError(http.StatusNotFound, "course not found")
			}
			return err
		}

		s, err := getSCM(c, course.Provider)
		if err != nil {
			return err
		}
		ctx, cancel := context.WithTimeout(c.Request().Context(), MaxWait)
		defer cancel()

		// Create and add repo to autograder group
		dir, err := s.GetDirectory(ctx, course.DirectoryID)
		if err != nil {
			return err
		}
		logger.WithField("course.DirID", course.DirectoryID).
			WithField("dir", dir.Path).
			Println("GetDir")
		repos, err := s.GetRepositories(ctx, dir)
		if err != nil {
			return err
		}
		existing := make(map[string]*scm.Repository)
		for _, repo := range repos {
			logger.WithField("path", oldgrp.Name).
				WithField("repoPath", repo.Path).
				Println("Existing repo")
			existing[repo.Path] = repo
		}
		repo, created := existing[oldgrp.Name]
		if !created {
			repo, err = s.CreateRepository(ctx, &scm.CreateRepositoryOptions{
				Directory: dir,
				Path:      oldgrp.Name,
				Private:   true,
			})
			if err != nil {
				logger.WithField("path", oldgrp.Name).WithError(err).Warn("Failed to create repository")
				//TODO(meling) this does not seem to hold group repos for unknown reasons
				repo = existing[oldgrp.Name]
				return err
			}
			logger.WithField("repo", repo).Println("Created new group repository")
		}

		// Add repo to DB
		dbRepo := models.Repository{
			DirectoryID:  course.DirectoryID,
			RepositoryID: repo.ID,
			HTMLURL:      repo.WebURL,
			Type:         models.UserRepo,
			UserID:       0,
			GroupID:      oldgrp.ID,
		}
		if err := db.CreateRepository(&dbRepo); err != nil {
			logger.WithField("url", repo.WebURL).WithField("gid", oldgrp.ID).WithError(err).Warn("Failed to create repository in database")
			return err
		}
		logger.WithField("repo", repo).Println("Created new group repository in database")

		if err := db.UpdateGroupStatus(&models.Group{
			ID:     oldgrp.ID,
			Status: ngrp.Status,
		}); err != nil {
			logger.WithField("status", ngrp.Status).WithField("gid", oldgrp.ID).WithError(err).Warn("Failed to update group status in database")
			return err
		}

		var gitUserNames []string
		for _, user := range users {
			remote, err := getRemoteIDFor(user, course.Provider)
			if err != nil {
				return err
			}
			// Note this requires one git call per user in the group
			userName, err := s.GetUserNameByID(ctx, remote.RemoteID)
			if err != nil {
				return err
			}
			gitUserNames = append(gitUserNames, userName)
		}

		// Create git-team
		team, err := s.CreateTeam(ctx, &scm.CreateTeamOptions{
			Directory: &scm.Directory{Path: dir.Path},
			TeamName:  oldgrp.Name,
			Users:     gitUserNames,
		})
		if err != nil {
			logger.WithField("path", dir.Path).WithField("team", oldgrp.Name).WithField("users", gitUserNames).WithError(err).Warn("Failed to create git-team")
			return err
		}
		// Adding Repo to git-team
		if err = s.AddTeamRepo(ctx, &scm.AddTeamRepoOptions{
			TeamID: team.ID,
			Owner:  repo.Owner,
			Repo:   repo.Path,
		}); err != nil {
			logger.WithField("repo", repo.Path).WithField("team", team.ID).WithField("owner", repo.Owner).WithError(err).Warn("Failed to add repo to git-team")
			return err
		}

		return c.NoContent(http.StatusOK)
	}
}

// GetGroup returns a group
func GetGroup(db database.Database) echo.HandlerFunc {
	return func(c echo.Context) error {
		gid, err := parseUint(c.Param("gid"))
		if err != nil {
			return err
		}
		group, err := db.GetGroup(false, gid)
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				return echo.NewHTTPError(http.StatusNotFound, "group not found")
			}
			return err
		}
		return c.JSONPretty(http.StatusOK, group, "\t")
	}
}

// DeleteGroup deletes a pending or rejected group
func DeleteGroup(db database.Database) echo.HandlerFunc {
	return func(c echo.Context) error {
		gid, err := parseUint(c.Param("gid"))
		if err != nil {
			return err
		}
		group, err := db.GetGroup(false, gid)
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				return echo.NewHTTPError(http.StatusNotFound, "group not found")
			}
			return err
		}
		if group.Status > models.Rejected {
			return echo.NewHTTPError(http.StatusForbidden, "accepted group cannot be deleted")
		}
		if err := db.DeleteGroup(gid); err != nil {
			return nil
		}
		return c.NoContent(http.StatusOK)
	}
}
