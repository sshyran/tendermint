// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: proto/privval/types.proto

package privval

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	keys "github.com/tendermint/tendermint/proto/crypto/keys"
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

type Errors int32

const (
	Errors_ERRORS_UNKNOWN             Errors = 0
	Errors_ERRORS_UNEXPECTED_RESPONSE Errors = 1
	Errors_ERRORS_NO_CONNECTION       Errors = 2
	Errors_ERRORS_CONNECTION_TIMEOUT  Errors = 3
	Errors_ERRORS_READ_TIMEOUT        Errors = 4
	Errors_ERRORS_WRITE_TIMEOUT       Errors = 5
)

var Errors_name = map[int32]string{
	0: "ERRORS_UNKNOWN",
	1: "ERRORS_UNEXPECTED_RESPONSE",
	2: "ERRORS_NO_CONNECTION",
	3: "ERRORS_CONNECTION_TIMEOUT",
	4: "ERRORS_READ_TIMEOUT",
	5: "ERRORS_WRITE_TIMEOUT",
}

var Errors_value = map[string]int32{
	"ERRORS_UNKNOWN":             0,
	"ERRORS_UNEXPECTED_RESPONSE": 1,
	"ERRORS_NO_CONNECTION":       2,
	"ERRORS_CONNECTION_TIMEOUT":  3,
	"ERRORS_READ_TIMEOUT":        4,
	"ERRORS_WRITE_TIMEOUT":       5,
}

func (x Errors) String() string {
	return proto.EnumName(Errors_name, int32(x))
}

func (Errors) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_a9d74c406df3ad93, []int{0}
}

