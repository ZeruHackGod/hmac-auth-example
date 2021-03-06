// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: envoy/config/filter/thrift/rate_limit/v2alpha1/rate_limit.proto

/*
	Package v2alpha1 is a generated protocol buffer package.

	It is generated from these files:
		envoy/config/filter/thrift/rate_limit/v2alpha1/rate_limit.proto

	It has these top-level messages:
		RateLimit
*/
package v2alpha1

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

// [#comment:next free field: 5]
type RateLimit struct {
	// The rate limit domain to use in the rate limit service request.
	Domain string `protobuf:"bytes,1,opt,name=domain,proto3" json:"domain,omitempty"`
	// Specifies the rate limit configuration stage. Each configured rate limit filter performs a
	// rate limit check using descriptors configured in the
	// :ref:`envoy_api_msg_config.filter.network.thrift_proxy.v2alpha1.RouteAction` for the request.
	// Only those entries with a matching stage number are used for a given filter. If not set, the
	// default stage number is 0.
	//
	// .. note::
	//
	//  The filter supports a range of 0 - 10 inclusively for stage numbers.
	Stage uint32 `protobuf:"varint,2,opt,name=stage,proto3" json:"stage,omitempty"`
	// The timeout in milliseconds for the rate limit service RPC. If not
	// set, this defaults to 20ms.
	Timeout *time.Duration `protobuf:"bytes,3,opt,name=timeout,stdduration" json:"timeout,omitempty"`
	// The filter's behaviour in case the rate limiting service does
	// not respond back. When it is set to true, Envoy will not allow traffic in case of
	// communication failure between rate limiting service and the proxy.
	// Defaults to false.
	FailureModeDeny bool `protobuf:"varint,4,opt,name=failure_mode_deny,json=failureModeDeny,proto3" json:"failure_mode_deny,omitempty"`
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

func init() {
	proto.RegisterType((*RateLimit)(nil), "envoy.config.filter.thrift.rate_limit.v2alpha1.RateLimit")
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
	if m.Timeout != nil {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintRateLimit(dAtA, i, uint64(types.SizeOfStdDuration(*m.Timeout)))
		n1, err := types.StdDurationMarshalTo(*m.Timeout, dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if m.FailureModeDeny {
		dAtA[i] = 0x20
		i++
		if m.FailureModeDeny {
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
	if m.Timeout != nil {
		l = types.SizeOfStdDuration(*m.Timeout)
		n += 1 + l + sovRateLimit(uint64(l))
	}
	if m.FailureModeDeny {
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
		case 4:
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
	proto.RegisterFile("envoy/config/filter/thrift/rate_limit/v2alpha1/rate_limit.proto", fileDescriptorRateLimit)
}

var fileDescriptorRateLimit = []byte{
	// 312 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x90, 0xb1, 0x4e, 0xf3, 0x30,
	0x14, 0x85, 0x7f, 0xb7, 0xfd, 0x4b, 0x6b, 0x84, 0x10, 0x11, 0x12, 0xa1, 0x43, 0x1a, 0x98, 0xa2,
	0x0e, 0xb6, 0x28, 0x13, 0x13, 0x52, 0xd5, 0x11, 0x96, 0x8c, 0x2c, 0x95, 0x4b, 0x6e, 0x52, 0x4b,
	0x89, 0x6f, 0x65, 0x6e, 0x2a, 0xf5, 0x4d, 0xfa, 0x24, 0x0c, 0x4c, 0x8c, 0x8c, 0xbc, 0x01, 0x28,
	0x4c, 0xbc, 0x05, 0x6a, 0x9c, 0x48, 0xdd, 0x8e, 0xce, 0x39, 0xfe, 0xee, 0x91, 0xf9, 0x3d, 0x98,
	0x0d, 0x6e, 0xe5, 0x33, 0x9a, 0x54, 0x67, 0x32, 0xd5, 0x39, 0x81, 0x95, 0xb4, 0xb2, 0x3a, 0x25,
	0x69, 0x15, 0xc1, 0x22, 0xd7, 0x85, 0x26, 0xb9, 0x99, 0xaa, 0x7c, 0xbd, 0x52, 0x37, 0x07, 0x9e,
	0x58, 0x5b, 0x24, 0xf4, 0x44, 0x0d, 0x10, 0x0e, 0x20, 0x1c, 0x40, 0x38, 0x80, 0x38, 0x28, 0xb7,
	0x80, 0x51, 0x90, 0x21, 0x66, 0x39, 0xc8, 0xfa, 0xf5, 0xb2, 0x4c, 0x65, 0x52, 0x5a, 0x45, 0x1a,
	0x8d, 0xe3, 0x8d, 0x2e, 0x36, 0x2a, 0xd7, 0x89, 0x22, 0x90, 0xad, 0x68, 0x82, 0xf3, 0x0c, 0x33,
	0xac, 0xa5, 0xdc, 0x2b, 0xe7, 0x5e, 0xbf, 0x32, 0x3e, 0x8c, 0x15, 0xc1, 0xc3, 0xfe, 0x8a, 0x77,
	0xc5, 0xfb, 0x09, 0x16, 0x4a, 0x1b, 0x9f, 0x85, 0x2c, 0x1a, 0xce, 0x86, 0x6f, 0xbf, 0xef, 0xdd,
	0x9e, 0xed, 0x84, 0x2c, 0x6e, 0x02, 0x6f, 0xcc, 0xff, 0xbf, 0x90, 0xca, 0xc0, 0xef, 0x84, 0x2c,
	0x3a, 0x69, 0x1a, 0x93, 0x8e, 0xcf, 0x63, 0xe7, 0x7b, 0x77, 0xfc, 0x88, 0x74, 0x01, 0x58, 0x92,
	0xdf, 0x0d, 0x59, 0x74, 0x3c, 0xbd, 0x14, 0x6e, 0xb2, 0x68, 0x27, 0x8b, 0x79, 0x33, 0x79, 0xd6,
	0xdb, 0x7d, 0x8d, 0x59, 0xdc, 0xf6, 0xbd, 0x09, 0x3f, 0x4b, 0x95, 0xce, 0x4b, 0x0b, 0x8b, 0x02,
	0x13, 0x58, 0x24, 0x60, 0xb6, 0x7e, 0x2f, 0x64, 0xd1, 0x20, 0x3e, 0x6d, 0x82, 0x47, 0x4c, 0x60,
	0x0e, 0x66, 0x3b, 0x1b, 0x7d, 0x54, 0x01, 0xfb, 0xac, 0x02, 0xf6, 0x5d, 0x05, 0x6c, 0xf7, 0x13,
	0xfc, 0x7b, 0x1a, 0xb4, 0x7f, 0xb4, 0xec, 0xd7, 0x97, 0x6e, 0xff, 0x02, 0x00, 0x00, 0xff, 0xff,
	0x17, 0xfe, 0x21, 0xe3, 0x9d, 0x01, 0x00, 0x00,
}
