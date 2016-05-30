// Code generated by protoc-gen-go.
// source: card.proto
// DO NOT EDIT!

/*
Package card is a generated protocol buffer package.

It is generated from these files:
	card.proto

It has these top-level messages:
	Empty
	Card
	DeleteRequest
	CardList
*/
package card

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.ProtoPackageIsVersion1

// The action to take along with the card
type Card_Action int32

const (
	Card_NONE     Card_Action = 0
	Card_VISITURL Card_Action = 1
)

var Card_Action_name = map[int32]string{
	0: "NONE",
	1: "VISITURL",
}
var Card_Action_value = map[string]int32{
	"NONE":     0,
	"VISITURL": 1,
}

func (x Card_Action) String() string {
	return proto.EnumName(Card_Action_name, int32(x))
}
func (Card_Action) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

type Empty struct {
}

func (m *Empty) Reset()                    { *m = Empty{} }
func (m *Empty) String() string            { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()               {}
func (*Empty) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Card struct {
	// A web link to the image to display
	Image string `protobuf:"bytes,1,opt,name=image" json:"image,omitempty"`
	// The text to display with the card
	Text   string      `protobuf:"bytes,2,opt,name=text" json:"text,omitempty"`
	Action Card_Action `protobuf:"varint,3,opt,name=action,enum=card.Card_Action" json:"action,omitempty"`
	// The date the card was created
	CreatedDate int64 `protobuf:"varint,4,opt,name=created_date,json=createdDate" json:"created_date,omitempty"`
	// The date from which the card should be applied
	ApplicationDate int64 `protobuf:"varint,5,opt,name=application_date,json=applicationDate" json:"application_date,omitempty"`
	// The date from which the card can be removed
	ExpirationDate int64 `protobuf:"varint,6,opt,name=expiration_date,json=expirationDate" json:"expiration_date,omitempty"`
	// The chosen priority of the card
	Priority int32 `protobuf:"varint,7,opt,name=priority" json:"priority,omitempty"`
	// A hash id of the card
	Hash string `protobuf:"bytes,8,opt,name=hash" json:"hash,omitempty"`
}

func (m *Card) Reset()                    { *m = Card{} }
func (m *Card) String() string            { return proto.CompactTextString(m) }
func (*Card) ProtoMessage()               {}
func (*Card) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type DeleteRequest struct {
	Hash string `protobuf:"bytes,1,opt,name=hash" json:"hash,omitempty"`
}

func (m *DeleteRequest) Reset()                    { *m = DeleteRequest{} }
func (m *DeleteRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteRequest) ProtoMessage()               {}
func (*DeleteRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type CardList struct {
	Cards []*Card `protobuf:"bytes,1,rep,name=cards" json:"cards,omitempty"`
}

func (m *CardList) Reset()                    { *m = CardList{} }
func (m *CardList) String() string            { return proto.CompactTextString(m) }
func (*CardList) ProtoMessage()               {}
func (*CardList) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *CardList) GetCards() []*Card {
	if m != nil {
		return m.Cards
	}
	return nil
}

func init() {
	proto.RegisterType((*Empty)(nil), "card.Empty")
	proto.RegisterType((*Card)(nil), "card.Card")
	proto.RegisterType((*DeleteRequest)(nil), "card.DeleteRequest")
	proto.RegisterType((*CardList)(nil), "card.CardList")
	proto.RegisterEnum("card.Card_Action", Card_Action_name, Card_Action_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion2

// Client API for CardService service

type CardServiceClient interface {
	// Gets the cards currently held in the system
	GetCards(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*CardList, error)
	// Adds the cards in the list to the system - returns the resulting card list
	AddCards(ctx context.Context, in *CardList, opts ...grpc.CallOption) (*CardList, error)
	// Deletes existing cards
	DeleteCards(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*CardList, error)
}

type cardServiceClient struct {
	cc *grpc.ClientConn
}

func NewCardServiceClient(cc *grpc.ClientConn) CardServiceClient {
	return &cardServiceClient{cc}
}

func (c *cardServiceClient) GetCards(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*CardList, error) {
	out := new(CardList)
	err := grpc.Invoke(ctx, "/card.CardService/GetCards", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cardServiceClient) AddCards(ctx context.Context, in *CardList, opts ...grpc.CallOption) (*CardList, error) {
	out := new(CardList)
	err := grpc.Invoke(ctx, "/card.CardService/AddCards", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cardServiceClient) DeleteCards(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*CardList, error) {
	out := new(CardList)
	err := grpc.Invoke(ctx, "/card.CardService/DeleteCards", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for CardService service

type CardServiceServer interface {
	// Gets the cards currently held in the system
	GetCards(context.Context, *Empty) (*CardList, error)
	// Adds the cards in the list to the system - returns the resulting card list
	AddCards(context.Context, *CardList) (*CardList, error)
	// Deletes existing cards
	DeleteCards(context.Context, *DeleteRequest) (*CardList, error)
}

func RegisterCardServiceServer(s *grpc.Server, srv CardServiceServer) {
	s.RegisterService(&_CardService_serviceDesc, srv)
}

func _CardService_GetCards_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CardServiceServer).GetCards(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/card.CardService/GetCards",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CardServiceServer).GetCards(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _CardService_AddCards_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CardList)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CardServiceServer).AddCards(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/card.CardService/AddCards",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CardServiceServer).AddCards(ctx, req.(*CardList))
	}
	return interceptor(ctx, in, info, handler)
}

func _CardService_DeleteCards_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CardServiceServer).DeleteCards(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/card.CardService/DeleteCards",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CardServiceServer).DeleteCards(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CardService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "card.CardService",
	HandlerType: (*CardServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCards",
			Handler:    _CardService_GetCards_Handler,
		},
		{
			MethodName: "AddCards",
			Handler:    _CardService_AddCards_Handler,
		},
		{
			MethodName: "DeleteCards",
			Handler:    _CardService_DeleteCards_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}

var fileDescriptor0 = []byte{
	// 346 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x6c, 0x92, 0xcd, 0x4e, 0xc2, 0x40,
	0x10, 0xc7, 0x29, 0xb4, 0xa5, 0x4e, 0x11, 0x70, 0xf5, 0xd0, 0x70, 0xc2, 0x7a, 0x10, 0x13, 0xc3,
	0x01, 0x7d, 0x01, 0x22, 0xc4, 0x90, 0x10, 0x4c, 0x16, 0xf5, 0x6a, 0xd6, 0x76, 0x22, 0x9b, 0x80,
	0xad, 0xdb, 0xd5, 0xc0, 0xbb, 0x78, 0xf4, 0x41, 0xdd, 0x0f, 0x02, 0xf5, 0xe3, 0xb2, 0x99, 0xf9,
	0xcf, 0x6f, 0xfe, 0xdd, 0x99, 0x2e, 0x40, 0xc2, 0x44, 0xda, 0xcf, 0x45, 0x26, 0x33, 0xe2, 0xea,
	0x38, 0xae, 0x83, 0x37, 0x5e, 0xe5, 0x72, 0x13, 0x7f, 0x55, 0xc1, 0xbd, 0x51, 0x0a, 0x39, 0x01,
	0x8f, 0xaf, 0xd8, 0x0b, 0x46, 0x4e, 0xd7, 0xe9, 0x1d, 0x50, 0x9b, 0x10, 0x02, 0xae, 0xc4, 0xb5,
	0x8c, 0xaa, 0x46, 0x34, 0x31, 0xb9, 0x00, 0x9f, 0x25, 0x92, 0x67, 0xaf, 0x51, 0x4d, 0xa9, 0xcd,
	0xc1, 0x51, 0xdf, 0xd8, 0x6b, 0x97, 0xfe, 0xd0, 0x14, 0xe8, 0x16, 0x20, 0xa7, 0xd0, 0x48, 0x04,
	0x32, 0x89, 0xe9, 0x53, 0xaa, 0xce, 0xc8, 0x55, 0x0d, 0x35, 0x1a, 0x6e, 0xb5, 0x91, 0x3a, 0x94,
	0x5b, 0x9b, 0xe5, 0xf9, 0x92, 0x27, 0x4c, 0x77, 0x58, 0xcc, 0x33, 0x58, 0xab, 0xa4, 0x1b, 0xf4,
	0x1c, 0x5a, 0xb8, 0xce, 0xb9, 0x28, 0x91, 0xbe, 0x21, 0x9b, 0x7b, 0xd9, 0x80, 0x1d, 0x08, 0x72,
	0xc1, 0x33, 0xc1, 0xe5, 0x26, 0xaa, 0x2b, 0xc2, 0xa3, 0xbb, 0x5c, 0x4f, 0xb4, 0x60, 0xc5, 0x22,
	0x0a, 0xec, 0x44, 0x3a, 0x8e, 0xbb, 0xe0, 0xdb, 0x8b, 0x93, 0x00, 0xdc, 0xd9, 0xdd, 0x6c, 0xdc,
	0xae, 0x90, 0x06, 0x04, 0x8f, 0x93, 0xf9, 0xe4, 0xfe, 0x81, 0x4e, 0xdb, 0x4e, 0x7c, 0x06, 0x87,
	0x23, 0x5c, 0xa2, 0x44, 0x8a, 0x6f, 0xef, 0x58, 0xc8, 0x9d, 0x8d, 0x53, 0xb2, 0xb9, 0x84, 0x40,
	0x2f, 0x61, 0xca, 0x55, 0xbd, 0x0b, 0x9e, 0xde, 0x4a, 0xa1, 0x80, 0x5a, 0x2f, 0x1c, 0xc0, 0x7e,
	0x47, 0xd4, 0x16, 0x06, 0x9f, 0x0e, 0x84, 0x3a, 0x9f, 0xa3, 0xf8, 0xe0, 0x89, 0x5e, 0x44, 0x70,
	0x8b, 0x52, 0x2b, 0x05, 0x09, 0x2d, 0x6e, 0x7e, 0x51, 0xa7, 0xb9, 0xef, 0xd5, 0xd6, 0x71, 0x85,
	0xa8, 0x0f, 0x0d, 0xd3, 0xd4, 0xa2, 0xbf, 0xaa, 0xff, 0xd0, 0xd7, 0x10, 0xda, 0xbb, 0xdb, 0x86,
	0x63, 0x0b, 0xfc, 0x18, 0xe7, 0x6f, 0xd7, 0xb3, 0x6f, 0x9e, 0xcb, 0xd5, 0x77, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x6b, 0x9d, 0x79, 0x2b, 0x3c, 0x02, 0x00, 0x00,
}
