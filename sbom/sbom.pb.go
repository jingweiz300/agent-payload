// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.12
// source: proto/sbom/sbom.proto

package sbom

import (
	cyclonedx_v1_4 "github.com/DataDog/agent-payload/v5/cyclonedx_v1_4"
	duration "github.com/golang/protobuf/ptypes/duration"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type SBOMSourceType int32

const (
	SBOMSourceType_UNSPECIFIED            SBOMSourceType = 0
	SBOMSourceType_CONTAINER_IMAGE_LAYERS SBOMSourceType = 1
	SBOMSourceType_CONTAINER_FILE_SYSTEM  SBOMSourceType = 2
	SBOMSourceType_HOST_FILE_SYSTEM       SBOMSourceType = 3
)

// Enum value maps for SBOMSourceType.
var (
	SBOMSourceType_name = map[int32]string{
		0: "UNSPECIFIED",
		1: "CONTAINER_IMAGE_LAYERS",
		2: "CONTAINER_FILE_SYSTEM",
		3: "HOST_FILE_SYSTEM",
	}
	SBOMSourceType_value = map[string]int32{
		"UNSPECIFIED":            0,
		"CONTAINER_IMAGE_LAYERS": 1,
		"CONTAINER_FILE_SYSTEM":  2,
		"HOST_FILE_SYSTEM":       3,
	}
)

func (x SBOMSourceType) Enum() *SBOMSourceType {
	p := new(SBOMSourceType)
	*p = x
	return p
}

func (x SBOMSourceType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SBOMSourceType) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_sbom_sbom_proto_enumTypes[0].Descriptor()
}

func (SBOMSourceType) Type() protoreflect.EnumType {
	return &file_proto_sbom_sbom_proto_enumTypes[0]
}

func (x SBOMSourceType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SBOMSourceType.Descriptor instead.
func (SBOMSourceType) EnumDescriptor() ([]byte, []int) {
	return file_proto_sbom_sbom_proto_rawDescGZIP(), []int{0}
}

// SBOMPayload represents the main SBOM payload
type SBOMPayload struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version  int32         `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty"`
	Host     string        `protobuf:"bytes,2,opt,name=host,proto3" json:"host,omitempty"`
	Source   *string       `protobuf:"bytes,3,opt,name=source,proto3,oneof" json:"source,omitempty"` // use to know the source of the message: agent, other
	Entities []*SBOMEntity `protobuf:"bytes,4,rep,name=entities,proto3" json:"entities,omitempty"`
	DdEnv    *string       `protobuf:"bytes,5,opt,name=dd_env,json=ddEnv,proto3,oneof" json:"dd_env,omitempty"`
}

