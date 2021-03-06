// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: config/mcp/v1alpha1/mcp.proto

package v1alpha1

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf2 "github.com/gogo/protobuf/types"
import google_rpc "github.com/gogo/googleapis/google/rpc"
import _ "github.com/gogo/protobuf/gogoproto"

import context "golang.org/x/net/context"
import grpc "google.golang.org/grpc"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Identifies a specific MCP client instance. The client identifier is
// presented to the management server, which may use this identifier
// to distinguish per client configuration for serving. This
// information is not authoriative. Authoritative identity should come
// from the underlying transport layer (e.g. rpc credentials).
type Client struct {
	// An opaque identifier for the MCP client.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Opaque metadata extending the client identifier.
	Metadata *google_protobuf2.Struct `protobuf:"bytes,2,opt,name=metadata" json:"metadata,omitempty"`
}

func (m *Client) Reset()                    { *m = Client{} }
func (m *Client) String() string            { return proto.CompactTextString(m) }
func (*Client) ProtoMessage()               {}
func (*Client) Descriptor() ([]byte, []int) { return fileDescriptorMcp, []int{0} }

func (m *Client) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Client) GetMetadata() *google_protobuf2.Struct {
	if m != nil {
		return m.Metadata
	}
	return nil
}

// A MeshConfigRequest requests a set of versioned resources of the
// same type for a given client.
type MeshConfigRequest struct {
	// The version_info provided in the request messages will be the
	// version_info received with the most recent successfully processed
	// response or empty on the first request. It is expected that no
	// new request is sent after a response is received until the client
	// instance is ready to ACK/NACK the new configuration. ACK/NACK
	// takes place by returning the new API config version as applied or
	// the previous API config version respectively. Each type_url (see
	// below) has an independent version associated with it.
	VersionInfo string `protobuf:"bytes,1,opt,name=version_info,json=versionInfo,proto3" json:"version_info,omitempty"`
	// The client making the request.
	Client *Client `protobuf:"bytes,2,opt,name=client" json:"client,omitempty"`
	// Type of the resource that is being requested, e.g.
	// "type.googleapis.com/istio.io.networking.v1alpha3.VirtualService".
	TypeUrl string `protobuf:"bytes,3,opt,name=type_url,json=typeUrl,proto3" json:"type_url,omitempty"`
	// The nonce corresponding to MeshConfigResponse being
	// ACK/NACKed. See above discussion on version_info and the
	// MeshConfigResponse nonce comment. This may be empty if no nonce is
	// available, e.g. at startup.
	ResponseNonce string `protobuf:"bytes,4,opt,name=response_nonce,json=responseNonce,proto3" json:"response_nonce,omitempty"`
	// This is populated when the previous MeshConfigResponse failed to
	// update configuration. The *message* field in *error_details*
	// provides the client internal exception related to the failure. It
	// is only intended for consumption during manual debugging, the
	// string provided is not guaranteed to be stable across client
	// versions.
	ErrorDetail *google_rpc.Status `protobuf:"bytes,5,opt,name=error_detail,json=errorDetail" json:"error_detail,omitempty"`
}

func (m *MeshConfigRequest) Reset()                    { *m = MeshConfigRequest{} }
func (m *MeshConfigRequest) String() string            { return proto.CompactTextString(m) }
func (*MeshConfigRequest) ProtoMessage()               {}
func (*MeshConfigRequest) Descriptor() ([]byte, []int) { return fileDescriptorMcp, []int{1} }

func (m *MeshConfigRequest) GetVersionInfo() string {
	if m != nil {
		return m.VersionInfo
	}
	return ""
}

func (m *MeshConfigRequest) GetClient() *Client {
	if m != nil {
		return m.Client
	}
	return nil
}

func (m *MeshConfigRequest) GetTypeUrl() string {
	if m != nil {
		return m.TypeUrl
	}
	return ""
}

func (m *MeshConfigRequest) GetResponseNonce() string {
	if m != nil {
		return m.ResponseNonce
	}
	return ""
}

