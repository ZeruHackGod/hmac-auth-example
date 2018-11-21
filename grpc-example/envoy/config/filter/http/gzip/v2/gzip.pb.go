// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: envoy/config/filter/http/gzip/v2/gzip.proto

/*
	Package v2 is a generated protocol buffer package.

	It is generated from these files:
		envoy/config/filter/http/gzip/v2/gzip.proto

	It has these top-level messages:
		Gzip
*/
package v2

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/gogo/protobuf/types"
import _ "github.com/lyft/protoc-gen-validate/validate"
import _ "github.com/gogo/protobuf/gogoproto"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Gzip_CompressionStrategy int32

const (
	Gzip_DEFAULT  Gzip_CompressionStrategy = 0
	Gzip_FILTERED Gzip_CompressionStrategy = 1
	Gzip_HUFFMAN  Gzip_CompressionStrategy = 2
	Gzip_RLE      Gzip_CompressionStrategy = 3
)

var Gzip_CompressionStrategy_name = map[int32]string{
	0: "DEFAULT",
	1: "FILTERED",
	2: "HUFFMAN",
	3: "RLE",
}
var Gzip_CompressionStrategy_value = map[string]int32{
	"DEFAULT":  0,
	"FILTERED": 1,
	"HUFFMAN":  2,
	"RLE":      3,
}

func (x Gzip_CompressionStrategy) String() string {
	return proto.EnumName(Gzip_CompressionStrategy_name, int32(x))
}
func (Gzip_CompressionStrategy) EnumDescriptor() ([]byte, []int) {
	return fileDescriptorGzip, []int{0, 0}
}

type Gzip_CompressionLevel_Enum int32

const (
	Gzip_CompressionLevel_DEFAULT Gzip_CompressionLevel_Enum = 0
	Gzip_CompressionLevel_BEST    Gzip_CompressionLevel_Enum = 1
	Gzip_CompressionLevel_SPEED   Gzip_CompressionLevel_Enum = 2
)

var Gzip_CompressionLevel_Enum_name = map[int32]string{
	0: "DEFAULT",
	1: "BEST",
	2: "SPEED",
}
var Gzip_CompressionLevel_Enum_value = map[string]int32{
	"DEFAULT": 0,
	"BEST":    1,
	"SPEED":   2,
}

func (x Gzip_CompressionLevel_Enum) String() string {
	return proto.EnumName(Gzip_CompressionLevel_Enum_name, int32(x))
}
func (Gzip_CompressionLevel_Enum) EnumDescriptor() ([]byte, []int) {
	return fileDescriptorGzip, []int{0, 0, 0}
}

