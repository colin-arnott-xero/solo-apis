// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/solo-apis/api/gloo/gateway/v1/route_table.proto

package v1

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"strings"

	"github.com/solo-io/protoc-gen-ext/pkg/clone"
	"google.golang.org/protobuf/proto"

	github_com_golang_protobuf_ptypes_struct "github.com/golang/protobuf/ptypes/struct"

	github_com_golang_protobuf_ptypes_wrappers "github.com/golang/protobuf/ptypes/wrappers"
)

// ensure the imports are used
var (
	_ = errors.New("")
	_ = fmt.Print
	_ = binary.LittleEndian
	_ = bytes.Compare
	_ = strings.Compare
	_ = clone.Cloner(nil)
	_ = proto.Message(nil)
)

// Clone function
func (m *RouteTableSpec) Clone() proto.Message {
	var target *RouteTableSpec
	if m == nil {
		return target
	}
	target = &RouteTableSpec{}

	if m.GetRoutes() != nil {
		target.Routes = make([]*Route, len(m.GetRoutes()))
		for idx, v := range m.GetRoutes() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.Routes[idx] = h.Clone().(*Route)
			} else {
				target.Routes[idx] = proto.Clone(v).(*Route)
			}

		}
	}

	if h, ok := interface{}(m.GetWeight()).(clone.Cloner); ok {
		target.Weight = h.Clone().(*github_com_golang_protobuf_ptypes_wrappers.Int32Value)
	} else {
		target.Weight = proto.Clone(m.GetWeight()).(*github_com_golang_protobuf_ptypes_wrappers.Int32Value)
	}

	return target
}

// Clone function
func (m *RouteTableStatus) Clone() proto.Message {
	var target *RouteTableStatus
	if m == nil {
		return target
	}
	target = &RouteTableStatus{}

	target.State = m.GetState()

	target.Reason = m.GetReason()

	target.ReportedBy = m.GetReportedBy()

	if m.GetSubresourceStatuses() != nil {
		target.SubresourceStatuses = make(map[string]*RouteTableStatus, len(m.GetSubresourceStatuses()))
		for k, v := range m.GetSubresourceStatuses() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.SubresourceStatuses[k] = h.Clone().(*RouteTableStatus)
			} else {
				target.SubresourceStatuses[k] = proto.Clone(v).(*RouteTableStatus)
			}

		}
	}

	if h, ok := interface{}(m.GetDetails()).(clone.Cloner); ok {
		target.Details = h.Clone().(*github_com_golang_protobuf_ptypes_struct.Struct)
	} else {
		target.Details = proto.Clone(m.GetDetails()).(*github_com_golang_protobuf_ptypes_struct.Struct)
	}

	return target
}

// Clone function
func (m *RouteTableNamespacedStatuses) Clone() proto.Message {
	var target *RouteTableNamespacedStatuses
	if m == nil {
		return target
	}
	target = &RouteTableNamespacedStatuses{}

	if m.GetStatuses() != nil {
		target.Statuses = make(map[string]*RouteTableStatus, len(m.GetStatuses()))
		for k, v := range m.GetStatuses() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.Statuses[k] = h.Clone().(*RouteTableStatus)
			} else {
				target.Statuses[k] = proto.Clone(v).(*RouteTableStatus)
			}

		}
	}

	return target
}