func (x *SBOMPayload) Reset() {
	*x = SBOMPayload{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_sbom_sbom_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SBOMPayload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SBOMPayload) ProtoMessage() {}

func (x *SBOMPayload) ProtoReflect() protoreflect.Message {
	mi := &file_proto_sbom_sbom_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SBOMPayload.ProtoReflect.Descriptor instead.
func (*SBOMPayload) Descriptor() ([]byte, []int) {
	return file_proto_sbom_sbom_proto_rawDescGZIP(), []int{0}
}

func (x *SBOMPayload) GetVersion() int32 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *SBOMPayload) GetHost() string {
	if x != nil {
		return x.Host
	}
	return ""
}

func (x *SBOMPayload) GetSource() string {
	if x != nil && x.Source != nil {
		return *x.Source
	}
	return ""
}

func (x *SBOMPayload) GetEntities() []*SBOMEntity {
	if x != nil {
		return x.Entities
	}
	return nil
}

func (x *SBOMPayload) GetDdEnv() string {
	if x != nil && x.DdEnv != nil {
		return *x.DdEnv
	}
	return ""
}

type SBOMEntity struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type               SBOMSourceType       `protobuf:"varint,1,opt,name=type,proto3,enum=datadog.sbom.SBOMSourceType" json:"type,omitempty"`
	Id                 string               `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`                                       // Unique identifier to be able to correlated and "deduplicate" SBOM
	DdTags             []string             `protobuf:"bytes,7,rep,name=dd_tags,json=ddTags,proto3" json:"dd_tags,omitempty"`                 // datadog tags that will be added by the agent depending of the SBOMSourceType
	GeneratedAt        *timestamp.Timestamp `protobuf:"bytes,3,opt,name=generatedAt,proto3,oneof" json:"generatedAt,omitempty"`               // the datetime of the SBOM generation
	RepoTags           []string             `protobuf:"bytes,4,rep,name=repo_tags,json=repoTags,proto3" json:"repo_tags,omitempty"`           // the tags of the container image
	InUse              bool                 `protobuf:"varint,5,opt,name=inUse,proto3" json:"inUse,omitempty"`                                // Whether the SBOM concerns a running entity (running container) or an inert entity (image not used by any container)
	GenerationDuration *duration.Duration   `protobuf:"bytes,6,opt,name=generationDuration,proto3,oneof" json:"generationDuration,omitempty"` // SBOM generation duration (how long it took to generate the SBOM report)
	// Types that are assignable to Sbom:
	//	*SBOMEntity_Cyclonedx
	Sbom isSBOMEntity_Sbom `protobuf_oneof:"sbom"`
}

