// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.6.1
// source: github.com/solo-io/solo-apis/api/skv2-enterprise/v1alpha1/multicluster.proto

package v1alpha1

import (
	reflect "reflect"
	sync "sync"

	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	_ "github.com/solo-io/protoc-gen-ext/extproto"
	v1 "github.com/solo-io/skv2/pkg/api/core.skv2.solo.io/v1"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type MultiClusterRoleSpec_Rule_Action int32

const (
	MultiClusterRoleSpec_Rule_ANY    MultiClusterRoleSpec_Rule_Action = 0
	MultiClusterRoleSpec_Rule_CREATE MultiClusterRoleSpec_Rule_Action = 1
	MultiClusterRoleSpec_Rule_UPDATE MultiClusterRoleSpec_Rule_Action = 2
	MultiClusterRoleSpec_Rule_DELETE MultiClusterRoleSpec_Rule_Action = 3
)

// Enum value maps for MultiClusterRoleSpec_Rule_Action.
var (
	MultiClusterRoleSpec_Rule_Action_name = map[int32]string{
		0: "ANY",
		1: "CREATE",
		2: "UPDATE",
		3: "DELETE",
	}
	MultiClusterRoleSpec_Rule_Action_value = map[string]int32{
		"ANY":    0,
		"CREATE": 1,
		"UPDATE": 2,
		"DELETE": 3,
	}
)

func (x MultiClusterRoleSpec_Rule_Action) Enum() *MultiClusterRoleSpec_Rule_Action {
	p := new(MultiClusterRoleSpec_Rule_Action)
	*p = x
	return p
}

func (x MultiClusterRoleSpec_Rule_Action) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MultiClusterRoleSpec_Rule_Action) Descriptor() protoreflect.EnumDescriptor {
	return file_github_com_solo_io_solo_apis_api_skv2_enterprise_v1alpha1_multicluster_proto_enumTypes[0].Descriptor()
}

func (MultiClusterRoleSpec_Rule_Action) Type() protoreflect.EnumType {
	return &file_github_com_solo_io_solo_apis_api_skv2_enterprise_v1alpha1_multicluster_proto_enumTypes[0]
}

func (x MultiClusterRoleSpec_Rule_Action) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MultiClusterRoleSpec_Rule_Action.Descriptor instead.
func (MultiClusterRoleSpec_Rule_Action) EnumDescriptor() ([]byte, []int) {
	return file_github_com_solo_io_solo_apis_api_skv2_enterprise_v1alpha1_multicluster_proto_rawDescGZIP(), []int{1, 0, 0}
}

//
//Object representing the clusters and namespaces on which resources are created/updated/deleted,
//computed as the cartesian product of all declared namespace and cluster values.
type Placement struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//
	//List of namespaces within each placement cluster in which to create/update/delete resources.
	//Wildcard ("*") represents any namespace.
	Namespaces []string `protobuf:"bytes,1,rep,name=namespaces,proto3" json:"namespaces,omitempty"`
	//
	//List of clusters (represented by a string) in which to create/update/delete resources.
	//Wildcard ("*") represents any cluster.
	Clusters []string `protobuf:"bytes,2,rep,name=clusters,proto3" json:"clusters,omitempty"`
}

func (x *Placement) Reset() {
	*x = Placement{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_solo_apis_api_skv2_enterprise_v1alpha1_multicluster_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Placement) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Placement) ProtoMessage() {}

func (x *Placement) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_solo_apis_api_skv2_enterprise_v1alpha1_multicluster_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Placement.ProtoReflect.Descriptor instead.
func (*Placement) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_solo_apis_api_skv2_enterprise_v1alpha1_multicluster_proto_rawDescGZIP(), []int{0}
}

func (x *Placement) GetNamespaces() []string {
	if x != nil {
		return x.Namespaces
	}
	return nil
}

func (x *Placement) GetClusters() []string {
	if x != nil {
		return x.Clusters
	}
	return nil
}

type MultiClusterRoleSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Rules []*MultiClusterRoleSpec_Rule `protobuf:"bytes,1,rep,name=rules,proto3" json:"rules,omitempty"`
}

