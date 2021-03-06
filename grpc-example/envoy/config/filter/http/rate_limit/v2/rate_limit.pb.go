// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: envoy/config/filter/http/rate_limit/v2/rate_limit.proto

/*
	Package v2 is a generated protocol buffer package.

	It is generated from these files:
		envoy/config/filter/http/rate_limit/v2/rate_limit.proto

	It has these top-level messages:
		RateLimit
*/
package v2

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/types"
import _ "github.com/lyft/protoc-gen-validate/validate"
import _ "github.com/gogo/protobuf/gogoproto"

import time "time"

import types "github.com/gogo/protobuf/types"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type RateLimit struct {
	// The rate limit domain to use when calling the rate limit service.
	Domain string `protobuf:"bytes,1,opt,name=domain,proto3" json:"domain,omitempty"`
	// Specifies the rate limit configurations to be applied with the same
	// stage number. If not set, the default stage number is 0.
	//
	// .. note::
	//
	//  The filter supports a range of 0 - 10 inclusively for stage numbers.
	Stage uint32 `protobuf:"varint,2,opt,name=stage,proto3" json:"stage,omitempty"`
	// The type of requests the filter should apply to. The supported
	// types are *internal*, *external* or *both*. A request is considered internal if
	// :ref:`x-envoy-internal<config_http_conn_man_headers_x-envoy-internal>` is set to true. If
	// :ref:`x-envoy-internal<config_http_conn_man_headers_x-envoy-internal>` is not set or false, a
	// request is considered external. The filter defaults to *both*, and it will apply to all request
	// types.
	RequestType string `protobuf:"bytes,3,opt,name=request_type,json=requestType,proto3" json:"request_type,omitempty"`
	// The timeout in milliseconds for the rate limit service RPC. If not
	// set, this defaults to 20ms.
	Timeout *time.Duration `protobuf:"bytes,4,opt,name=timeout,stdduration" json:"timeout,omitempty"`
	// The filter's behaviour in case the rate limiting service does
	// not respond back. When it is set to true, Envoy will not allow traffic in case of
	// communication failure between rate limiting service and the proxy.
	// Defaults to false.
	FailureModeDeny bool `protobuf:"varint,5,opt,name=failure_mode_deny,json=failureModeDeny,proto3" json:"failure_mode_deny,omitempty"`
	// Specifies whether a `RESOURCE_EXHAUSTED` gRPC code must be returned instead
	// of the default `UNAVAILABLE` gRPC code for a rate limited gRPC call. The
	// HTTP code will be 200 for a gRPC response.
	RateLimitedAsResourceExhausted bool `protobuf:"varint,6,opt,name=rate_limited_as_resource_exhausted,json=rateLimitedAsResourceExhausted,proto3" json:"rate_limited_as_resource_exhausted,omitempty"`
}

func (m *RateLimit) Reset()                    { *m = RateLimit{} }
func (m *RateLimit) String() string            { return proto.CompactTextString(m) }
func (*RateLimit) ProtoMessage()               {}
func (*RateLimit) Descriptor() ([]byte, []int) { return fileDescriptorRateLimit, []int{0} }

func (m *RateLimit) GetDomain() string {
	if m != nil {
		return m.Domain
	}
	return ""
}

func (m *RateLimit) GetStage() uint32 {
	if m != nil {
		return m.Stage
	}
	return 0
}

func (m *RateLimit) GetRequestType() string {
	if m != nil {
		return m.RequestType
	}
	return ""
}

func (m *RateLimit) GetTimeout() *time.Duration {
	if m != nil {
		return m.Timeout
	}
	return nil
}

func (m *RateLimit) GetFailureModeDeny() bool {
	if m != nil {
		return m.FailureModeDeny
	}
	return false
}

func (m *RateLimit) GetRateLimitedAsResourceExhausted() bool {
	if m != nil {
		return m.RateLimitedAsResourceExhausted
	}
	return false
}

