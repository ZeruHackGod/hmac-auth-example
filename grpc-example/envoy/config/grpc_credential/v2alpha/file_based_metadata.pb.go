// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: envoy/config/grpc_credential/v2alpha/file_based_metadata.proto

/*
	Package v2alpha is a generated protocol buffer package.

	It is generated from these files:
		envoy/config/grpc_credential/v2alpha/file_based_metadata.proto

	It has these top-level messages:
		FileBasedMetadataConfig
*/
package v2alpha

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import envoy_api_v2_core "github.com/envoyproxy/go-control-plane/envoy/api/v2/core"

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

type FileBasedMetadataConfig struct {
	// Location or inline data of secret to use for authentication of the Google gRPC connection
	// this secret will be attached to a header of the gRPC connection
	SecretData *envoy_api_v2_core.DataSource `protobuf:"bytes,1,opt,name=secret_data,json=secretData" json:"secret_data,omitempty"`
	// Metadata header key to use for sending the secret data
	// if no header key is set, "authorization" header will be used
	HeaderKey string `protobuf:"bytes,2,opt,name=header_key,json=headerKey,proto3" json:"header_key,omitempty"`
	// Prefix to prepend to the secret in the metadata header
	// if no prefix is set, the default is to use no prefix
	HeaderPrefix string `protobuf:"bytes,3,opt,name=header_prefix,json=headerPrefix,proto3" json:"header_prefix,omitempty"`
}

func (m *FileBasedMetadataConfig) Reset()         { *m = FileBasedMetadataConfig{} }
func (m *FileBasedMetadataConfig) String() string { return proto.CompactTextString(m) }
func (*FileBasedMetadataConfig) ProtoMessage()    {}
func (*FileBasedMetadataConfig) Descriptor() ([]byte, []int) {
	return fileDescriptorFileBasedMetadata, []int{0}
}

func (m *FileBasedMetadataConfig) GetSecretData() *envoy_api_v2_core.DataSource {
	if m != nil {
		return m.SecretData
	}
	return nil
}

func (m *FileBasedMetadataConfig) GetHeaderKey() string {
	if m != nil {
		return m.HeaderKey
	}
	return ""
}

func (m *FileBasedMetadataConfig) GetHeaderPrefix() string {
	if m != nil {
		return m.HeaderPrefix
	}
	return ""
}

func init() {
	proto.RegisterType((*FileBasedMetadataConfig)(nil), "envoy.config.grpc_credential.v2alpha.FileBasedMetadataConfig")
}
func (m *FileBasedMetadataConfig) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *FileBasedMetadataConfig) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.SecretData != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintFileBasedMetadata(dAtA, i, uint64(m.SecretData.Size()))
		n1, err := m.SecretData.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if len(m.HeaderKey) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintFileBasedMetadata(dAtA, i, uint64(len(m.HeaderKey)))
		i += copy(dAtA[i:], m.HeaderKey)
	}
	if len(m.HeaderPrefix) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintFileBasedMetadata(dAtA, i, uint64(len(m.HeaderPrefix)))
		i += copy(dAtA[i:], m.HeaderPrefix)
	}
	return i, nil
}

func encodeVarintFileBasedMetadata(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *FileBasedMetadataConfig) Size() (n int) {
	var l int
	_ = l
	if m.SecretData != nil {
		l = m.SecretData.Size()
		n += 1 + l + sovFileBasedMetadata(uint64(l))
	}
	l = len(m.HeaderKey)
	if l > 0 {
		n += 1 + l + sovFileBasedMetadata(uint64(l))
	}
	l = len(m.HeaderPrefix)
	if l > 0 {
		n += 1 + l + sovFileBasedMetadata(uint64(l))
	}
	return n
}

