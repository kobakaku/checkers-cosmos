// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: checkers/checkers/stored_game.proto

package types

import (
	fmt "fmt"
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

type StoredGame struct {
	Index    string `protobuf:"bytes,1,opt,name=index,proto3" json:"index,omitempty"`
	Board    string `protobuf:"bytes,2,opt,name=board,proto3" json:"board,omitempty"`
	Turn     string `protobuf:"bytes,3,opt,name=turn,proto3" json:"turn,omitempty"`
	Black    string `protobuf:"bytes,4,opt,name=black,proto3" json:"black,omitempty"`
	Red      string `protobuf:"bytes,5,opt,name=red,proto3" json:"red,omitempty"`
	Winner   string `protobuf:"bytes,6,opt,name=winner,proto3" json:"winner,omitempty"`
	Deadline string `protobuf:"bytes,7,opt,name=deadline,proto3" json:"deadline,omitempty"`
}

func (m *StoredGame) Reset()         { *m = StoredGame{} }
func (m *StoredGame) String() string { return proto.CompactTextString(m) }
func (*StoredGame) ProtoMessage()    {}
func (*StoredGame) Descriptor() ([]byte, []int) {
	return fileDescriptor_6a777ebb9b26769b, []int{0}
}
func (m *StoredGame) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *StoredGame) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_StoredGame.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *StoredGame) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StoredGame.Merge(m, src)
}
func (m *StoredGame) XXX_Size() int {
	return m.Size()
}
func (m *StoredGame) XXX_DiscardUnknown() {
	xxx_messageInfo_StoredGame.DiscardUnknown(m)
}

var xxx_messageInfo_StoredGame proto.InternalMessageInfo

func (m *StoredGame) GetIndex() string {
	if m != nil {
		return m.Index
	}
	return ""
}

func (m *StoredGame) GetBoard() string {
	if m != nil {
		return m.Board
	}
	return ""
}

func (m *StoredGame) GetTurn() string {
	if m != nil {
		return m.Turn
	}
	return ""
}

func (m *StoredGame) GetBlack() string {
	if m != nil {
		return m.Black
	}
	return ""
}

func (m *StoredGame) GetRed() string {
	if m != nil {
		return m.Red
	}
	return ""
}

func (m *StoredGame) GetWinner() string {
	if m != nil {
		return m.Winner
	}
	return ""
}

func (m *StoredGame) GetDeadline() string {
	if m != nil {
		return m.Deadline
	}
	return ""
}

func init() {
	proto.RegisterType((*StoredGame)(nil), "alice.checkers.checkers.StoredGame")
}

func init() {
	proto.RegisterFile("checkers/checkers/stored_game.proto", fileDescriptor_6a777ebb9b26769b)
}

var fileDescriptor_6a777ebb9b26769b = []byte{
	// 237 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x44, 0x90, 0x31, 0x4e, 0xc3, 0x30,
	0x14, 0x86, 0x63, 0xda, 0x06, 0x78, 0x13, 0xb2, 0x10, 0x58, 0x0c, 0x16, 0x82, 0x05, 0x31, 0xa4,
	0x03, 0x37, 0x40, 0x48, 0xec, 0xb0, 0xb1, 0x20, 0xc7, 0x7e, 0x6a, 0xad, 0x26, 0x76, 0xe5, 0xb8,
	0xa2, 0xdc, 0x82, 0x63, 0x70, 0x14, 0xc6, 0x8c, 0x8c, 0x28, 0xb9, 0x08, 0xca, 0x4b, 0x9a, 0x6c,
	0xdf, 0xf7, 0xf9, 0x5f, 0xfc, 0xe0, 0x56, 0xaf, 0x51, 0x6f, 0x30, 0x54, 0xcb, 0x11, 0xaa, 0xe8,
	0x03, 0x9a, 0xf7, 0x95, 0x2a, 0x31, 0xdb, 0x06, 0x1f, 0x3d, 0xbf, 0x54, 0x85, 0xd5, 0x98, 0x1d,
	0x16, 0x23, 0xdc, 0x7c, 0x33, 0x80, 0x57, 0x9a, 0x3f, 0xab, 0x12, 0xf9, 0x39, 0x2c, 0xac, 0x33,
	0xb8, 0x17, 0xec, 0x9a, 0xdd, 0x9d, 0xbe, 0xf4, 0xd2, 0xd5, 0xdc, 0xab, 0x60, 0xc4, 0x51, 0x5f,
	0x49, 0x38, 0x87, 0x79, 0xdc, 0x05, 0x27, 0x66, 0x14, 0x89, 0x69, 0x59, 0x28, 0xbd, 0x11, 0xf3,
	0x61, 0xd9, 0x09, 0x3f, 0x83, 0x59, 0x40, 0x23, 0x16, 0xd4, 0x3a, 0xe4, 0x17, 0x90, 0x7e, 0x58,
	0xe7, 0x30, 0x88, 0x94, 0xe2, 0x60, 0xfc, 0x0a, 0x4e, 0x0c, 0x2a, 0x53, 0x58, 0x87, 0xe2, 0x98,
	0x5e, 0x46, 0x7f, 0x7c, 0xfa, 0x69, 0x24, 0xab, 0x1b, 0xc9, 0xfe, 0x1a, 0xc9, 0xbe, 0x5a, 0x99,
	0xd4, 0xad, 0x4c, 0x7e, 0x5b, 0x99, 0xbc, 0xdd, 0xaf, 0x6c, 0x5c, 0xef, 0xf2, 0x4c, 0xfb, 0x72,
	0x49, 0x1f, 0x9d, 0x4e, 0xb1, 0x9f, 0x30, 0x7e, 0x6e, 0xb1, 0xca, 0x53, 0x3a, 0xc8, 0xc3, 0x7f,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x30, 0x25, 0xc9, 0x16, 0x37, 0x01, 0x00, 0x00,
}