func (x *SBOMEntity) Reset() {
	*x = SBOMEntity{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_sbom_sbom_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SBOMEntity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SBOMEntity) ProtoMessage() {}

func (x *SBOMEntity) ProtoReflect() protoreflect.Message {
	mi := &file_proto_sbom_sbom_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SBOMEntity.ProtoReflect.Descriptor instead.
func (*SBOMEntity) Descriptor() ([]byte, []int) {
	return file_proto_sbom_sbom_proto_rawDescGZIP(), []int{1}
}

func (x *SBOMEntity) GetType() SBOMSourceType {
	if x != nil {
		return x.Type
	}
	return SBOMSourceType_UNSPECIFIED
}

func (x *SBOMEntity) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *SBOMEntity) GetDdTags() []string {
	if x != nil {
		return x.DdTags
	}
	return nil
}

func (x *SBOMEntity) GetGeneratedAt() *timestamp.Timestamp {
	if x != nil {
		return x.GeneratedAt
	}
	return nil
}

func (x *SBOMEntity) GetRepoTags() []string {
	if x != nil {
		return x.RepoTags
	}
	return nil
}

func (x *SBOMEntity) GetInUse() bool {
	if x != nil {
		return x.InUse
	}
	return false
}

func (x *SBOMEntity) GetGenerationDuration() *duration.Duration {
	if x != nil {
		return x.GenerationDuration
	}
	return nil
}

func (m *SBOMEntity) GetSbom() isSBOMEntity_Sbom {
	if m != nil {
		return m.Sbom
	}
	return nil
}

func (x *SBOMEntity) GetCyclonedx() *cyclonedx_v1_4.Bom {
	if x, ok := x.GetSbom().(*SBOMEntity_Cyclonedx); ok {
		return x.Cyclonedx
	}
	return nil
}

type isSBOMEntity_Sbom interface {
	isSBOMEntity_Sbom()
}

type SBOMEntity_Cyclonedx struct {
	Cyclonedx *cyclonedx_v1_4.Bom `protobuf:"bytes,10,opt,name=cyclonedx,proto3,oneof"` // only cyclonedx will be supported initially but putting it optional will allow us to move to another format later
}

func (*SBOMEntity_Cyclonedx) isSBOMEntity_Sbom() {}

var File_proto_sbom_sbom_proto protoreflect.FileDescriptor

var file_proto_sbom_sbom_proto_rawDesc = []byte{
	0x0a, 0x15, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x62, 0x6f, 0x6d, 0x2f, 0x73, 0x62, 0x6f,
	0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x64, 0x61, 0x74, 0x61, 0x64, 0x6f, 0x67,
	0x2e, 0x73, 0x62, 0x6f, 0x6d, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x42, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x64, 0x65,
	0x70, 0x73, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x43, 0x79,
	0x63, 0x6c, 0x6f, 0x6e, 0x65, 0x44, 0x58, 0x2f, 0x73, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2f, 0x62, 0x6f, 0x6d,
	0x2d, 0x31, 0x2e, 0x34, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc0, 0x01, 0x0a, 0x0b, 0x53,
	0x42, 0x4f, 0x4d, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x06, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x88, 0x01, 0x01, 0x12, 0x34, 0x0a, 0x08, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65,
	0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x64, 0x6f,
	0x67, 0x2e, 0x73, 0x62, 0x6f, 0x6d, 0x2e, 0x53, 0x42, 0x4f, 0x4d, 0x45, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x52, 0x08, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x12, 0x1a, 0x0a, 0x06, 0x64,
	0x64, 0x5f, 0x65, 0x6e, 0x76, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x05, 0x64,
	0x64, 0x45, 0x6e, 0x76, 0x88, 0x01, 0x01, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x64, 0x64, 0x5f, 0x65, 0x6e, 0x76, 0x22, 0x91, 0x03,
	0x0a, 0x0a, 0x53, 0x42, 0x4f, 0x4d, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x30, 0x0a, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1c, 0x2e, 0x64, 0x61, 0x74,
	0x61, 0x64, 0x6f, 0x67, 0x2e, 0x73, 0x62, 0x6f, 0x6d, 0x2e, 0x53, 0x42, 0x4f, 0x4d, 0x53, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x17,
	0x0a, 0x07, 0x64, 0x64, 0x5f, 0x74, 0x61, 0x67, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x06, 0x64, 0x64, 0x54, 0x61, 0x67, 0x73, 0x12, 0x41, 0x0a, 0x0b, 0x67, 0x65, 0x6e, 0x65, 0x72,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x48, 0x01, 0x52, 0x0b, 0x67, 0x65, 0x6e, 0x65,
	0x72, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x88, 0x01, 0x01, 0x12, 0x1b, 0x0a, 0x09, 0x72, 0x65,
	0x70, 0x6f, 0x5f, 0x74, 0x61, 0x67, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x72,
	0x65, 0x70, 0x6f, 0x54, 0x61, 0x67, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e, 0x55, 0x73, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x69, 0x6e, 0x55, 0x73, 0x65, 0x12, 0x4e, 0x0a,
	0x12, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x75, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x48, 0x02, 0x52, 0x12, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x88, 0x01, 0x01, 0x12, 0x33, 0x0a,
	0x09, 0x63, 0x79, 0x63, 0x6c, 0x6f, 0x6e, 0x65, 0x64, 0x78, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x13, 0x2e, 0x63, 0x79, 0x63, 0x6c, 0x6f, 0x6e, 0x65, 0x64, 0x78, 0x2e, 0x76, 0x31, 0x5f,
	0x34, 0x2e, 0x42, 0x6f, 0x6d, 0x48, 0x00, 0x52, 0x09, 0x63, 0x79, 0x63, 0x6c, 0x6f, 0x6e, 0x65,
	0x64, 0x78, 0x42, 0x06, 0x0a, 0x04, 0x73, 0x62, 0x6f, 0x6d, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x67,
	0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x42, 0x15, 0x0a, 0x13, 0x5f, 0x67,
	0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2a, 0x6e, 0x0a, 0x0e, 0x53, 0x42, 0x4f, 0x4d, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49,
	0x45, 0x44, 0x10, 0x00, 0x12, 0x1a, 0x0a, 0x16, 0x43, 0x4f, 0x4e, 0x54, 0x41, 0x49, 0x4e, 0x45,
	0x52, 0x5f, 0x49, 0x4d, 0x41, 0x47, 0x45, 0x5f, 0x4c, 0x41, 0x59, 0x45, 0x52, 0x53, 0x10, 0x01,
	0x12, 0x19, 0x0a, 0x15, 0x43, 0x4f, 0x4e, 0x54, 0x41, 0x49, 0x4e, 0x45, 0x52, 0x5f, 0x46, 0x49,
	0x4c, 0x45, 0x5f, 0x53, 0x59, 0x53, 0x54, 0x45, 0x4d, 0x10, 0x02, 0x12, 0x14, 0x0a, 0x10, 0x48,
	0x4f, 0x53, 0x54, 0x5f, 0x46, 0x49, 0x4c, 0x45, 0x5f, 0x53, 0x59, 0x53, 0x54, 0x45, 0x4d, 0x10,
	0x03, 0x42, 0x2a, 0x5a, 0x28, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x44, 0x61, 0x74, 0x61, 0x44, 0x6f, 0x67, 0x2f, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2d, 0x70, 0x61,
	0x79, 0x6c, 0x6f, 0x61, 0x64, 0x2f, 0x76, 0x35, 0x2f, 0x73, 0x62, 0x6f, 0x6d, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_sbom_sbom_proto_rawDescOnce sync.Once
	file_proto_sbom_sbom_proto_rawDescData = file_proto_sbom_sbom_proto_rawDesc
)

func file_proto_sbom_sbom_proto_rawDescGZIP() []byte {
	file_proto_sbom_sbom_proto_rawDescOnce.Do(func() {
		file_proto_sbom_sbom_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_sbom_sbom_proto_rawDescData)
	})
	return file_proto_sbom_sbom_proto_rawDescData
}

