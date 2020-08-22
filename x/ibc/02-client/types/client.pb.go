// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ibc/client/client.proto

package types

import (
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/codec/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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

// IdentifiedClientState defines a client state with additional client identifier field.
type IdentifiedClientState struct {
	// client identifier
	ClientId string `protobuf:"bytes,1,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty" yaml:"client_id"`
	// client state
	ClientState *types.Any `protobuf:"bytes,2,opt,name=client_state,json=clientState,proto3" json:"client_state,omitempty" yaml:"client_state"`
}

func (m *IdentifiedClientState) Reset()         { *m = IdentifiedClientState{} }
func (m *IdentifiedClientState) String() string { return proto.CompactTextString(m) }
func (*IdentifiedClientState) ProtoMessage()    {}
func (*IdentifiedClientState) Descriptor() ([]byte, []int) {
	return fileDescriptor_226f80e576f20abd, []int{0}
}
func (m *IdentifiedClientState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *IdentifiedClientState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_IdentifiedClientState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *IdentifiedClientState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IdentifiedClientState.Merge(m, src)
}
func (m *IdentifiedClientState) XXX_Size() int {
	return m.Size()
}
func (m *IdentifiedClientState) XXX_DiscardUnknown() {
	xxx_messageInfo_IdentifiedClientState.DiscardUnknown(m)
}

var xxx_messageInfo_IdentifiedClientState proto.InternalMessageInfo

func (m *IdentifiedClientState) GetClientId() string {
	if m != nil {
		return m.ClientId
	}
	return ""
}

func (m *IdentifiedClientState) GetClientState() *types.Any {
	if m != nil {
		return m.ClientState
	}
	return nil
}

// ClientConsensusStates defines all the stored consensus states for a given client.
type ClientConsensusStates struct {
	// client identifier
	ClientId string `protobuf:"bytes,1,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty" yaml:"client_id"`
	// consensus states associated with the client
	ConsensusStates []*types.Any `protobuf:"bytes,2,rep,name=consensus_states,json=consensusStates,proto3" json:"consensus_states,omitempty" yaml:"consensus_states"`
}

func (m *ClientConsensusStates) Reset()         { *m = ClientConsensusStates{} }
func (m *ClientConsensusStates) String() string { return proto.CompactTextString(m) }
func (*ClientConsensusStates) ProtoMessage()    {}
func (*ClientConsensusStates) Descriptor() ([]byte, []int) {
	return fileDescriptor_226f80e576f20abd, []int{1}
}
func (m *ClientConsensusStates) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ClientConsensusStates) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ClientConsensusStates.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ClientConsensusStates) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClientConsensusStates.Merge(m, src)
}
func (m *ClientConsensusStates) XXX_Size() int {
	return m.Size()
}
func (m *ClientConsensusStates) XXX_DiscardUnknown() {
	xxx_messageInfo_ClientConsensusStates.DiscardUnknown(m)
}

var xxx_messageInfo_ClientConsensusStates proto.InternalMessageInfo

func (m *ClientConsensusStates) GetClientId() string {
	if m != nil {
		return m.ClientId
	}
	return ""
}

func (m *ClientConsensusStates) GetConsensusStates() []*types.Any {
	if m != nil {
		return m.ConsensusStates
	}
	return nil
}

// ClientUpdateProposal is a governance proposal.
// If it passes, the client is updated with the provided header
type ClientUpdateProposal struct {
	Title       string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty" yaml:"title"`
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty" yaml:"description"`
	ClientId    string `protobuf:"bytes,3,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty" yaml:"client_id"`
	Header      []byte `protobuf:"bytes,4,opt,name=header,proto3" json:"header,omitempty" yaml:"header"`
}

func (m *ClientUpdateProposal) Reset()      { *m = ClientUpdateProposal{} }
func (*ClientUpdateProposal) ProtoMessage() {}
func (*ClientUpdateProposal) Descriptor() ([]byte, []int) {
	return fileDescriptor_226f80e576f20abd, []int{2}
}
func (m *ClientUpdateProposal) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ClientUpdateProposal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ClientUpdateProposal.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ClientUpdateProposal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClientUpdateProposal.Merge(m, src)
}
func (m *ClientUpdateProposal) XXX_Size() int {
	return m.Size()
}
func (m *ClientUpdateProposal) XXX_DiscardUnknown() {
	xxx_messageInfo_ClientUpdateProposal.DiscardUnknown(m)
}

var xxx_messageInfo_ClientUpdateProposal proto.InternalMessageInfo

func init() {
	proto.RegisterType((*IdentifiedClientState)(nil), "ibc.client.IdentifiedClientState")
	proto.RegisterType((*ClientConsensusStates)(nil), "ibc.client.ClientConsensusStates")
	proto.RegisterType((*ClientUpdateProposal)(nil), "ibc.client.ClientUpdateProposal")
}

func init() { proto.RegisterFile("ibc/client/client.proto", fileDescriptor_226f80e576f20abd) }