func (x *MultiClusterRoleSpec) Reset() {
	*x = MultiClusterRoleSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_solo_apis_api_skv2_enterprise_v1alpha1_multicluster_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MultiClusterRoleSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MultiClusterRoleSpec) ProtoMessage() {}

func (x *MultiClusterRoleSpec) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_solo_apis_api_skv2_enterprise_v1alpha1_multicluster_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MultiClusterRoleSpec.ProtoReflect.Descriptor instead.
func (*MultiClusterRoleSpec) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_solo_apis_api_skv2_enterprise_v1alpha1_multicluster_proto_rawDescGZIP(), []int{1}
}

func (x *MultiClusterRoleSpec) GetRules() []*MultiClusterRoleSpec_Rule {
	if x != nil {
		return x.Rules
	}
	return nil
}

type MultiClusterRoleStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *MultiClusterRoleStatus) Reset() {
	*x = MultiClusterRoleStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_solo_apis_api_skv2_enterprise_v1alpha1_multicluster_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MultiClusterRoleStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MultiClusterRoleStatus) ProtoMessage() {}

func (x *MultiClusterRoleStatus) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_solo_apis_api_skv2_enterprise_v1alpha1_multicluster_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MultiClusterRoleStatus.ProtoReflect.Descriptor instead.
func (*MultiClusterRoleStatus) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_solo_apis_api_skv2_enterprise_v1alpha1_multicluster_proto_rawDescGZIP(), []int{2}
}

type MultiClusterRoleBindingSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// reference to users or groups to apply the MultiClusterRole to
	Subjects []*v1.TypedObjectRef `protobuf:"bytes,1,rep,name=subjects,proto3" json:"subjects,omitempty"`
	// reference to a MultiClusterRole
	RoleRef *v1.ObjectRef `protobuf:"bytes,2,opt,name=role_ref,json=roleRef,proto3" json:"role_ref,omitempty"`
}

func (x *MultiClusterRoleBindingSpec) Reset() {
	*x = MultiClusterRoleBindingSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_solo_apis_api_skv2_enterprise_v1alpha1_multicluster_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MultiClusterRoleBindingSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MultiClusterRoleBindingSpec) ProtoMessage() {}

func (x *MultiClusterRoleBindingSpec) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_solo_apis_api_skv2_enterprise_v1alpha1_multicluster_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MultiClusterRoleBindingSpec.ProtoReflect.Descriptor instead.
func (*MultiClusterRoleBindingSpec) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_solo_apis_api_skv2_enterprise_v1alpha1_multicluster_proto_rawDescGZIP(), []int{3}
}

func (x *MultiClusterRoleBindingSpec) GetSubjects() []*v1.TypedObjectRef {
	if x != nil {
		return x.Subjects
	}
	return nil
}

func (x *MultiClusterRoleBindingSpec) GetRoleRef() *v1.ObjectRef {
	if x != nil {
		return x.RoleRef
	}
	return nil
}

type MultiClusterRoleBindingStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *MultiClusterRoleBindingStatus) Reset() {
	*x = MultiClusterRoleBindingStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_solo_apis_api_skv2_enterprise_v1alpha1_multicluster_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MultiClusterRoleBindingStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MultiClusterRoleBindingStatus) ProtoMessage() {}

func (x *MultiClusterRoleBindingStatus) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_solo_apis_api_skv2_enterprise_v1alpha1_multicluster_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MultiClusterRoleBindingStatus.ProtoReflect.Descriptor instead.
func (*MultiClusterRoleBindingStatus) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_solo_apis_api_skv2_enterprise_v1alpha1_multicluster_proto_rawDescGZIP(), []int{4}
}

type MultiClusterRoleSpec_Rule struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ApiGroup string `protobuf:"bytes,1,opt,name=api_group,json=apiGroup,proto3" json:"api_group,omitempty"`
	// The kind of the object to apply to, if left empty will apply to all kinds in a group.
	Kind   *wrappers.StringValue            `protobuf:"bytes,2,opt,name=kind,proto3" json:"kind,omitempty"`
	Action MultiClusterRoleSpec_Rule_Action `protobuf:"varint,3,opt,name=action,proto3,enum=multicluster.solo.io.MultiClusterRoleSpec_Rule_Action" json:"action,omitempty"`
	// The clusters and namespaces this rule applies to.
	Placements []*Placement `protobuf:"bytes,4,rep,name=placements,proto3" json:"placements,omitempty"`
}

