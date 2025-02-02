// Code generated by skv2. DO NOT EDIT.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.6.1
// source: github.com/solo-io/solo-apis/api/gloo-fed/fed.rpc/v1/enterprise_gloo_resources.proto

package v1

import (
	context "context"
	reflect "reflect"
	sync "sync"

	_ "github.com/solo-io/protoc-gen-ext/extproto"
	v11 "github.com/solo-io/skv2/pkg/api/core.skv2.solo.io/v1"
	v1 "github.com/solo-io/solo-apis/pkg/api/enterprise.gloo.solo.io/v1"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type AuthConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Metadata     *ObjectMeta          `protobuf:"bytes,1,opt,name=metadata,proto3" json:"metadata,omitempty"`
	Spec         *v1.AuthConfigSpec   `protobuf:"bytes,2,opt,name=spec,proto3" json:"spec,omitempty"`
	Status       *v1.AuthConfigStatus `protobuf:"bytes,3,opt,name=status,proto3" json:"status,omitempty"`
	GlooInstance *v11.ObjectRef       `protobuf:"bytes,4,opt,name=gloo_instance,json=glooInstance,proto3" json:"gloo_instance,omitempty"`
}

func (x *AuthConfig) Reset() {
	*x = AuthConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_solo_apis_api_gloo_fed_fed_rpc_v1_enterprise_gloo_resources_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthConfig) ProtoMessage() {}

func (x *AuthConfig) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_solo_apis_api_gloo_fed_fed_rpc_v1_enterprise_gloo_resources_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthConfig.ProtoReflect.Descriptor instead.
func (*AuthConfig) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_solo_apis_api_gloo_fed_fed_rpc_v1_enterprise_gloo_resources_proto_rawDescGZIP(), []int{0}
}

func (x *AuthConfig) GetMetadata() *ObjectMeta {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *AuthConfig) GetSpec() *v1.AuthConfigSpec {
	if x != nil {
		return x.Spec
	}
	return nil
}

func (x *AuthConfig) GetStatus() *v1.AuthConfigStatus {
	if x != nil {
		return x.Status
	}
	return nil
}

func (x *AuthConfig) GetGlooInstance() *v11.ObjectRef {
	if x != nil {
		return x.GlooInstance
	}
	return nil
}

type ListAuthConfigsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GlooInstanceRef *v11.ObjectRef `protobuf:"bytes,1,opt,name=gloo_instance_ref,json=glooInstanceRef,proto3" json:"gloo_instance_ref,omitempty"`
}

func (x *ListAuthConfigsRequest) Reset() {
	*x = ListAuthConfigsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_solo_apis_api_gloo_fed_fed_rpc_v1_enterprise_gloo_resources_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListAuthConfigsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAuthConfigsRequest) ProtoMessage() {}

func (x *ListAuthConfigsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_solo_apis_api_gloo_fed_fed_rpc_v1_enterprise_gloo_resources_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListAuthConfigsRequest.ProtoReflect.Descriptor instead.
func (*ListAuthConfigsRequest) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_solo_apis_api_gloo_fed_fed_rpc_v1_enterprise_gloo_resources_proto_rawDescGZIP(), []int{1}
}

func (x *ListAuthConfigsRequest) GetGlooInstanceRef() *v11.ObjectRef {
	if x != nil {
		return x.GlooInstanceRef
	}
	return nil
}

type ListAuthConfigsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AuthConfigs []*AuthConfig `protobuf:"bytes,1,rep,name=auth_configs,json=authConfigs,proto3" json:"auth_configs,omitempty"`
}

func (x *ListAuthConfigsResponse) Reset() {
	*x = ListAuthConfigsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_solo_apis_api_gloo_fed_fed_rpc_v1_enterprise_gloo_resources_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListAuthConfigsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAuthConfigsResponse) ProtoMessage() {}

func (x *ListAuthConfigsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_solo_apis_api_gloo_fed_fed_rpc_v1_enterprise_gloo_resources_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListAuthConfigsResponse.ProtoReflect.Descriptor instead.
func (*ListAuthConfigsResponse) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_solo_apis_api_gloo_fed_fed_rpc_v1_enterprise_gloo_resources_proto_rawDescGZIP(), []int{2}
}