// FilePVKey stores the immutable part of PrivValidator.
type FilePVKey struct {
	Address  []byte          `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	PubKey   keys.PublicKey  `protobuf:"bytes,2,opt,name=pub_key,json=pubKey,proto3" json:"pub_key"`
	PrivKey  keys.PrivateKey `protobuf:"bytes,3,opt,name=priv_key,json=privKey,proto3" json:"priv_key"`
	FilePath string          `protobuf:"bytes,4,opt,name=file_path,json=filePath,proto3" json:"file_path,omitempty"`
}

func (m *FilePVKey) Reset()         { *m = FilePVKey{} }
func (m *FilePVKey) String() string { return proto.CompactTextString(m) }
func (*FilePVKey) ProtoMessage()    {}
func (*FilePVKey) Descriptor() ([]byte, []int) {
	return fileDescriptor_a9d74c406df3ad93, []int{0}
}
func (m *FilePVKey) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *FilePVKey) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_FilePVKey.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *FilePVKey) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FilePVKey.Merge(m, src)
}
func (m *FilePVKey) XXX_Size() int {
	return m.Size()
}
func (m *FilePVKey) XXX_DiscardUnknown() {
	xxx_messageInfo_FilePVKey.DiscardUnknown(m)
}

var xxx_messageInfo_FilePVKey proto.InternalMessageInfo

func (m *FilePVKey) GetAddress() []byte {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *FilePVKey) GetPubKey() keys.PublicKey {
	if m != nil {
		return m.PubKey
	}
	return keys.PublicKey{}
}

func (m *FilePVKey) GetPrivKey() keys.PrivateKey {
	if m != nil {
		return m.PrivKey
	}
	return keys.PrivateKey{}
}

func (m *FilePVKey) GetFilePath() string {
	if m != nil {
		return m.FilePath
	}
	return ""
}

// FilePVLastSignState stores the mutable part of PrivValidator.
type FilePVLastSignState struct {
	Height    int64  `protobuf:"varint,1,opt,name=height,proto3" json:"height,omitempty"`
	Round     int32  `protobuf:"varint,2,opt,name=round,proto3" json:"round,omitempty"`
	Step      int32  `protobuf:"varint,3,opt,name=step,proto3" json:"step,omitempty"`
	Signature []byte `protobuf:"bytes,4,opt,name=signature,proto3" json:"signature,omitempty"`
	SignBytes []byte `protobuf:"bytes,5,opt,name=sign_bytes,json=signBytes,proto3" json:"sign_bytes,omitempty"`
	FilePath  string `protobuf:"bytes,6,opt,name=file_path,json=filePath,proto3" json:"file_path,omitempty"`
}

func (m *FilePVLastSignState) Reset()         { *m = FilePVLastSignState{} }
func (m *FilePVLastSignState) String() string { return proto.CompactTextString(m) }
func (*FilePVLastSignState) ProtoMessage()    {}
func (*FilePVLastSignState) Descriptor() ([]byte, []int) {
	return fileDescriptor_a9d74c406df3ad93, []int{1}
}
func (m *FilePVLastSignState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *FilePVLastSignState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_FilePVLastSignState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *FilePVLastSignState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FilePVLastSignState.Merge(m, src)
}
func (m *FilePVLastSignState) XXX_Size() int {
	return m.Size()
}
func (m *FilePVLastSignState) XXX_DiscardUnknown() {
	xxx_messageInfo_FilePVLastSignState.DiscardUnknown(m)
}

var xxx_messageInfo_FilePVLastSignState proto.InternalMessageInfo

func (m *FilePVLastSignState) GetHeight() int64 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *FilePVLastSignState) GetRound() int32 {
	if m != nil {
		return m.Round
	}
	return 0
}

func (m *FilePVLastSignState) GetStep() int32 {
	if m != nil {
		return m.Step
	}
	return 0
}

func (m *FilePVLastSignState) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

func (m *FilePVLastSignState) GetSignBytes() []byte {
	if m != nil {
		return m.SignBytes
	}
	return nil
}

func (m *FilePVLastSignState) GetFilePath() string {
	if m != nil {
		return m.FilePath
	}
	return ""
}

func init() {
	proto.RegisterEnum("tendermint.proto.privval.Errors", Errors_name, Errors_value)
	proto.RegisterType((*FilePVKey)(nil), "tendermint.proto.privval.FilePVKey")
	proto.RegisterType((*FilePVLastSignState)(nil), "tendermint.proto.privval.FilePVLastSignState")
}

func init() { proto.RegisterFile("proto/privval/types.proto", fileDescriptor_a9d74c406df3ad93) }

var fileDescriptor_a9d74c406df3ad93 = []byte{
	// 493 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x52, 0xcf, 0x6e, 0xd3, 0x4e,
	0x18, 0xcc, 0x36, 0x89, 0xd3, 0xec, 0xaf, 0xfa, 0x29, 0xda, 0x56, 0xe0, 0x06, 0x62, 0xa2, 0x1e,
	0x20, 0xe2, 0x60, 0x4b, 0xf0, 0x04, 0x24, 0xdd, 0x8a, 0x28, 0x60, 0x5b, 0x6b, 0x97, 0x22, 0x2e,
	0x96, 0x1d, 0x2f, 0xf6, 0xaa, 0xa9, 0x6d, 0xad, 0xd7, 0x91, 0xfc, 0x16, 0x3c, 0x06, 0x57, 0xde,
	0xa2, 0xc7, 0x5e, 0x90, 0x38, 0x21, 0x94, 0xbc, 0x08, 0xf2, 0x1f, 0xe2, 0x44, 0x1c, 0xb8, 0x7d,
	0x33, 0xf3, 0x79, 0xe6, 0x1b, 0x79, 0xe1, 0x79, 0xc2, 0x63, 0x11, 0x6b, 0x09, 0x67, 0xeb, 0xb5,
	0xbb, 0xd2, 0x44, 0x9e, 0xd0, 0x54, 0x2d, 0x39, 0x24, 0x0b, 0x1a, 0xf9, 0x94, 0xdf, 0xb1, 0x48,
	0x54, 0x8c, 0x5a, 0x6f, 0x0d, 0x9f, 0x8b, 0x90, 0x71, 0xdf, 0x49, 0x5c, 0x2e, 0x72, 0xad, 0x32,
	0x08, 0xe2, 0x20, 0x6e, 0xa6, 0x6a, 0x7f, 0x38, 0xaa, 0x98, 0x25, 0xcf, 0x13, 0x11, 0x6b, 0xb7,
	0x34, 0x4f, 0xf7, 0x03, 0x2e, 0xbe, 0x03, 0xd8, 0xbf, 0x62, 0x2b, 0x6a, 0x7e, 0x58, 0xd0, 0x1c,
	0xc9, 0xb0, 0xe7, 0xfa, 0x3e, 0xa7, 0x69, 0x2a, 0x83, 0x31, 0x98, 0x9c, 0x90, 0x3f, 0x10, 0x5d,
	0xc1, 0x5e, 0x92, 0x79, 0xce, 0x2d, 0xcd, 0xe5, 0xa3, 0x31, 0x98, 0xfc, 0xf7, 0xea, 0x85, 0xfa,
	0xd7, 0x69, 0x55, 0x86, 0x5a, 0x64, 0xa8, 0x66, 0xe6, 0xad, 0xd8, 0x72, 0x41, 0xf3, 0x69, 0xe7,
	0xfe, 0xe7, 0xb3, 0x16, 0x91, 0x92, 0xcc, 0x2b, 0x12, 0xe6, 0xf0, 0xb8, 0x68, 0x50, 0x1a, 0xb5,
	0x4b, 0xa3, 0xc9, 0x3f, 0x8c, 0x38, 0x5b, 0xbb, 0x82, 0x36, 0x4e, 0xbd, 0xe2, 0xfb, 0xc2, 0xea,
	0x09, 0xec, 0x7f, 0x66, 0x2b, 0xea, 0x24, 0xae, 0x08, 0xe5, 0xce, 0x18, 0x4c, 0xfa, 0xe4, 0xb8,
	0x20, 0x4c, 0x57, 0x84, 0x17, 0xdf, 0x00, 0x3c, 0xad, 0x7a, 0xbd, 0x73, 0x53, 0x61, 0xb1, 0x20,
	0xb2, 0x84, 0x2b, 0x28, 0x7a, 0x04, 0xa5, 0x90, 0xb2, 0x20, 0x14, 0x65, 0xc1, 0x36, 0xa9, 0x11,
	0x3a, 0x83, 0x5d, 0x1e, 0x67, 0x91, 0x5f, 0xb6, 0xeb, 0x92, 0x0a, 0x20, 0x04, 0x3b, 0xa9, 0xa0,
	0x49, 0x79, 0x69, 0x97, 0x94, 0x33, 0x7a, 0x0a, 0xfb, 0x29, 0x0b, 0x22, 0x57, 0x64, 0x9c, 0x96,
	0xb1, 0x27, 0xa4, 0x21, 0xd0, 0x08, 0xc2, 0x02, 0x38, 0x5e, 0x2e, 0x68, 0x2a, 0x77, 0x1b, 0x79,
	0x5a, 0x10, 0x87, 0x37, 0x4b, 0x87, 0x37, 0xbf, 0xfc, 0x0a, 0xa0, 0x84, 0x39, 0x8f, 0x79, 0x8a,
	0x10, 0xfc, 0x1f, 0x13, 0x62, 0x10, 0xcb, 0xb9, 0xd6, 0x17, 0xba, 0x71, 0xa3, 0x0f, 0x5a, 0x48,
	0x81, 0xc3, 0x1d, 0x87, 0x3f, 0x9a, 0x78, 0x66, 0xe3, 0x4b, 0x87, 0x60, 0xcb, 0x34, 0x74, 0x0b,
	0x0f, 0x00, 0x92, 0xe1, 0x59, 0xad, 0xeb, 0x86, 0x33, 0x33, 0x74, 0x1d, 0xcf, 0xec, 0xb9, 0xa1,
	0x0f, 0x8e, 0xd0, 0x08, 0x9e, 0xd7, 0x4a, 0x43, 0x3b, 0xf6, 0xfc, 0x3d, 0x36, 0xae, 0xed, 0x41,
	0x1b, 0x3d, 0x86, 0xa7, 0xb5, 0x4c, 0xf0, 0x9b, 0xcb, 0x9d, 0xd0, 0xd9, 0x73, 0xbc, 0x21, 0x73,
	0x1b, 0xef, 0x94, 0xee, 0xf4, 0xed, 0xfd, 0x46, 0x01, 0x0f, 0x1b, 0x05, 0xfc, 0xda, 0x28, 0xe0,
	0xcb, 0x56, 0x69, 0x3d, 0x6c, 0x95, 0xd6, 0x8f, 0xad, 0xd2, 0xfa, 0xa4, 0x06, 0x4c, 0x84, 0x99,
	0xa7, 0x2e, 0xe3, 0x3b, 0xad, 0xf9, 0xb1, 0xfb, 0xe3, 0xc1, 0x6b, 0xf7, 0xa4, 0x12, 0xbe, 0xfe,
	0x1d, 0x00, 0x00, 0xff, 0xff, 0x60, 0xba, 0xbc, 0x27, 0x05, 0x03, 0x00, 0x00,
}

func (m *FilePVKey) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *FilePVKey) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *FilePVKey) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.FilePath) > 0 {
		i -= len(m.FilePath)
		copy(dAtA[i:], m.FilePath)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.FilePath)))
		i--
		dAtA[i] = 0x22
	}
	{
		size, err := m.PrivKey.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTypes(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size, err := m.PubKey.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTypes(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *FilePVLastSignState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *FilePVLastSignState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *FilePVLastSignState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.FilePath) > 0 {
		i -= len(m.FilePath)
		copy(dAtA[i:], m.FilePath)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.FilePath)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.SignBytes) > 0 {
		i -= len(m.SignBytes)
		copy(dAtA[i:], m.SignBytes)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.SignBytes)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Signature) > 0 {
		i -= len(m.Signature)
		copy(dAtA[i:], m.Signature)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Signature)))
		i--
		dAtA[i] = 0x22
	}
	if m.Step != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.Step))
		i--
		dAtA[i] = 0x18
	}
	if m.Round != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.Round))
		i--
		dAtA[i] = 0x10
	}
	if m.Height != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.Height))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintTypes(dAtA []byte, offset int, v uint64) int {
	offset -= sovTypes(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *FilePVKey) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = m.PubKey.Size()
	n += 1 + l + sovTypes(uint64(l))
	l = m.PrivKey.Size()
	n += 1 + l + sovTypes(uint64(l))
	l = len(m.FilePath)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}

func (m *FilePVLastSignState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Height != 0 {
		n += 1 + sovTypes(uint64(m.Height))
	}
	if m.Round != 0 {
		n += 1 + sovTypes(uint64(m.Round))
	}
	if m.Step != 0 {
		n += 1 + sovTypes(uint64(m.Step))
	}
	l = len(m.Signature)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = len(m.SignBytes)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = len(m.FilePath)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}

func sovTypes(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTypes(x uint64) (n int) {
	return sovTypes(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *FilePVKey) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: FilePVKey: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: FilePVKey: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = append(m.Address[:0], dAtA[iNdEx:postIndex]...)
			if m.Address == nil {
				m.Address = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PubKey", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.PubKey.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PrivKey", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.PrivKey.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FilePath", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FilePath = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTypes
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
func (m *FilePVLastSignState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: FilePVLastSignState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: FilePVLastSignState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Height", wireType)
			}
			m.Height = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Height |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Round", wireType)
			}
			m.Round = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Round |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Step", wireType)
			}
			m.Step = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Step |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Signature", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Signature = append(m.Signature[:0], dAtA[iNdEx:postIndex]...)
			if m.Signature == nil {
				m.Signature = []byte{}
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SignBytes", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SignBytes = append(m.SignBytes[:0], dAtA[iNdEx:postIndex]...)
			if m.SignBytes == nil {
				m.SignBytes = []byte{}
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FilePath", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FilePath = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTypes
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
func skipTypes(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTypes
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
					return 0, ErrIntOverflowTypes
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
					return 0, ErrIntOverflowTypes
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
				return 0, ErrInvalidLengthTypes
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTypes
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTypes
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTypes        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTypes          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTypes = fmt.Errorf("proto: unexpected end of group")
)