func (x *MultiClusterRoleSpec_Rule) Reset() {
	*x = MultiClusterRoleSpec_Rule{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_solo_apis_api_skv2_enterprise_v1alpha1_multicluster_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MultiClusterRoleSpec_Rule) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MultiClusterRoleSpec_Rule) ProtoMessage() {}

func (x *MultiClusterRoleSpec_Rule) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_solo_apis_api_skv2_enterprise_v1alpha1_multicluster_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MultiClusterRoleSpec_Rule.ProtoReflect.Descriptor instead.
func (*MultiClusterRoleSpec_Rule) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_solo_apis_api_skv2_enterprise_v1alpha1_multicluster_proto_rawDescGZIP(), []int{1, 0}
}

func (x *MultiClusterRoleSpec_Rule) GetApiGroup() string {
	if x != nil {
		return x.ApiGroup
	}
	return ""
}

func (x *MultiClusterRoleSpec_Rule) GetKind() *wrappers.StringValue {
	if x != nil {
		return x.Kind
	}
	return nil
}

func (x *MultiClusterRoleSpec_Rule) GetAction() MultiClusterRoleSpec_Rule_Action {
	if x != nil {
		return x.Action
	}
	return MultiClusterRoleSpec_Rule_ANY
}

func (x *MultiClusterRoleSpec_Rule) GetPlacements() []*Placement {
	if x != nil {
		return x.Placements
	}
	return nil
}

var File_github_com_solo_io_solo_apis_api_skv2_enterprise_v1alpha1_multicluster_proto protoreflect.FileDescriptor

