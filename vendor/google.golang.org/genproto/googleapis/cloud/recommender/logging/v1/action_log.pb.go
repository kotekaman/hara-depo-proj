// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/cloud/recommender/logging/v1/action_log.proto

package logging

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	v1 "google.golang.org/genproto/googleapis/cloud/recommender/v1"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Log content of an action on a recommendation. This includes Mark* actions, as
// well as ApplyRule actions.
type ActionLog struct {
	// Required. User that executed this action. Eg, foo@gmail.com
	Actor string `protobuf:"bytes,1,opt,name=actor,proto3" json:"actor,omitempty"`
	// Required. Action that was taken by the actor. Eg, MarkCompleted.
	State v1.RecommendationStateInfo_State `protobuf:"varint,2,opt,name=state,proto3,enum=google.cloud.recommender.v1.RecommendationStateInfo_State" json:"state,omitempty"`
	// Optional. Metadata that was included with the action that was taken.
	StateMetadata map[string]string `protobuf:"bytes,3,rep,name=state_metadata,json=stateMetadata,proto3" json:"state_metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Required. Name of the recommendation which was acted on. Eg, :
	// 'projects/foo/locations/global/recommenders/roleReco/recommendations/r1'
	RecommendationName   string   `protobuf:"bytes,4,opt,name=recommendation_name,json=recommendationName,proto3" json:"recommendation_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ActionLog) Reset()         { *m = ActionLog{} }
func (m *ActionLog) String() string { return proto.CompactTextString(m) }
func (*ActionLog) ProtoMessage()    {}
func (*ActionLog) Descriptor() ([]byte, []int) {
	return fileDescriptor_e38d9a8d58fa2358, []int{0}
}

func (m *ActionLog) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ActionLog.Unmarshal(m, b)
}
func (m *ActionLog) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ActionLog.Marshal(b, m, deterministic)
}
func (m *ActionLog) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ActionLog.Merge(m, src)
}
func (m *ActionLog) XXX_Size() int {
	return xxx_messageInfo_ActionLog.Size(m)
}
func (m *ActionLog) XXX_DiscardUnknown() {
	xxx_messageInfo_ActionLog.DiscardUnknown(m)
}

var xxx_messageInfo_ActionLog proto.InternalMessageInfo

func (m *ActionLog) GetActor() string {
	if m != nil {
		return m.Actor
	}
	return ""
}

func (m *ActionLog) GetState() v1.RecommendationStateInfo_State {
	if m != nil {
		return m.State
	}
	return v1.RecommendationStateInfo_STATE_UNSPECIFIED
}

func (m *ActionLog) GetStateMetadata() map[string]string {
	if m != nil {
		return m.StateMetadata
	}
	return nil
}

func (m *ActionLog) GetRecommendationName() string {
	if m != nil {
		return m.RecommendationName
	}
	return ""
}

func init() {
	proto.RegisterType((*ActionLog)(nil), "google.cloud.recommender.logging.v1.ActionLog")
	proto.RegisterMapType((map[string]string)(nil), "google.cloud.recommender.logging.v1.ActionLog.StateMetadataEntry")
}

func init() {
	proto.RegisterFile("google/cloud/recommender/logging/v1/action_log.proto", fileDescriptor_e38d9a8d58fa2358)
}

var fileDescriptor_e38d9a8d58fa2358 = []byte{
	// 323 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xc1, 0x4e, 0x02, 0x31,
	0x10, 0x86, 0x53, 0x10, 0x13, 0x6a, 0x24, 0xa6, 0x7a, 0xd8, 0x70, 0x22, 0x7a, 0x90, 0x53, 0xeb,
	0xa2, 0x07, 0x83, 0x17, 0x21, 0xf1, 0xa0, 0x51, 0x43, 0xd6, 0x9b, 0x17, 0x32, 0x2e, 0xb5, 0x12,
	0xdb, 0x0e, 0x29, 0x65, 0x13, 0x9e, 0xc0, 0xd7, 0xf3, 0x91, 0xcc, 0xb6, 0x1b, 0x84, 0x18, 0x0c,
	0xb7, 0x7f, 0xb6, 0xf9, 0xbf, 0x99, 0x7f, 0x66, 0xe9, 0x95, 0x42, 0x54, 0x5a, 0x8a, 0x5c, 0xe3,
	0x62, 0x22, 0x9c, 0xcc, 0xd1, 0x18, 0x69, 0x27, 0xd2, 0x09, 0x8d, 0x4a, 0x4d, 0xad, 0x12, 0x45,
	0x2a, 0x20, 0xf7, 0x53, 0xb4, 0x63, 0x8d, 0x8a, 0xcf, 0x1c, 0x7a, 0x64, 0x67, 0xd1, 0xc5, 0x83,
	0x8b, 0xaf, 0xb9, 0x78, 0xe5, 0xe2, 0x45, 0xda, 0xbe, 0xd8, 0x8a, 0x2e, 0xd2, 0xdf, 0x12, 0x4a,
	0x74, 0xc4, 0x9e, 0x7e, 0xd7, 0x68, 0x73, 0x10, 0x7a, 0x3d, 0xa2, 0x62, 0x27, 0xb4, 0x01, 0xb9,
	0x47, 0x97, 0x90, 0x0e, 0xe9, 0x36, 0xb3, 0x58, 0xb0, 0x11, 0x6d, 0xcc, 0x3d, 0x78, 0x99, 0xd4,
	0x3a, 0xa4, 0xdb, 0xea, 0xf5, 0xf9, 0xd6, 0x51, 0x8a, 0x94, 0x67, 0x1b, 0x5d, 0x5e, 0x4a, 0xdf,
	0xbd, 0x7d, 0x47, 0x1e, 0x54, 0x16, 0x41, 0xec, 0x83, 0xb6, 0x82, 0x18, 0x1b, 0xe9, 0x61, 0x02,
	0x1e, 0x92, 0x7a, 0xa7, 0xde, 0x3d, 0xe8, 0x0d, 0xf8, 0x0e, 0x29, 0xf9, 0x6a, 0xde, 0x08, 0x7d,
	0xaa, 0x18, 0x77, 0xd6, 0xbb, 0x65, 0x76, 0x38, 0x5f, 0xff, 0xc6, 0x04, 0x3d, 0xde, 0xcc, 0x3d,
	0xb6, 0x60, 0x64, 0xb2, 0x17, 0xf2, 0xb1, 0xcd, 0xa7, 0x67, 0x30, 0xb2, 0x7d, 0x4b, 0xd9, 0x5f,
	0x2a, 0x3b, 0xa2, 0xf5, 0x4f, 0xb9, 0xac, 0xd6, 0x52, 0xca, 0x72, 0x55, 0x05, 0xe8, 0x45, 0x5c,
	0x4a, 0x33, 0x8b, 0x45, 0xbf, 0x76, 0x4d, 0x86, 0x5f, 0x84, 0x9e, 0xe7, 0x68, 0x76, 0x89, 0x32,
	0x6c, 0xad, 0xb2, 0x8c, 0xca, 0x73, 0x8c, 0xc8, 0xeb, 0x43, 0x65, 0x53, 0xa8, 0xc1, 0x2a, 0x8e,
	0x4e, 0x09, 0x25, 0x6d, 0x38, 0x96, 0x88, 0x4f, 0x30, 0x9b, 0xce, 0xff, 0xfd, 0x79, 0x6e, 0x2a,
	0xf9, 0xb6, 0x1f, 0x6c, 0x97, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xe0, 0x16, 0x89, 0x26, 0x72,
	0x02, 0x00, 0x00,
}