// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: message.proto

package protogo

import (
	context "context"
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type Message struct {
	Msg string `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{0}
}
func (m *Message) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Message.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message.Merge(m, src)
}
func (m *Message) XXX_Size() int {
	return m.Size()
}
func (m *Message) XXX_DiscardUnknown() {
	xxx_messageInfo_Message.DiscardUnknown(m)
}

var xxx_messageInfo_Message proto.InternalMessageInfo

func (m *Message) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func init() {
	proto.RegisterType((*Message)(nil), "proto.Message")
}

func init() { proto.RegisterFile("message.proto", fileDescriptor_33c57e4bae7b9afd) }

var fileDescriptor_33c57e4bae7b9afd = []byte{
	// 184 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xcd, 0x4d, 0x2d, 0x2e,
	0x4e, 0x4c, 0x4f, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0x53, 0x4a, 0xd2, 0x5c,
	0xec, 0xbe, 0x10, 0x71, 0x21, 0x01, 0x2e, 0xe6, 0xdc, 0xe2, 0x74, 0x09, 0x46, 0x05, 0x46, 0x0d,
	0xce, 0x20, 0x10, 0xd3, 0x68, 0x37, 0x23, 0x17, 0x17, 0x54, 0x36, 0x28, 0xc0, 0x59, 0xc8, 0x9a,
	0x4b, 0x18, 0xca, 0x0b, 0x29, 0xcf, 0x77, 0xc9, 0x2c, 0x4a, 0x4d, 0x2e, 0xc9, 0xcc, 0xcf, 0x13,
	0xe2, 0x83, 0x98, 0xa8, 0x07, 0x95, 0x93, 0x42, 0xe3, 0x2b, 0x31, 0x68, 0x30, 0x1a, 0x30, 0x0a,
	0x99, 0x73, 0xf1, 0xc3, 0x34, 0xe7, 0x07, 0xa7, 0x16, 0x95, 0xa5, 0x16, 0x91, 0xa1, 0xd1, 0x39,
	0x27, 0x33, 0x35, 0xaf, 0x84, 0x38, 0x8d, 0x4e, 0x7a, 0x27, 0x1e, 0xc9, 0x31, 0x5e, 0x78, 0x24,
	0xc7, 0xf8, 0xe0, 0x91, 0x1c, 0xe3, 0x84, 0xc7, 0x72, 0x0c, 0x17, 0x1e, 0xcb, 0x31, 0xdc, 0x78,
	0x2c, 0xc7, 0x10, 0x25, 0x92, 0x5e, 0x54, 0x90, 0xac, 0x9b, 0x91, 0x9f, 0x9b, 0xaa, 0x5f, 0x90,
	0xa4, 0x0f, 0xd6, 0x9b, 0x9e, 0x9f, 0xc4, 0x06, 0x66, 0x18, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff,
	0x08, 0x11, 0xd5, 0x76, 0x29, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MessageRPCClient is the client API for MessageRPC service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MessageRPCClient interface {
	MessageTwoDirection(ctx context.Context, opts ...grpc.CallOption) (MessageRPC_MessageTwoDirectionClient, error)
	MessageToServer(ctx context.Context, opts ...grpc.CallOption) (MessageRPC_MessageToServerClient, error)
	MessageToClient(ctx context.Context, opts ...grpc.CallOption) (MessageRPC_MessageToClientClient, error)
}

type messageRPCClient struct {
	cc *grpc.ClientConn
}

func NewMessageRPCClient(cc *grpc.ClientConn) MessageRPCClient {
	return &messageRPCClient{cc}
}

func (c *messageRPCClient) MessageTwoDirection(ctx context.Context, opts ...grpc.CallOption) (MessageRPC_MessageTwoDirectionClient, error) {
	stream, err := c.cc.NewStream(ctx, &_MessageRPC_serviceDesc.Streams[0], "/proto.MessageRPC/MessageTwoDirection", opts...)
	if err != nil {
		return nil, err
	}
	x := &messageRPCMessageTwoDirectionClient{stream}
	return x, nil
}

type MessageRPC_MessageTwoDirectionClient interface {
	Send(*Message) error
	Recv() (*Message, error)
	grpc.ClientStream
}

type messageRPCMessageTwoDirectionClient struct {
	grpc.ClientStream
}

func (x *messageRPCMessageTwoDirectionClient) Send(m *Message) error {
	return x.ClientStream.SendMsg(m)
}

func (x *messageRPCMessageTwoDirectionClient) Recv() (*Message, error) {
	m := new(Message)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *messageRPCClient) MessageToServer(ctx context.Context, opts ...grpc.CallOption) (MessageRPC_MessageToServerClient, error) {
	stream, err := c.cc.NewStream(ctx, &_MessageRPC_serviceDesc.Streams[1], "/proto.MessageRPC/MessageToServer", opts...)
	if err != nil {
		return nil, err
	}
	x := &messageRPCMessageToServerClient{stream}
	return x, nil
}

type MessageRPC_MessageToServerClient interface {
	Send(*Message) error
	Recv() (*Message, error)
	grpc.ClientStream
}

type messageRPCMessageToServerClient struct {
	grpc.ClientStream
}

func (x *messageRPCMessageToServerClient) Send(m *Message) error {
	return x.ClientStream.SendMsg(m)
}

func (x *messageRPCMessageToServerClient) Recv() (*Message, error) {
	m := new(Message)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *messageRPCClient) MessageToClient(ctx context.Context, opts ...grpc.CallOption) (MessageRPC_MessageToClientClient, error) {
	stream, err := c.cc.NewStream(ctx, &_MessageRPC_serviceDesc.Streams[2], "/proto.MessageRPC/MessageToClient", opts...)
	if err != nil {
		return nil, err
	}
	x := &messageRPCMessageToClientClient{stream}
	return x, nil
}

type MessageRPC_MessageToClientClient interface {
	Send(*Message) error
	Recv() (*Message, error)
	grpc.ClientStream
}

type messageRPCMessageToClientClient struct {
	grpc.ClientStream
}

func (x *messageRPCMessageToClientClient) Send(m *Message) error {
	return x.ClientStream.SendMsg(m)
}

func (x *messageRPCMessageToClientClient) Recv() (*Message, error) {
	m := new(Message)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// MessageRPCServer is the server API for MessageRPC service.
type MessageRPCServer interface {
	MessageTwoDirection(MessageRPC_MessageTwoDirectionServer) error
	MessageToServer(MessageRPC_MessageToServerServer) error
	MessageToClient(MessageRPC_MessageToClientServer) error
}

// UnimplementedMessageRPCServer can be embedded to have forward compatible implementations.
type UnimplementedMessageRPCServer struct {
}

func (*UnimplementedMessageRPCServer) MessageTwoDirection(srv MessageRPC_MessageTwoDirectionServer) error {
	return status.Errorf(codes.Unimplemented, "method MessageTwoDirection not implemented")
}
func (*UnimplementedMessageRPCServer) MessageToServer(srv MessageRPC_MessageToServerServer) error {
	return status.Errorf(codes.Unimplemented, "method MessageToServer not implemented")
}
func (*UnimplementedMessageRPCServer) MessageToClient(srv MessageRPC_MessageToClientServer) error {
	return status.Errorf(codes.Unimplemented, "method MessageToClient not implemented")
}

func RegisterMessageRPCServer(s *grpc.Server, srv MessageRPCServer) {
	s.RegisterService(&_MessageRPC_serviceDesc, srv)
}

func _MessageRPC_MessageTwoDirection_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(MessageRPCServer).MessageTwoDirection(&messageRPCMessageTwoDirectionServer{stream})
}

type MessageRPC_MessageTwoDirectionServer interface {
	Send(*Message) error
	Recv() (*Message, error)
	grpc.ServerStream
}

type messageRPCMessageTwoDirectionServer struct {
	grpc.ServerStream
}

func (x *messageRPCMessageTwoDirectionServer) Send(m *Message) error {
	return x.ServerStream.SendMsg(m)
}

func (x *messageRPCMessageTwoDirectionServer) Recv() (*Message, error) {
	m := new(Message)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _MessageRPC_MessageToServer_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(MessageRPCServer).MessageToServer(&messageRPCMessageToServerServer{stream})
}

type MessageRPC_MessageToServerServer interface {
	Send(*Message) error
	Recv() (*Message, error)
	grpc.ServerStream
}

type messageRPCMessageToServerServer struct {
	grpc.ServerStream
}

func (x *messageRPCMessageToServerServer) Send(m *Message) error {
	return x.ServerStream.SendMsg(m)
}

func (x *messageRPCMessageToServerServer) Recv() (*Message, error) {
	m := new(Message)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _MessageRPC_MessageToClient_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(MessageRPCServer).MessageToClient(&messageRPCMessageToClientServer{stream})
}

type MessageRPC_MessageToClientServer interface {
	Send(*Message) error
	Recv() (*Message, error)
	grpc.ServerStream
}

type messageRPCMessageToClientServer struct {
	grpc.ServerStream
}

func (x *messageRPCMessageToClientServer) Send(m *Message) error {
	return x.ServerStream.SendMsg(m)
}

func (x *messageRPCMessageToClientServer) Recv() (*Message, error) {
	m := new(Message)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _MessageRPC_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.MessageRPC",
	HandlerType: (*MessageRPCServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "MessageTwoDirection",
			Handler:       _MessageRPC_MessageTwoDirection_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "MessageToServer",
			Handler:       _MessageRPC_MessageToServer_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "MessageToClient",
			Handler:       _MessageRPC_MessageToClient_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "message.proto",
}

func (m *Message) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Message) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Message) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Msg) > 0 {
		i -= len(m.Msg)
		copy(dAtA[i:], m.Msg)
		i = encodeVarintMessage(dAtA, i, uint64(len(m.Msg)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintMessage(dAtA []byte, offset int, v uint64) int {
	offset -= sovMessage(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Message) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Msg)
	if l > 0 {
		n += 1 + l + sovMessage(uint64(l))
	}
	return n
}

func sovMessage(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMessage(x uint64) (n int) {
	return sovMessage(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Message) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMessage
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Message: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Message: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Msg", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMessage
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMessage
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Msg = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMessage(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMessage
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
func skipMessage(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMessage
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
					return 0, ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowMessage
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
			if length < 0 {
				return 0, ErrInvalidLengthMessage
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupMessage
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthMessage
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthMessage        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMessage          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupMessage = fmt.Errorf("proto: unexpected end of group")
)
