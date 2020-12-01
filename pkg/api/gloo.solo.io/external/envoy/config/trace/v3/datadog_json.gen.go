// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/solo-apis/api/gloo/gloo/external/envoy/config/trace/v3/datadog.proto

package v3

import (
	bytes "bytes"
	fmt "fmt"
	math "math"

	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "github.com/gogo/protobuf/gogoproto"
	github_com_gogo_protobuf_jsonpb "github.com/gogo/protobuf/jsonpb"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/solo-io/solo-apis/pkg/api/gloo.solo.io/external/udpa/annotations"
	_ "github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// MarshalJSON is a custom marshaler for DatadogConfig
func (this *DatadogConfig) MarshalJSON() ([]byte, error) {
	str, err := DatadogMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for DatadogConfig
func (this *DatadogConfig) UnmarshalJSON(b []byte) error {
	return DatadogUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

var (
	DatadogMarshaler   = &github_com_gogo_protobuf_jsonpb.Marshaler{}
	DatadogUnmarshaler = &github_com_gogo_protobuf_jsonpb.Unmarshaler{}
)