func (m *MeshConfigRequest) GetErrorDetail() *google_rpc.Status {
	if m != nil {
		return m.ErrorDetail
	}
	return nil
}

// A MeshConfigResponse delivers a set of versioned resources of the
// same type in response to a MeshConfigRequest.
type MeshConfigResponse struct {
	// The version of the response data.
	VersionInfo string `protobuf:"bytes,1,opt,name=version_info,json=versionInfo,proto3" json:"version_info,omitempty"`
	// The response resources wrapped in the common MCP *Envelope*
	// message.
	Envelopes []Envelope `protobuf:"bytes,2,rep,name=envelopes" json:"envelopes"`
	// Type URL for resources wrapped in the provided envelope(s). This
	// must be consistent with the type_url in the wrapper messages if
	// envelopes is non-empty.
	TypeUrl string `protobuf:"bytes,3,opt,name=type_url,json=typeUrl,proto3" json:"type_url,omitempty"`
	// The nonce provides a way to explicitly ack a specific
	// MeshConfigResponse in a following MeshConfigRequest. Additional
	// messages may have been sent by client to the management server for
	// the previous version on the stream prior to this
	// MeshConfigResponse, that were unprocessed at response send
	// time. The nonce allows the management server to ignore any
	// further MeshConfigRequests for the previous version until a
	// MeshConfigRequest bearing the nonce.
	Nonce string `protobuf:"bytes,4,opt,name=nonce,proto3" json:"nonce,omitempty"`
}

func (m *MeshConfigResponse) Reset()                    { *m = MeshConfigResponse{} }
func (m *MeshConfigResponse) String() string            { return proto.CompactTextString(m) }
func (*MeshConfigResponse) ProtoMessage()               {}
func (*MeshConfigResponse) Descriptor() ([]byte, []int) { return fileDescriptorMcp, []int{2} }

func (m *MeshConfigResponse) GetVersionInfo() string {
	if m != nil {
		return m.VersionInfo
	}
	return ""
}

func (m *MeshConfigResponse) GetEnvelopes() []Envelope {
	if m != nil {
		return m.Envelopes
	}
	return nil
}

func (m *MeshConfigResponse) GetTypeUrl() string {
	if m != nil {
		return m.TypeUrl
	}
	return ""
}

func (m *MeshConfigResponse) GetNonce() string {
	if m != nil {
		return m.Nonce
	}
	return ""
}

func init() {
	proto.RegisterType((*Client)(nil), "istio.config.mcp.v1alpha1.Client")
	proto.RegisterType((*MeshConfigRequest)(nil), "istio.config.mcp.v1alpha1.MeshConfigRequest")
	proto.RegisterType((*MeshConfigResponse)(nil), "istio.config.mcp.v1alpha1.MeshConfigResponse")
}
func (this *Client) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Client)
	if !ok {
		that2, ok := that.(Client)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Id != that1.Id {
		return false
	}
	if !this.Metadata.Equal(that1.Metadata) {
		return false
	}
	return true
}
func (this *MeshConfigRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*MeshConfigRequest)
	if !ok {
		that2, ok := that.(MeshConfigRequest)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.VersionInfo != that1.VersionInfo {
		return false
	}
	if !this.Client.Equal(that1.Client) {
		return false
	}
	if this.TypeUrl != that1.TypeUrl {
		return false
	}
	if this.ResponseNonce != that1.ResponseNonce {
		return false
	}
	if !this.ErrorDetail.Equal(that1.ErrorDetail) {
		return false
	}
	return true
}
func (this *MeshConfigResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*MeshConfigResponse)
	if !ok {
		that2, ok := that.(MeshConfigResponse)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.VersionInfo != that1.VersionInfo {
		return false
	}
	if len(this.Envelopes) != len(that1.Envelopes) {
		return false
	}
	for i := range this.Envelopes {
		if !this.Envelopes[i].Equal(&that1.Envelopes[i]) {
			return false
		}
	}
	if this.TypeUrl != that1.TypeUrl {
		return false
	}
	if this.Nonce != that1.Nonce {
		return false
	}
	return true
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for AggregatedMeshConfigService service

type AggregatedMeshConfigServiceClient interface {
	// StreamAggregatedResources provides the ability to carefully
	// sequence updates across multiple resource types. A single stream
	// is used with multiple independent MeshConfigRequest /
	// MeshConfigResponses sequences multiplexed via the type URL.
	StreamAggregatedResources(ctx context.Context, opts ...grpc.CallOption) (AggregatedMeshConfigService_StreamAggregatedResourcesClient, error)
}

type aggregatedMeshConfigServiceClient struct {
	cc *grpc.ClientConn
}

func NewAggregatedMeshConfigServiceClient(cc *grpc.ClientConn) AggregatedMeshConfigServiceClient {
	return &aggregatedMeshConfigServiceClient{cc}
}

func (c *aggregatedMeshConfigServiceClient) StreamAggregatedResources(ctx context.Context, opts ...grpc.CallOption) (AggregatedMeshConfigService_StreamAggregatedResourcesClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_AggregatedMeshConfigService_serviceDesc.Streams[0], c.cc, "/istio.config.mcp.v1alpha1.AggregatedMeshConfigService/StreamAggregatedResources", opts...)
	if err != nil {
		return nil, err
	}
	x := &aggregatedMeshConfigServiceStreamAggregatedResourcesClient{stream}
	return x, nil
}

