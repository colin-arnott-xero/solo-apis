// Code generated by skv2. DO NOT EDIT.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.6.1
// source: github.com/solo-io/solo-apis/api/gloo-fed/fed.rpc/v1/federated_enterprise_gloo_resources.proto

package v1

import (
	context "context"
	reflect "reflect"
	sync "sync"

	_ "github.com/solo-io/protoc-gen-ext/extproto"
	v1 "github.com/solo-io/skv2/pkg/api/core.skv2.solo.io/v1"
	types "github.com/solo-io/solo-apis/pkg/api/fed.enterprise.gloo.solo.io/v1/types"
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

type FederatedAuthConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Metadata *ObjectMeta                      `protobuf:"bytes,1,opt,name=metadata,proto3" json:"metadata,omitempty"`
	Spec     *types.FederatedAuthConfigSpec   `protobuf:"bytes,2,opt,name=spec,proto3" json:"spec,omitempty"`
	Status   *types.FederatedAuthConfigStatus `protobuf:"bytes,3,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *FederatedAuthConfig) Reset() {
	*x = FederatedAuthConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_solo_apis_api_gloo_fed_fed_rpc_v1_federated_enterprise_gloo_resources_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FederatedAuthConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FederatedAuthConfig) ProtoMessage() {}

func (x *FederatedAuthConfig) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_solo_apis_api_gloo_fed_fed_rpc_v1_federated_enterprise_gloo_resources_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FederatedAuthConfig.ProtoReflect.Descriptor instead.
func (*FederatedAuthConfig) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_solo_apis_api_gloo_fed_fed_rpc_v1_federated_enterprise_gloo_resources_proto_rawDescGZIP(), []int{0}
}

func (x *FederatedAuthConfig) GetMetadata() *ObjectMeta {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *FederatedAuthConfig) GetSpec() *types.FederatedAuthConfigSpec {
	if x != nil {
		return x.Spec
	}
	return nil
}

func (x *FederatedAuthConfig) GetStatus() *types.FederatedAuthConfigStatus {
	if x != nil {
		return x.Status
	}
	return nil
}

type ListFederatedAuthConfigsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListFederatedAuthConfigsRequest) Reset() {
	*x = ListFederatedAuthConfigsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_solo_apis_api_gloo_fed_fed_rpc_v1_federated_enterprise_gloo_resources_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListFederatedAuthConfigsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListFederatedAuthConfigsRequest) ProtoMessage() {}

func (x *ListFederatedAuthConfigsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_solo_apis_api_gloo_fed_fed_rpc_v1_federated_enterprise_gloo_resources_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListFederatedAuthConfigsRequest.ProtoReflect.Descriptor instead.
func (*ListFederatedAuthConfigsRequest) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_solo_apis_api_gloo_fed_fed_rpc_v1_federated_enterprise_gloo_resources_proto_rawDescGZIP(), []int{1}
}

type ListFederatedAuthConfigsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FederatedAuthConfigs []*FederatedAuthConfig `protobuf:"bytes,1,rep,name=federated_auth_configs,json=federatedAuthConfigs,proto3" json:"federated_auth_configs,omitempty"`
}

func (x *ListFederatedAuthConfigsResponse) Reset() {
	*x = ListFederatedAuthConfigsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_solo_apis_api_gloo_fed_fed_rpc_v1_federated_enterprise_gloo_resources_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListFederatedAuthConfigsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListFederatedAuthConfigsResponse) ProtoMessage() {}

func (x *ListFederatedAuthConfigsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_solo_apis_api_gloo_fed_fed_rpc_v1_federated_enterprise_gloo_resources_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListFederatedAuthConfigsResponse.ProtoReflect.Descriptor instead.
func (*ListFederatedAuthConfigsResponse) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_solo_apis_api_gloo_fed_fed_rpc_v1_federated_enterprise_gloo_resources_proto_rawDescGZIP(), []int{2}
}

func (x *ListFederatedAuthConfigsResponse) GetFederatedAuthConfigs() []*FederatedAuthConfig {
	if x != nil {
		return x.FederatedAuthConfigs
	}
	return nil
}

type GetFederatedAuthConfigYamlRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FederatedAuthConfigRef *v1.ObjectRef `protobuf:"bytes,1,opt,name=federated_auth_config_ref,json=federatedAuthConfigRef,proto3" json:"federated_auth_config_ref,omitempty"`
}

func (x *GetFederatedAuthConfigYamlRequest) Reset() {
	*x = GetFederatedAuthConfigYamlRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_solo_apis_api_gloo_fed_fed_rpc_v1_federated_enterprise_gloo_resources_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFederatedAuthConfigYamlRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFederatedAuthConfigYamlRequest) ProtoMessage() {}

func (x *GetFederatedAuthConfigYamlRequest) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_solo_apis_api_gloo_fed_fed_rpc_v1_federated_enterprise_gloo_resources_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFederatedAuthConfigYamlRequest.ProtoReflect.Descriptor instead.
func (*GetFederatedAuthConfigYamlRequest) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_solo_apis_api_gloo_fed_fed_rpc_v1_federated_enterprise_gloo_resources_proto_rawDescGZIP(), []int{3}
}

func (x *GetFederatedAuthConfigYamlRequest) GetFederatedAuthConfigRef() *v1.ObjectRef {
	if x != nil {
		return x.FederatedAuthConfigRef
	}
	return nil
}

type GetFederatedAuthConfigYamlResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	YamlData *ResourceYaml `protobuf:"bytes,1,opt,name=yaml_data,json=yamlData,proto3" json:"yaml_data,omitempty"`
}

func (x *GetFederatedAuthConfigYamlResponse) Reset() {
	*x = GetFederatedAuthConfigYamlResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_solo_apis_api_gloo_fed_fed_rpc_v1_federated_enterprise_gloo_resources_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFederatedAuthConfigYamlResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFederatedAuthConfigYamlResponse) ProtoMessage() {}

func (x *GetFederatedAuthConfigYamlResponse) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_solo_apis_api_gloo_fed_fed_rpc_v1_federated_enterprise_gloo_resources_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFederatedAuthConfigYamlResponse.ProtoReflect.Descriptor instead.
func (*GetFederatedAuthConfigYamlResponse) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_solo_apis_api_gloo_fed_fed_rpc_v1_federated_enterprise_gloo_resources_proto_rawDescGZIP(), []int{4}
}

func (x *GetFederatedAuthConfigYamlResponse) GetYamlData() *ResourceYaml {
	if x != nil {
		return x.YamlData
	}
	return nil
}

var File_github_com_solo_io_solo_apis_api_gloo_fed_fed_rpc_v1_federated_enterprise_gloo_resources_proto protoreflect.FileDescriptor

var file_github_com_solo_io_solo_apis_api_gloo_fed_fed_rpc_v1_federated_enterprise_gloo_resources_proto_rawDesc = []byte{
	0x0a, 0x5e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c,
	0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2d, 0x66, 0x65, 0x64, 0x2f, 0x66, 0x65, 0x64, 0x2e,
	0x72, 0x70, 0x63, 0x2f, 0x76, 0x31, 0x2f, 0x66, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64,
	0x5f, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x72, 0x69, 0x73, 0x65, 0x5f, 0x67, 0x6c, 0x6f, 0x6f,
	0x5f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0f, 0x66, 0x65, 0x64, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69,
	0x6f, 0x1a, 0x12, 0x65, 0x78, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x78, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x52, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x61,
	0x70, 0x69, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2d, 0x66, 0x65, 0x64,
	0x2f, 0x66, 0x65, 0x64, 0x2e, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x72, 0x69, 0x73, 0x65, 0x2e,
	0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x73, 0x6b,
	0x76, 0x32, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x63,
	0x6f, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x41, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x73, 0x6f,
	0x6c, 0x6f, 0x2d, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x6c, 0x6f, 0x6f,
	0x2d, 0x66, 0x65, 0x64, 0x2f, 0x66, 0x65, 0x64, 0x2e, 0x72, 0x70, 0x63, 0x2f, 0x76, 0x31, 0x2f,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe8, 0x01, 0x0a,
	0x13, 0x46, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x41, 0x75, 0x74, 0x68, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x12, 0x37, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x66, 0x65, 0x64, 0x2e, 0x72, 0x70, 0x63,
	0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x4d,
	0x65, 0x74, 0x61, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x48, 0x0a,
	0x04, 0x73, 0x70, 0x65, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x34, 0x2e, 0x66, 0x65,
	0x64, 0x2e, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x72, 0x69, 0x73, 0x65, 0x2e, 0x67, 0x6c, 0x6f,
	0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x46, 0x65, 0x64, 0x65, 0x72, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x75, 0x74, 0x68, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x53, 0x70, 0x65,
	0x63, 0x52, 0x04, 0x73, 0x70, 0x65, 0x63, 0x12, 0x4e, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x36, 0x2e, 0x66, 0x65, 0x64, 0x2e, 0x65, 0x6e,
	0x74, 0x65, 0x72, 0x70, 0x72, 0x69, 0x73, 0x65, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f,
	0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x46, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x75, 0x74, 0x68, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x21, 0x0a, 0x1f, 0x4c, 0x69, 0x73, 0x74, 0x46,
	0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x41, 0x75, 0x74, 0x68, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x7e, 0x0a, 0x20, 0x4c, 0x69,
	0x73, 0x74, 0x46, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x41, 0x75, 0x74, 0x68, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5a,
	0x0a, 0x16, 0x66, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x75, 0x74, 0x68,
	0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24,
	0x2e, 0x66, 0x65, 0x64, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f,
	0x2e, 0x46, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x41, 0x75, 0x74, 0x68, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x52, 0x14, 0x66, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x75, 0x74, 0x68, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x22, 0x7c, 0x0a, 0x21, 0x47, 0x65,
	0x74, 0x46, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x41, 0x75, 0x74, 0x68, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x59, 0x61, 0x6d, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x57, 0x0a, 0x19, 0x66, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x75, 0x74,
	0x68, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f, 0x72, 0x65, 0x66, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x73, 0x6b, 0x76, 0x32, 0x2e, 0x73,
	0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x66,
	0x52, 0x16, 0x66, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x41, 0x75, 0x74, 0x68, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x66, 0x22, 0x60, 0x0a, 0x22, 0x47, 0x65, 0x74, 0x46,
	0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x41, 0x75, 0x74, 0x68, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x59, 0x61, 0x6d, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3a,
	0x0a, 0x09, 0x79, 0x61, 0x6d, 0x6c, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1d, 0x2e, 0x66, 0x65, 0x64, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x73, 0x6f, 0x6c, 0x6f,
	0x2e, 0x69, 0x6f, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x59, 0x61, 0x6d, 0x6c,
	0x52, 0x08, 0x79, 0x61, 0x6d, 0x6c, 0x44, 0x61, 0x74, 0x61, 0x32, 0xb2, 0x02, 0x0a, 0x22, 0x46,
	0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x45, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x72, 0x69,
	0x73, 0x65, 0x47, 0x6c, 0x6f, 0x6f, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x41, 0x70,
	0x69, 0x12, 0x81, 0x01, 0x0a, 0x18, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x65, 0x64, 0x65, 0x72, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x75, 0x74, 0x68, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x12, 0x30,
	0x2e, 0x66, 0x65, 0x64, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x41, 0x75,
	0x74, 0x68, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x31, 0x2e, 0x66, 0x65, 0x64, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e,
	0x69, 0x6f, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x75, 0x74, 0x68, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x87, 0x01, 0x0a, 0x1a, 0x47, 0x65, 0x74, 0x46, 0x65, 0x64,
	0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x41, 0x75, 0x74, 0x68, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x59, 0x61, 0x6d, 0x6c, 0x12, 0x32, 0x2e, 0x66, 0x65, 0x64, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x73,
	0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x65, 0x64, 0x65, 0x72, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x75, 0x74, 0x68, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x59, 0x61, 0x6d,
	0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x33, 0x2e, 0x66, 0x65, 0x64, 0x2e, 0x72,
	0x70, 0x63, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x65,
	0x64, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x41, 0x75, 0x74, 0x68, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x59, 0x61, 0x6d, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42,
	0x39, 0x5a, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f,
	0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x61, 0x70, 0x69, 0x73, 0x2f,
	0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x65, 0x64, 0x2e, 0x72, 0x70, 0x63, 0x2f,
	0x76, 0x31, 0xb8, 0xf5, 0x04, 0x01, 0xc0, 0xf5, 0x04, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_github_com_solo_io_solo_apis_api_gloo_fed_fed_rpc_v1_federated_enterprise_gloo_resources_proto_rawDescOnce sync.Once
	file_github_com_solo_io_solo_apis_api_gloo_fed_fed_rpc_v1_federated_enterprise_gloo_resources_proto_rawDescData = file_github_com_solo_io_solo_apis_api_gloo_fed_fed_rpc_v1_federated_enterprise_gloo_resources_proto_rawDesc
)

func file_github_com_solo_io_solo_apis_api_gloo_fed_fed_rpc_v1_federated_enterprise_gloo_resources_proto_rawDescGZIP() []byte {
	file_github_com_solo_io_solo_apis_api_gloo_fed_fed_rpc_v1_federated_enterprise_gloo_resources_proto_rawDescOnce.Do(func() {
		file_github_com_solo_io_solo_apis_api_gloo_fed_fed_rpc_v1_federated_enterprise_gloo_resources_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_solo_io_solo_apis_api_gloo_fed_fed_rpc_v1_federated_enterprise_gloo_resources_proto_rawDescData)
	})
	return file_github_com_solo_io_solo_apis_api_gloo_fed_fed_rpc_v1_federated_enterprise_gloo_resources_proto_rawDescData
}

var file_github_com_solo_io_solo_apis_api_gloo_fed_fed_rpc_v1_federated_enterprise_gloo_resources_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_github_com_solo_io_solo_apis_api_gloo_fed_fed_rpc_v1_federated_enterprise_gloo_resources_proto_goTypes = []interface{}{
	(*FederatedAuthConfig)(nil),                // 0: fed.rpc.solo.io.FederatedAuthConfig
	(*ListFederatedAuthConfigsRequest)(nil),    // 1: fed.rpc.solo.io.ListFederatedAuthConfigsRequest
	(*ListFederatedAuthConfigsResponse)(nil),   // 2: fed.rpc.solo.io.ListFederatedAuthConfigsResponse
	(*GetFederatedAuthConfigYamlRequest)(nil),  // 3: fed.rpc.solo.io.GetFederatedAuthConfigYamlRequest
	(*GetFederatedAuthConfigYamlResponse)(nil), // 4: fed.rpc.solo.io.GetFederatedAuthConfigYamlResponse
	(*ObjectMeta)(nil),                         // 5: fed.rpc.solo.io.ObjectMeta
	(*types.FederatedAuthConfigSpec)(nil),      // 6: fed.enterprise.gloo.solo.io.FederatedAuthConfigSpec
	(*types.FederatedAuthConfigStatus)(nil),    // 7: fed.enterprise.gloo.solo.io.FederatedAuthConfigStatus
	(*v1.ObjectRef)(nil),                       // 8: core.skv2.solo.io.ObjectRef
	(*ResourceYaml)(nil),                       // 9: fed.rpc.solo.io.ResourceYaml
}
var file_github_com_solo_io_solo_apis_api_gloo_fed_fed_rpc_v1_federated_enterprise_gloo_resources_proto_depIdxs = []int32{
	5, // 0: fed.rpc.solo.io.FederatedAuthConfig.metadata:type_name -> fed.rpc.solo.io.ObjectMeta
	6, // 1: fed.rpc.solo.io.FederatedAuthConfig.spec:type_name -> fed.enterprise.gloo.solo.io.FederatedAuthConfigSpec
	7, // 2: fed.rpc.solo.io.FederatedAuthConfig.status:type_name -> fed.enterprise.gloo.solo.io.FederatedAuthConfigStatus
	0, // 3: fed.rpc.solo.io.ListFederatedAuthConfigsResponse.federated_auth_configs:type_name -> fed.rpc.solo.io.FederatedAuthConfig
	8, // 4: fed.rpc.solo.io.GetFederatedAuthConfigYamlRequest.federated_auth_config_ref:type_name -> core.skv2.solo.io.ObjectRef
	9, // 5: fed.rpc.solo.io.GetFederatedAuthConfigYamlResponse.yaml_data:type_name -> fed.rpc.solo.io.ResourceYaml
	1, // 6: fed.rpc.solo.io.FederatedEnterpriseGlooResourceApi.ListFederatedAuthConfigs:input_type -> fed.rpc.solo.io.ListFederatedAuthConfigsRequest
	3, // 7: fed.rpc.solo.io.FederatedEnterpriseGlooResourceApi.GetFederatedAuthConfigYaml:input_type -> fed.rpc.solo.io.GetFederatedAuthConfigYamlRequest
	2, // 8: fed.rpc.solo.io.FederatedEnterpriseGlooResourceApi.ListFederatedAuthConfigs:output_type -> fed.rpc.solo.io.ListFederatedAuthConfigsResponse
	4, // 9: fed.rpc.solo.io.FederatedEnterpriseGlooResourceApi.GetFederatedAuthConfigYaml:output_type -> fed.rpc.solo.io.GetFederatedAuthConfigYamlResponse
	8, // [8:10] is the sub-list for method output_type
	6, // [6:8] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() {
	file_github_com_solo_io_solo_apis_api_gloo_fed_fed_rpc_v1_federated_enterprise_gloo_resources_proto_init()
}
func file_github_com_solo_io_solo_apis_api_gloo_fed_fed_rpc_v1_federated_enterprise_gloo_resources_proto_init() {
	if File_github_com_solo_io_solo_apis_api_gloo_fed_fed_rpc_v1_federated_enterprise_gloo_resources_proto != nil {
		return
	}
	file_github_com_solo_io_solo_apis_api_gloo_fed_fed_rpc_v1_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_github_com_solo_io_solo_apis_api_gloo_fed_fed_rpc_v1_federated_enterprise_gloo_resources_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FederatedAuthConfig); i {
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
		file_github_com_solo_io_solo_apis_api_gloo_fed_fed_rpc_v1_federated_enterprise_gloo_resources_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListFederatedAuthConfigsRequest); i {
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
		file_github_com_solo_io_solo_apis_api_gloo_fed_fed_rpc_v1_federated_enterprise_gloo_resources_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListFederatedAuthConfigsResponse); i {
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
		file_github_com_solo_io_solo_apis_api_gloo_fed_fed_rpc_v1_federated_enterprise_gloo_resources_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFederatedAuthConfigYamlRequest); i {
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
		file_github_com_solo_io_solo_apis_api_gloo_fed_fed_rpc_v1_federated_enterprise_gloo_resources_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFederatedAuthConfigYamlResponse); i {
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
			RawDescriptor: file_github_com_solo_io_solo_apis_api_gloo_fed_fed_rpc_v1_federated_enterprise_gloo_resources_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_github_com_solo_io_solo_apis_api_gloo_fed_fed_rpc_v1_federated_enterprise_gloo_resources_proto_goTypes,
		DependencyIndexes: file_github_com_solo_io_solo_apis_api_gloo_fed_fed_rpc_v1_federated_enterprise_gloo_resources_proto_depIdxs,
		MessageInfos:      file_github_com_solo_io_solo_apis_api_gloo_fed_fed_rpc_v1_federated_enterprise_gloo_resources_proto_msgTypes,
	}.Build()
	File_github_com_solo_io_solo_apis_api_gloo_fed_fed_rpc_v1_federated_enterprise_gloo_resources_proto = out.File
	file_github_com_solo_io_solo_apis_api_gloo_fed_fed_rpc_v1_federated_enterprise_gloo_resources_proto_rawDesc = nil
	file_github_com_solo_io_solo_apis_api_gloo_fed_fed_rpc_v1_federated_enterprise_gloo_resources_proto_goTypes = nil
	file_github_com_solo_io_solo_apis_api_gloo_fed_fed_rpc_v1_federated_enterprise_gloo_resources_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// FederatedEnterpriseGlooResourceApiClient is the client API for FederatedEnterpriseGlooResourceApi service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type FederatedEnterpriseGlooResourceApiClient interface {
	ListFederatedAuthConfigs(ctx context.Context, in *ListFederatedAuthConfigsRequest, opts ...grpc.CallOption) (*ListFederatedAuthConfigsResponse, error)
	GetFederatedAuthConfigYaml(ctx context.Context, in *GetFederatedAuthConfigYamlRequest, opts ...grpc.CallOption) (*GetFederatedAuthConfigYamlResponse, error)
}

type federatedEnterpriseGlooResourceApiClient struct {
	cc grpc.ClientConnInterface
}

func NewFederatedEnterpriseGlooResourceApiClient(cc grpc.ClientConnInterface) FederatedEnterpriseGlooResourceApiClient {
	return &federatedEnterpriseGlooResourceApiClient{cc}
}

func (c *federatedEnterpriseGlooResourceApiClient) ListFederatedAuthConfigs(ctx context.Context, in *ListFederatedAuthConfigsRequest, opts ...grpc.CallOption) (*ListFederatedAuthConfigsResponse, error) {
	out := new(ListFederatedAuthConfigsResponse)
	err := c.cc.Invoke(ctx, "/fed.rpc.solo.io.FederatedEnterpriseGlooResourceApi/ListFederatedAuthConfigs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *federatedEnterpriseGlooResourceApiClient) GetFederatedAuthConfigYaml(ctx context.Context, in *GetFederatedAuthConfigYamlRequest, opts ...grpc.CallOption) (*GetFederatedAuthConfigYamlResponse, error) {
	out := new(GetFederatedAuthConfigYamlResponse)
	err := c.cc.Invoke(ctx, "/fed.rpc.solo.io.FederatedEnterpriseGlooResourceApi/GetFederatedAuthConfigYaml", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FederatedEnterpriseGlooResourceApiServer is the server API for FederatedEnterpriseGlooResourceApi service.
type FederatedEnterpriseGlooResourceApiServer interface {
	ListFederatedAuthConfigs(context.Context, *ListFederatedAuthConfigsRequest) (*ListFederatedAuthConfigsResponse, error)
	GetFederatedAuthConfigYaml(context.Context, *GetFederatedAuthConfigYamlRequest) (*GetFederatedAuthConfigYamlResponse, error)
}

// UnimplementedFederatedEnterpriseGlooResourceApiServer can be embedded to have forward compatible implementations.
type UnimplementedFederatedEnterpriseGlooResourceApiServer struct {
}

func (*UnimplementedFederatedEnterpriseGlooResourceApiServer) ListFederatedAuthConfigs(context.Context, *ListFederatedAuthConfigsRequest) (*ListFederatedAuthConfigsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListFederatedAuthConfigs not implemented")
}
func (*UnimplementedFederatedEnterpriseGlooResourceApiServer) GetFederatedAuthConfigYaml(context.Context, *GetFederatedAuthConfigYamlRequest) (*GetFederatedAuthConfigYamlResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFederatedAuthConfigYaml not implemented")
}

func RegisterFederatedEnterpriseGlooResourceApiServer(s *grpc.Server, srv FederatedEnterpriseGlooResourceApiServer) {
	s.RegisterService(&_FederatedEnterpriseGlooResourceApi_serviceDesc, srv)
}

func _FederatedEnterpriseGlooResourceApi_ListFederatedAuthConfigs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListFederatedAuthConfigsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FederatedEnterpriseGlooResourceApiServer).ListFederatedAuthConfigs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fed.rpc.solo.io.FederatedEnterpriseGlooResourceApi/ListFederatedAuthConfigs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FederatedEnterpriseGlooResourceApiServer).ListFederatedAuthConfigs(ctx, req.(*ListFederatedAuthConfigsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FederatedEnterpriseGlooResourceApi_GetFederatedAuthConfigYaml_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFederatedAuthConfigYamlRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FederatedEnterpriseGlooResourceApiServer).GetFederatedAuthConfigYaml(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fed.rpc.solo.io.FederatedEnterpriseGlooResourceApi/GetFederatedAuthConfigYaml",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FederatedEnterpriseGlooResourceApiServer).GetFederatedAuthConfigYaml(ctx, req.(*GetFederatedAuthConfigYamlRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _FederatedEnterpriseGlooResourceApi_serviceDesc = grpc.ServiceDesc{
	ServiceName: "fed.rpc.solo.io.FederatedEnterpriseGlooResourceApi",
	HandlerType: (*FederatedEnterpriseGlooResourceApiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListFederatedAuthConfigs",
			Handler:    _FederatedEnterpriseGlooResourceApi_ListFederatedAuthConfigs_Handler,
		},
		{
			MethodName: "GetFederatedAuthConfigYaml",
			Handler:    _FederatedEnterpriseGlooResourceApi_GetFederatedAuthConfigYaml_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/solo-io/solo-apis/api/gloo-fed/fed.rpc/v1/federated_enterprise_gloo_resources.proto",
}
