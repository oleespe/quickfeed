// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: kit/score/score.proto

package score

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Score give the score for a single test named TestName.
type Score struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID           uint64 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	SubmissionID uint64 `protobuf:"varint,2,opt,name=SubmissionID,proto3" json:"SubmissionID,omitempty"`
	Secret       string `protobuf:"bytes,3,opt,name=Secret,proto3" json:"Secret,omitempty"`           // the unique identifier for a scoring session
	TestName     string `protobuf:"bytes,4,opt,name=TestName,proto3" json:"TestName,omitempty"`       // name of the test
	Score        int32  `protobuf:"varint,5,opt,name=Score,proto3" json:"Score,omitempty"`            // the score obtained
	MaxScore     int32  `protobuf:"varint,6,opt,name=MaxScore,proto3" json:"MaxScore,omitempty"`      // max score possible to get on this specific test
	Weight       int32  `protobuf:"varint,7,opt,name=Weight,proto3" json:"Weight,omitempty"`          // the weight of this test; used to compute final grade
	TestDetails  string `protobuf:"bytes,8,opt,name=TestDetails,proto3" json:"TestDetails,omitempty"` // if populated, the frontend may display additional details (TODO(meling) adapt to output from go test -json)
}

func (x *Score) Reset() {
	*x = Score{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kit_score_score_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Score) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Score) ProtoMessage() {}

func (x *Score) ProtoReflect() protoreflect.Message {
	mi := &file_kit_score_score_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Score.ProtoReflect.Descriptor instead.
func (*Score) Descriptor() ([]byte, []int) {
	return file_kit_score_score_proto_rawDescGZIP(), []int{0}
}

func (x *Score) GetID() uint64 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *Score) GetSubmissionID() uint64 {
	if x != nil {
		return x.SubmissionID
	}
	return 0
}

func (x *Score) GetSecret() string {
	if x != nil {
		return x.Secret
	}
	return ""
}

func (x *Score) GetTestName() string {
	if x != nil {
		return x.TestName
	}
	return ""
}

func (x *Score) GetScore() int32 {
	if x != nil {
		return x.Score
	}
	return 0
}

func (x *Score) GetMaxScore() int32 {
	if x != nil {
		return x.MaxScore
	}
	return 0
}

func (x *Score) GetWeight() int32 {
	if x != nil {
		return x.Weight
	}
	return 0
}

func (x *Score) GetTestDetails() string {
	if x != nil {
		return x.TestDetails
	}
	return ""
}

// BuildInfo holds build data for an assignment's test execution.
type BuildInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID           uint64                 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	SubmissionID uint64                 `protobuf:"varint,2,opt,name=SubmissionID,proto3" json:"SubmissionID,omitempty"`
	BuildDate    *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=BuildDate,proto3" json:"BuildDate,omitempty"`
	BuildLog     string                 `protobuf:"bytes,4,opt,name=BuildLog,proto3" json:"BuildLog,omitempty"`
	ExecTime     int64                  `protobuf:"varint,5,opt,name=ExecTime,proto3" json:"ExecTime,omitempty"`
}

func (x *BuildInfo) Reset() {
	*x = BuildInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kit_score_score_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BuildInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BuildInfo) ProtoMessage() {}

func (x *BuildInfo) ProtoReflect() protoreflect.Message {
	mi := &file_kit_score_score_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BuildInfo.ProtoReflect.Descriptor instead.
func (*BuildInfo) Descriptor() ([]byte, []int) {
	return file_kit_score_score_proto_rawDescGZIP(), []int{1}
}

func (x *BuildInfo) GetID() uint64 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *BuildInfo) GetSubmissionID() uint64 {
	if x != nil {
		return x.SubmissionID
	}
	return 0
}

func (x *BuildInfo) GetBuildDate() *timestamppb.Timestamp {
	if x != nil {
		return x.BuildDate
	}
	return nil
}

func (x *BuildInfo) GetBuildLog() string {
	if x != nil {
		return x.BuildLog
	}
	return ""
}

func (x *BuildInfo) GetExecTime() int64 {
	if x != nil {
		return x.ExecTime
	}
	return 0
}

// Results holds a map from test names to score objects and build info.
// Not to be stored in database.
type Results struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BuildInfo *BuildInfo `protobuf:"bytes,1,opt,name=BuildInfo,proto3" json:"BuildInfo,omitempty"` // build info for tests
	Scores    []*Score   `protobuf:"bytes,2,rep,name=Scores,proto3" json:"Scores,omitempty"`       // list of scores for different tests
}

func (x *Results) Reset() {
	*x = Results{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kit_score_score_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Results) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Results) ProtoMessage() {}

