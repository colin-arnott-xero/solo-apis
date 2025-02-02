// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.6.1
// source: github.com/solo-io/solo-apis/api/gloo-fed/fed.gateway/v1/route_table.proto

package types

import (
	reflect "reflect"
	sync "sync"

	_ "github.com/golang/protobuf/ptypes/wrappers"
	_ "github.com/solo-io/protoc-gen-ext/extproto"
	v1 "github.com/solo-io/solo-apis/pkg/api/fed.solo.io/core/v1"
	v11 "github.com/solo-io/solo-apis/pkg/api/gateway.solo.io/v1"
	v1alpha1 "github.com/solo-io/solo-apis/pkg/api/multicluster.solo.io/v1alpha1"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// The Federated[Resource] CRDs are a gloo-fed wrapper around Gloo Edge CRDs, with a Placement field indicating which
// clusters and namespaces to federate the object to.
// For more, see: https://docs.solo.io/gloo-edge/latest/guides/gloo_federation/federated_configuration/
type FederatedRouteTableSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Template  *FederatedRouteTableSpec_Template `protobuf:"bytes,1,opt,name=template,proto3" json:"template,omitempty"`
	Placement *v1alpha1.Placement               `protobuf:"bytes,2,opt,name=placement,proto3" json:"placement,omitempty"`
}

func (x *FederatedRouteTableSpec) Reset() {
	*x = FederatedRouteTableSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_solo_apis_api_gloo_fed_fed_gateway_v1_route_table_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FederatedRouteTableSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FederatedRouteTableSpec) ProtoMessage() {}

func (x *FederatedRouteTableSpec) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_solo_apis_api_gloo_fed_fed_gateway_v1_route_table_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FederatedRouteTableSpec.ProtoReflect.Descriptor instead.
func (*FederatedRouteTableSpec) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_solo_apis_api_gloo_fed_fed_gateway_v1_route_table_proto_rawDescGZIP(), []int{0}
}

func (x *FederatedRouteTableSpec) GetTemplate() *FederatedRouteTableSpec_Template {
	if x != nil {
		return x.Template
	}
	return nil
}

func (x *FederatedRouteTableSpec) GetPlacement() *v1alpha1.Placement {
	if x != nil {
		return x.Placement
	}
	return nil
}

type FederatedRouteTableStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlacementStatus *v1.PlacementStatus `protobuf:"bytes,1,opt,name=placement_status,json=placementStatus,proto3" json:"placement_status,omitempty"`
}

func (x *FederatedRouteTableStatus) Reset() {
	*x = FederatedRouteTableStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_solo_apis_api_gloo_fed_fed_gateway_v1_route_table_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FederatedRouteTableStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FederatedRouteTableStatus) ProtoMessage() {}

func (x *FederatedRouteTableStatus) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_solo_apis_api_gloo_fed_fed_gateway_v1_route_table_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FederatedRouteTableStatus.ProtoReflect.Descriptor instead.
func (*FederatedRouteTableStatus) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_solo_apis_api_gloo_fed_fed_gateway_v1_route_table_proto_rawDescGZIP(), []int{1}
}

func (x *FederatedRouteTableStatus) GetPlacementStatus() *v1.PlacementStatus {
	if x != nil {
		return x.PlacementStatus
	}
	return nil
}

type FederatedRouteTableSpec_Template struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Spec     *v11.RouteTableSpec  `protobuf:"bytes,1,opt,name=spec,proto3" json:"spec,omitempty"`
	Metadata *v1.TemplateMetadata `protobuf:"bytes,2,opt,name=metadata,proto3" json:"metadata,omitempty"`
}

func (x *FederatedRouteTableSpec_Template) Reset() {
	*x = FederatedRouteTableSpec_Template{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_solo_apis_api_gloo_fed_fed_gateway_v1_route_table_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FederatedRouteTableSpec_Template) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FederatedRouteTableSpec_Template) ProtoMessage() {}

func (x *FederatedRouteTableSpec_Template) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_solo_apis_api_gloo_fed_fed_gateway_v1_route_table_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FederatedRouteTableSpec_Template.ProtoReflect.Descriptor instead.
func (*FederatedRouteTableSpec_Template) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_solo_apis_api_gloo_fed_fed_gateway_v1_route_table_proto_rawDescGZIP(), []int{0, 0}
}

