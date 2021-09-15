// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/solo-apis/api/gloo/gloo/v1/options/aws/aws.proto

package aws

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"strings"

	"github.com/golang/protobuf/proto"
	equality "github.com/solo-io/protoc-gen-ext/pkg/equality"
)

// ensure the imports are used
var (
	_ = errors.New("")
	_ = fmt.Print
	_ = binary.LittleEndian
	_ = bytes.Compare
	_ = strings.Compare
	_ = equality.Equalizer(nil)
	_ = proto.Message(nil)
)

// Equal function
func (m *UpstreamSpec) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*UpstreamSpec)
	if !ok {
		that2, ok := that.(UpstreamSpec)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if strings.Compare(m.GetRegion(), target.GetRegion()) != 0 {
		return false
	}

	if h, ok := interface{}(m.GetSecretRef()).(equality.Equalizer); ok {
		if !h.Equal(target.GetSecretRef()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetSecretRef(), target.GetSecretRef()) {
			return false
		}
	}

	if len(m.GetLambdaFunctions()) != len(target.GetLambdaFunctions()) {
		return false
	}
	for idx, v := range m.GetLambdaFunctions() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetLambdaFunctions()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetLambdaFunctions()[idx]) {
				return false
			}
		}

	}

	if strings.Compare(m.GetRoleArn(), target.GetRoleArn()) != 0 {
		return false
	}

	return true
}

// Equal function
func (m *LambdaFunctionSpec) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*LambdaFunctionSpec)
	if !ok {
		that2, ok := that.(LambdaFunctionSpec)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if strings.Compare(m.GetLogicalName(), target.GetLogicalName()) != 0 {
		return false
	}

	if strings.Compare(m.GetLambdaFunctionName(), target.GetLambdaFunctionName()) != 0 {
		return false
	}

	if strings.Compare(m.GetQualifier(), target.GetQualifier()) != 0 {
		return false
	}

	return true
}

// Equal function
func (m *DestinationSpec) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*DestinationSpec)
	if !ok {
		that2, ok := that.(DestinationSpec)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if strings.Compare(m.GetLogicalName(), target.GetLogicalName()) != 0 {
		return false
	}

	if m.GetInvocationStyle() != target.GetInvocationStyle() {
		return false
	}

	if m.GetRequestTransformation() != target.GetRequestTransformation() {
		return false
	}

	if m.GetResponseTransformation() != target.GetResponseTransformation() {
		return false
	}

	return true
}