func (x *ListAuthConfigsResponse) GetAuthConfigs() []*AuthConfig {
	if x != nil {
		return x.AuthConfigs
	}
	return nil
}

type GetAuthConfigYamlRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AuthConfigRef *v11.ClusterObjectRef `protobuf:"bytes,1,opt,name=auth_config_ref,json=authConfigRef,proto3" json:"auth_config_ref,omitempty"`
}

func (x *GetAuthConfigYamlRequest) Reset() {
	*x = GetAuthConfigYamlRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_solo_apis_api_gloo_fed_fed_rpc_v1_enterprise_gloo_resources_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAuthConfigYamlRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAuthConfigYamlRequest) ProtoMessage() {}

func (x *GetAuthConfigYamlRequest) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_solo_apis_api_gloo_fed_fed_rpc_v1_enterprise_gloo_resources_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAuthConfigYamlRequest.ProtoReflect.Descriptor instead.
func (*GetAuthConfigYamlRequest) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_solo_apis_api_gloo_fed_fed_rpc_v1_enterprise_gloo_resources_proto_rawDescGZIP(), []int{3}
}

func (x *GetAuthConfigYamlRequest) GetAuthConfigRef() *v11.ClusterObjectRef {
	if x != nil {
		return x.AuthConfigRef
	}
	return nil
}

type GetAuthConfigYamlResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	YamlData *ResourceYaml `protobuf:"bytes,1,opt,name=yaml_data,json=yamlData,proto3" json:"yaml_data,omitempty"`
}

func (x *GetAuthConfigYamlResponse) Reset() {
	*x = GetAuthConfigYamlResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_solo_apis_api_gloo_fed_fed_rpc_v1_enterprise_gloo_resources_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAuthConfigYamlResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAuthConfigYamlResponse) ProtoMessage() {}

func (x *GetAuthConfigYamlResponse) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_solo_apis_api_gloo_fed_fed_rpc_v1_enterprise_gloo_resources_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAuthConfigYamlResponse.ProtoReflect.Descriptor instead.
func (*GetAuthConfigYamlResponse) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_solo_apis_api_gloo_fed_fed_rpc_v1_enterprise_gloo_resources_proto_rawDescGZIP(), []int{4}
}

func (x *GetAuthConfigYamlResponse) GetYamlData() *ResourceYaml {
	if x != nil {
		return x.YamlData
	}
	return nil
}

var File_github_com_solo_io_solo_apis_api_gloo_fed_fed_rpc_v1_enterprise_gloo_resources_proto protoreflect.FileDescriptor