type AggregatedMeshConfigService_StreamAggregatedResourcesClient interface {
	Send(*MeshConfigRequest) error
	Recv() (*MeshConfigResponse, error)
	grpc.ClientStream
}

type aggregatedMeshConfigServiceStreamAggregatedResourcesClient struct {
	grpc.ClientStream
}

func (x *aggregatedMeshConfigServiceStreamAggregatedResourcesClient) Send(m *MeshConfigRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *aggregatedMeshConfigServiceStreamAggregatedResourcesClient) Recv() (*MeshConfigResponse, error) {
	m := new(MeshConfigResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for AggregatedMeshConfigService service

type AggregatedMeshConfigServiceServer interface {
	// StreamAggregatedResources provides the ability to carefully
	// sequence updates across multiple resource types. A single stream
	// is used with multiple independent MeshConfigRequest /
	// MeshConfigResponses sequences multiplexed via the type URL.
	StreamAggregatedResources(AggregatedMeshConfigService_StreamAggregatedResourcesServer) error
}

func RegisterAggregatedMeshConfigServiceServer(s *grpc.Server, srv AggregatedMeshConfigServiceServer) {
	s.RegisterService(&_AggregatedMeshConfigService_serviceDesc, srv)
}

func _AggregatedMeshConfigService_StreamAggregatedResources_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(AggregatedMeshConfigServiceServer).StreamAggregatedResources(&aggregatedMeshConfigServiceStreamAggregatedResourcesServer{stream})
}

type AggregatedMeshConfigService_StreamAggregatedResourcesServer interface {
	Send(*MeshConfigResponse) error
	Recv() (*MeshConfigRequest, error)
	grpc.ServerStream
}

type aggregatedMeshConfigServiceStreamAggregatedResourcesServer struct {
	grpc.ServerStream
}

func (x *aggregatedMeshConfigServiceStreamAggregatedResourcesServer) Send(m *MeshConfigResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *aggregatedMeshConfigServiceStreamAggregatedResourcesServer) Recv() (*MeshConfigRequest, error) {
	m := new(MeshConfigRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _AggregatedMeshConfigService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "istio.config.mcp.v1alpha1.AggregatedMeshConfigService",
	HandlerType: (*AggregatedMeshConfigServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamAggregatedResources",
			Handler:       _AggregatedMeshConfigService_StreamAggregatedResources_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "config/mcp/v1alpha1/mcp.proto",
}

func (m *Client) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Client) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Id) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintMcp(dAtA, i, uint64(len(m.Id)))
		i += copy(dAtA[i:], m.Id)
	}
	if m.Metadata != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintMcp(dAtA, i, uint64(m.Metadata.Size()))
		n1, err := m.Metadata.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	return i, nil
}

