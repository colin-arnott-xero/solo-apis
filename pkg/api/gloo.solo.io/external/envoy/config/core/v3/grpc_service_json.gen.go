// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/solo-apis/api/gloo/gloo/external/envoy/config/core/v3/grpc_service.proto

package v3

import (
	bytes "bytes"
	fmt "fmt"
	math "math"

	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "github.com/gogo/protobuf/gogoproto"
	github_com_gogo_protobuf_jsonpb "github.com/gogo/protobuf/jsonpb"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/gogo/protobuf/types"
	_ "github.com/solo-io/solo-apis/pkg/api/gloo.solo.io/external/udpa/annotations"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// MarshalJSON is a custom marshaler for GrpcService
func (this *GrpcService) MarshalJSON() ([]byte, error) {
	str, err := GrpcServiceMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for GrpcService
func (this *GrpcService) UnmarshalJSON(b []byte) error {
	return GrpcServiceUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for GrpcService_EnvoyGrpc
func (this *GrpcService_EnvoyGrpc) MarshalJSON() ([]byte, error) {
	str, err := GrpcServiceMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for GrpcService_EnvoyGrpc
func (this *GrpcService_EnvoyGrpc) UnmarshalJSON(b []byte) error {
	return GrpcServiceUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for GrpcService_GoogleGrpc
func (this *GrpcService_GoogleGrpc) MarshalJSON() ([]byte, error) {
	str, err := GrpcServiceMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for GrpcService_GoogleGrpc
func (this *GrpcService_GoogleGrpc) UnmarshalJSON(b []byte) error {
	return GrpcServiceUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for GrpcService_GoogleGrpc_SslCredentials
func (this *GrpcService_GoogleGrpc_SslCredentials) MarshalJSON() ([]byte, error) {
	str, err := GrpcServiceMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for GrpcService_GoogleGrpc_SslCredentials
func (this *GrpcService_GoogleGrpc_SslCredentials) UnmarshalJSON(b []byte) error {
	return GrpcServiceUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for GrpcService_GoogleGrpc_GoogleLocalCredentials
func (this *GrpcService_GoogleGrpc_GoogleLocalCredentials) MarshalJSON() ([]byte, error) {
	str, err := GrpcServiceMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for GrpcService_GoogleGrpc_GoogleLocalCredentials
func (this *GrpcService_GoogleGrpc_GoogleLocalCredentials) UnmarshalJSON(b []byte) error {
	return GrpcServiceUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for GrpcService_GoogleGrpc_ChannelCredentials
func (this *GrpcService_GoogleGrpc_ChannelCredentials) MarshalJSON() ([]byte, error) {
	str, err := GrpcServiceMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for GrpcService_GoogleGrpc_ChannelCredentials
func (this *GrpcService_GoogleGrpc_ChannelCredentials) UnmarshalJSON(b []byte) error {
	return GrpcServiceUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for GrpcService_GoogleGrpc_CallCredentials
func (this *GrpcService_GoogleGrpc_CallCredentials) MarshalJSON() ([]byte, error) {
	str, err := GrpcServiceMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for GrpcService_GoogleGrpc_CallCredentials
func (this *GrpcService_GoogleGrpc_CallCredentials) UnmarshalJSON(b []byte) error {
	return GrpcServiceUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for GrpcService_GoogleGrpc_CallCredentials_ServiceAccountJWTAccessCredentials
func (this *GrpcService_GoogleGrpc_CallCredentials_ServiceAccountJWTAccessCredentials) MarshalJSON() ([]byte, error) {
	str, err := GrpcServiceMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for GrpcService_GoogleGrpc_CallCredentials_ServiceAccountJWTAccessCredentials
func (this *GrpcService_GoogleGrpc_CallCredentials_ServiceAccountJWTAccessCredentials) UnmarshalJSON(b []byte) error {
	return GrpcServiceUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for GrpcService_GoogleGrpc_CallCredentials_GoogleIAMCredentials
func (this *GrpcService_GoogleGrpc_CallCredentials_GoogleIAMCredentials) MarshalJSON() ([]byte, error) {
	str, err := GrpcServiceMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for GrpcService_GoogleGrpc_CallCredentials_GoogleIAMCredentials
func (this *GrpcService_GoogleGrpc_CallCredentials_GoogleIAMCredentials) UnmarshalJSON(b []byte) error {
	return GrpcServiceUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for GrpcService_GoogleGrpc_CallCredentials_MetadataCredentialsFromPlugin
func (this *GrpcService_GoogleGrpc_CallCredentials_MetadataCredentialsFromPlugin) MarshalJSON() ([]byte, error) {
	str, err := GrpcServiceMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for GrpcService_GoogleGrpc_CallCredentials_MetadataCredentialsFromPlugin
func (this *GrpcService_GoogleGrpc_CallCredentials_MetadataCredentialsFromPlugin) UnmarshalJSON(b []byte) error {
	return GrpcServiceUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for GrpcService_GoogleGrpc_CallCredentials_StsService
func (this *GrpcService_GoogleGrpc_CallCredentials_StsService) MarshalJSON() ([]byte, error) {
	str, err := GrpcServiceMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for GrpcService_GoogleGrpc_CallCredentials_StsService
func (this *GrpcService_GoogleGrpc_CallCredentials_StsService) UnmarshalJSON(b []byte) error {
	return GrpcServiceUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for GrpcService_GoogleGrpc_ChannelArgs
func (this *GrpcService_GoogleGrpc_ChannelArgs) MarshalJSON() ([]byte, error) {
	str, err := GrpcServiceMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for GrpcService_GoogleGrpc_ChannelArgs
func (this *GrpcService_GoogleGrpc_ChannelArgs) UnmarshalJSON(b []byte) error {
	return GrpcServiceUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for GrpcService_GoogleGrpc_ChannelArgs_Value
func (this *GrpcService_GoogleGrpc_ChannelArgs_Value) MarshalJSON() ([]byte, error) {
	str, err := GrpcServiceMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for GrpcService_GoogleGrpc_ChannelArgs_Value
func (this *GrpcService_GoogleGrpc_ChannelArgs_Value) UnmarshalJSON(b []byte) error {
	return GrpcServiceUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

var (
	GrpcServiceMarshaler   = &github_com_gogo_protobuf_jsonpb.Marshaler{}
	GrpcServiceUnmarshaler = &github_com_gogo_protobuf_jsonpb.Unmarshaler{}
)