type Gzip struct {
	// Value from 1 to 9 that controls the amount of internal memory used by zlib. Higher values
	// use more memory, but are faster and produce better compression results. The default value is 5.
	MemoryLevel *google_protobuf.UInt32Value `protobuf:"bytes,1,opt,name=memory_level,json=memoryLevel" json:"memory_level,omitempty"`
	// Minimum response length, in bytes, which will trigger compression. The default value is 30.
	ContentLength *google_protobuf.UInt32Value `protobuf:"bytes,2,opt,name=content_length,json=contentLength" json:"content_length,omitempty"`
	// A value used for selecting the zlib compression level. This setting will affect speed and
	// amount of compression applied to the content. "BEST" provides higher compression at the cost of
	// higher latency, "SPEED" provides lower compression with minimum impact on response time.
	// "DEFAULT" provides an optimal result between speed and compression. This field will be set to
	// "DEFAULT" if not specified.
	CompressionLevel Gzip_CompressionLevel_Enum `protobuf:"varint,3,opt,name=compression_level,json=compressionLevel,proto3,enum=envoy.config.filter.http.gzip.v2.Gzip_CompressionLevel_Enum" json:"compression_level,omitempty"`
	// A value used for selecting the zlib compression strategy which is directly related to the
	// characteristics of the content. Most of the time "DEFAULT" will be the best choice, though
	// there are situations which changing this parameter might produce better results. For example,
	// run-length encoding (RLE) is typically used when the content is known for having sequences
	// which same data occurs many consecutive times. For more information about each strategy, please
	// refer to zlib manual.
	CompressionStrategy Gzip_CompressionStrategy `protobuf:"varint,4,opt,name=compression_strategy,json=compressionStrategy,proto3,enum=envoy.config.filter.http.gzip.v2.Gzip_CompressionStrategy" json:"compression_strategy,omitempty"`
	// Set of strings that allows specifying which mime-types yield compression; e.g.,
	// application/json, text/html, etc. When this field is not defined, compression will be applied
	// to the following mime-types: "application/javascript", "application/json",
	// "application/xhtml+xml", "image/svg+xml", "text/css", "text/html", "text/plain", "text/xml".
	ContentType []string `protobuf:"bytes,6,rep,name=content_type,json=contentType" json:"content_type,omitempty"`
	// If true, disables compression when the response contains an etag header. When it is false, the
	// filter will preserve weak etags and remove the ones that require strong validation.
	DisableOnEtagHeader bool `protobuf:"varint,7,opt,name=disable_on_etag_header,json=disableOnEtagHeader,proto3" json:"disable_on_etag_header,omitempty"`
	// If true, removes accept-encoding from the request headers before dispatching it to the upstream
	// so that responses do not get compressed before reaching the filter.
	RemoveAcceptEncodingHeader bool `protobuf:"varint,8,opt,name=remove_accept_encoding_header,json=removeAcceptEncodingHeader,proto3" json:"remove_accept_encoding_header,omitempty"`
	// Value from 9 to 15 that represents the base two logarithmic of the compressor's window size.
	// Larger window results in better compression at the expense of memory usage. The default is 12
	// which will produce a 4096 bytes window. For more details about this parameter, please refer to
	// zlib manual > deflateInit2.
	WindowBits *google_protobuf.UInt32Value `protobuf:"bytes,9,opt,name=window_bits,json=windowBits" json:"window_bits,omitempty"`
}

func (m *Gzip) Reset()                    { *m = Gzip{} }
func (m *Gzip) String() string            { return proto.CompactTextString(m) }
func (*Gzip) ProtoMessage()               {}
func (*Gzip) Descriptor() ([]byte, []int) { return fileDescriptorGzip, []int{0} }

func (m *Gzip) GetMemoryLevel() *google_protobuf.UInt32Value {
	if m != nil {
		return m.MemoryLevel
	}
	return nil
}

func (m *Gzip) GetContentLength() *google_protobuf.UInt32Value {
	if m != nil {
		return m.ContentLength
	}
	return nil
}

func (m *Gzip) GetCompressionLevel() Gzip_CompressionLevel_Enum {
	if m != nil {
		return m.CompressionLevel
	}
	return Gzip_CompressionLevel_DEFAULT
}

func (m *Gzip) GetCompressionStrategy() Gzip_CompressionStrategy {
	if m != nil {
		return m.CompressionStrategy
	}
	return Gzip_DEFAULT
}

func (m *Gzip) GetContentType() []string {
	if m != nil {
		return m.ContentType
	}
	return nil
}

func (m *Gzip) GetDisableOnEtagHeader() bool {
	if m != nil {
		return m.DisableOnEtagHeader
	}
	return false
}

func (m *Gzip) GetRemoveAcceptEncodingHeader() bool {
	if m != nil {
		return m.RemoveAcceptEncodingHeader
	}
	return false
}

func (m *Gzip) GetWindowBits() *google_protobuf.UInt32Value {
	if m != nil {
		return m.WindowBits
	}
	return nil
}

type Gzip_CompressionLevel struct {
}

func (m *Gzip_CompressionLevel) Reset()                    { *m = Gzip_CompressionLevel{} }
func (m *Gzip_CompressionLevel) String() string            { return proto.CompactTextString(m) }
func (*Gzip_CompressionLevel) ProtoMessage()               {}
func (*Gzip_CompressionLevel) Descriptor() ([]byte, []int) { return fileDescriptorGzip, []int{0, 0} }

