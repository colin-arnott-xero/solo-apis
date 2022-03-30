// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/solo-apis/api/gloo/gloo/v1/options/hcm/hcm.proto

package hcm

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
func (m *HttpConnectionManagerSettings) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*HttpConnectionManagerSettings)
	if !ok {
		that2, ok := that.(HttpConnectionManagerSettings)
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

	if m.GetSkipXffAppend() != target.GetSkipXffAppend() {
		return false
	}

	if strings.Compare(m.GetVia(), target.GetVia()) != 0 {
		return false
	}

	if m.GetXffNumTrustedHops() != target.GetXffNumTrustedHops() {
		return false
	}

	if h, ok := interface{}(m.GetUseRemoteAddress()).(equality.Equalizer); ok {
		if !h.Equal(target.GetUseRemoteAddress()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetUseRemoteAddress(), target.GetUseRemoteAddress()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetGenerateRequestId()).(equality.Equalizer); ok {
		if !h.Equal(target.GetGenerateRequestId()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetGenerateRequestId(), target.GetGenerateRequestId()) {
			return false
		}
	}

	if m.GetProxy_100Continue() != target.GetProxy_100Continue() {
		return false
	}

	if h, ok := interface{}(m.GetStreamIdleTimeout()).(equality.Equalizer); ok {
		if !h.Equal(target.GetStreamIdleTimeout()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetStreamIdleTimeout(), target.GetStreamIdleTimeout()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetIdleTimeout()).(equality.Equalizer); ok {
		if !h.Equal(target.GetIdleTimeout()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetIdleTimeout(), target.GetIdleTimeout()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetMaxRequestHeadersKb()).(equality.Equalizer); ok {
		if !h.Equal(target.GetMaxRequestHeadersKb()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetMaxRequestHeadersKb(), target.GetMaxRequestHeadersKb()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetRequestTimeout()).(equality.Equalizer); ok {
		if !h.Equal(target.GetRequestTimeout()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetRequestTimeout(), target.GetRequestTimeout()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetDrainTimeout()).(equality.Equalizer); ok {
		if !h.Equal(target.GetDrainTimeout()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetDrainTimeout(), target.GetDrainTimeout()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetDelayedCloseTimeout()).(equality.Equalizer); ok {
		if !h.Equal(target.GetDelayedCloseTimeout()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetDelayedCloseTimeout(), target.GetDelayedCloseTimeout()) {
			return false
		}
	}

	if strings.Compare(m.GetServerName(), target.GetServerName()) != 0 {
		return false
	}

	if m.GetStripAnyHostPort() != target.GetStripAnyHostPort() {
		return false
	}

	if m.GetAcceptHttp_10() != target.GetAcceptHttp_10() {
		return false
	}

	if strings.Compare(m.GetDefaultHostForHttp_10(), target.GetDefaultHostForHttp_10()) != 0 {
		return false
	}

	if m.GetAllowChunkedLength() != target.GetAllowChunkedLength() {
		return false
	}

	if m.GetEnableTrailers() != target.GetEnableTrailers() {
		return false
	}

	if h, ok := interface{}(m.GetTracing()).(equality.Equalizer); ok {
		if !h.Equal(target.GetTracing()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetTracing(), target.GetTracing()) {
			return false
		}
	}

	if m.GetForwardClientCertDetails() != target.GetForwardClientCertDetails() {
		return false
	}

	if h, ok := interface{}(m.GetSetCurrentClientCertDetails()).(equality.Equalizer); ok {
		if !h.Equal(target.GetSetCurrentClientCertDetails()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetSetCurrentClientCertDetails(), target.GetSetCurrentClientCertDetails()) {
			return false
		}
	}

	if m.GetPreserveExternalRequestId() != target.GetPreserveExternalRequestId() {
		return false
	}

	if len(m.GetUpgrades()) != len(target.GetUpgrades()) {
		return false
	}
	for idx, v := range m.GetUpgrades() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetUpgrades()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetUpgrades()[idx]) {
				return false
			}
		}

	}

	if h, ok := interface{}(m.GetMaxConnectionDuration()).(equality.Equalizer); ok {
		if !h.Equal(target.GetMaxConnectionDuration()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetMaxConnectionDuration(), target.GetMaxConnectionDuration()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetMaxStreamDuration()).(equality.Equalizer); ok {
		if !h.Equal(target.GetMaxStreamDuration()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetMaxStreamDuration(), target.GetMaxStreamDuration()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetMaxHeadersCount()).(equality.Equalizer); ok {
		if !h.Equal(target.GetMaxHeadersCount()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetMaxHeadersCount(), target.GetMaxHeadersCount()) {
			return false
		}
	}

	if m.GetHeadersWithUnderscoresAction() != target.GetHeadersWithUnderscoresAction() {
		return false
	}

	if h, ok := interface{}(m.GetMaxRequestsPerConnection()).(equality.Equalizer); ok {
		if !h.Equal(target.GetMaxRequestsPerConnection()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetMaxRequestsPerConnection(), target.GetMaxRequestsPerConnection()) {
			return false
		}
	}

	if m.GetServerHeaderTransformation() != target.GetServerHeaderTransformation() {
		return false
	}

	if m.GetPathWithEscapedSlashesAction() != target.GetPathWithEscapedSlashesAction() {
		return false
	}

	if m.GetCodecType() != target.GetCodecType() {
		return false
	}

	if m.GetMergeSlashes() != target.GetMergeSlashes() {
		return false
	}

	if h, ok := interface{}(m.GetNormalizePath()).(equality.Equalizer); ok {
		if !h.Equal(target.GetNormalizePath()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetNormalizePath(), target.GetNormalizePath()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetUuidRequestIdConfig()).(equality.Equalizer); ok {
		if !h.Equal(target.GetUuidRequestIdConfig()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetUuidRequestIdConfig(), target.GetUuidRequestIdConfig()) {
			return false
		}
	}

	switch m.HeaderFormat.(type) {

	case *HttpConnectionManagerSettings_ProperCaseHeaderKeyFormat:
		if _, ok := target.HeaderFormat.(*HttpConnectionManagerSettings_ProperCaseHeaderKeyFormat); !ok {
			return false
		}

		if m.GetProperCaseHeaderKeyFormat() != target.GetProperCaseHeaderKeyFormat() {
			return false
		}

	case *HttpConnectionManagerSettings_PreserveCaseHeaderKeyFormat:
		if _, ok := target.HeaderFormat.(*HttpConnectionManagerSettings_PreserveCaseHeaderKeyFormat); !ok {
			return false
		}

		if m.GetPreserveCaseHeaderKeyFormat() != target.GetPreserveCaseHeaderKeyFormat() {
			return false
		}

	default:
		// m is nil but target is not nil
		if m.HeaderFormat != target.HeaderFormat {
			return false
		}
	}

	return true
}

// Equal function
func (m *HttpConnectionManagerSettings_SetCurrentClientCertDetails) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*HttpConnectionManagerSettings_SetCurrentClientCertDetails)
	if !ok {
		that2, ok := that.(HttpConnectionManagerSettings_SetCurrentClientCertDetails)
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

	if h, ok := interface{}(m.GetSubject()).(equality.Equalizer); ok {
		if !h.Equal(target.GetSubject()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetSubject(), target.GetSubject()) {
			return false
		}
	}

	if m.GetCert() != target.GetCert() {
		return false
	}

	if m.GetChain() != target.GetChain() {
		return false
	}

	if m.GetDns() != target.GetDns() {
		return false
	}

	if m.GetUri() != target.GetUri() {
		return false
	}

	return true
}

// Equal function
func (m *HttpConnectionManagerSettings_UuidRequestIdConfigSettings) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*HttpConnectionManagerSettings_UuidRequestIdConfigSettings)
	if !ok {
		that2, ok := that.(HttpConnectionManagerSettings_UuidRequestIdConfigSettings)
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

	if h, ok := interface{}(m.GetPackTraceReason()).(equality.Equalizer); ok {
		if !h.Equal(target.GetPackTraceReason()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetPackTraceReason(), target.GetPackTraceReason()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetUseRequestIdForTraceSampling()).(equality.Equalizer); ok {
		if !h.Equal(target.GetUseRequestIdForTraceSampling()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetUseRequestIdForTraceSampling(), target.GetUseRequestIdForTraceSampling()) {
			return false
		}
	}

	return true
}