var file_github_com_solo_io_solo_apis_api_skv2_enterprise_v1alpha1_multicluster_proto_rawDesc = []byte{
	0x0a, 0x4c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c,
	0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x73, 0x6b, 0x76, 0x32, 0x2d, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x72, 0x69,
	0x73, 0x65, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x6d, 0x75, 0x6c, 0x74,
	0x69, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x14,
	0x6d, 0x75, 0x6c, 0x74, 0x69, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x73, 0x6f, 0x6c,
	0x6f, 0x2e, 0x69, 0x6f, 0x1a, 0x12, 0x65, 0x78, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65,
	0x78, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65,
	0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x73, 0x6b, 0x76,
	0x32, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f,
	0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x47, 0x0a, 0x09, 0x50, 0x6c, 0x61, 0x63,
	0x65, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0a, 0x6e, 0x61, 0x6d, 0x65, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x73, 0x22, 0xfd, 0x02, 0x0a, 0x14, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x43, 0x6c, 0x75, 0x73, 0x74,
	0x65, 0x72, 0x52, 0x6f, 0x6c, 0x65, 0x53, 0x70, 0x65, 0x63, 0x12, 0x45, 0x0a, 0x05, 0x72, 0x75,
	0x6c, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x6d, 0x75, 0x6c, 0x74,
	0x69, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f,
	0x2e, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x52, 0x6f, 0x6c,
	0x65, 0x53, 0x70, 0x65, 0x63, 0x2e, 0x52, 0x75, 0x6c, 0x65, 0x52, 0x05, 0x72, 0x75, 0x6c, 0x65,
	0x73, 0x1a, 0x9d, 0x02, 0x0a, 0x04, 0x52, 0x75, 0x6c, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x61, 0x70,
	0x69, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x61,
	0x70, 0x69, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x30, 0x0a, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x52, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x12, 0x4e, 0x0a, 0x06, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x36, 0x2e, 0x6d, 0x75, 0x6c, 0x74,
	0x69, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f,
	0x2e, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x52, 0x6f, 0x6c,
	0x65, 0x53, 0x70, 0x65, 0x63, 0x2e, 0x52, 0x75, 0x6c, 0x65, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x3f, 0x0a, 0x0a, 0x70, 0x6c, 0x61,
	0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e,
	0x6d, 0x75, 0x6c, 0x74, 0x69, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x73, 0x6f, 0x6c,
	0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x0a,
	0x70, 0x6c, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x22, 0x35, 0x0a, 0x06, 0x41, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x07, 0x0a, 0x03, 0x41, 0x4e, 0x59, 0x10, 0x00, 0x12, 0x0a, 0x0a,
	0x06, 0x43, 0x52, 0x45, 0x41, 0x54, 0x45, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x55, 0x50, 0x44,
	0x41, 0x54, 0x45, 0x10, 0x02, 0x12, 0x0a, 0x0a, 0x06, 0x44, 0x45, 0x4c, 0x45, 0x54, 0x45, 0x10,
	0x03, 0x22, 0x18, 0x0a, 0x16, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65,
	0x72, 0x52, 0x6f, 0x6c, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x95, 0x01, 0x0a, 0x1b,
	0x4d, 0x75, 0x6c, 0x74, 0x69, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x52, 0x6f, 0x6c, 0x65,
	0x42, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x53, 0x70, 0x65, 0x63, 0x12, 0x3d, 0x0a, 0x08, 0x73,
	0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e,
	0x63, 0x6f, 0x72, 0x65, 0x2e, 0x73, 0x6b, 0x76, 0x32, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69,
	0x6f, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x64, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x66,
	0x52, 0x08, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x12, 0x37, 0x0a, 0x08, 0x72, 0x6f,
	0x6c, 0x65, 0x5f, 0x72, 0x65, 0x66, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x63,
	0x6f, 0x72, 0x65, 0x2e, 0x73, 0x6b, 0x76, 0x32, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f,
	0x2e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x66, 0x52, 0x07, 0x72, 0x6f, 0x6c, 0x65,
	0x52, 0x65, 0x66, 0x22, 0x1f, 0x0a, 0x1d, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x43, 0x6c, 0x75, 0x73,
	0x74, 0x65, 0x72, 0x52, 0x6f, 0x6c, 0x65, 0x42, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x42, 0x4c, 0x5a, 0x42, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d,
	0x61, 0x70, 0x69, 0x73, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x75, 0x6c,
	0x74, 0x69, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69,
	0x6f, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xb8, 0xf5, 0x04, 0x01, 0xc0, 0xf5,
	0x04, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_github_com_solo_io_solo_apis_api_skv2_enterprise_v1alpha1_multicluster_proto_rawDescOnce sync.Once
	file_github_com_solo_io_solo_apis_api_skv2_enterprise_v1alpha1_multicluster_proto_rawDescData = file_github_com_solo_io_solo_apis_api_skv2_enterprise_v1alpha1_multicluster_proto_rawDesc
)

func file_github_com_solo_io_solo_apis_api_skv2_enterprise_v1alpha1_multicluster_proto_rawDescGZIP() []byte {
	file_github_com_solo_io_solo_apis_api_skv2_enterprise_v1alpha1_multicluster_proto_rawDescOnce.Do(func() {
		file_github_com_solo_io_solo_apis_api_skv2_enterprise_v1alpha1_multicluster_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_solo_io_solo_apis_api_skv2_enterprise_v1alpha1_multicluster_proto_rawDescData)
	})
	return file_github_com_solo_io_solo_apis_api_skv2_enterprise_v1alpha1_multicluster_proto_rawDescData
}

