// Code generated by skv2. DO NOT EDIT.

// This file contains generated Deepcopy methods for proto-based Spec and Status fields

package types

import (
    proto "github.com/golang/protobuf/proto"
)

// DeepCopyInto for the FederatedGateway.Spec
func (in *FederatedGatewaySpec) DeepCopyInto(out *FederatedGatewaySpec) {
    p := proto.Clone(in).(*FederatedGatewaySpec)
    *out = *p
}
// DeepCopyInto for the FederatedGateway.Status
func (in *FederatedGatewayStatus) DeepCopyInto(out *FederatedGatewayStatus) {
    p := proto.Clone(in).(*FederatedGatewayStatus)
    *out = *p
}

// DeepCopyInto for the FederatedRouteTable.Spec
func (in *FederatedRouteTableSpec) DeepCopyInto(out *FederatedRouteTableSpec) {
    p := proto.Clone(in).(*FederatedRouteTableSpec)
    *out = *p
}
// DeepCopyInto for the FederatedRouteTable.Status
func (in *FederatedRouteTableStatus) DeepCopyInto(out *FederatedRouteTableStatus) {
    p := proto.Clone(in).(*FederatedRouteTableStatus)
    *out = *p
}

// DeepCopyInto for the FederatedVirtualService.Spec
func (in *FederatedVirtualServiceSpec) DeepCopyInto(out *FederatedVirtualServiceSpec) {
    p := proto.Clone(in).(*FederatedVirtualServiceSpec)
    *out = *p
}
// DeepCopyInto for the FederatedVirtualService.Status
func (in *FederatedVirtualServiceStatus) DeepCopyInto(out *FederatedVirtualServiceStatus) {
    p := proto.Clone(in).(*FederatedVirtualServiceStatus)
    *out = *p
}