func (m *MeshConfigRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MeshConfigRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.VersionInfo) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintMcp(dAtA, i, uint64(len(m.VersionInfo)))
		i += copy(dAtA[i:], m.VersionInfo)
	}
	if m.Client != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintMcp(dAtA, i, uint64(m.Client.Size()))
		n2, err := m.Client.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	if len(m.TypeUrl) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintMcp(dAtA, i, uint64(len(m.TypeUrl)))
		i += copy(dAtA[i:], m.TypeUrl)
	}
	if len(m.ResponseNonce) > 0 {
		dAtA[i] = 0x22
		i++
		i = encodeVarintMcp(dAtA, i, uint64(len(m.ResponseNonce)))
		i += copy(dAtA[i:], m.ResponseNonce)
	}
	if m.ErrorDetail != nil {
		dAtA[i] = 0x2a
		i++
		i = encodeVarintMcp(dAtA, i, uint64(m.ErrorDetail.Size()))
		n3, err := m.ErrorDetail.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n3
	}
	return i, nil
}

func (m *MeshConfigResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MeshConfigResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.VersionInfo) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintMcp(dAtA, i, uint64(len(m.VersionInfo)))
		i += copy(dAtA[i:], m.VersionInfo)
	}
	if len(m.Envelopes) > 0 {
		for _, msg := range m.Envelopes {
			dAtA[i] = 0x12
			i++
			i = encodeVarintMcp(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if len(m.TypeUrl) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintMcp(dAtA, i, uint64(len(m.TypeUrl)))
		i += copy(dAtA[i:], m.TypeUrl)
	}
	if len(m.Nonce) > 0 {
		dAtA[i] = 0x22
		i++
		i = encodeVarintMcp(dAtA, i, uint64(len(m.Nonce)))
		i += copy(dAtA[i:], m.Nonce)
	}
	return i, nil
}

func encodeVarintMcp(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Client) Size() (n int) {
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovMcp(uint64(l))
	}
	if m.Metadata != nil {
		l = m.Metadata.Size()
		n += 1 + l + sovMcp(uint64(l))
	}
	return n
}

func (m *MeshConfigRequest) Size() (n int) {
	var l int
	_ = l
	l = len(m.VersionInfo)
	if l > 0 {
		n += 1 + l + sovMcp(uint64(l))
	}
	if m.Client != nil {
		l = m.Client.Size()
		n += 1 + l + sovMcp(uint64(l))
	}
	l = len(m.TypeUrl)
	if l > 0 {
		n += 1 + l + sovMcp(uint64(l))
	}
	l = len(m.ResponseNonce)
	if l > 0 {
		n += 1 + l + sovMcp(uint64(l))
	}
	if m.ErrorDetail != nil {
		l = m.ErrorDetail.Size()
		n += 1 + l + sovMcp(uint64(l))
	}
	return n
}

func (m *MeshConfigResponse) Size() (n int) {
	var l int
	_ = l
	l = len(m.VersionInfo)
	if l > 0 {
		n += 1 + l + sovMcp(uint64(l))
	}
	if len(m.Envelopes) > 0 {
		for _, e := range m.Envelopes {
			l = e.Size()
			n += 1 + l + sovMcp(uint64(l))
		}
	}
	l = len(m.TypeUrl)
	if l > 0 {
		n += 1 + l + sovMcp(uint64(l))
	}
	l = len(m.Nonce)
	if l > 0 {
		n += 1 + l + sovMcp(uint64(l))
	}
	return n
}

