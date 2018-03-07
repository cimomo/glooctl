// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: envoy/service/accesslog/v2/als.proto

/*
	Package v2 is a generated protocol buffer package.

	It is generated from these files:
		envoy/service/accesslog/v2/als.proto

	It has these top-level messages:
		StreamAccessLogsResponse
		StreamAccessLogsMessage
*/
package v2

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import envoy_api_v2_core "github.com/envoyproxy/go-control-plane/envoy/api/v2/core"
import envoy_config_filter_accesslog_v2 "github.com/envoyproxy/go-control-plane/envoy/config/filter/accesslog/v2"
import _ "github.com/lyft/protoc-gen-validate/validate"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// Empty response for the StreamAccessLogs API. Will never be sent. See below.
// [#not-implemented-hide:] Not configuration. TBD how to doc proto APIs.
type StreamAccessLogsResponse struct {
}

func (m *StreamAccessLogsResponse) Reset()                    { *m = StreamAccessLogsResponse{} }
func (m *StreamAccessLogsResponse) String() string            { return proto.CompactTextString(m) }
func (*StreamAccessLogsResponse) ProtoMessage()               {}
func (*StreamAccessLogsResponse) Descriptor() ([]byte, []int) { return fileDescriptorAls, []int{0} }

// [#proto-status: experimental]
// [#not-implemented-hide:] Not configuration. TBD how to doc proto APIs.
// Stream message for the StreamAccessLogs API. Envoy will open a stream to the server and stream
// access logs without ever expecting a response.
type StreamAccessLogsMessage struct {
	// Identifier data that will only be sent in the first message on the stream. This is effectively
	// structured metadata and is a performance optimization.
	Identifier *StreamAccessLogsMessage_Identifier `protobuf:"bytes,1,opt,name=identifier" json:"identifier,omitempty"`
	// Batches of log entries of a single type. Generally speaking, a given stream should only
	// ever incude one type of log entry.
	//
	// Types that are valid to be assigned to LogEntries:
	//	*StreamAccessLogsMessage_HttpLogs
	//	*StreamAccessLogsMessage_TcpLogs
	LogEntries isStreamAccessLogsMessage_LogEntries `protobuf_oneof:"log_entries"`
}

func (m *StreamAccessLogsMessage) Reset()                    { *m = StreamAccessLogsMessage{} }
func (m *StreamAccessLogsMessage) String() string            { return proto.CompactTextString(m) }
func (*StreamAccessLogsMessage) ProtoMessage()               {}
func (*StreamAccessLogsMessage) Descriptor() ([]byte, []int) { return fileDescriptorAls, []int{1} }

type isStreamAccessLogsMessage_LogEntries interface {
	isStreamAccessLogsMessage_LogEntries()
	MarshalTo([]byte) (int, error)
	Size() int
}

type StreamAccessLogsMessage_HttpLogs struct {
	HttpLogs *StreamAccessLogsMessage_HTTPAccessLogEntries `protobuf:"bytes,2,opt,name=http_logs,json=httpLogs,oneof"`
}
type StreamAccessLogsMessage_TcpLogs struct {
	TcpLogs *StreamAccessLogsMessage_TCPAccessLogEntries `protobuf:"bytes,3,opt,name=tcp_logs,json=tcpLogs,oneof"`
}

func (*StreamAccessLogsMessage_HttpLogs) isStreamAccessLogsMessage_LogEntries() {}
func (*StreamAccessLogsMessage_TcpLogs) isStreamAccessLogsMessage_LogEntries()  {}

func (m *StreamAccessLogsMessage) GetLogEntries() isStreamAccessLogsMessage_LogEntries {
	if m != nil {
		return m.LogEntries
	}
	return nil
}

func (m *StreamAccessLogsMessage) GetIdentifier() *StreamAccessLogsMessage_Identifier {
	if m != nil {
		return m.Identifier
	}
	return nil
}