var file_proto_sbom_sbom_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_proto_sbom_sbom_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_sbom_sbom_proto_goTypes = []interface{}{
	(SBOMSourceType)(0),         // 0: datadog.sbom.SBOMSourceType
	(*SBOMPayload)(nil),         // 1: datadog.sbom.SBOMPayload
	(*SBOMEntity)(nil),          // 2: datadog.sbom.SBOMEntity
	(*timestamp.Timestamp)(nil), // 3: google.protobuf.Timestamp
	(*duration.Duration)(nil),   // 4: google.protobuf.Duration
	(*cyclonedx_v1_4.Bom)(nil),  // 5: cyclonedx.v1_4.Bom
}
var file_proto_sbom_sbom_proto_depIdxs = []int32{
	2, // 0: datadog.sbom.SBOMPayload.entities:type_name -> datadog.sbom.SBOMEntity
	0, // 1: datadog.sbom.SBOMEntity.type:type_name -> datadog.sbom.SBOMSourceType
	3, // 2: datadog.sbom.SBOMEntity.generatedAt:type_name -> google.protobuf.Timestamp
	4, // 3: datadog.sbom.SBOMEntity.generationDuration:type_name -> google.protobuf.Duration
	5, // 4: datadog.sbom.SBOMEntity.cyclonedx:type_name -> cyclonedx.v1_4.Bom
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_proto_sbom_sbom_proto_init() }
func file_proto_sbom_sbom_proto_init() {
	if File_proto_sbom_sbom_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_sbom_sbom_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SBOMPayload); i {
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
		file_proto_sbom_sbom_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SBOMEntity); i {
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
	file_proto_sbom_sbom_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_proto_sbom_sbom_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*SBOMEntity_Cyclonedx)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_sbom_sbom_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_sbom_sbom_proto_goTypes,
		DependencyIndexes: file_proto_sbom_sbom_proto_depIdxs,
		EnumInfos:         file_proto_sbom_sbom_proto_enumTypes,
		MessageInfos:      file_proto_sbom_sbom_proto_msgTypes,
	}.Build()
	File_proto_sbom_sbom_proto = out.File
	file_proto_sbom_sbom_proto_rawDesc = nil
	file_proto_sbom_sbom_proto_goTypes = nil
	file_proto_sbom_sbom_proto_depIdxs = nil
}