func (m *StoredGame) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *StoredGame) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *StoredGame) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Deadline) > 0 {
		i -= len(m.Deadline)
		copy(dAtA[i:], m.Deadline)
		i = encodeVarintStoredGame(dAtA, i, uint64(len(m.Deadline)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.Winner) > 0 {
		i -= len(m.Winner)
		copy(dAtA[i:], m.Winner)
		i = encodeVarintStoredGame(dAtA, i, uint64(len(m.Winner)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.Red) > 0 {
		i -= len(m.Red)
		copy(dAtA[i:], m.Red)
		i = encodeVarintStoredGame(dAtA, i, uint64(len(m.Red)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Black) > 0 {
		i -= len(m.Black)
		copy(dAtA[i:], m.Black)
		i = encodeVarintStoredGame(dAtA, i, uint64(len(m.Black)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Turn) > 0 {
		i -= len(m.Turn)
		copy(dAtA[i:], m.Turn)
		i = encodeVarintStoredGame(dAtA, i, uint64(len(m.Turn)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Board) > 0 {
		i -= len(m.Board)
		copy(dAtA[i:], m.Board)
		i = encodeVarintStoredGame(dAtA, i, uint64(len(m.Board)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Index) > 0 {
		i -= len(m.Index)
		copy(dAtA[i:], m.Index)
		i = encodeVarintStoredGame(dAtA, i, uint64(len(m.Index)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintStoredGame(dAtA []byte, offset int, v uint64) int {
	offset -= sovStoredGame(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *StoredGame) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Index)
	if l > 0 {
		n += 1 + l + sovStoredGame(uint64(l))
	}
	l = len(m.Board)
	if l > 0 {
		n += 1 + l + sovStoredGame(uint64(l))
	}
	l = len(m.Turn)
	if l > 0 {
		n += 1 + l + sovStoredGame(uint64(l))
	}
	l = len(m.Black)
	if l > 0 {
		n += 1 + l + sovStoredGame(uint64(l))
	}
	l = len(m.Red)
	if l > 0 {
		n += 1 + l + sovStoredGame(uint64(l))
	}
	l = len(m.Winner)
	if l > 0 {
		n += 1 + l + sovStoredGame(uint64(l))
	}
	l = len(m.Deadline)
	if l > 0 {
		n += 1 + l + sovStoredGame(uint64(l))
	}
	return n
}

func sovStoredGame(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozStoredGame(x uint64) (n int) {
	return sovStoredGame(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *StoredGame) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStoredGame
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
			return fmt.Errorf("proto: StoredGame: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: StoredGame: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStoredGame
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
				return ErrInvalidLengthStoredGame
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStoredGame
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Index = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Board", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStoredGame
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
				return ErrInvalidLengthStoredGame
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStoredGame
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Board = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Turn", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStoredGame
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
				return ErrInvalidLengthStoredGame
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStoredGame
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Turn = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Black", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStoredGame
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
				return ErrInvalidLengthStoredGame
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStoredGame
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Black = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Red", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStoredGame
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
				return ErrInvalidLengthStoredGame
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStoredGame
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Red = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Winner", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStoredGame
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
				return ErrInvalidLengthStoredGame
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStoredGame
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Winner = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Deadline", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStoredGame
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
				return ErrInvalidLengthStoredGame
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStoredGame
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Deadline = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipStoredGame(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthStoredGame
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
func skipStoredGame(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowStoredGame
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
					return 0, ErrIntOverflowStoredGame
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
					return 0, ErrIntOverflowStoredGame
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
				return 0, ErrInvalidLengthStoredGame
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupStoredGame
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthStoredGame
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthStoredGame        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowStoredGame          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupStoredGame = fmt.Errorf("proto: unexpected end of group")
)