func (m *StreamAccessLogsMessage) GetHttpLogs() *StreamAccessLogsMessage_HTTPAccessLogEntries {
	if x, ok := m.GetLogEntries().(*StreamAccessLogsMessage_HttpLogs); ok {
		return x.HttpLogs
	}
	return nil
}

func (m *StreamAccessLogsMessage) GetTcpLogs() *StreamAccessLogsMessage_TCPAccessLogEntries {
	if x, ok := m.GetLogEntries().(*StreamAccessLogsMessage_TcpLogs); ok {
		return x.TcpLogs
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*StreamAccessLogsMessage) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _StreamAccessLogsMessage_OneofMarshaler, _StreamAccessLogsMessage_OneofUnmarshaler, _StreamAccessLogsMessage_OneofSizer, []interface{}{
		(*StreamAccessLogsMessage_HttpLogs)(nil),
		(*StreamAccessLogsMessage_TcpLogs)(nil),
	}
}

func _StreamAccessLogsMessage_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*StreamAccessLogsMessage)
	// log_entries
	switch x := m.LogEntries.(type) {
	case *StreamAccessLogsMessage_HttpLogs:
		_ = b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.HttpLogs); err != nil {
			return err
		}
	case *StreamAccessLogsMessage_TcpLogs:
		_ = b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.TcpLogs); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("StreamAccessLogsMessage.LogEntries has unexpected type %T", x)
	}
	return nil
}

func _StreamAccessLogsMessage_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*StreamAccessLogsMessage)
	switch tag {
	case 2: // log_entries.http_logs
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(StreamAccessLogsMessage_HTTPAccessLogEntries)
		err := b.DecodeMessage(msg)
		m.LogEntries = &StreamAccessLogsMessage_HttpLogs{msg}
		return true, err
	case 3: // log_entries.tcp_logs
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(StreamAccessLogsMessage_TCPAccessLogEntries)
		err := b.DecodeMessage(msg)
		m.LogEntries = &StreamAccessLogsMessage_TcpLogs{msg}
		return true, err
	default:
		return false, nil
	}
}