func (x *Results) ProtoReflect() protoreflect.Message {
	mi := &file_kit_score_score_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Results.ProtoReflect.Descriptor instead.
func (*Results) Descriptor() ([]byte, []int) {
	return file_kit_score_score_proto_rawDescGZIP(), []int{2}
}

func (x *Results) GetBuildInfo() *BuildInfo {
	if x != nil {
		return x.BuildInfo
	}
	return nil
}

func (x *Results) GetScores() []*Score {
	if x != nil {
		return x.Scores
	}
	return nil
}

var File_kit_score_score_proto protoreflect.FileDescriptor

var file_kit_score_score_proto_rawDesc = []byte{
	0x0a, 0x15, 0x6b, 0x69, 0x74, 0x2f, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x73, 0x63, 0x6f, 0x72,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x1a, 0x1f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xdb, 0x01, 0x0a, 0x05, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x49, 0x44, 0x12, 0x22, 0x0a, 0x0c, 0x53, 0x75, 0x62,
	0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x0c, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x12, 0x16, 0x0a,
	0x06, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x53,
	0x65, 0x63, 0x72, 0x65, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x54, 0x65, 0x73, 0x74, 0x4e, 0x61, 0x6d,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x54, 0x65, 0x73, 0x74, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x05, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x4d, 0x61, 0x78, 0x53, 0x63,
	0x6f, 0x72, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x4d, 0x61, 0x78, 0x53, 0x63,
	0x6f, 0x72, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x57, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x06, 0x57, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x54,
	0x65, 0x73, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x54, 0x65, 0x73, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x22, 0xb1, 0x01,
	0x0a, 0x09, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x49,
	0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x49, 0x44, 0x12, 0x22, 0x0a, 0x0c, 0x53,
	0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x0c, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x12,
	0x38, 0x0a, 0x09, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x44, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09,
	0x42, 0x75, 0x69, 0x6c, 0x64, 0x44, 0x61, 0x74, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x42, 0x75, 0x69,
	0x6c, 0x64, 0x4c, 0x6f, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x42, 0x75, 0x69,
	0x6c, 0x64, 0x4c, 0x6f, 0x67, 0x12, 0x1a, 0x0a, 0x08, 0x45, 0x78, 0x65, 0x63, 0x54, 0x69, 0x6d,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x45, 0x78, 0x65, 0x63, 0x54, 0x69, 0x6d,
	0x65, 0x22, 0x5f, 0x0a, 0x07, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x12, 0x2e, 0x0a, 0x09,
	0x42, 0x75, 0x69, 0x6c, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x10, 0x2e, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x09, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x24, 0x0a, 0x06,
	0x53, 0x63, 0x6f, 0x72, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x73,
	0x63, 0x6f, 0x72, 0x65, 0x2e, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x52, 0x06, 0x53, 0x63, 0x6f, 0x72,
	0x65, 0x73, 0x42, 0x2a, 0x5a, 0x28, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x61, 0x75, 0x74, 0x6f, 0x67, 0x72, 0x61, 0x64, 0x65, 0x2f, 0x71, 0x75, 0x69, 0x63, 0x6b,
	0x66, 0x65, 0x65, 0x64, 0x2f, 0x6b, 0x69, 0x74, 0x2f, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_kit_score_score_proto_rawDescOnce sync.Once
	file_kit_score_score_proto_rawDescData = file_kit_score_score_proto_rawDesc
)

func file_kit_score_score_proto_rawDescGZIP() []byte {
	file_kit_score_score_proto_rawDescOnce.Do(func() {
		file_kit_score_score_proto_rawDescData = protoimpl.X.CompressGZIP(file_kit_score_score_proto_rawDescData)
	})
	return file_kit_score_score_proto_rawDescData
}

var file_kit_score_score_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_kit_score_score_proto_goTypes = []interface{}{
	(*Score)(nil),                 // 0: score.Score
	(*BuildInfo)(nil),             // 1: score.BuildInfo
	(*Results)(nil),               // 2: score.Results
	(*timestamppb.Timestamp)(nil), // 3: google.protobuf.Timestamp
}
var file_kit_score_score_proto_depIdxs = []int32{
	3, // 0: score.BuildInfo.BuildDate:type_name -> google.protobuf.Timestamp
	1, // 1: score.Results.BuildInfo:type_name -> score.BuildInfo
	0, // 2: score.Results.Scores:type_name -> score.Score
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_kit_score_score_proto_init() }
func file_kit_score_score_proto_init() {
	if File_kit_score_score_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_kit_score_score_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Score); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_kit_score_score_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BuildInfo); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_kit_score_score_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Results); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_kit_score_score_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_kit_score_score_proto_goTypes,
		DependencyIndexes: file_kit_score_score_proto_depIdxs,
		MessageInfos:      file_kit_score_score_proto_msgTypes,
	}.Build()
	File_kit_score_score_proto = out.File
	file_kit_score_score_proto_rawDesc = nil
	file_kit_score_score_proto_goTypes = nil
	file_kit_score_score_proto_depIdxs = nil
}