func init() {
	proto.RegisterType((*RateLimit)(nil), "envoy.config.filter.http.rate_limit.v2.RateLimit")
}
func (m *RateLimit) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RateLimit) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Domain) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintRateLimit(dAtA, i, uint64(len(m.Domain)))
		i += copy(dAtA[i:], m.Domain)
	}
	if m.Stage != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintRateLimit(dAtA, i, uint64(m.Stage))
	}
	if len(m.RequestType) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintRateLimit(dAtA, i, uint64(len(m.RequestType)))
		i += copy(dAtA[i:], m.RequestType)
	}
	if m.Timeout != nil {
		dAtA[i] = 0x22
		i++
		i = encodeVarintRateLimit(dAtA, i, uint64(types.SizeOfStdDuration(*m.Timeout)))
		n1, err := types.StdDurationMarshalTo(*m.Timeout, dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if m.FailureModeDeny {
		dAtA[i] = 0x28
		i++
		if m.FailureModeDeny {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.RateLimitedAsResourceExhausted {
		dAtA[i] = 0x30
		i++
		if m.RateLimitedAsResourceExhausted {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	return i, nil
}

func encodeVarintRateLimit(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *RateLimit) Size() (n int) {
	var l int
	_ = l
	l = len(m.Domain)
	if l > 0 {
		n += 1 + l + sovRateLimit(uint64(l))
	}
	if m.Stage != 0 {
		n += 1 + sovRateLimit(uint64(m.Stage))
	}
	l = len(m.RequestType)
	if l > 0 {
		n += 1 + l + sovRateLimit(uint64(l))
	}
	if m.Timeout != nil {
		l = types.SizeOfStdDuration(*m.Timeout)
		n += 1 + l + sovRateLimit(uint64(l))
	}
	if m.FailureModeDeny {
		n += 2
	}
	if m.RateLimitedAsResourceExhausted {
		n += 2
	}
	return n
}

func sovRateLimit(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozRateLimit(x uint64) (n int) {
	return sovRateLimit(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *RateLimit) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRateLimit
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: RateLimit: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RateLimit: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Domain", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRateLimit
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthRateLimit
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Domain = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Stage", wireType)
			}
			m.Stage = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRateLimit
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Stage |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RequestType", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRateLimit
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthRateLimit
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RequestType = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timeout", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRateLimit
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRateLimit
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Timeout == nil {
				m.Timeout = new(time.Duration)
			}
			if err := types.StdDurationUnmarshal(m.Timeout, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field FailureModeDeny", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRateLimit
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.FailureModeDeny = bool(v != 0)
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RateLimitedAsResourceExhausted", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRateLimit
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.RateLimitedAsResourceExhausted = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipRateLimit(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRateLimit
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipRateLimit(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowRateLimit
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowRateLimit
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowRateLimit
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthRateLimit
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowRateLimit
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipRateLimit(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthRateLimit = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowRateLimit   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("envoy/config/filter/http/rate_limit/v2/rate_limit.proto", fileDescriptorRateLimit)
}

var fileDescriptorRateLimit = []byte{
	// 372 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x90, 0xbd, 0x6e, 0xdb, 0x30,
	0x14, 0x85, 0x4b, 0xf9, 0xa7, 0x35, 0xdd, 0xa2, 0xa8, 0x60, 0xa0, 0xaa, 0x07, 0x59, 0xf6, 0x50,
	0x08, 0x1e, 0x48, 0x40, 0x1d, 0x8a, 0x8e, 0x35, 0xdc, 0xa5, 0x68, 0x17, 0x21, 0x53, 0x16, 0x81,
	0x36, 0xaf, 0x64, 0x02, 0x92, 0xa8, 0x50, 0x94, 0x10, 0xbd, 0x89, 0x5f, 0x20, 0x2f, 0x91, 0x29,
	0x63, 0xc6, 0xbc, 0x41, 0x02, 0x67, 0xca, 0x5b, 0x04, 0xfa, 0x43, 0xbc, 0x5d, 0xde, 0x73, 0xbe,
	0x03, 0x9e, 0x8b, 0x7f, 0x42, 0x5a, 0xca, 0x8a, 0xee, 0x65, 0x1a, 0x8a, 0x88, 0x86, 0x22, 0xd6,
	0xa0, 0xe8, 0x41, 0xeb, 0x8c, 0x2a, 0xa6, 0x21, 0x88, 0x45, 0x22, 0x34, 0x2d, 0xbd, 0xb3, 0x17,
	0xc9, 0x94, 0xd4, 0xd2, 0xfc, 0xde, 0x80, 0xa4, 0x05, 0x49, 0x0b, 0x92, 0x1a, 0x24, 0x67, 0xd6,
	0xd2, 0x9b, 0xdb, 0x91, 0x94, 0x51, 0x0c, 0xb4, 0xa1, 0x76, 0x45, 0x48, 0x79, 0xa1, 0x98, 0x16,
	0x32, 0x6d, 0x73, 0xe6, 0x5f, 0x4b, 0x16, 0x0b, 0xce, 0x34, 0xd0, 0x7e, 0xe8, 0x84, 0x59, 0x24,
	0x23, 0xd9, 0x8c, 0xb4, 0x9e, 0xda, 0xed, 0xea, 0xc6, 0xc0, 0x13, 0x9f, 0x69, 0xf8, 0x57, 0xe7,
	0x9b, 0x4b, 0x3c, 0xe6, 0x32, 0x61, 0x22, 0xb5, 0x90, 0x83, 0xdc, 0xc9, 0x66, 0x72, 0xfb, 0x72,
	0x37, 0x18, 0x2a, 0xc3, 0x41, 0x7e, 0x27, 0x98, 0x0b, 0x3c, 0xca, 0x35, 0x8b, 0xc0, 0x32, 0x1c,
	0xe4, 0x7e, 0xea, 0x1c, 0x6b, 0xc3, 0xc2, 0x7e, 0xbb, 0x37, 0x97, 0xf8, 0xa3, 0x82, 0xab, 0x02,
	0x72, 0x1d, 0xe8, 0x2a, 0x03, 0x6b, 0x50, 0x27, 0xf9, 0xd3, 0x6e, 0x77, 0x51, 0x65, 0x60, 0xfe,
	0xc2, 0xef, 0xb5, 0x48, 0x40, 0x16, 0xda, 0x1a, 0x3a, 0xc8, 0x9d, 0x7a, 0xdf, 0x48, 0xdb, 0x8a,
	0xf4, 0xad, 0xc8, 0xb6, 0x6b, 0xb5, 0x19, 0x1e, 0x1f, 0x17, 0xc8, 0xef, 0xfd, 0xe6, 0x1a, 0x7f,
	0x09, 0x99, 0x88, 0x0b, 0x05, 0x41, 0x22, 0x39, 0x04, 0x1c, 0xd2, 0xca, 0x1a, 0x39, 0xc8, 0xfd,
	0xe0, 0x7f, 0xee, 0x84, 0xff, 0x92, 0xc3, 0x16, 0xd2, 0xca, 0xfc, 0x8b, 0x57, 0x6f, 0xb7, 0x03,
	0x1e, 0xb0, 0x3c, 0x50, 0x90, 0xcb, 0x42, 0xed, 0x21, 0x80, 0xeb, 0x03, 0x2b, 0x72, 0x0d, 0xdc,
	0x1a, 0x37, 0xb0, 0xad, 0xfa, 0x23, 0x00, 0xff, 0x9d, 0xfb, 0x9d, 0xed, 0x4f, 0xef, 0xda, 0xcc,
	0xee, 0x4f, 0x36, 0x7a, 0x38, 0xd9, 0xe8, 0xe9, 0x64, 0xa3, 0xe3, 0xb3, 0xfd, 0xee, 0xd2, 0x28,
	0xbd, 0xdd, 0xb8, 0xf9, 0xef, 0x8f, 0xd7, 0x00, 0x00, 0x00, 0xff, 0xff, 0x67, 0x54, 0x00, 0xc5,
	0xf6, 0x01, 0x00, 0x00,
}
