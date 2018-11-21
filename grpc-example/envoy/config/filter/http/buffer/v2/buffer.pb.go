// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: envoy/config/filter/http/buffer/v2/buffer.proto

/*
	Package v2 is a generated protocol buffer package.

	It is generated from these files:
		envoy/config/filter/http/buffer/v2/buffer.proto

	It has these top-level messages:
		Buffer
		BufferPerRoute
*/
package v2

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/types"
import google_protobuf1 "github.com/gogo/protobuf/types"
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

type Buffer struct {
	// The maximum request size that the filter will buffer before the connection
	// manager will stop buffering and return a 413 response.
	MaxRequestBytes *google_protobuf1.UInt32Value `protobuf:"bytes,1,opt,name=max_request_bytes,json=maxRequestBytes" json:"max_request_bytes,omitempty"`
	// The maximum number of seconds that the filter will wait for a complete
	// request before returning a 408 response.
	MaxRequestTime *time.Duration `protobuf:"bytes,2,opt,name=max_request_time,json=maxRequestTime,stdduration" json:"max_request_time,omitempty"`
}

func (m *Buffer) Reset()                    { *m = Buffer{} }
func (m *Buffer) String() string            { return proto.CompactTextString(m) }
func (*Buffer) ProtoMessage()               {}
func (*Buffer) Descriptor() ([]byte, []int) { return fileDescriptorBuffer, []int{0} }

func (m *Buffer) GetMaxRequestBytes() *google_protobuf1.UInt32Value {
	if m != nil {
		return m.MaxRequestBytes
	}
	return nil
}

func (m *Buffer) GetMaxRequestTime() *time.Duration {
	if m != nil {
		return m.MaxRequestTime
	}
	return nil
}

type BufferPerRoute struct {
	// Types that are valid to be assigned to Override:
	//	*BufferPerRoute_Disabled
	//	*BufferPerRoute_Buffer
	Override isBufferPerRoute_Override `protobuf_oneof:"override"`
}

func (m *BufferPerRoute) Reset()                    { *m = BufferPerRoute{} }
func (m *BufferPerRoute) String() string            { return proto.CompactTextString(m) }
func (*BufferPerRoute) ProtoMessage()               {}
func (*BufferPerRoute) Descriptor() ([]byte, []int) { return fileDescriptorBuffer, []int{1} }

type isBufferPerRoute_Override interface {
	isBufferPerRoute_Override()
	MarshalTo([]byte) (int, error)
	Size() int
}

type BufferPerRoute_Disabled struct {
	Disabled bool `protobuf:"varint,1,opt,name=disabled,proto3,oneof"`
}
type BufferPerRoute_Buffer struct {
	Buffer *Buffer `protobuf:"bytes,2,opt,name=buffer,oneof"`
}

func (*BufferPerRoute_Disabled) isBufferPerRoute_Override() {}
func (*BufferPerRoute_Buffer) isBufferPerRoute_Override()   {}

func (m *BufferPerRoute) GetOverride() isBufferPerRoute_Override {
	if m != nil {
		return m.Override
	}
	return nil
}

func (m *BufferPerRoute) GetDisabled() bool {
	if x, ok := m.GetOverride().(*BufferPerRoute_Disabled); ok {
		return x.Disabled
	}
	return false
}

func (m *BufferPerRoute) GetBuffer() *Buffer {
	if x, ok := m.GetOverride().(*BufferPerRoute_Buffer); ok {
		return x.Buffer
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*BufferPerRoute) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _BufferPerRoute_OneofMarshaler, _BufferPerRoute_OneofUnmarshaler, _BufferPerRoute_OneofSizer, []interface{}{
		(*BufferPerRoute_Disabled)(nil),
		(*BufferPerRoute_Buffer)(nil),
	}
}

func _BufferPerRoute_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*BufferPerRoute)
	// override
	switch x := m.Override.(type) {
	case *BufferPerRoute_Disabled:
		t := uint64(0)
		if x.Disabled {
			t = 1
		}
		_ = b.EncodeVarint(1<<3 | proto.WireVarint)
		_ = b.EncodeVarint(t)
	case *BufferPerRoute_Buffer:
		_ = b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Buffer); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("BufferPerRoute.Override has unexpected type %T", x)
	}
	return nil
}

func _BufferPerRoute_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*BufferPerRoute)
	switch tag {
	case 1: // override.disabled
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Override = &BufferPerRoute_Disabled{x != 0}
		return true, err
	case 2: // override.buffer
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Buffer)
		err := b.DecodeMessage(msg)
		m.Override = &BufferPerRoute_Buffer{msg}
		return true, err
	default:
		return false, nil
	}
}