var fileDescriptor_226f80e576f20abd = []byte{
	// 429 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x52, 0x31, 0x8f, 0xd3, 0x30,
	0x18, 0x8d, 0xef, 0xe0, 0x74, 0x75, 0x8b, 0x28, 0x21, 0x47, 0xc3, 0x21, 0xc5, 0x91, 0x07, 0x54,
	0x86, 0x4b, 0xa0, 0x2c, 0xe8, 0x36, 0x72, 0xd3, 0x49, 0x0c, 0xa7, 0x20, 0x16, 0x84, 0x84, 0x12,
	0xdb, 0x97, 0xb3, 0x48, 0xe3, 0x28, 0x76, 0x25, 0xf2, 0x0f, 0x18, 0x19, 0x19, 0x18, 0x3a, 0xf2,
	0x53, 0x18, 0x3b, 0x32, 0x45, 0x55, 0xfb, 0x0f, 0xf2, 0x0b, 0x50, 0xed, 0x14, 0x5a, 0x24, 0x24,
	0xc4, 0x64, 0xfb, 0xbd, 0xe7, 0xf7, 0x3d, 0x3d, 0x7d, 0x70, 0xc4, 0x53, 0x12, 0x92, 0x9c, 0xb3,
	0x42, 0x75, 0x47, 0x50, 0x56, 0x42, 0x09, 0x1b, 0xf2, 0x94, 0x04, 0x06, 0x39, 0x75, 0x32, 0x91,
	0x09, 0x0d, 0x87, 0x9b, 0x9b, 0x51, 0x9c, 0x3e, 0xcc, 0x84, 0xc8, 0x72, 0x16, 0xea, 0x57, 0x3a,
	0xbb, 0x0e, 0x93, 0xa2, 0x36, 0x14, 0xfe, 0x0a, 0xe0, 0xc9, 0x25, 0x65, 0x85, 0xe2, 0xd7, 0x9c,
	0xd1, 0x0b, 0xed, 0xf2, 0x5a, 0x25, 0x8a, 0xd9, 0xcf, 0x60, 0xcf, 0x98, 0xbe, 0xe7, 0xd4, 0x05,
	0x3e, 0x18, 0xf7, 0x22, 0xa7, 0x6d, 0xd0, 0xb0, 0x4e, 0xa6, 0xf9, 0x39, 0xfe, 0x45, 0xe1, 0xf8,
	0xd8, 0xdc, 0x2f, 0xa9, 0x7d, 0x05, 0x07, 0x1d, 0x2e, 0x37, 0x16, 0xee, 0x81, 0x0f, 0xc6, 0xfd,
	0x89, 0x13, 0x98, 0xf1, 0xc1, 0x76, 0x7c, 0xf0, 0xb2, 0xa8, 0xa3, 0x51, 0xdb, 0xa0, 0xfb, 0x7b,
	0x5e, 0xfa, 0x0f, 0x8e, 0xfb, 0xe4, 0x77, 0x08, 0xfc, 0x0d, 0xc0, 0x13, 0x13, 0xea, 0x42, 0x14,
	0x92, 0x15, 0x72, 0x26, 0x35, 0x21, 0xff, 0x27, 0xde, 0x3b, 0x38, 0x24, 0x5b, 0x17, 0x33, 0x4d,
	0xba, 0x07, 0xfe, 0xe1, 0x5f, 0x23, 0x3e, 0x6a, 0x1b, 0x34, 0xea, 0xfc, 0xfe, 0xf8, 0x87, 0xe3,
	0xbb, 0x64, 0x3f, 0x10, 0x5e, 0x02, 0xe8, 0x98, 0xa8, 0x6f, 0x4a, 0x9a, 0x28, 0x76, 0x55, 0x89,
	0x52, 0xc8, 0x24, 0xb7, 0x1f, 0xc3, 0xdb, 0x8a, 0xab, 0x9c, 0x75, 0x29, 0x87, 0x6d, 0x83, 0x06,
	0xc6, 0x55, 0xc3, 0x38, 0x36, 0xb4, 0xfd, 0x02, 0xf6, 0x29, 0x93, 0xa4, 0xe2, 0xa5, 0xe2, 0xa2,
	0xd0, 0xe5, 0xf5, 0xa2, 0x07, 0x6d, 0x83, 0x6c, 0xa3, 0xde, 0x21, 0x71, 0xbc, 0x2b, 0xdd, 0xef,
	0xe2, 0xf0, 0x9f, 0xba, 0x78, 0x02, 0x8f, 0x6e, 0x58, 0x42, 0x59, 0xe5, 0xde, 0xf2, 0xc1, 0x78,
	0x10, 0xdd, 0x6b, 0x1b, 0x74, 0xc7, 0xe8, 0x0d, 0x8e, 0xe3, 0x4e, 0x70, 0x7e, 0xfc, 0x69, 0x8e,
	0xac, 0x2f, 0x73, 0x64, 0x45, 0xaf, 0xbe, 0xaf, 0x3c, 0xb0, 0x58, 0x79, 0x60, 0xb9, 0xf2, 0xc0,
	0xe7, 0xb5, 0x67, 0x2d, 0xd6, 0x9e, 0xf5, 0x63, 0xed, 0x59, 0x6f, 0x27, 0x19, 0x57, 0x37, 0xb3,
	0x34, 0x20, 0x62, 0x1a, 0x12, 0x21, 0xa7, 0x42, 0x76, 0xc7, 0x99, 0xa4, 0x1f, 0xc2, 0x8f, 0xe1,
	0x66, 0x77, 0x9f, 0x4e, 0xce, 0xba, 0xf5, 0x55, 0x75, 0xc9, 0x64, 0x7a, 0xa4, 0xcb, 0x7e, 0xfe,
	0x33, 0x00, 0x00, 0xff, 0xff, 0x94, 0xc8, 0x07, 0x8a, 0xd9, 0x02, 0x00, 0x00,
}