func (x *FederatedRouteTableSpec_Template) GetSpec() *v11.RouteTableSpec {
	if x != nil {
		return x.Spec
	}
	return nil
}

func (x *FederatedRouteTableSpec_Template) GetMetadata() *v1.TemplateMetadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

var File_github_com_solo_io_solo_apis_api_gloo_fed_fed_gateway_v1_route_table_proto protoreflect.FileDescriptor

var file_github_com_solo_io_solo_apis_api_gloo_fed_fed_gateway_v1_route_table_proto_rawDesc = []byte{
	0x0a, 0x4a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c,
	0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2d, 0x66, 0x65, 0x64, 0x2f, 0x66, 0x65, 0x64, 0x2e,
	0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x6f, 0x75, 0x74, 0x65,
	0x5f, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x66, 0x65,
	0x64, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69,
	0x6f, 0x1a, 0x12, 0x65, 0x78, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x78, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x4c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x61,
	0x70, 0x69, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x6b, 0x76, 0x32, 0x2d, 0x65, 0x6e, 0x74,
	0x65, 0x72, 0x70, 0x72, 0x69, 0x73, 0x65, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x2f, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x45, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x61, 0x70, 0x69,
	0x73, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2d, 0x66, 0x65, 0x64, 0x2f, 0x66,
	0x65, 0x64, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x6c, 0x61, 0x63, 0x65,
	0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x42, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x73,
	0x6f, 0x6c, 0x6f, 0x2d, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x6c, 0x6f,
	0x6f, 0x2f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x6f, 0x75,
	0x74, 0x65, 0x5f, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xac,
	0x02, 0x0a, 0x17, 0x46, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x52, 0x6f, 0x75, 0x74,
	0x65, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x53, 0x70, 0x65, 0x63, 0x12, 0x51, 0x0a, 0x08, 0x74, 0x65,
	0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x35, 0x2e, 0x66,
	0x65, 0x64, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e,
	0x69, 0x6f, 0x2e, 0x46, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x52, 0x6f, 0x75, 0x74,
	0x65, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x53, 0x70, 0x65, 0x63, 0x2e, 0x54, 0x65, 0x6d, 0x70, 0x6c,
	0x61, 0x74, 0x65, 0x52, 0x08, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x12, 0x3d, 0x0a,
	0x09, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1f, 0x2e, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e,
	0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x52, 0x09, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x1a, 0x7f, 0x0a, 0x08,
	0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x12, 0x33, 0x0a, 0x04, 0x73, 0x70, 0x65, 0x63,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79,
	0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x54, 0x61,
	0x62, 0x6c, 0x65, 0x53, 0x70, 0x65, 0x63, 0x52, 0x04, 0x73, 0x70, 0x65, 0x63, 0x12, 0x3e, 0x0a,
	0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x22, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x66, 0x65, 0x64, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e,
	0x69, 0x6f, 0x2e, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x22, 0x69, 0x0a,
	0x19, 0x46, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x54,
	0x61, 0x62, 0x6c, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x4c, 0x0a, 0x10, 0x70, 0x6c,
	0x61, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x66, 0x65, 0x64, 0x2e,
	0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x0f, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x6d, 0x65,
	0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x42, 0x4b, 0x5a, 0x41, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x73,
	0x6f, 0x6c, 0x6f, 0x2d, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x66, 0x65, 0x64, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x73, 0x6f, 0x6c,
	0x6f, 0x2e, 0x69, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0xc0, 0xf5, 0x04,
	0x01, 0xb8, 0xf5, 0x04, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_github_com_solo_io_solo_apis_api_gloo_fed_fed_gateway_v1_route_table_proto_rawDescOnce sync.Once
	file_github_com_solo_io_solo_apis_api_gloo_fed_fed_gateway_v1_route_table_proto_rawDescData = file_github_com_solo_io_solo_apis_api_gloo_fed_fed_gateway_v1_route_table_proto_rawDesc
)