func _BufferPerRoute_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*BufferPerRoute)
	// override
	switch x := m.Override.(type) {
	case *BufferPerRoute_Disabled:
		n += proto.SizeVarint(1<<3 | proto.WireVarint)
		n += 1
	case *BufferPerRoute_Buffer:
		s := proto.Size(x.Buffer)
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

func init() {
	proto.RegisterType((*Buffer)(nil), "envoy.config.filter.http.buffer.v2.Buffer")
	proto.RegisterType((*BufferPerRoute)(nil), "envoy.config.filter.http.buffer.v2.BufferPerRoute")
}
func (m *Buffer) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Buffer) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.MaxRequestBytes != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintBuffer(dAtA, i, uint64(m.MaxRequestBytes.Size()))
		n1, err := m.MaxRequestBytes.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if m.MaxRequestTime != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintBuffer(dAtA, i, uint64(types.SizeOfStdDuration(*m.MaxRequestTime)))
		n2, err := types.StdDurationMarshalTo(*m.MaxRequestTime, dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	return i, nil
}

func (m *BufferPerRoute) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BufferPerRoute) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Override != nil {
		nn3, err := m.Override.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += nn3
	}
	return i, nil
}

func (m *BufferPerRoute_Disabled) MarshalTo(dAtA []byte) (int, error) {
	i := 0
	dAtA[i] = 0x8
	i++
	if m.Disabled {
		dAtA[i] = 1
	} else {
		dAtA[i] = 0
	}
	i++
	return i, nil
}
func (m *BufferPerRoute_Buffer) MarshalTo(dAtA []byte) (int, error) {
	i := 0
	if m.Buffer != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintBuffer(dAtA, i, uint64(m.Buffer.Size()))
		n4, err := m.Buffer.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n4
	}
	return i, nil
}
func encodeVarintBuffer(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Buffer) Size() (n int) {
	var l int
	_ = l
	if m.MaxRequestBytes != nil {
		l = m.MaxRequestBytes.Size()
		n += 1 + l + sovBuffer(uint64(l))
	}
	if m.MaxRequestTime != nil {
		l = types.SizeOfStdDuration(*m.MaxRequestTime)
		n += 1 + l + sovBuffer(uint64(l))
	}
	return n
}

func (m *BufferPerRoute) Size() (n int) {
	var l int
	_ = l
	if m.Override != nil {
		n += m.Override.Size()
	}
	return n
}

func (m *BufferPerRoute_Disabled) Size() (n int) {
	var l int
	_ = l
	n += 2
	return n
}
func (m *BufferPerRoute_Buffer) Size() (n int) {
	var l int
	_ = l
	if m.Buffer != nil {
		l = m.Buffer.Size()
		n += 1 + l + sovBuffer(uint64(l))
	}
	return n
}