var file_github_com_solo_io_solo_apis_api_gloo_fed_fed_rpc_v1_enterprise_gloo_resources_proto_rawDesc = []byte{
	0x0a, 0x54, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c,
	0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2d, 0x66, 0x65, 0x64, 0x2f, 0x66, 0x65, 0x64, 0x2e,
	0x72, 0x70, 0x63, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x72, 0x69, 0x73,
	0x65, 0x5f, 0x67, 0x6c, 0x6f, 0x6f, 0x5f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x66, 0x65, 0x64, 0x2e, 0x72, 0x70, 0x63, 0x2e,
	0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x1a, 0x12, 0x65, 0x78, 0x74, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x65, 0x78, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x4a, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f,
	0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x6c,
	0x6f, 0x6f, 0x2f, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x72, 0x69, 0x73, 0x65, 0x2e, 0x67, 0x6c,
	0x6f, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x73, 0x6b, 0x76, 0x32,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x72,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x41, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x73, 0x6f, 0x6c, 0x6f,
	0x2d, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2d, 0x66,
	0x65, 0x64, 0x2f, 0x66, 0x65, 0x64, 0x2e, 0x72, 0x70, 0x63, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x88, 0x02, 0x0a, 0x0a, 0x41,
	0x75, 0x74, 0x68, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x37, 0x0a, 0x08, 0x6d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x66, 0x65,
	0x64, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x4f, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x12, 0x3b, 0x0a, 0x04, 0x73, 0x70, 0x65, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x27, 0x2e, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x72, 0x69, 0x73, 0x65, 0x2e, 0x67, 0x6c,
	0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x53, 0x70, 0x65, 0x63, 0x52, 0x04, 0x73, 0x70, 0x65, 0x63, 0x12,
	0x41, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x29, 0x2e, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x72, 0x69, 0x73, 0x65, 0x2e, 0x67, 0x6c, 0x6f,
	0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x41, 0x0a, 0x0d, 0x67, 0x6c, 0x6f, 0x6f, 0x5f, 0x69, 0x6e, 0x73, 0x74, 0x61,
	0x6e, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x63, 0x6f, 0x72, 0x65,
	0x2e, 0x73, 0x6b, 0x76, 0x32, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x4f, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x66, 0x52, 0x0c, 0x67, 0x6c, 0x6f, 0x6f, 0x49, 0x6e, 0x73,
	0x74, 0x61, 0x6e, 0x63, 0x65, 0x22, 0x62, 0x0a, 0x16, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x75, 0x74,
	0x68, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x48, 0x0a, 0x11, 0x67, 0x6c, 0x6f, 0x6f, 0x5f, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65,
	0x5f, 0x72, 0x65, 0x66, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x63, 0x6f, 0x72,
	0x65, 0x2e, 0x73, 0x6b, 0x76, 0x32, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x4f,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x66, 0x52, 0x0f, 0x67, 0x6c, 0x6f, 0x6f, 0x49, 0x6e,
	0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x66, 0x22, 0x59, 0x0a, 0x17, 0x4c, 0x69, 0x73,
	0x74, 0x41, 0x75, 0x74, 0x68, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3e, 0x0a, 0x0c, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x66, 0x65, 0x64,
	0x2e, 0x72, 0x70, 0x63, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x41, 0x75, 0x74,
	0x68, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x0b, 0x61, 0x75, 0x74, 0x68, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x73, 0x22, 0x67, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x41, 0x75, 0x74, 0x68, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x59, 0x61, 0x6d, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x4b, 0x0a, 0x0f, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f,
	0x72, 0x65, 0x66, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x63, 0x6f, 0x72, 0x65,
	0x2e, 0x73, 0x6b, 0x76, 0x32, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x43, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x66, 0x52, 0x0d,
	0x61, 0x75, 0x74, 0x68, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x66, 0x22, 0x57, 0x0a,
	0x19, 0x47, 0x65, 0x74, 0x41, 0x75, 0x74, 0x68, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x59, 0x61,
	0x6d, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3a, 0x0a, 0x09, 0x79, 0x61,
	0x6d, 0x6c, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e,
	0x66, 0x65, 0x64, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e,
	0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x59, 0x61, 0x6d, 0x6c, 0x52, 0x08, 0x79, 0x61,
	0x6d, 0x6c, 0x44, 0x61, 0x74, 0x61, 0x32, 0xf1, 0x01, 0x0a, 0x19, 0x45, 0x6e, 0x74, 0x65, 0x72,
	0x70, 0x72, 0x69, 0x73, 0x65, 0x47, 0x6c, 0x6f, 0x6f, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x41, 0x70, 0x69, 0x12, 0x66, 0x0a, 0x0f, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x75, 0x74, 0x68,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x12, 0x27, 0x2e, 0x66, 0x65, 0x64, 0x2e, 0x72, 0x70,
	0x63, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x75,
	0x74, 0x68, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x28, 0x2e, 0x66, 0x65, 0x64, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e,
	0x69, 0x6f, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x75, 0x74, 0x68, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x6c, 0x0a, 0x11,
	0x47, 0x65, 0x74, 0x41, 0x75, 0x74, 0x68, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x59, 0x61, 0x6d,
	0x6c, 0x12, 0x29, 0x2e, 0x66, 0x65, 0x64, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x73, 0x6f, 0x6c, 0x6f,
	0x2e, 0x69, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x75, 0x74, 0x68, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x59, 0x61, 0x6d, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2a, 0x2e, 0x66,
	0x65, 0x64, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x47,
	0x65, 0x74, 0x41, 0x75, 0x74, 0x68, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x59, 0x61, 0x6d, 0x6c,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x39, 0x5a, 0x2f, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f,
	0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x66, 0x65, 0x64, 0x2e, 0x72, 0x70, 0x63, 0x2f, 0x76, 0x31, 0xb8, 0xf5, 0x04,
	0x01, 0xc0, 0xf5, 0x04, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_github_com_solo_io_solo_apis_api_gloo_fed_fed_rpc_v1_enterprise_gloo_resources_proto_rawDescOnce sync.Once
	file_github_com_solo_io_solo_apis_api_gloo_fed_fed_rpc_v1_enterprise_gloo_resources_proto_rawDescData = file_github_com_solo_io_solo_apis_api_gloo_fed_fed_rpc_v1_enterprise_gloo_resources_proto_rawDesc
)

