// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: envoy/service/auth/v2alpha/external_auth.proto

package v2alpha

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import envoy_api_v2_core "github.com/envoyproxy/go-control-plane/envoy/api/v2/core"
import envoy_type1 "github.com/envoyproxy/go-control-plane/envoy/type"
import google_rpc "github.com/gogo/googleapis/google/rpc"
import _ "github.com/lyft/protoc-gen-validate/validate"

import context "golang.org/x/net/context"
import grpc "google.golang.org/grpc"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type CheckRequest struct {
	// The request attributes.
	Attributes *AttributeContext `protobuf:"bytes,1,opt,name=attributes" json:"attributes,omitempty"`
}

func (m *CheckRequest) Reset()                    { *m = CheckRequest{} }
func (m *CheckRequest) String() string            { return proto.CompactTextString(m) }
func (*CheckRequest) ProtoMessage()               {}
func (*CheckRequest) Descriptor() ([]byte, []int) { return fileDescriptorExternalAuth, []int{0} }

func (m *CheckRequest) GetAttributes() *AttributeContext {
	if m != nil {
		return m.Attributes
	}
	return nil
}

// HTTP attributes for a denied response.
type DeniedHttpResponse struct {
	// This field allows the authorization service to send a HTTP response status
	// code to the downstream client other than 403 (Forbidden).
	Status *envoy_type1.HttpStatus `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	// This field allows the authorization service to send HTTP response headers
	// to the the downstream client.
	Headers []*envoy_api_v2_core.HeaderValueOption `protobuf:"bytes,2,rep,name=headers" json:"headers,omitempty"`
	// This field allows the authorization service to send a response body data
	// to the the downstream client.
	Body string `protobuf:"bytes,3,opt,name=body,proto3" json:"body,omitempty"`
}

func (m *DeniedHttpResponse) Reset()                    { *m = DeniedHttpResponse{} }
func (m *DeniedHttpResponse) String() string            { return proto.CompactTextString(m) }
func (*DeniedHttpResponse) ProtoMessage()               {}
func (*DeniedHttpResponse) Descriptor() ([]byte, []int) { return fileDescriptorExternalAuth, []int{1} }

func (m *DeniedHttpResponse) GetStatus() *envoy_type1.HttpStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *DeniedHttpResponse) GetHeaders() []*envoy_api_v2_core.HeaderValueOption {
	if m != nil {
		return m.Headers
	}
	return nil
}

func (m *DeniedHttpResponse) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

// HTTP attributes for an ok response.
type OkHttpResponse struct {
	// HTTP entity headers in addition to the original request headers. This allows the authorization
	// service to append, to add or to override headers from the original request before
	// dispatching it to the upstream. By setting `append` field to `true` in the `HeaderValueOption`,
	// the filter will append the correspondent header value to the matched request header. Note that
	// by Leaving `append` as false, the filter will either add a new header, or override an existing
	// one if there is a match.
	Headers []*envoy_api_v2_core.HeaderValueOption `protobuf:"bytes,2,rep,name=headers" json:"headers,omitempty"`
}

func (m *OkHttpResponse) Reset()                    { *m = OkHttpResponse{} }
func (m *OkHttpResponse) String() string            { return proto.CompactTextString(m) }
func (*OkHttpResponse) ProtoMessage()               {}
func (*OkHttpResponse) Descriptor() ([]byte, []int) { return fileDescriptorExternalAuth, []int{2} }

func (m *OkHttpResponse) GetHeaders() []*envoy_api_v2_core.HeaderValueOption {
	if m != nil {
		return m.Headers
	}
	return nil
}

// Intended for gRPC and Network Authorization servers `only`.
type CheckResponse struct {
	// Status `OK` allows the request. Any other status indicates the request should be denied.
	Status *google_rpc.Status `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	// An message that contains HTTP response attributes. This message is
	// used when the authorization service needs to send custom responses to the
	// downstream client or, to modify/add request headers being dispatched to the upstream.
	//
	// Types that are valid to be assigned to HttpResponse:
	//	*CheckResponse_DeniedResponse
	//	*CheckResponse_OkResponse
	HttpResponse isCheckResponse_HttpResponse `protobuf_oneof:"http_response"`
}

func (m *CheckResponse) Reset()                    { *m = CheckResponse{} }
func (m *CheckResponse) String() string            { return proto.CompactTextString(m) }
func (*CheckResponse) ProtoMessage()               {}
func (*CheckResponse) Descriptor() ([]byte, []int) { return fileDescriptorExternalAuth, []int{3} }