func sovBuffer(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozBuffer(x uint64) (n int) {
	return sovBuffer(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Buffer) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBuffer
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
			return fmt.Errorf("proto: Buffer: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Buffer: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxRequestBytes", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBuffer
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
				return ErrInvalidLengthBuffer
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.MaxRequestBytes == nil {
				m.MaxRequestBytes = &google_protobuf1.UInt32Value{}
			}
			if err := m.MaxRequestBytes.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxRequestTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBuffer
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
				return ErrInvalidLengthBuffer
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.MaxRequestTime == nil {
				m.MaxRequestTime = new(time.Duration)
			}
			if err := types.StdDurationUnmarshal(m.MaxRequestTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipBuffer(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthBuffer
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
func (m *BufferPerRoute) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBuffer
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
			return fmt.Errorf("proto: BufferPerRoute: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BufferPerRoute: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Disabled", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBuffer
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
			b := bool(v != 0)
			m.Override = &BufferPerRoute_Disabled{b}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Buffer", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBuffer
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
				return ErrInvalidLengthBuffer
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &Buffer{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Override = &BufferPerRoute_Buffer{v}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipBuffer(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthBuffer
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
func skipBuffer(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowBuffer
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
					return 0, ErrIntOverflowBuffer
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
					return 0, ErrIntOverflowBuffer
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
				return 0, ErrInvalidLengthBuffer
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowBuffer
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
				next, err := skipBuffer(dAtA[start:])
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
	ErrInvalidLengthBuffer = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowBuffer   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("envoy/config/filter/http/buffer/v2/buffer.proto", fileDescriptorBuffer)
}

var fileDescriptorBuffer = []byte{
	// 380 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x91, 0x3f, 0xee, 0xd3, 0x30,
	0x1c, 0xc5, 0xe3, 0xf4, 0x0f, 0xc1, 0x48, 0xa5, 0x8d, 0x2a, 0x51, 0x10, 0x0a, 0x55, 0x17, 0x50,
	0x07, 0x5b, 0x4a, 0x6f, 0x10, 0x31, 0xc0, 0x56, 0x05, 0xca, 0xc0, 0x52, 0x39, 0xe4, 0x9b, 0x60,
	0x94, 0xc4, 0xc1, 0x71, 0x42, 0x7b, 0x05, 0x4e, 0xd0, 0x89, 0x03, 0x30, 0x33, 0x20, 0xa6, 0x8e,
	0x8c, 0xdc, 0x00, 0x54, 0xa6, 0xde, 0x02, 0x25, 0x76, 0x01, 0xa9, 0xc3, 0x6f, 0x7b, 0xca, 0xcb,
	0xfb, 0xbc, 0xf7, 0x4d, 0x30, 0x85, 0xa2, 0x11, 0x7b, 0xfa, 0x46, 0x14, 0x09, 0x4f, 0x69, 0xc2,
	0x33, 0x05, 0x92, 0xbe, 0x55, 0xaa, 0xa4, 0x51, 0x9d, 0x24, 0x20, 0x69, 0xe3, 0x1b, 0x45, 0x4a,
	0x29, 0x94, 0x70, 0x17, 0x5d, 0x80, 0xe8, 0x00, 0xd1, 0x01, 0xd2, 0x06, 0x88, 0x79, 0xad, 0xf1,
	0x1f, 0x78, 0xa9, 0x10, 0x69, 0x06, 0xb4, 0x4b, 0x44, 0x75, 0x42, 0xe3, 0x5a, 0x32, 0xc5, 0x45,
	0xa1, 0x19, 0xd7, 0xfe, 0x07, 0xc9, 0xca, 0x12, 0x64, 0x65, 0xfc, 0x7b, 0x0d, 0xcb, 0x78, 0xcc,
	0x14, 0xd0, 0x8b, 0x30, 0xc6, 0x34, 0x15, 0xa9, 0xe8, 0x24, 0x6d, 0x95, 0x7e, 0xba, 0xf8, 0x82,
	0xf0, 0x30, 0xe8, 0xca, 0xdd, 0x17, 0x78, 0x92, 0xb3, 0xdd, 0x56, 0xc2, 0xfb, 0x1a, 0x2a, 0xb5,
	0x8d, 0xf6, 0x0a, 0xaa, 0x19, 0x9a, 0xa3, 0x27, 0x77, 0xfc, 0x87, 0x44, 0xb7, 0x92, 0x4b, 0x2b,
	0xd9, 0x3c, 0x2f, 0xd4, 0xca, 0x7f, 0xc5, 0xb2, 0x1a, 0x82, 0xdb, 0xdf, 0xce, 0xc7, 0x5e, 0x7f,
	0x69, 0xcf, 0xad, 0xf0, 0x6e, 0xce, 0x76, 0xa1, 0x06, 0x04, 0x6d, 0xde, 0xdd, 0xe0, 0xf1, 0xff,
	0x50, 0xc5, 0x73, 0x98, 0xd9, 0x1d, 0xf3, 0xfe, 0x15, 0xf3, 0xa9, 0xb9, 0x34, 0x18, 0x1f, 0x7e,
	0x3e, 0x42, 0x2d, 0xf4, 0xd6, 0x67, 0xd4, 0x77, 0xd0, 0xd2, 0x0a, 0x47, 0xff, 0xb8, 0x2f, 0x79,
	0x0e, 0x8b, 0x4f, 0x08, 0x8f, 0xf4, 0xec, 0x35, 0xc8, 0x50, 0xd4, 0x0a, 0xdc, 0xc7, 0xd8, 0x89,
	0x79, 0xc5, 0xa2, 0x0c, 0xe2, 0x6e, 0xb5, 0x63, 0x76, 0xbd, 0xb3, 0x1d, 0xf4, 0xcc, 0x0a, 0xff,
	0x9a, 0xee, 0x1a, 0x0f, 0xf5, 0xe7, 0x36, 0x43, 0x96, 0xe4, 0xe6, 0xdf, 0x42, 0x74, 0x59, 0x80,
	0x5b, 0xe4, 0xe0, 0x23, 0xb2, 0xc7, 0x2d, 0xd3, 0x70, 0x82, 0x09, 0x76, 0x44, 0x03, 0x52, 0xf2,
	0x18, 0xdc, 0xc1, 0xd7, 0xf3, 0xb1, 0x87, 0x82, 0xe9, 0xf7, 0x93, 0x87, 0x7e, 0x9c, 0x3c, 0xf4,
	0xeb, 0xe4, 0xa1, 0xc3, 0x6f, 0xcf, 0x7a, 0x6d, 0x37, 0x7e, 0x34, 0xec, 0x6e, 0x5d, 0xfd, 0x09,
	0x00, 0x00, 0xff, 0xff, 0xee, 0xc3, 0x36, 0x04, 0x3a, 0x02, 0x00, 0x00,
}