func file_github_com_solo_io_solo_apis_api_gloo_fed_fed_rpc_v1_enterprise_gloo_resources_proto_rawDescGZIP() []byte {
	file_github_com_solo_io_solo_apis_api_gloo_fed_fed_rpc_v1_enterprise_gloo_resources_proto_rawDescOnce.Do(func() {
		file_github_com_solo_io_solo_apis_api_gloo_fed_fed_rpc_v1_enterprise_gloo_resources_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_solo_io_solo_apis_api_gloo_fed_fed_rpc_v1_enterprise_gloo_resources_proto_rawDescData)
	})
	return file_github_com_solo_io_solo_apis_api_gloo_fed_fed_rpc_v1_enterprise_gloo_resources_proto_rawDescData
}

var file_github_com_solo_io_solo_apis_api_gloo_fed_fed_rpc_v1_enterprise_gloo_resources_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_github_com_solo_io_solo_apis_api_gloo_fed_fed_rpc_v1_enterprise_gloo_resources_proto_goTypes = []interface{}{
	(*AuthConfig)(nil),                // 0: fed.rpc.solo.io.AuthConfig
	(*ListAuthConfigsRequest)(nil),    // 1: fed.rpc.solo.io.ListAuthConfigsRequest
	(*ListAuthConfigsResponse)(nil),   // 2: fed.rpc.solo.io.ListAuthConfigsResponse
	(*GetAuthConfigYamlRequest)(nil),  // 3: fed.rpc.solo.io.GetAuthConfigYamlRequest
	(*GetAuthConfigYamlResponse)(nil), // 4: fed.rpc.solo.io.GetAuthConfigYamlResponse
	(*ObjectMeta)(nil),                // 5: fed.rpc.solo.io.ObjectMeta
	(*v1.AuthConfigSpec)(nil),         // 6: enterprise.gloo.solo.io.AuthConfigSpec
	(*v1.AuthConfigStatus)(nil),       // 7: enterprise.gloo.solo.io.AuthConfigStatus
	(*v11.ObjectRef)(nil),             // 8: core.skv2.solo.io.ObjectRef
	(*v11.ClusterObjectRef)(nil),      // 9: core.skv2.solo.io.ClusterObjectRef
	(*ResourceYaml)(nil),              // 10: fed.rpc.solo.io.ResourceYaml
}
var file_github_com_solo_io_solo_apis_api_gloo_fed_fed_rpc_v1_enterprise_gloo_resources_proto_depIdxs = []int32{
	5,  // 0: fed.rpc.solo.io.AuthConfig.metadata:type_name -> fed.rpc.solo.io.ObjectMeta
	6,  // 1: fed.rpc.solo.io.AuthConfig.spec:type_name -> enterprise.gloo.solo.io.AuthConfigSpec
	7,  // 2: fed.rpc.solo.io.AuthConfig.status:type_name -> enterprise.gloo.solo.io.AuthConfigStatus
	8,  // 3: fed.rpc.solo.io.AuthConfig.gloo_instance:type_name -> core.skv2.solo.io.ObjectRef
	8,  // 4: fed.rpc.solo.io.ListAuthConfigsRequest.gloo_instance_ref:type_name -> core.skv2.solo.io.ObjectRef
	0,  // 5: fed.rpc.solo.io.ListAuthConfigsResponse.auth_configs:type_name -> fed.rpc.solo.io.AuthConfig
	9,  // 6: fed.rpc.solo.io.GetAuthConfigYamlRequest.auth_config_ref:type_name -> core.skv2.solo.io.ClusterObjectRef
	10, // 7: fed.rpc.solo.io.GetAuthConfigYamlResponse.yaml_data:type_name -> fed.rpc.solo.io.ResourceYaml
	1,  // 8: fed.rpc.solo.io.EnterpriseGlooResourceApi.ListAuthConfigs:input_type -> fed.rpc.solo.io.ListAuthConfigsRequest
	3,  // 9: fed.rpc.solo.io.EnterpriseGlooResourceApi.GetAuthConfigYaml:input_type -> fed.rpc.solo.io.GetAuthConfigYamlRequest
	2,  // 10: fed.rpc.solo.io.EnterpriseGlooResourceApi.ListAuthConfigs:output_type -> fed.rpc.solo.io.ListAuthConfigsResponse
	4,  // 11: fed.rpc.solo.io.EnterpriseGlooResourceApi.GetAuthConfigYaml:output_type -> fed.rpc.solo.io.GetAuthConfigYamlResponse
	10, // [10:12] is the sub-list for method output_type
	8,  // [8:10] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() {
	file_github_com_solo_io_solo_apis_api_gloo_fed_fed_rpc_v1_enterprise_gloo_resources_proto_init()
}
func file_github_com_solo_io_solo_apis_api_gloo_fed_fed_rpc_v1_enterprise_gloo_resources_proto_init() {
	if File_github_com_solo_io_solo_apis_api_gloo_fed_fed_rpc_v1_enterprise_gloo_resources_proto != nil {
		return
	}
	file_github_com_solo_io_solo_apis_api_gloo_fed_fed_rpc_v1_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_github_com_solo_io_solo_apis_api_gloo_fed_fed_rpc_v1_enterprise_gloo_resources_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthConfig); i {
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
		file_github_com_solo_io_solo_apis_api_gloo_fed_fed_rpc_v1_enterprise_gloo_resources_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListAuthConfigsRequest); i {
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
		file_github_com_solo_io_solo_apis_api_gloo_fed_fed_rpc_v1_enterprise_gloo_resources_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListAuthConfigsResponse); i {
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
		file_github_com_solo_io_solo_apis_api_gloo_fed_fed_rpc_v1_enterprise_gloo_resources_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAuthConfigYamlRequest); i {
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
		file_github_com_solo_io_solo_apis_api_gloo_fed_fed_rpc_v1_enterprise_gloo_resources_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAuthConfigYamlResponse); i {
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
			RawDescriptor: file_github_com_solo_io_solo_apis_api_gloo_fed_fed_rpc_v1_enterprise_gloo_resources_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_github_com_solo_io_solo_apis_api_gloo_fed_fed_rpc_v1_enterprise_gloo_resources_proto_goTypes,
		DependencyIndexes: file_github_com_solo_io_solo_apis_api_gloo_fed_fed_rpc_v1_enterprise_gloo_resources_proto_depIdxs,
		MessageInfos:      file_github_com_solo_io_solo_apis_api_gloo_fed_fed_rpc_v1_enterprise_gloo_resources_proto_msgTypes,
	}.Build()
	File_github_com_solo_io_solo_apis_api_gloo_fed_fed_rpc_v1_enterprise_gloo_resources_proto = out.File
	file_github_com_solo_io_solo_apis_api_gloo_fed_fed_rpc_v1_enterprise_gloo_resources_proto_rawDesc = nil
	file_github_com_solo_io_solo_apis_api_gloo_fed_fed_rpc_v1_enterprise_gloo_resources_proto_goTypes = nil
	file_github_com_solo_io_solo_apis_api_gloo_fed_fed_rpc_v1_enterprise_gloo_resources_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// EnterpriseGlooResourceApiClient is the client API for EnterpriseGlooResourceApi service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type EnterpriseGlooResourceApiClient interface {
	ListAuthConfigs(ctx context.Context, in *ListAuthConfigsRequest, opts ...grpc.CallOption) (*ListAuthConfigsResponse, error)
	GetAuthConfigYaml(ctx context.Context, in *GetAuthConfigYamlRequest, opts ...grpc.CallOption) (*GetAuthConfigYamlResponse, error)
}