type isCheckResponse_HttpResponse interface {
	isCheckResponse_HttpResponse()
	MarshalTo([]byte) (int, error)
	Size() int
}

type CheckResponse_DeniedResponse struct {
	DeniedResponse *DeniedHttpResponse `protobuf:"bytes,2,opt,name=denied_response,json=deniedResponse,oneof"`
}
type CheckResponse_OkResponse struct {
	OkResponse *OkHttpResponse `protobuf:"bytes,3,opt,name=ok_response,json=okResponse,oneof"`
}

func (*CheckResponse_DeniedResponse) isCheckResponse_HttpResponse() {}
func (*CheckResponse_OkResponse) isCheckResponse_HttpResponse()     {}

func (m *CheckResponse) GetHttpResponse() isCheckResponse_HttpResponse {
	if m != nil {
		return m.HttpResponse
	}
	return nil
}

func (m *CheckResponse) GetStatus() *google_rpc.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *CheckResponse) GetDeniedResponse() *DeniedHttpResponse {
	if x, ok := m.GetHttpResponse().(*CheckResponse_DeniedResponse); ok {
		return x.DeniedResponse
	}
	return nil
}

func (m *CheckResponse) GetOkResponse() *OkHttpResponse {
	if x, ok := m.GetHttpResponse().(*CheckResponse_OkResponse); ok {
		return x.OkResponse
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*CheckResponse) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _CheckResponse_OneofMarshaler, _CheckResponse_OneofUnmarshaler, _CheckResponse_OneofSizer, []interface{}{
		(*CheckResponse_DeniedResponse)(nil),
		(*CheckResponse_OkResponse)(nil),
	}
}

func _CheckResponse_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*CheckResponse)
	// http_response
	switch x := m.HttpResponse.(type) {
	case *CheckResponse_DeniedResponse:
		_ = b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.DeniedResponse); err != nil {
			return err
		}
	case *CheckResponse_OkResponse:
		_ = b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.OkResponse); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("CheckResponse.HttpResponse has unexpected type %T", x)
	}
	return nil
}

func _CheckResponse_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*CheckResponse)
	switch tag {
	case 2: // http_response.denied_response
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(DeniedHttpResponse)
		err := b.DecodeMessage(msg)
		m.HttpResponse = &CheckResponse_DeniedResponse{msg}
		return true, err
	case 3: // http_response.ok_response
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(OkHttpResponse)
		err := b.DecodeMessage(msg)
		m.HttpResponse = &CheckResponse_OkResponse{msg}
		return true, err
	default:
		return false, nil
	}
}