func (m *IdentifiedClientState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *IdentifiedClientState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *IdentifiedClientState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.ClientState != nil {
		{
			size, err := m.ClientState.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintClient(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.ClientId) > 0 {
		i -= len(m.ClientId)
		copy(dAtA[i:], m.ClientId)
		i = encodeVarintClient(dAtA, i, uint64(len(m.ClientId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ClientConsensusStates) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ClientConsensusStates) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ClientConsensusStates) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ConsensusStates) > 0 {
		for iNdEx := len(m.ConsensusStates) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ConsensusStates[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintClient(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.ClientId) > 0 {
		i -= len(m.ClientId)
		copy(dAtA[i:], m.ClientId)
		i = encodeVarintClient(dAtA, i, uint64(len(m.ClientId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ClientUpdateProposal) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ClientUpdateProposal) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ClientUpdateProposal) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Header) > 0 {
		i -= len(m.Header)
		copy(dAtA[i:], m.Header)
		i = encodeVarintClient(dAtA, i, uint64(len(m.Header)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.ClientId) > 0 {
		i -= len(m.ClientId)
		copy(dAtA[i:], m.ClientId)
		i = encodeVarintClient(dAtA, i, uint64(len(m.ClientId)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintClient(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Title) > 0 {
		i -= len(m.Title)
		copy(dAtA[i:], m.Title)
		i = encodeVarintClient(dAtA, i, uint64(len(m.Title)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintClient(dAtA []byte, offset int, v uint64) int {
	offset -= sovClient(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *IdentifiedClientState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ClientId)
	if l > 0 {
		n += 1 + l + sovClient(uint64(l))
	}
	if m.ClientState != nil {
		l = m.ClientState.Size()
		n += 1 + l + sovClient(uint64(l))
	}
	return n
}

func (m *ClientConsensusStates) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ClientId)
	if l > 0 {
		n += 1 + l + sovClient(uint64(l))
	}
	if len(m.ConsensusStates) > 0 {
		for _, e := range m.ConsensusStates {
			l = e.Size()
			n += 1 + l + sovClient(uint64(l))
		}
	}
	return n
}

func (m *ClientUpdateProposal) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Title)
	if l > 0 {
		n += 1 + l + sovClient(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovClient(uint64(l))
	}
	l = len(m.ClientId)
	if l > 0 {
		n += 1 + l + sovClient(uint64(l))
	}
	l = len(m.Header)
	if l > 0 {
		n += 1 + l + sovClient(uint64(l))
	}
	return n
}

func sovClient(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozClient(x uint64) (n int) {
	return sovClient(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *IdentifiedClientState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowClient
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
			return fmt.Errorf("proto: IdentifiedClientState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: IdentifiedClientState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClientId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClient
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
				return ErrInvalidLengthClient
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthClient
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ClientId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClientState", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClient
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthClient
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthClient
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ClientState == nil {
				m.ClientState = &types.Any{}
			}
			if err := m.ClientState.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipClient(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthClient
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthClient
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
func (m *ClientConsensusStates) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowClient
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
			return fmt.Errorf("proto: ClientConsensusStates: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ClientConsensusStates: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClientId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClient
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
				return ErrInvalidLengthClient
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthClient
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ClientId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ConsensusStates", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClient
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthClient
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthClient
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ConsensusStates = append(m.ConsensusStates, &types.Any{})
			if err := m.ConsensusStates[len(m.ConsensusStates)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipClient(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthClient
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthClient
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
func (m *ClientUpdateProposal) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowClient
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
			return fmt.Errorf("proto: ClientUpdateProposal: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ClientUpdateProposal: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Title", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClient
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
				return ErrInvalidLengthClient
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthClient
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Title = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClient
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
				return ErrInvalidLengthClient
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthClient
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClientId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClient
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
				return ErrInvalidLengthClient
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthClient
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ClientId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Header", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClient
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthClient
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthClient
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Header = append(m.Header[:0], dAtA[iNdEx:postIndex]...)
			if m.Header == nil {
				m.Header = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipClient(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthClient
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthClient
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
func skipClient(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowClient
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
					return 0, ErrIntOverflowClient
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
					return 0, ErrIntOverflowClient
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
				return 0, ErrInvalidLengthClient
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupClient
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthClient
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthClient        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowClient          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupClient = fmt.Errorf("proto: unexpected end of group")
)