func _StreamAccessLogsMessage_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*StreamAccessLogsMessage)
	// log_entries
	switch x := m.LogEntries.(type) {
	case *StreamAccessLogsMessage_HttpLogs:
		s := proto.Size(x.HttpLogs)
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *StreamAccessLogsMessage_TcpLogs:
		s := proto.Size(x.TcpLogs)
		n += proto.SizeVarint(3<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type StreamAccessLogsMessage_Identifier struct {
	// The node sending the access log messages over the stream.
	Node *envoy_api_v2_core.Node `protobuf:"bytes,1,opt,name=node" json:"node,omitempty"`
	// The friendly name of the log configured in AccessLogServiceConfig.
	LogName string `protobuf:"bytes,2,opt,name=log_name,json=logName,proto3" json:"log_name,omitempty"`
}

func (m *StreamAccessLogsMessage_Identifier) Reset()         { *m = StreamAccessLogsMessage_Identifier{} }
func (m *StreamAccessLogsMessage_Identifier) String() string { return proto.CompactTextString(m) }
func (*StreamAccessLogsMessage_Identifier) ProtoMessage()    {}
func (*StreamAccessLogsMessage_Identifier) Descriptor() ([]byte, []int) {
	return fileDescriptorAls, []int{1, 0}
}

func (m *StreamAccessLogsMessage_Identifier) GetNode() *envoy_api_v2_core.Node {
	if m != nil {
		return m.Node
	}
	return nil
}

func (m *StreamAccessLogsMessage_Identifier) GetLogName() string {
	if m != nil {
		return m.LogName
	}
	return ""
}

// Wrapper for batches of HTTP access log entries.
type StreamAccessLogsMessage_HTTPAccessLogEntries struct {
	LogEntry []*envoy_config_filter_accesslog_v2.HTTPAccessLogEntry `protobuf:"bytes,1,rep,name=log_entry,json=logEntry" json:"log_entry,omitempty"`
}

func (m *StreamAccessLogsMessage_HTTPAccessLogEntries) Reset() {
	*m = StreamAccessLogsMessage_HTTPAccessLogEntries{}
}
func (m *StreamAccessLogsMessage_HTTPAccessLogEntries) String() string {
	return proto.CompactTextString(m)
}
func (*StreamAccessLogsMessage_HTTPAccessLogEntries) ProtoMessage() {}
func (*StreamAccessLogsMessage_HTTPAccessLogEntries) Descriptor() ([]byte, []int) {
	return fileDescriptorAls, []int{1, 1}
}

func (m *StreamAccessLogsMessage_HTTPAccessLogEntries) GetLogEntry() []*envoy_config_filter_accesslog_v2.HTTPAccessLogEntry {
	if m != nil {
		return m.LogEntry
	}
	return nil
}

// Wrapper for batches of TCP access log entries.
type StreamAccessLogsMessage_TCPAccessLogEntries struct {
	LogEntry []*envoy_config_filter_accesslog_v2.TCPAccessLogEntry `protobuf:"bytes,1,rep,name=log_entry,json=logEntry" json:"log_entry,omitempty"`
}

func (m *StreamAccessLogsMessage_TCPAccessLogEntries) Reset() {
	*m = StreamAccessLogsMessage_TCPAccessLogEntries{}
}
func (m *StreamAccessLogsMessage_TCPAccessLogEntries) String() string {
	return proto.CompactTextString(m)
}
func (*StreamAccessLogsMessage_TCPAccessLogEntries) ProtoMessage() {}
func (*StreamAccessLogsMessage_TCPAccessLogEntries) Descriptor() ([]byte, []int) {
	return fileDescriptorAls, []int{1, 2}
}

func (m *StreamAccessLogsMessage_TCPAccessLogEntries) GetLogEntry() []*envoy_config_filter_accesslog_v2.TCPAccessLogEntry {
	if m != nil {
		return m.LogEntry
	}
	return nil
}

func init() {
	proto.RegisterType((*StreamAccessLogsResponse)(nil), "envoy.service.accesslog.v2.StreamAccessLogsResponse")
	proto.RegisterType((*StreamAccessLogsMessage)(nil), "envoy.service.accesslog.v2.StreamAccessLogsMessage")
	proto.RegisterType((*StreamAccessLogsMessage_Identifier)(nil), "envoy.service.accesslog.v2.StreamAccessLogsMessage.Identifier")
	proto.RegisterType((*StreamAccessLogsMessage_HTTPAccessLogEntries)(nil), "envoy.service.accesslog.v2.StreamAccessLogsMessage.HTTPAccessLogEntries")
	proto.RegisterType((*StreamAccessLogsMessage_TCPAccessLogEntries)(nil), "envoy.service.accesslog.v2.StreamAccessLogsMessage.TCPAccessLogEntries")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for AccessLogService service

type AccessLogServiceClient interface {
	// Envoy will connect and send StreamAccessLogsMessage messages forever. It does not expect any
	// response to be sent as nothing would be done in the case of failure. The server should
	// disconnect if it expects Envoy to reconnect. In the future we may decide to add a different
	// API for "critical" access logs in which Envoy will buffer access logs for some period of time
	// until it gets an ACK so it could then retry. This API is designed for high throughput with the
	// expectation that it might be lossy.
	StreamAccessLogs(ctx context.Context, opts ...grpc.CallOption) (AccessLogService_StreamAccessLogsClient, error)
}

type accessLogServiceClient struct {
	cc *grpc.ClientConn
}

func NewAccessLogServiceClient(cc *grpc.ClientConn) AccessLogServiceClient {
	return &accessLogServiceClient{cc}
}

func (c *accessLogServiceClient) StreamAccessLogs(ctx context.Context, opts ...grpc.CallOption) (AccessLogService_StreamAccessLogsClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_AccessLogService_serviceDesc.Streams[0], c.cc, "/envoy.service.accesslog.v2.AccessLogService/StreamAccessLogs", opts...)
	if err != nil {
		return nil, err
	}
	x := &accessLogServiceStreamAccessLogsClient{stream}
	return x, nil
}

type AccessLogService_StreamAccessLogsClient interface {
	Send(*StreamAccessLogsMessage) error
	CloseAndRecv() (*StreamAccessLogsResponse, error)
	grpc.ClientStream
}

type accessLogServiceStreamAccessLogsClient struct {
	grpc.ClientStream
}

func (x *accessLogServiceStreamAccessLogsClient) Send(m *StreamAccessLogsMessage) error {
	return x.ClientStream.SendMsg(m)
}

func (x *accessLogServiceStreamAccessLogsClient) CloseAndRecv() (*StreamAccessLogsResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(StreamAccessLogsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for AccessLogService service

type AccessLogServiceServer interface {
	// Envoy will connect and send StreamAccessLogsMessage messages forever. It does not expect any
	// response to be sent as nothing would be done in the case of failure. The server should
	// disconnect if it expects Envoy to reconnect. In the future we may decide to add a different
	// API for "critical" access logs in which Envoy will buffer access logs for some period of time
	// until it gets an ACK so it could then retry. This API is designed for high throughput with the
	// expectation that it might be lossy.
	StreamAccessLogs(AccessLogService_StreamAccessLogsServer) error
}

func RegisterAccessLogServiceServer(s *grpc.Server, srv AccessLogServiceServer) {
	s.RegisterService(&_AccessLogService_serviceDesc, srv)
}

func _AccessLogService_StreamAccessLogs_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(AccessLogServiceServer).StreamAccessLogs(&accessLogServiceStreamAccessLogsServer{stream})
}

type AccessLogService_StreamAccessLogsServer interface {
	SendAndClose(*StreamAccessLogsResponse) error
	Recv() (*StreamAccessLogsMessage, error)
	grpc.ServerStream
}

type accessLogServiceStreamAccessLogsServer struct {
	grpc.ServerStream
}

func (x *accessLogServiceStreamAccessLogsServer) SendAndClose(m *StreamAccessLogsResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *accessLogServiceStreamAccessLogsServer) Recv() (*StreamAccessLogsMessage, error) {
	m := new(StreamAccessLogsMessage)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _AccessLogService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "envoy.service.accesslog.v2.AccessLogService",
	HandlerType: (*AccessLogServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamAccessLogs",
			Handler:       _AccessLogService_StreamAccessLogs_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "envoy/service/accesslog/v2/als.proto",
}

func (m *StreamAccessLogsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *StreamAccessLogsResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	return i, nil
}

func (m *StreamAccessLogsMessage) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *StreamAccessLogsMessage) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Identifier != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintAls(dAtA, i, uint64(m.Identifier.Size()))
		n1, err := m.Identifier.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if m.LogEntries != nil {
		nn2, err := m.LogEntries.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += nn2
	}
	return i, nil
}

func (m *StreamAccessLogsMessage_HttpLogs) MarshalTo(dAtA []byte) (int, error) {
	i := 0
	if m.HttpLogs != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintAls(dAtA, i, uint64(m.HttpLogs.Size()))
		n3, err := m.HttpLogs.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n3
	}
	return i, nil
}
func (m *StreamAccessLogsMessage_TcpLogs) MarshalTo(dAtA []byte) (int, error) {
	i := 0
	if m.TcpLogs != nil {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintAls(dAtA, i, uint64(m.TcpLogs.Size()))
		n4, err := m.TcpLogs.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n4
	}
	return i, nil
}
func (m *StreamAccessLogsMessage_Identifier) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *StreamAccessLogsMessage_Identifier) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Node != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintAls(dAtA, i, uint64(m.Node.Size()))
		n5, err := m.Node.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n5
	}
	if len(m.LogName) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintAls(dAtA, i, uint64(len(m.LogName)))
		i += copy(dAtA[i:], m.LogName)
	}
	return i, nil
}