func sovMcp(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozMcp(x uint64) (n int) {
	return sovMcp(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Client) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMcp
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
			return fmt.Errorf("proto: Client: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Client: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMcp
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
				return ErrInvalidLengthMcp
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Metadata", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMcp
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
				return ErrInvalidLengthMcp
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Metadata == nil {
				m.Metadata = &google_protobuf2.Struct{}
			}
			if err := m.Metadata.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMcp(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMcp
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
func (m *MeshConfigRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMcp
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
			return fmt.Errorf("proto: MeshConfigRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MeshConfigRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field VersionInfo", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMcp
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
				return ErrInvalidLengthMcp
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.VersionInfo = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Client", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMcp
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
				return ErrInvalidLengthMcp
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Client == nil {
				m.Client = &Client{}
			}
			if err := m.Client.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TypeUrl", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMcp
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
				return ErrInvalidLengthMcp
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TypeUrl = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ResponseNonce", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMcp
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
				return ErrInvalidLengthMcp
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ResponseNonce = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ErrorDetail", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMcp
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
				return ErrInvalidLengthMcp
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ErrorDetail == nil {
				m.ErrorDetail = &google_rpc.Status{}
			}
			if err := m.ErrorDetail.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMcp(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMcp
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
func (m *MeshConfigResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMcp
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
			return fmt.Errorf("proto: MeshConfigResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MeshConfigResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field VersionInfo", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMcp
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
				return ErrInvalidLengthMcp
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.VersionInfo = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Envelopes", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMcp
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
				return ErrInvalidLengthMcp
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Envelopes = append(m.Envelopes, Envelope{})
			if err := m.Envelopes[len(m.Envelopes)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TypeUrl", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMcp
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
				return ErrInvalidLengthMcp
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TypeUrl = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Nonce", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMcp
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
				return ErrInvalidLengthMcp
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Nonce = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMcp(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMcp
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
func skipMcp(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMcp
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
					return 0, ErrIntOverflowMcp
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
					return 0, ErrIntOverflowMcp
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
				return 0, ErrInvalidLengthMcp
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowMcp
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
				next, err := skipMcp(dAtA[start:])
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
	ErrInvalidLengthMcp = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMcp   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("config/mcp/v1alpha1/mcp.proto", fileDescriptorMcp) }

var fileDescriptorMcp = []byte{
	// 470 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xc1, 0x6e, 0x13, 0x31,
	0x10, 0x86, 0xeb, 0xb4, 0x0d, 0xad, 0x53, 0x2a, 0x61, 0x55, 0xea, 0x26, 0x40, 0x48, 0x17, 0x21,
	0xe5, 0x00, 0x5e, 0x9a, 0x8a, 0x03, 0x47, 0x5a, 0x10, 0xe2, 0x50, 0x0e, 0x1b, 0x71, 0xe1, 0x12,
	0xb9, 0xde, 0xc9, 0xd6, 0xd2, 0xc6, 0x36, 0xb6, 0x37, 0x12, 0x17, 0xde, 0x84, 0x3b, 0x57, 0xde,
	0xa2, 0x47, 0x9e, 0x00, 0xa1, 0x3d, 0xf2, 0x14, 0x68, 0xed, 0x5d, 0x52, 0x89, 0x12, 0xf5, 0xe6,
	0x19, 0x7f, 0x33, 0xf3, 0xff, 0x63, 0xe3, 0x87, 0x5c, 0xc9, 0xb9, 0xc8, 0x93, 0x05, 0xd7, 0xc9,
	0xf2, 0x98, 0x15, 0xfa, 0x92, 0x1d, 0xd7, 0x01, 0xd5, 0x46, 0x39, 0x45, 0xfa, 0xc2, 0x3a, 0xa1,
	0x68, 0x80, 0x68, 0x9d, 0x6f, 0xa1, 0xc1, 0x83, 0x5c, 0xa9, 0xbc, 0x80, 0xc4, 0x83, 0x17, 0xe5,
	0x3c, 0xb1, 0xce, 0x94, 0xdc, 0x85, 0xc2, 0xc1, 0x61, 0x73, 0x6b, 0x34, 0x4f, 0xac, 0x63, 0xae,
	0xb4, 0xcd, 0xc5, 0x41, 0xae, 0x72, 0xe5, 0x8f, 0x49, 0x7d, 0x6a, 0xb2, 0xf1, 0x4d, 0x32, 0x40,
	0x2e, 0xa1, 0x50, 0x1a, 0x02, 0x13, 0x9f, 0xe3, 0xee, 0x59, 0x21, 0x40, 0x3a, 0xb2, 0x8f, 0x3b,
	0x22, 0x8b, 0xd0, 0x08, 0x8d, 0x77, 0xd3, 0x8e, 0xc8, 0xc8, 0x09, 0xde, 0x59, 0x80, 0x63, 0x19,
	0x73, 0x2c, 0xea, 0x8c, 0xd0, 0xb8, 0x37, 0x39, 0xa4, 0x61, 0x3e, 0x6d, 0xd5, 0xd1, 0xa9, 0x57,
	0x97, 0xfe, 0x05, 0xe3, 0xdf, 0x08, 0xdf, 0x3b, 0x07, 0x7b, 0x79, 0xe6, 0x27, 0xa7, 0xf0, 0xa9,
	0x04, 0xeb, 0xc8, 0x11, 0xde, 0x5b, 0x82, 0xb1, 0x42, 0xc9, 0x99, 0x90, 0x73, 0xd5, 0x0c, 0xe9,
	0x35, 0xb9, 0x77, 0x72, 0xae, 0xc8, 0x4b, 0xdc, 0xe5, 0x5e, 0x47, 0x33, 0xeb, 0x88, 0xfe, 0x77,
	0x49, 0x34, 0x08, 0x4e, 0x9b, 0x02, 0xd2, 0xc7, 0x3b, 0xee, 0xb3, 0x86, 0x59, 0x69, 0x8a, 0x68,
	0xd3, 0x77, 0xbe, 0x53, 0xc7, 0x1f, 0x4c, 0x41, 0x9e, 0xe0, 0x7d, 0x03, 0x56, 0x2b, 0x69, 0x61,
	0x26, 0x95, 0xe4, 0x10, 0x6d, 0x79, 0xe0, 0x6e, 0x9b, 0x7d, 0x5f, 0x27, 0xc9, 0x0b, 0xbc, 0x07,
	0xc6, 0x28, 0x33, 0xcb, 0xc0, 0x31, 0x51, 0x44, 0xdb, 0x5e, 0x02, 0x69, 0xed, 0x1a, 0xcd, 0xe9,
	0xd4, 0xaf, 0x3b, 0xed, 0x79, 0xee, 0xb5, 0xc7, 0xe2, 0xef, 0x08, 0x93, 0xeb, 0x66, 0x43, 0xcb,
	0xdb, 0xb8, 0x7d, 0x8b, 0x77, 0xdb, 0x77, 0xb0, 0x51, 0x67, 0xb4, 0x39, 0xee, 0x4d, 0x1e, 0xaf,
	0x31, 0xfc, 0xa6, 0x61, 0x4f, 0xb7, 0xae, 0x7e, 0x3e, 0xda, 0x48, 0x57, 0xb5, 0xeb, 0xbc, 0x1f,
	0xe0, 0xed, 0xeb, 0x96, 0x43, 0x30, 0xf9, 0x8a, 0xf0, 0xfd, 0x57, 0x79, 0x6e, 0x20, 0x67, 0x0e,
	0xb2, 0x95, 0xfa, 0x29, 0x98, 0xa5, 0xe0, 0x40, 0xbe, 0xe0, 0xfe, 0xd4, 0x19, 0x60, 0x8b, 0x15,
	0x94, 0x82, 0x55, 0xa5, 0xe1, 0x60, 0xc9, 0xd3, 0x35, 0x1a, 0xff, 0x79, 0xf5, 0xc1, 0xb3, 0x5b,
	0xd2, 0x61, 0x6d, 0xf1, 0xc6, 0x18, 0x3d, 0x47, 0xa7, 0x93, 0x6f, 0xd5, 0x10, 0x5d, 0x55, 0x43,
	0xf4, 0xa3, 0x1a, 0xa2, 0x5f, 0xd5, 0x10, 0x7d, 0x1c, 0x85, 0x2e, 0x42, 0x25, 0x4c, 0x8b, 0xe4,
	0x86, 0x2f, 0x7d, 0xd1, 0xf5, 0xff, 0xf1, 0xe4, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xd5, 0x7a,
	0xb5, 0xf9, 0x77, 0x03, 0x00, 0x00,
}