var file_github_com_solo_io_solo_apis_api_skv2_enterprise_v1alpha1_multicluster_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_github_com_solo_io_solo_apis_api_skv2_enterprise_v1alpha1_multicluster_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_github_com_solo_io_solo_apis_api_skv2_enterprise_v1alpha1_multicluster_proto_goTypes = []interface{}{
	(MultiClusterRoleSpec_Rule_Action)(0), // 0: multicluster.solo.io.MultiClusterRoleSpec.Rule.Action
	(*Placement)(nil),                     // 1: multicluster.solo.io.Placement
	(*MultiClusterRoleSpec)(nil),          // 2: multicluster.solo.io.MultiClusterRoleSpec
	(*MultiClusterRoleStatus)(nil),        // 3: multicluster.solo.io.MultiClusterRoleStatus
	(*MultiClusterRoleBindingSpec)(nil),   // 4: multicluster.solo.io.MultiClusterRoleBindingSpec
	(*MultiClusterRoleBindingStatus)(nil), // 5: multicluster.solo.io.MultiClusterRoleBindingStatus
	(*MultiClusterRoleSpec_Rule)(nil),     // 6: multicluster.solo.io.MultiClusterRoleSpec.Rule
	(*v1.TypedObjectRef)(nil),             // 7: core.skv2.solo.io.TypedObjectRef
	(*v1.ObjectRef)(nil),                  // 8: core.skv2.solo.io.ObjectRef
	(*wrappers.StringValue)(nil),          // 9: google.protobuf.StringValue
}
var file_github_com_solo_io_solo_apis_api_skv2_enterprise_v1alpha1_multicluster_proto_depIdxs = []int32{
	6, // 0: multicluster.solo.io.MultiClusterRoleSpec.rules:type_name -> multicluster.solo.io.MultiClusterRoleSpec.Rule
	7, // 1: multicluster.solo.io.MultiClusterRoleBindingSpec.subjects:type_name -> core.skv2.solo.io.TypedObjectRef
	8, // 2: multicluster.solo.io.MultiClusterRoleBindingSpec.role_ref:type_name -> core.skv2.solo.io.ObjectRef
	9, // 3: multicluster.solo.io.MultiClusterRoleSpec.Rule.kind:type_name -> google.protobuf.StringValue
	0, // 4: multicluster.solo.io.MultiClusterRoleSpec.Rule.action:type_name -> multicluster.solo.io.MultiClusterRoleSpec.Rule.Action
	1, // 5: multicluster.solo.io.MultiClusterRoleSpec.Rule.placements:type_name -> multicluster.solo.io.Placement
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_github_com_solo_io_solo_apis_api_skv2_enterprise_v1alpha1_multicluster_proto_init() }
func file_github_com_solo_io_solo_apis_api_skv2_enterprise_v1alpha1_multicluster_proto_init() {
	if File_github_com_solo_io_solo_apis_api_skv2_enterprise_v1alpha1_multicluster_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_github_com_solo_io_solo_apis_api_skv2_enterprise_v1alpha1_multicluster_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Placement); i {
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
		file_github_com_solo_io_solo_apis_api_skv2_enterprise_v1alpha1_multicluster_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MultiClusterRoleSpec); i {
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
		file_github_com_solo_io_solo_apis_api_skv2_enterprise_v1alpha1_multicluster_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MultiClusterRoleStatus); i {
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
		file_github_com_solo_io_solo_apis_api_skv2_enterprise_v1alpha1_multicluster_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MultiClusterRoleBindingSpec); i {
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
		file_github_com_solo_io_solo_apis_api_skv2_enterprise_v1alpha1_multicluster_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MultiClusterRoleBindingStatus); i {
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
		file_github_com_solo_io_solo_apis_api_skv2_enterprise_v1alpha1_multicluster_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MultiClusterRoleSpec_Rule); i {
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
			RawDescriptor: file_github_com_solo_io_solo_apis_api_skv2_enterprise_v1alpha1_multicluster_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_github_com_solo_io_solo_apis_api_skv2_enterprise_v1alpha1_multicluster_proto_goTypes,
		DependencyIndexes: file_github_com_solo_io_solo_apis_api_skv2_enterprise_v1alpha1_multicluster_proto_depIdxs,
		EnumInfos:         file_github_com_solo_io_solo_apis_api_skv2_enterprise_v1alpha1_multicluster_proto_enumTypes,
		MessageInfos:      file_github_com_solo_io_solo_apis_api_skv2_enterprise_v1alpha1_multicluster_proto_msgTypes,
	}.Build()
	File_github_com_solo_io_solo_apis_api_skv2_enterprise_v1alpha1_multicluster_proto = out.File
	file_github_com_solo_io_solo_apis_api_skv2_enterprise_v1alpha1_multicluster_proto_rawDesc = nil
	file_github_com_solo_io_solo_apis_api_skv2_enterprise_v1alpha1_multicluster_proto_goTypes = nil
	file_github_com_solo_io_solo_apis_api_skv2_enterprise_v1alpha1_multicluster_proto_depIdxs = nil
}