func (m *StreamAccessLogsMessage_HTTPAccessLogEntries) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *StreamAccessLogsMessage_HTTPAccessLogEntries) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.LogEntry) > 0 {
		for _, msg := range m.LogEntry {
			dAtA[i] = 0xa
			i++
			i = encodeVarintAls(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func (m *StreamAccessLogsMessage_TCPAccessLogEntries) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *StreamAccessLogsMessage_TCPAccessLogEntries) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.LogEntry) > 0 {
		for _, msg := range m.LogEntry {
			dAtA[i] = 0xa
			i++
			i = encodeVarintAls(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func encodeVarintAls(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *StreamAccessLogsResponse) Size() (n int) {
	var l int
	_ = l
	return n
}

func (m *StreamAccessLogsMessage) Size() (n int) {
	var l int
	_ = l
	if m.Identifier != nil {
		l = m.Identifier.Size()
		n += 1 + l + sovAls(uint64(l))
	}
	if m.LogEntries != nil {
		n += m.LogEntries.Size()
	}
	return n
}

func (m *StreamAccessLogsMessage_HttpLogs) Size() (n int) {
	var l int
	_ = l
	if m.HttpLogs != nil {
		l = m.HttpLogs.Size()
		n += 1 + l + sovAls(uint64(l))
	}
	return n
}
func (m *StreamAccessLogsMessage_TcpLogs) Size() (n int) {
	var l int
	_ = l
	if m.TcpLogs != nil {
		l = m.TcpLogs.Size()
		n += 1 + l + sovAls(uint64(l))
	}
	return n
}
func (m *StreamAccessLogsMessage_Identifier) Size() (n int) {
	var l int
	_ = l
	if m.Node != nil {
		l = m.Node.Size()
		n += 1 + l + sovAls(uint64(l))
	}
	l = len(m.LogName)
	if l > 0 {
		n += 1 + l + sovAls(uint64(l))
	}
	return n
}

func (m *StreamAccessLogsMessage_HTTPAccessLogEntries) Size() (n int) {
	var l int
	_ = l
	if len(m.LogEntry) > 0 {
		for _, e := range m.LogEntry {
			l = e.Size()
			n += 1 + l + sovAls(uint64(l))
		}
	}
	return n
}

func (m *StreamAccessLogsMessage_TCPAccessLogEntries) Size() (n int) {
	var l int
	_ = l
	if len(m.LogEntry) > 0 {
		for _, e := range m.LogEntry {
			l = e.Size()
			n += 1 + l + sovAls(uint64(l))
		}
	}
	return n
}

func sovAls(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozAls(x uint64) (n int) {
	return sovAls(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *StreamAccessLogsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAls
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
			return fmt.Errorf("proto: StreamAccessLogsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: StreamAccessLogsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipAls(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthAls
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
func (m *StreamAccessLogsMessage) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAls
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
			return fmt.Errorf("proto: StreamAccessLogsMessage: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: StreamAccessLogsMessage: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Identifier", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAls
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
				return ErrInvalidLengthAls
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Identifier == nil {
				m.Identifier = &StreamAccessLogsMessage_Identifier{}
			}
			if err := m.Identifier.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field HttpLogs", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAls
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
				return ErrInvalidLengthAls
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &StreamAccessLogsMessage_HTTPAccessLogEntries{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.LogEntries = &StreamAccessLogsMessage_HttpLogs{v}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TcpLogs", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAls
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
				return ErrInvalidLengthAls
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &StreamAccessLogsMessage_TCPAccessLogEntries{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.LogEntries = &StreamAccessLogsMessage_TcpLogs{v}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAls(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthAls
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
func (m *StreamAccessLogsMessage_Identifier) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAls
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
			return fmt.Errorf("proto: Identifier: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Identifier: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Node", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAls
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
				return ErrInvalidLengthAls
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Node == nil {
				m.Node = &envoy_api_v2_core.Node{}
			}
			if err := m.Node.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LogName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAls
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
				return ErrInvalidLengthAls
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.LogName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAls(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthAls
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
func (m *StreamAccessLogsMessage_HTTPAccessLogEntries) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAls
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
			return fmt.Errorf("proto: HTTPAccessLogEntries: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: HTTPAccessLogEntries: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LogEntry", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAls
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
				return ErrInvalidLengthAls
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.LogEntry = append(m.LogEntry, &envoy_config_filter_accesslog_v2.HTTPAccessLogEntry{})
			if err := m.LogEntry[len(m.LogEntry)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAls(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthAls
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
func (m *StreamAccessLogsMessage_TCPAccessLogEntries) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAls
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
			return fmt.Errorf("proto: TCPAccessLogEntries: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TCPAccessLogEntries: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LogEntry", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAls
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
				return ErrInvalidLengthAls
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.LogEntry = append(m.LogEntry, &envoy_config_filter_accesslog_v2.TCPAccessLogEntry{})
			if err := m.LogEntry[len(m.LogEntry)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAls(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthAls
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
func skipAls(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowAls
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
					return 0, ErrIntOverflowAls
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
					return 0, ErrIntOverflowAls
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
				return 0, ErrInvalidLengthAls
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowAls
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
				next, err := skipAls(dAtA[start:])
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
	ErrInvalidLengthAls = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowAls   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("envoy/service/accesslog/v2/als.proto", fileDescriptorAls) }

var fileDescriptorAls = []byte{
	// 476 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x93, 0x41, 0x8b, 0xd3, 0x40,
	0x14, 0xc7, 0x77, 0xd2, 0xad, 0xdb, 0xbe, 0x5e, 0xca, 0xb8, 0xd0, 0x12, 0xa4, 0x94, 0x65, 0x0f,
	0x3d, 0x25, 0x92, 0x2e, 0x78, 0x13, 0xac, 0x88, 0x15, 0x74, 0x91, 0x6c, 0x4f, 0xa2, 0x2e, 0xb3,
	0xc9, 0x6b, 0x1c, 0x4c, 0x33, 0x61, 0x66, 0x0c, 0xf4, 0xe8, 0xd5, 0xa3, 0x07, 0xbf, 0x8a, 0xe2,
	0x69, 0x8f, 0x1e, 0xfd, 0x08, 0xd2, 0xdb, 0x7e, 0x0b, 0x99, 0x49, 0x36, 0x6a, 0xb7, 0x45, 0xb6,
	0xb7, 0x90, 0x79, 0xff, 0xdf, 0x6f, 0xde, 0xbc, 0x19, 0x38, 0xc6, 0xac, 0x10, 0x4b, 0x5f, 0xa1,
	0x2c, 0x78, 0x84, 0x3e, 0x8b, 0x22, 0x54, 0x2a, 0x15, 0x89, 0x5f, 0x04, 0x3e, 0x4b, 0x95, 0x97,
	0x4b, 0xa1, 0x05, 0x75, 0x6d, 0x95, 0x57, 0x55, 0x79, 0x75, 0x95, 0x57, 0x04, 0xee, 0xbd, 0x92,
	0xc0, 0x72, 0x6e, 0x32, 0x91, 0x90, 0xe8, 0x5f, 0x30, 0x85, 0x65, 0xd2, 0xbd, 0x5f, 0xae, 0x46,
	0x22, 0x9b, 0xf3, 0xc4, 0x9f, 0xf3, 0x54, 0xa3, 0x5c, 0xb3, 0xd4, 0xb0, 0x32, 0xd1, 0x2b, 0x58,
	0xca, 0x63, 0xa6, 0xd1, 0xbf, 0xfe, 0x28, 0x17, 0x8e, 0x5c, 0xe8, 0x9f, 0x69, 0x89, 0x6c, 0xf1,
	0xc8, 0x26, 0x9e, 0x8b, 0x44, 0x85, 0xa8, 0x72, 0x91, 0x29, 0x3c, 0xfa, 0xda, 0x84, 0xde, 0xfa,
	0xe2, 0x0b, 0x54, 0x8a, 0x25, 0x48, 0xdf, 0x02, 0xf0, 0x18, 0x33, 0xcd, 0xe7, 0x1c, 0x65, 0x9f,
	0x0c, 0xc9, 0xa8, 0x13, 0x3c, 0xf4, 0xb6, 0x77, 0xe4, 0x6d, 0x01, 0x79, 0xcf, 0x6a, 0x4a, 0xf8,
	0x17, 0x91, 0x26, 0xd0, 0x7e, 0xa7, 0x75, 0x7e, 0x9e, 0x8a, 0x44, 0xf5, 0x1d, 0x8b, 0x9f, 0xee,
	0x82, 0x9f, 0xce, 0x66, 0x2f, 0xeb, 0xbf, 0x4f, 0x32, 0x2d, 0x39, 0xaa, 0xe9, 0x5e, 0xd8, 0x32,
	0x70, 0x53, 0x47, 0x63, 0x68, 0xe9, 0xa8, 0xf2, 0x34, 0xac, 0xe7, 0xe9, 0x2e, 0x9e, 0xd9, 0xe3,
	0x4d, 0x9a, 0x03, 0x1d, 0x59, 0x8b, 0xfb, 0x1e, 0xe0, 0x4f, 0xa3, 0xf4, 0x01, 0xec, 0x67, 0x22,
	0xc6, 0xea, 0xd8, 0x7a, 0x95, 0x8f, 0xe5, 0xdc, 0x18, 0xcc, 0xb0, 0xbd, 0x53, 0x11, 0xe3, 0x04,
	0xbe, 0x5f, 0x5d, 0x36, 0x9a, 0x9f, 0x88, 0xd3, 0x25, 0xa1, 0x0d, 0xd0, 0x63, 0x68, 0xa5, 0x22,
	0x39, 0xcf, 0xd8, 0x02, 0xed, 0xa1, 0xb4, 0x27, 0x6d, 0x53, 0xb3, 0x2f, 0x9d, 0x21, 0x09, 0x0f,
	0x52, 0x91, 0x9c, 0xb2, 0x05, 0xba, 0x1f, 0xe0, 0x70, 0x53, 0xdb, 0xf4, 0x0d, 0xb4, 0x4d, 0x1a,
	0x33, 0x2d, 0x97, 0x7d, 0x32, 0x6c, 0x8c, 0x3a, 0xc1, 0x49, 0xe5, 0x2e, 0xaf, 0x92, 0x57, 0x5e,
	0xa5, 0x7f, 0x3b, 0xbe, 0x81, 0x5a, 0x56, 0x1b, 0xfb, 0x4c, 0x9c, 0x16, 0x09, 0xcd, 0x86, 0xec,
	0x5f, 0x57, 0xc1, 0xdd, 0x0d, 0xa7, 0x40, 0x5f, 0xdf, 0xb4, 0x8e, 0xff, 0x6f, 0x5d, 0x27, 0x6d,
	0x91, 0x4e, 0x0e, 0xa1, 0x73, 0x4d, 0x37, 0xb2, 0xe6, 0xb7, 0xab, 0xcb, 0x06, 0x09, 0xbe, 0x10,
	0xe8, 0xd6, 0xf1, 0xb3, 0x72, 0x8e, 0xf4, 0x23, 0x81, 0xee, 0xfa, 0xf8, 0xe8, 0x78, 0x87, 0x61,
	0xbb, 0x27, 0xb7, 0x09, 0xd5, 0xcf, 0x69, 0x6f, 0x44, 0x26, 0xdd, 0x1f, 0xab, 0x01, 0xf9, 0xb9,
	0x1a, 0x90, 0x5f, 0xab, 0x01, 0x79, 0xe5, 0x14, 0xc1, 0xc5, 0x1d, 0xfb, 0x0e, 0xc7, 0xbf, 0x03,
	0x00, 0x00, 0xff, 0xff, 0x39, 0x05, 0x98, 0xa6, 0x34, 0x04, 0x00, 0x00,
}