func _CheckResponse_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*CheckResponse)
	// http_response
	switch x := m.HttpResponse.(type) {
	case *CheckResponse_DeniedResponse:
		s := proto.Size(x.DeniedResponse)
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *CheckResponse_OkResponse:
		s := proto.Size(x.OkResponse)
		n += proto.SizeVarint(3<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

func init() {
	proto.RegisterType((*CheckRequest)(nil), "envoy.service.auth.v2alpha.CheckRequest")
	proto.RegisterType((*DeniedHttpResponse)(nil), "envoy.service.auth.v2alpha.DeniedHttpResponse")
	proto.RegisterType((*OkHttpResponse)(nil), "envoy.service.auth.v2alpha.OkHttpResponse")
	proto.RegisterType((*CheckResponse)(nil), "envoy.service.auth.v2alpha.CheckResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Authorization service

type AuthorizationClient interface {
	// Performs authorization check based on the attributes associated with the
	// incoming request, and returns status `OK` or not `OK`.
	Check(ctx context.Context, in *CheckRequest, opts ...grpc.CallOption) (*CheckResponse, error)
}

type authorizationClient struct {
	cc *grpc.ClientConn
}

func NewAuthorizationClient(cc *grpc.ClientConn) AuthorizationClient {
	return &authorizationClient{cc}
}

func (c *authorizationClient) Check(ctx context.Context, in *CheckRequest, opts ...grpc.CallOption) (*CheckResponse, error) {
	out := new(CheckResponse)
	err := grpc.Invoke(ctx, "/envoy.service.auth.v2alpha.Authorization/Check", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Authorization service

type AuthorizationServer interface {
	// Performs authorization check based on the attributes associated with the
	// incoming request, and returns status `OK` or not `OK`.
	Check(context.Context, *CheckRequest) (*CheckResponse, error)
}

func RegisterAuthorizationServer(s *grpc.Server, srv AuthorizationServer) {
	s.RegisterService(&_Authorization_serviceDesc, srv)
}

func _Authorization_Check_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthorizationServer).Check(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/envoy.service.auth.v2alpha.Authorization/Check",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthorizationServer).Check(ctx, req.(*CheckRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Authorization_serviceDesc = grpc.ServiceDesc{
	ServiceName: "envoy.service.auth.v2alpha.Authorization",
	HandlerType: (*AuthorizationServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Check",
			Handler:    _Authorization_Check_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "envoy/service/auth/v2alpha/external_auth.proto",
}

func (m *CheckRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CheckRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Attributes != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintExternalAuth(dAtA, i, uint64(m.Attributes.Size()))
		n1, err := m.Attributes.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	return i, nil
}

func (m *DeniedHttpResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DeniedHttpResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Status != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintExternalAuth(dAtA, i, uint64(m.Status.Size()))
		n2, err := m.Status.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	if len(m.Headers) > 0 {
		for _, msg := range m.Headers {
			dAtA[i] = 0x12
			i++
			i = encodeVarintExternalAuth(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if len(m.Body) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintExternalAuth(dAtA, i, uint64(len(m.Body)))
		i += copy(dAtA[i:], m.Body)
	}
	return i, nil
}

func (m *OkHttpResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *OkHttpResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Headers) > 0 {
		for _, msg := range m.Headers {
			dAtA[i] = 0x12
			i++
			i = encodeVarintExternalAuth(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func (m *CheckResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CheckResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Status != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintExternalAuth(dAtA, i, uint64(m.Status.Size()))
		n3, err := m.Status.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n3
	}
	if m.HttpResponse != nil {
		nn4, err := m.HttpResponse.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += nn4
	}
	return i, nil
}

func (m *CheckResponse_DeniedResponse) MarshalTo(dAtA []byte) (int, error) {
	i := 0
	if m.DeniedResponse != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintExternalAuth(dAtA, i, uint64(m.DeniedResponse.Size()))
		n5, err := m.DeniedResponse.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n5
	}
	return i, nil
}
func (m *CheckResponse_OkResponse) MarshalTo(dAtA []byte) (int, error) {
	i := 0
	if m.OkResponse != nil {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintExternalAuth(dAtA, i, uint64(m.OkResponse.Size()))
		n6, err := m.OkResponse.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n6
	}
	return i, nil
}
func encodeVarintExternalAuth(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *CheckRequest) Size() (n int) {
	var l int
	_ = l
	if m.Attributes != nil {
		l = m.Attributes.Size()
		n += 1 + l + sovExternalAuth(uint64(l))
	}
	return n
}

func (m *DeniedHttpResponse) Size() (n int) {
	var l int
	_ = l
	if m.Status != nil {
		l = m.Status.Size()
		n += 1 + l + sovExternalAuth(uint64(l))
	}
	if len(m.Headers) > 0 {
		for _, e := range m.Headers {
			l = e.Size()
			n += 1 + l + sovExternalAuth(uint64(l))
		}
	}
	l = len(m.Body)
	if l > 0 {
		n += 1 + l + sovExternalAuth(uint64(l))
	}
	return n
}

func (m *OkHttpResponse) Size() (n int) {
	var l int
	_ = l
	if len(m.Headers) > 0 {
		for _, e := range m.Headers {
			l = e.Size()
			n += 1 + l + sovExternalAuth(uint64(l))
		}
	}
	return n
}

func (m *CheckResponse) Size() (n int) {
	var l int
	_ = l
	if m.Status != nil {
		l = m.Status.Size()
		n += 1 + l + sovExternalAuth(uint64(l))
	}
	if m.HttpResponse != nil {
		n += m.HttpResponse.Size()
	}
	return n
}

func (m *CheckResponse_DeniedResponse) Size() (n int) {
	var l int
	_ = l
	if m.DeniedResponse != nil {
		l = m.DeniedResponse.Size()
		n += 1 + l + sovExternalAuth(uint64(l))
	}
	return n
}
func (m *CheckResponse_OkResponse) Size() (n int) {
	var l int
	_ = l
	if m.OkResponse != nil {
		l = m.OkResponse.Size()
		n += 1 + l + sovExternalAuth(uint64(l))
	}
	return n
}

func sovExternalAuth(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozExternalAuth(x uint64) (n int) {
	return sovExternalAuth(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *CheckRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowExternalAuth
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
			return fmt.Errorf("proto: CheckRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CheckRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Attributes", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExternalAuth
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
				return ErrInvalidLengthExternalAuth
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Attributes == nil {
				m.Attributes = &AttributeContext{}
			}
			if err := m.Attributes.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipExternalAuth(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthExternalAuth
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
func (m *DeniedHttpResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowExternalAuth
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
			return fmt.Errorf("proto: DeniedHttpResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DeniedHttpResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExternalAuth
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
				return ErrInvalidLengthExternalAuth
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Status == nil {
				m.Status = &envoy_type1.HttpStatus{}
			}
			if err := m.Status.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Headers", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExternalAuth
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
				return ErrInvalidLengthExternalAuth
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Headers = append(m.Headers, &envoy_api_v2_core.HeaderValueOption{})
			if err := m.Headers[len(m.Headers)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Body", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExternalAuth
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
				return ErrInvalidLengthExternalAuth
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Body = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipExternalAuth(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthExternalAuth
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
func (m *OkHttpResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowExternalAuth
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
			return fmt.Errorf("proto: OkHttpResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: OkHttpResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Headers", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExternalAuth
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
				return ErrInvalidLengthExternalAuth
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Headers = append(m.Headers, &envoy_api_v2_core.HeaderValueOption{})
			if err := m.Headers[len(m.Headers)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipExternalAuth(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthExternalAuth
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
func (m *CheckResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowExternalAuth
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
			return fmt.Errorf("proto: CheckResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CheckResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExternalAuth
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
				return ErrInvalidLengthExternalAuth
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Status == nil {
				m.Status = &google_rpc.Status{}
			}
			if err := m.Status.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DeniedResponse", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExternalAuth
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
				return ErrInvalidLengthExternalAuth
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &DeniedHttpResponse{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.HttpResponse = &CheckResponse_DeniedResponse{v}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OkResponse", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExternalAuth
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
				return ErrInvalidLengthExternalAuth
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &OkHttpResponse{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.HttpResponse = &CheckResponse_OkResponse{v}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipExternalAuth(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthExternalAuth
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
func skipExternalAuth(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowExternalAuth
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
					return 0, ErrIntOverflowExternalAuth
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
					return 0, ErrIntOverflowExternalAuth
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
				return 0, ErrInvalidLengthExternalAuth
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowExternalAuth
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
				next, err := skipExternalAuth(dAtA[start:])
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
	ErrInvalidLengthExternalAuth = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowExternalAuth   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("envoy/service/auth/v2alpha/external_auth.proto", fileDescriptorExternalAuth)
}

var fileDescriptorExternalAuth = []byte{
	// 485 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x52, 0xc1, 0x8a, 0x13, 0x4b,
	0x14, 0x4d, 0x25, 0x6f, 0x66, 0x78, 0x15, 0x33, 0x23, 0xb5, 0x70, 0x42, 0x18, 0x42, 0x08, 0x2e,
	0xe2, 0x20, 0x55, 0xd0, 0xee, 0x5c, 0x08, 0x33, 0xe3, 0x22, 0x0b, 0x65, 0xa4, 0x05, 0x41, 0x19,
	0x08, 0x95, 0xee, 0xcb, 0x74, 0x93, 0xb6, 0xab, 0xac, 0xba, 0xdd, 0x4c, 0xfc, 0x02, 0xf1, 0x0b,
	0xfc, 0x00, 0xbf, 0xc2, 0x95, 0x4b, 0x97, 0x7e, 0x82, 0xc4, 0x95, 0x3f, 0xe0, 0x5a, 0xba, 0xba,
	0x3a, 0x4e, 0x14, 0x83, 0xe0, 0xae, 0xe9, 0x7b, 0xce, 0xb9, 0xe7, 0x9c, 0xba, 0x94, 0x43, 0x5e,
	0xaa, 0xa5, 0xb0, 0x60, 0xca, 0x34, 0x02, 0x21, 0x0b, 0x4c, 0x44, 0x19, 0xc8, 0x4c, 0x27, 0x52,
	0xc0, 0x15, 0x82, 0xc9, 0x65, 0x36, 0xab, 0xfe, 0x72, 0x6d, 0x14, 0x2a, 0x36, 0x70, 0x78, 0xee,
	0xf1, 0xdc, 0x4d, 0x3c, 0x7e, 0x70, 0x54, 0x6b, 0x49, 0x9d, 0x8a, 0x32, 0x10, 0x91, 0x32, 0x20,
	0xe6, 0xd2, 0x42, 0xcd, 0x6c, 0xa6, 0xb8, 0xd4, 0x20, 0x12, 0x44, 0x3d, 0xb3, 0x28, 0xb1, 0xb0,
	0x7e, 0x1a, 0x6c, 0xf1, 0x21, 0x11, 0x4d, 0x3a, 0x2f, 0x10, 0x66, 0x91, 0xca, 0x11, 0xae, 0xd0,
	0x73, 0x0e, 0x2f, 0x95, 0xba, 0xcc, 0x40, 0x18, 0x1d, 0x89, 0x0d, 0xb1, 0xc3, 0x52, 0x66, 0x69,
	0x2c, 0x11, 0x44, 0xf3, 0x51, 0x0f, 0xc6, 0x17, 0xf4, 0xc6, 0x59, 0x02, 0xd1, 0x22, 0x84, 0x57,
	0x05, 0x58, 0x64, 0x8f, 0x28, 0x5d, 0x8b, 0xdb, 0x3e, 0x19, 0x91, 0x49, 0x37, 0xb8, 0xcb, 0xff,
	0x1c, 0x91, 0x9f, 0x34, 0xe8, 0xb3, 0xda, 0x49, 0x78, 0x8d, 0x3f, 0x7e, 0x4f, 0x28, 0x7b, 0x08,
	0x79, 0x0a, 0xf1, 0x14, 0x51, 0x87, 0x60, 0xb5, 0xca, 0x2d, 0xb0, 0xfb, 0x74, 0xb7, 0x76, 0xe7,
	0x17, 0xdc, 0xf2, 0x0b, 0xaa, 0x26, 0x78, 0x85, 0x7c, 0xea, 0xa6, 0xa7, 0xf4, 0xc3, 0xb7, 0x8f,
	0x9d, 0x9d, 0xb7, 0xa4, 0x7d, 0x93, 0x84, 0x9e, 0xc1, 0x1e, 0xd0, 0xbd, 0x04, 0x64, 0x0c, 0xc6,
	0xf6, 0xdb, 0xa3, 0xce, 0xa4, 0x1b, 0xdc, 0xf6, 0x64, 0xa9, 0x53, 0x5e, 0x06, 0xbc, 0x2a, 0x99,
	0x4f, 0x1d, 0xe2, 0x99, 0xcc, 0x0a, 0x38, 0xd7, 0x98, 0xaa, 0x3c, 0x6c, 0x48, 0x8c, 0xd1, 0xff,
	0xe6, 0x2a, 0x5e, 0xf6, 0x3b, 0x23, 0x32, 0xf9, 0x3f, 0x74, 0xdf, 0xe3, 0x27, 0x74, 0xff, 0x7c,
	0xb1, 0xe1, 0xf0, 0x1f, 0xb7, 0x8c, 0xbf, 0x13, 0xda, 0xf3, 0xbd, 0x7a, 0xc5, 0xe3, 0x5f, 0x32,
	0x33, 0x5e, 0xbf, 0x15, 0x37, 0x3a, 0xe2, 0x75, 0xde, 0x75, 0xc6, 0xe7, 0xf4, 0x20, 0x76, 0xad,
	0xcd, 0x8c, 0xa7, 0xf7, 0xdb, 0x8e, 0xc4, 0xb7, 0xbd, 0xc4, 0xef, 0x45, 0x4f, 0x5b, 0xe1, 0x7e,
	0x2d, 0xb4, 0xb6, 0xf1, 0x98, 0x76, 0xd5, 0xe2, 0xa7, 0x6c, 0xc7, 0xc9, 0x1e, 0x6f, 0x93, 0xdd,
	0x6c, 0x66, 0xda, 0x0a, 0xa9, 0x5a, 0xa7, 0x3a, 0x3d, 0xa0, 0x3d, 0x77, 0xb9, 0x8d, 0x60, 0xf0,
	0x92, 0xf6, 0x4e, 0x0a, 0x4c, 0x94, 0x49, 0x5f, 0xcb, 0xaa, 0x12, 0x76, 0x41, 0x77, 0x5c, 0x11,
	0x6c, 0xb2, 0x6d, 0xc9, 0xf5, 0x1b, 0x1c, 0xdc, 0xf9, 0x0b, 0xa4, 0xdf, 0x7f, 0xf4, 0x69, 0x35,
	0x24, 0x9f, 0x57, 0x43, 0xf2, 0x65, 0x35, 0x24, 0xef, 0xbe, 0x0e, 0x5b, 0x2f, 0xf6, 0x3c, 0xf0,
	0x0d, 0x21, 0xf3, 0x5d, 0x77, 0xe3, 0xf7, 0x7e, 0x04, 0x00, 0x00, 0xff, 0xff, 0x51, 0x9b, 0x19,
	0x73, 0xd3, 0x03, 0x00, 0x00,
}