func file_github_com_solo_io_solo_apis_api_gloo_fed_fed_gateway_v1_route_table_proto_rawDescGZIP() []byte {
	file_github_com_solo_io_solo_apis_api_gloo_fed_fed_gateway_v1_route_table_proto_rawDescOnce.Do(func() {
		file_github_com_solo_io_solo_apis_api_gloo_fed_fed_gateway_v1_route_table_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_solo_io_solo_apis_api_gloo_fed_fed_gateway_v1_route_table_proto_rawDescData)
	})
	return file_github_com_solo_io_solo_apis_api_gloo_fed_fed_gateway_v1_route_table_proto_rawDescData
}

var file_github_com_solo_io_solo_apis_api_gloo_fed_fed_gateway_v1_route_table_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_github_com_solo_io_solo_apis_api_gloo_fed_fed_gateway_v1_route_table_proto_goTypes = []interface{}{
	(*FederatedRouteTableSpec)(nil),          // 0: fed.gateway.solo.io.FederatedRouteTableSpec
	(*FederatedRouteTableStatus)(nil),        // 1: fed.gateway.solo.io.FederatedRouteTableStatus
	(*FederatedRouteTableSpec_Template)(nil), // 2: fed.gateway.solo.io.FederatedRouteTableSpec.Template
	(*v1alpha1.Placement)(nil),               // 3: multicluster.solo.io.Placement
	(*v1.PlacementStatus)(nil),               // 4: core.fed.solo.io.PlacementStatus
	(*v11.RouteTableSpec)(nil),               // 5: gateway.solo.io.RouteTableSpec
	(*v1.TemplateMetadata)(nil),              // 6: core.fed.solo.io.TemplateMetadata
}
var file_github_com_solo_io_solo_apis_api_gloo_fed_fed_gateway_v1_route_table_proto_depIdxs = []int32{
	2, // 0: fed.gateway.solo.io.FederatedRouteTableSpec.template:type_name -> fed.gateway.solo.io.FederatedRouteTableSpec.Template
	3, // 1: fed.gateway.solo.io.FederatedRouteTableSpec.placement:type_name -> multicluster.solo.io.Placement
	4, // 2: fed.gateway.solo.io.FederatedRouteTableStatus.placement_status:type_name -> core.fed.solo.io.PlacementStatus
	5, // 3: fed.gateway.solo.io.FederatedRouteTableSpec.Template.spec:type_name -> gateway.solo.io.RouteTableSpec
	6, // 4: fed.gateway.solo.io.FederatedRouteTableSpec.Template.metadata:type_name -> core.fed.solo.io.TemplateMetadata
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_github_com_solo_io_solo_apis_api_gloo_fed_fed_gateway_v1_route_table_proto_init() }
func file_github_com_solo_io_solo_apis_api_gloo_fed_fed_gateway_v1_route_table_proto_init() {
	if File_github_com_solo_io_solo_apis_api_gloo_fed_fed_gateway_v1_route_table_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_github_com_solo_io_solo_apis_api_gloo_fed_fed_gateway_v1_route_table_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FederatedRouteTableSpec); i {
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
		file_github_com_solo_io_solo_apis_api_gloo_fed_fed_gateway_v1_route_table_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FederatedRouteTableStatus); i {
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
		file_github_com_solo_io_solo_apis_api_gloo_fed_fed_gateway_v1_route_table_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FederatedRouteTableSpec_Template); i {
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
			RawDescriptor: file_github_com_solo_io_solo_apis_api_gloo_fed_fed_gateway_v1_route_table_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_github_com_solo_io_solo_apis_api_gloo_fed_fed_gateway_v1_route_table_proto_goTypes,
		DependencyIndexes: file_github_com_solo_io_solo_apis_api_gloo_fed_fed_gateway_v1_route_table_proto_depIdxs,
		MessageInfos:      file_github_com_solo_io_solo_apis_api_gloo_fed_fed_gateway_v1_route_table_proto_msgTypes,
	}.Build()
	File_github_com_solo_io_solo_apis_api_gloo_fed_fed_gateway_v1_route_table_proto = out.File
	file_github_com_solo_io_solo_apis_api_gloo_fed_fed_gateway_v1_route_table_proto_rawDesc = nil
	file_github_com_solo_io_solo_apis_api_gloo_fed_fed_gateway_v1_route_table_proto_goTypes = nil
	file_github_com_solo_io_solo_apis_api_gloo_fed_fed_gateway_v1_route_table_proto_depIdxs = nil
}