func init() {
	proto.RegisterType((*Gzip)(nil), "envoy.config.filter.http.gzip.v2.Gzip")
	proto.RegisterType((*Gzip_CompressionLevel)(nil), "envoy.config.filter.http.gzip.v2.Gzip.CompressionLevel")
	proto.RegisterEnum("envoy.config.filter.http.gzip.v2.Gzip_CompressionStrategy", Gzip_CompressionStrategy_name, Gzip_CompressionStrategy_value)
	proto.RegisterEnum("envoy.config.filter.http.gzip.v2.Gzip_CompressionLevel_Enum", Gzip_CompressionLevel_Enum_name, Gzip_CompressionLevel_Enum_value)
}
func (m *Gzip) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Gzip) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.MemoryLevel != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintGzip(dAtA, i, uint64(m.MemoryLevel.Size()))
		n1, err := m.MemoryLevel.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if m.ContentLength != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintGzip(dAtA, i, uint64(m.ContentLength.Size()))
		n2, err := m.ContentLength.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	if m.CompressionLevel != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintGzip(dAtA, i, uint64(m.CompressionLevel))
	}
	if m.CompressionStrategy != 0 {
		dAtA[i] = 0x20
		i++
		i = encodeVarintGzip(dAtA, i, uint64(m.CompressionStrategy))
	}
	if len(m.ContentType) > 0 {
		for _, s := range m.ContentType {
			dAtA[i] = 0x32
			i++
			l = len(s)
			for l >= 1<<7 {
				dAtA[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			dAtA[i] = uint8(l)
			i++
			i += copy(dAtA[i:], s)
		}
	}
	if m.DisableOnEtagHeader {
		dAtA[i] = 0x38
		i++
		if m.DisableOnEtagHeader {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.RemoveAcceptEncodingHeader {
		dAtA[i] = 0x40
		i++
		if m.RemoveAcceptEncodingHeader {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.WindowBits != nil {
		dAtA[i] = 0x4a
		i++
		i = encodeVarintGzip(dAtA, i, uint64(m.WindowBits.Size()))
		n3, err := m.WindowBits.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n3
	}
	return i, nil
}

func (m *Gzip_CompressionLevel) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Gzip_CompressionLevel) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	return i, nil
}

func encodeVarintGzip(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Gzip) Size() (n int) {
	var l int
	_ = l
	if m.MemoryLevel != nil {
		l = m.MemoryLevel.Size()
		n += 1 + l + sovGzip(uint64(l))
	}
	if m.ContentLength != nil {
		l = m.ContentLength.Size()
		n += 1 + l + sovGzip(uint64(l))
	}
	if m.CompressionLevel != 0 {
		n += 1 + sovGzip(uint64(m.CompressionLevel))
	}
	if m.CompressionStrategy != 0 {
		n += 1 + sovGzip(uint64(m.CompressionStrategy))
	}
	if len(m.ContentType) > 0 {
		for _, s := range m.ContentType {
			l = len(s)
			n += 1 + l + sovGzip(uint64(l))
		}
	}
	if m.DisableOnEtagHeader {
		n += 2
	}
	if m.RemoveAcceptEncodingHeader {
		n += 2
	}
	if m.WindowBits != nil {
		l = m.WindowBits.Size()
		n += 1 + l + sovGzip(uint64(l))
	}
	return n
}

func (m *Gzip_CompressionLevel) Size() (n int) {
	var l int
	_ = l
	return n
}

func sovGzip(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozGzip(x uint64) (n int) {
	return sovGzip(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Gzip) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGzip
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
			return fmt.Errorf("proto: Gzip: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Gzip: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MemoryLevel", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGzip
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
				return ErrInvalidLengthGzip
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.MemoryLevel == nil {
				m.MemoryLevel = &google_protobuf.UInt32Value{}
			}
			if err := m.MemoryLevel.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ContentLength", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGzip
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
				return ErrInvalidLengthGzip
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ContentLength == nil {
				m.ContentLength = &google_protobuf.UInt32Value{}
			}
			if err := m.ContentLength.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CompressionLevel", wireType)
			}
			m.CompressionLevel = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGzip
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CompressionLevel |= (Gzip_CompressionLevel_Enum(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CompressionStrategy", wireType)
			}
			m.CompressionStrategy = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGzip
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CompressionStrategy |= (Gzip_CompressionStrategy(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ContentType", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGzip
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
				return ErrInvalidLengthGzip
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ContentType = append(m.ContentType, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DisableOnEtagHeader", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGzip
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
			m.DisableOnEtagHeader = bool(v != 0)
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RemoveAcceptEncodingHeader", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGzip
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
			m.RemoveAcceptEncodingHeader = bool(v != 0)
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field WindowBits", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGzip
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
				return ErrInvalidLengthGzip
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.WindowBits == nil {
				m.WindowBits = &google_protobuf.UInt32Value{}
			}
			if err := m.WindowBits.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGzip(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGzip
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
func (m *Gzip_CompressionLevel) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGzip
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
			return fmt.Errorf("proto: CompressionLevel: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CompressionLevel: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipGzip(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGzip
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
func skipGzip(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGzip
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
					return 0, ErrIntOverflowGzip
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
					return 0, ErrIntOverflowGzip
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
				return 0, ErrInvalidLengthGzip
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowGzip
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
				next, err := skipGzip(dAtA[start:])
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
	ErrInvalidLengthGzip = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGzip   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("envoy/config/filter/http/gzip/v2/gzip.proto", fileDescriptorGzip) }

var fileDescriptorGzip = []byte{
	// 561 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0xc1, 0x6e, 0xd3, 0x4a,
	0x14, 0x86, 0x3b, 0x8e, 0x9b, 0xc6, 0xe3, 0xde, 0x5e, 0x33, 0xad, 0xc0, 0x8a, 0xc0, 0x8a, 0xba,
	0xb2, 0x8a, 0x18, 0x4b, 0xee, 0x0e, 0x75, 0x93, 0x50, 0x87, 0x16, 0x99, 0x82, 0xdc, 0x84, 0x05,
	0x1b, 0xcb, 0x71, 0x26, 0x8e, 0x25, 0x67, 0xc6, 0xb2, 0x27, 0x8e, 0xdc, 0x25, 0x2f, 0x80, 0xc4,
	0x8a, 0x67, 0x61, 0xc5, 0x92, 0x25, 0x8f, 0x80, 0xc2, 0x8a, 0xb7, 0x40, 0x1e, 0x3b, 0x40, 0x0b,
	0x12, 0x74, 0x95, 0xa3, 0x9c, 0xef, 0x3b, 0xff, 0xcc, 0x19, 0xc3, 0x87, 0x84, 0x16, 0xac, 0xb4,
	0x42, 0x46, 0x67, 0x71, 0x64, 0xcd, 0xe2, 0x84, 0x93, 0xcc, 0x9a, 0x73, 0x9e, 0x5a, 0xd1, 0x55,
	0x9c, 0x5a, 0x85, 0x2d, 0x7e, 0x71, 0x9a, 0x31, 0xce, 0x50, 0x4f, 0xc0, 0xb8, 0x86, 0x71, 0x0d,
	0xe3, 0x0a, 0xc6, 0x02, 0x2a, 0xec, 0xae, 0x11, 0x31, 0x16, 0x25, 0xc4, 0x12, 0xfc, 0x64, 0x39,
	0xb3, 0x56, 0x59, 0x90, 0xa6, 0x24, 0xcb, 0xeb, 0x09, 0xdd, 0x7b, 0x45, 0x90, 0xc4, 0xd3, 0x80,
	0x13, 0x6b, 0x53, 0x34, 0x8d, 0x83, 0x88, 0x45, 0x4c, 0x94, 0x56, 0x55, 0xd5, 0xff, 0x1e, 0xbe,
	0x6d, 0x43, 0xf9, 0xe9, 0x55, 0x9c, 0x22, 0x17, 0xee, 0x2e, 0xc8, 0x82, 0x65, 0xa5, 0x9f, 0x90,
	0x82, 0x24, 0x3a, 0xe8, 0x01, 0x53, 0xb5, 0xef, 0xe3, 0x3a, 0x0e, 0x6f, 0xe2, 0xf0, 0xf8, 0x9c,
	0xf2, 0x63, 0xfb, 0x55, 0x90, 0x2c, 0xc9, 0x40, 0xfd, 0xf0, 0xed, 0x63, 0xab, 0x7d, 0x24, 0xeb,
	0x8a, 0x09, 0x3c, 0xb5, 0xd6, 0xdd, 0xca, 0x46, 0x17, 0x70, 0x2f, 0x64, 0x94, 0x13, 0xca, 0xfd,
	0x84, 0xd0, 0x88, 0xcf, 0x75, 0xe9, 0x1f, 0xe6, 0x29, 0xd5, 0x3c, 0xf9, 0x48, 0x32, 0x0d, 0xef,
	0xbf, 0x46, 0x77, 0x85, 0x8d, 0x96, 0xf0, 0x4e, 0xc8, 0x16, 0x69, 0x46, 0xf2, 0x3c, 0x66, 0xb4,
	0x39, 0x62, 0xab, 0x07, 0xcc, 0x3d, 0xfb, 0x04, 0xff, 0x6d, 0x67, 0xb8, 0xba, 0x20, 0x7e, 0xf2,
	0xd3, 0x17, 0x67, 0xc4, 0x0e, 0x5d, 0x2e, 0x06, 0xb0, 0x8a, 0xdc, 0x7e, 0x03, 0x24, 0x0d, 0x78,
	0x5a, 0x78, 0x03, 0x41, 0x25, 0x3c, 0xf8, 0x35, 0x36, 0xe7, 0x59, 0xc0, 0x49, 0x54, 0xea, 0xb2,
	0x48, 0x7e, 0x7c, 0xfb, 0xe4, 0xcb, 0x66, 0xc2, 0xb5, 0xdc, 0xfd, 0xf0, 0x77, 0x00, 0x3d, 0x82,
	0xbb, 0x9b, 0x0d, 0xf2, 0x32, 0x25, 0x7a, 0xbb, 0xd7, 0x32, 0x95, 0x46, 0x7b, 0x07, 0x24, 0xcd,
	0xf6, 0xd4, 0xa6, 0x3f, 0x2a, 0x53, 0x82, 0x8e, 0xe1, 0xdd, 0x69, 0x9c, 0x07, 0x93, 0x84, 0xf8,
	0x8c, 0xfa, 0x84, 0x07, 0x91, 0x3f, 0x27, 0xc1, 0x94, 0x64, 0xfa, 0x4e, 0x0f, 0x98, 0x1d, 0x6f,
	0xbf, 0xe9, 0xbe, 0xa0, 0x0e, 0x0f, 0xa2, 0x33, 0xd1, 0x42, 0x7d, 0xf8, 0x20, 0x23, 0x0b, 0x56,
	0x10, 0x3f, 0x08, 0x43, 0x92, 0x72, 0x9f, 0xd0, 0x90, 0x4d, 0x63, 0xfa, 0xc3, 0xed, 0x08, 0xb7,
	0x5b, 0x43, 0x7d, 0xc1, 0x38, 0x0d, 0xd2, 0x8c, 0x78, 0x06, 0xd5, 0x55, 0x4c, 0xa7, 0x6c, 0xe5,
	0x4f, 0x62, 0x9e, 0xeb, 0xca, 0x6d, 0xbe, 0x9a, 0xff, 0x4d, 0xc5, 0x83, 0xb5, 0x3d, 0x88, 0x79,
	0xde, 0x3d, 0x81, 0xda, 0xcd, 0x47, 0x3a, 0x34, 0xa1, 0x5c, 0xbd, 0x13, 0x52, 0xe1, 0xce, 0xa9,
	0x33, 0xec, 0x8f, 0xdd, 0x91, 0xb6, 0x85, 0x3a, 0x50, 0x1e, 0x38, 0x97, 0x23, 0x0d, 0x20, 0x05,
	0x6e, 0x5f, 0xbe, 0x74, 0x9c, 0x53, 0x4d, 0x3a, 0x1c, 0xc2, 0xfd, 0x3f, 0x2c, 0xfa, 0xba, 0xb8,
	0x0b, 0x3b, 0xc3, 0x73, 0x77, 0xe4, 0x78, 0xce, 0xa9, 0x06, 0xaa, 0xd6, 0xd9, 0x78, 0x38, 0x7c,
	0xde, 0xbf, 0xd0, 0x24, 0xb4, 0x03, 0x5b, 0x9e, 0xeb, 0x68, 0xad, 0xc1, 0xc1, 0xa7, 0xb5, 0x01,
	0x3e, 0xaf, 0x0d, 0xf0, 0x65, 0x6d, 0x80, 0xf7, 0x5f, 0x8d, 0xad, 0xd7, 0x52, 0x61, 0x4f, 0xda,
	0xe2, 0x2a, 0xc7, 0xdf, 0x03, 0x00, 0x00, 0xff, 0xff, 0xf5, 0xfe, 0xc3, 0x2e, 0xce, 0x03, 0x00,
	0x00,
}