func sovFileBasedMetadata(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozFileBasedMetadata(x uint64) (n int) {
	return sovFileBasedMetadata(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *FileBasedMetadataConfig) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFileBasedMetadata
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
			return fmt.Errorf("proto: FileBasedMetadataConfig: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: FileBasedMetadataConfig: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SecretData", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFileBasedMetadata
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
				return ErrInvalidLengthFileBasedMetadata
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.SecretData == nil {
				m.SecretData = &envoy_api_v2_core.DataSource{}
			}
			if err := m.SecretData.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field HeaderKey", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFileBasedMetadata
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
				return ErrInvalidLengthFileBasedMetadata
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.HeaderKey = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field HeaderPrefix", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFileBasedMetadata
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
				return ErrInvalidLengthFileBasedMetadata
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.HeaderPrefix = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipFileBasedMetadata(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthFileBasedMetadata
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
func skipFileBasedMetadata(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowFileBasedMetadata
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
					return 0, ErrIntOverflowFileBasedMetadata
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
					return 0, ErrIntOverflowFileBasedMetadata
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
				return 0, ErrInvalidLengthFileBasedMetadata
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowFileBasedMetadata
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
				next, err := skipFileBasedMetadata(dAtA[start:])
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
	ErrInvalidLengthFileBasedMetadata = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowFileBasedMetadata   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("envoy/config/grpc_credential/v2alpha/file_based_metadata.proto", fileDescriptorFileBasedMetadata)
}

var fileDescriptorFileBasedMetadata = []byte{
	// 259 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x8f, 0xbd, 0x4e, 0xc3, 0x30,
	0x14, 0x46, 0x31, 0x48, 0xa0, 0xba, 0xb0, 0x64, 0x21, 0x20, 0x1a, 0x55, 0xc0, 0xd0, 0xc9, 0x96,
	0xc2, 0xde, 0xa1, 0x20, 0x16, 0x84, 0x84, 0xca, 0xc6, 0x62, 0xdd, 0x3a, 0x37, 0xad, 0x45, 0x88,
	0x2d, 0xd7, 0x44, 0xe4, 0x4d, 0x58, 0x78, 0x1f, 0x46, 0x1e, 0x01, 0x85, 0x17, 0x41, 0xfe, 0x99,
	0xba, 0x1e, 0xfb, 0x9c, 0x4f, 0x97, 0xce, 0xb1, 0xed, 0x74, 0xcf, 0xa5, 0x6e, 0x6b, 0xb5, 0xe6,
	0x6b, 0x6b, 0xa4, 0x90, 0x16, 0x2b, 0x6c, 0x9d, 0x82, 0x86, 0x77, 0x25, 0x34, 0x66, 0x03, 0xbc,
	0x56, 0x0d, 0x8a, 0x15, 0x6c, 0xb1, 0x12, 0x6f, 0xe8, 0xa0, 0x02, 0x07, 0xcc, 0x58, 0xed, 0x74,
	0x76, 0x1d, 0x7c, 0x16, 0x7d, 0xb6, 0xe3, 0xb3, 0xe4, 0x9f, 0x5f, 0xc4, 0x15, 0x30, 0x8a, 0x77,
	0x25, 0x97, 0xda, 0x22, 0xf7, 0xb5, 0xd8, 0xb8, 0xfc, 0x22, 0xf4, 0xf4, 0x5e, 0x35, 0xb8, 0xf0,
	0x03, 0x8f, 0xa9, 0x7f, 0x1b, 0x8a, 0xd9, 0x9c, 0x8e, 0xb7, 0x28, 0x2d, 0x3a, 0xe1, 0x61, 0x4e,
	0xa6, 0x64, 0x36, 0x2e, 0x27, 0x2c, 0xae, 0x82, 0x51, 0xac, 0x2b, 0x99, 0xef, 0xb1, 0x3b, 0x70,
	0xf0, 0xac, 0xdf, 0xad, 0xc4, 0x25, 0x8d, 0x86, 0x27, 0xd9, 0x84, 0xd2, 0x0d, 0x42, 0x85, 0x56,
	0xbc, 0x62, 0x9f, 0xef, 0x4f, 0xc9, 0x6c, 0xb4, 0x1c, 0x45, 0xf2, 0x80, 0x7d, 0x76, 0x45, 0x4f,
	0xd2, 0xb3, 0xb1, 0x58, 0xab, 0x8f, 0xfc, 0x20, 0xfc, 0x38, 0x8e, 0xf0, 0x29, 0xb0, 0xc5, 0xd9,
	0xf7, 0x50, 0x90, 0x9f, 0xa1, 0x20, 0xbf, 0x43, 0x41, 0x3e, 0xff, 0x8a, 0xbd, 0x97, 0xa3, 0x74,
	0xd8, 0xea, 0x30, 0x5c, 0x70, 0xf3, 0x1f, 0x00, 0x00, 0xff, 0xff, 0xe6, 0xc9, 0xd9, 0x4f, 0x47,
	0x01, 0x00, 0x00,
}