type enterpriseGlooResourceApiClient struct {
	cc grpc.ClientConnInterface
}

func NewEnterpriseGlooResourceApiClient(cc grpc.ClientConnInterface) EnterpriseGlooResourceApiClient {
	return &enterpriseGlooResourceApiClient{cc}
}

func (c *enterpriseGlooResourceApiClient) ListAuthConfigs(ctx context.Context, in *ListAuthConfigsRequest, opts ...grpc.CallOption) (*ListAuthConfigsResponse, error) {
	out := new(ListAuthConfigsResponse)
	err := c.cc.Invoke(ctx, "/fed.rpc.solo.io.EnterpriseGlooResourceApi/ListAuthConfigs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *enterpriseGlooResourceApiClient) GetAuthConfigYaml(ctx context.Context, in *GetAuthConfigYamlRequest, opts ...grpc.CallOption) (*GetAuthConfigYamlResponse, error) {
	out := new(GetAuthConfigYamlResponse)
	err := c.cc.Invoke(ctx, "/fed.rpc.solo.io.EnterpriseGlooResourceApi/GetAuthConfigYaml", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EnterpriseGlooResourceApiServer is the server API for EnterpriseGlooResourceApi service.
type EnterpriseGlooResourceApiServer interface {
	ListAuthConfigs(context.Context, *ListAuthConfigsRequest) (*ListAuthConfigsResponse, error)
	GetAuthConfigYaml(context.Context, *GetAuthConfigYamlRequest) (*GetAuthConfigYamlResponse, error)
}

// UnimplementedEnterpriseGlooResourceApiServer can be embedded to have forward compatible implementations.
type UnimplementedEnterpriseGlooResourceApiServer struct {
}

func (*UnimplementedEnterpriseGlooResourceApiServer) ListAuthConfigs(context.Context, *ListAuthConfigsRequest) (*ListAuthConfigsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAuthConfigs not implemented")
}
func (*UnimplementedEnterpriseGlooResourceApiServer) GetAuthConfigYaml(context.Context, *GetAuthConfigYamlRequest) (*GetAuthConfigYamlResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAuthConfigYaml not implemented")
}

func RegisterEnterpriseGlooResourceApiServer(s *grpc.Server, srv EnterpriseGlooResourceApiServer) {
	s.RegisterService(&_EnterpriseGlooResourceApi_serviceDesc, srv)
}

func _EnterpriseGlooResourceApi_ListAuthConfigs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListAuthConfigsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EnterpriseGlooResourceApiServer).ListAuthConfigs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fed.rpc.solo.io.EnterpriseGlooResourceApi/ListAuthConfigs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EnterpriseGlooResourceApiServer).ListAuthConfigs(ctx, req.(*ListAuthConfigsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EnterpriseGlooResourceApi_GetAuthConfigYaml_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAuthConfigYamlRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EnterpriseGlooResourceApiServer).GetAuthConfigYaml(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fed.rpc.solo.io.EnterpriseGlooResourceApi/GetAuthConfigYaml",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EnterpriseGlooResourceApiServer).GetAuthConfigYaml(ctx, req.(*GetAuthConfigYamlRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _EnterpriseGlooResourceApi_serviceDesc = grpc.ServiceDesc{
	ServiceName: "fed.rpc.solo.io.EnterpriseGlooResourceApi",
	HandlerType: (*EnterpriseGlooResourceApiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListAuthConfigs",
			Handler:    _EnterpriseGlooResourceApi_ListAuthConfigs_Handler,
		},
		{
			MethodName: "GetAuthConfigYaml",
			Handler:    _EnterpriseGlooResourceApi_GetAuthConfigYaml_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/solo-io/solo-apis/api/gloo-fed/fed.rpc/v1/enterprise_gloo_resources.proto",
}
