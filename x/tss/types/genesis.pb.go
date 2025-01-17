// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: tss/v1beta1/genesis.proto

package types

import (
	fmt "fmt"
	exported "github.com/axelarnetwork/axelar-core/x/tss/exported"
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

type GenesisState struct {
	Params            Params               `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	KeyRecoveryInfos  []KeyRecoveryInfo    `protobuf:"bytes,2,rep,name=key_recovery_infos,json=keyRecoveryInfos,proto3" json:"key_recovery_infos"`
	Keys              []exported.Key       `protobuf:"bytes,3,rep,name=keys,proto3" json:"keys"`
	MultisigInfos     []MultisigInfo       `protobuf:"bytes,4,rep,name=multisig_infos,json=multisigInfos,proto3" json:"multisig_infos"`
	ExternalKeys      []ExternalKeys       `protobuf:"bytes,5,rep,name=external_keys,json=externalKeys,proto3" json:"external_keys"`
	Signatures        []exported.Signature `protobuf:"bytes,6,rep,name=signatures,proto3" json:"signatures"`
	ValidatorStatuses []ValidatorStatus    `protobuf:"bytes,7,rep,name=validator_statuses,json=validatorStatuses,proto3" json:"validator_statuses"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_eb5f1c2be1950e47, []int{0}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func init() {
	proto.RegisterType((*GenesisState)(nil), "tss.v1beta1.GenesisState")
}

func init() { proto.RegisterFile("tss/v1beta1/genesis.proto", fileDescriptor_eb5f1c2be1950e47) }

var fileDescriptor_eb5f1c2be1950e47 = []byte{
	// 408 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0xcf, 0x6e, 0xd3, 0x40,
	0x10, 0x87, 0x6d, 0x12, 0x82, 0xb4, 0x69, 0x11, 0x2c, 0x48, 0xb8, 0x11, 0x72, 0x23, 0x4e, 0xbd,
	0x60, 0x93, 0xf6, 0x0d, 0x2a, 0x0a, 0x42, 0x15, 0xa8, 0x34, 0x12, 0x07, 0x2e, 0xd6, 0x26, 0x9d,
	0x1a, 0xcb, 0x7f, 0xd6, 0xda, 0x19, 0x1b, 0xfb, 0x2d, 0x78, 0xac, 0x1c, 0x73, 0xe4, 0x84, 0x20,
	0x79, 0x01, 0x1e, 0xa1, 0xca, 0x66, 0x2d, 0x6d, 0x24, 0xdf, 0xec, 0x99, 0xef, 0xf7, 0xed, 0xec,
	0x6a, 0xd8, 0x09, 0x21, 0x86, 0xf5, 0x6c, 0x01, 0x24, 0x66, 0x61, 0x0c, 0x05, 0x60, 0x82, 0x41,
	0xa9, 0x24, 0x49, 0x3e, 0x26, 0xc4, 0xc0, 0xb4, 0x26, 0x2f, 0x63, 0x19, 0x4b, 0x5d, 0x0f, 0x77,
	0x5f, 0x7b, 0x64, 0xe2, 0xd9, 0xe9, 0x52, 0x28, 0x91, 0x9b, 0xf0, 0xe4, 0x95, 0xdd, 0xa1, 0xb6,
	0x84, 0xae, 0x31, 0xdd, 0x35, 0xa0, 0x29, 0xa5, 0x22, 0xb8, 0xeb, 0x23, 0xde, 0xfc, 0x1f, 0xb0,
	0xa3, 0x8f, 0xfb, 0x49, 0xe6, 0x24, 0x08, 0xf8, 0x8c, 0x8d, 0xf6, 0x6e, 0xcf, 0x9d, 0xba, 0x67,
	0xe3, 0xf3, 0x17, 0x81, 0x35, 0x59, 0x70, 0xa3, 0x5b, 0x97, 0xc3, 0xd5, 0x9f, 0x53, 0xe7, 0xd6,
	0x80, 0xfc, 0x86, 0xf1, 0x14, 0xda, 0x48, 0xc1, 0x52, 0xd6, 0xa0, 0xda, 0x28, 0x29, 0xee, 0x25,
	0x7a, 0x8f, 0xa6, 0x83, 0xb3, 0xf1, 0xf9, 0xeb, 0x83, 0xf8, 0x35, 0xb4, 0xb7, 0x86, 0xfa, 0x54,
	0xdc, 0x4b, 0xe3, 0x79, 0x96, 0x1e, 0x96, 0x91, 0x5f, 0xb0, 0x61, 0x0a, 0x2d, 0x7a, 0x03, 0xed,
	0x38, 0xd1, 0x8e, 0xee, 0x1a, 0xb6, 0xcc, 0x08, 0x34, 0xcc, 0x3f, 0xb0, 0xa7, 0x79, 0x95, 0x51,
	0x82, 0x49, 0x6c, 0x46, 0x18, 0x5a, 0xf1, 0x2e, 0xf5, 0xd9, 0x20, 0xd6, 0xf9, 0xc7, 0xb9, 0x55,
	0x43, 0xfe, 0x9e, 0x1d, 0x43, 0x43, 0xa0, 0x0a, 0x91, 0x45, 0x7a, 0x8a, 0xc7, 0x3d, 0x9a, 0x2b,
	0x43, 0x5c, 0x43, 0xdb, 0x3d, 0xc7, 0x11, 0x58, 0x35, 0x7e, 0xc5, 0x18, 0x26, 0x71, 0x21, 0xa8,
	0x52, 0x80, 0xde, 0x48, 0x2b, 0x4e, 0xfb, 0x2f, 0x32, 0xef, 0x38, 0x23, 0xb2, 0x82, 0xfc, 0x2b,
	0xe3, 0xb5, 0xc8, 0x92, 0x3b, 0x41, 0x52, 0x45, 0x48, 0x82, 0x2a, 0x04, 0xf4, 0x9e, 0xf4, 0xbc,
	0xed, 0xb7, 0x0e, 0x9b, 0x6b, 0xca, 0xb8, 0x9e, 0xd7, 0x87, 0x65, 0xc0, 0xcb, 0x2f, 0xab, 0x7f,
	0xbe, 0xb3, 0xda, 0xf8, 0xee, 0x7a, 0xe3, 0xbb, 0x7f, 0x37, 0xbe, 0xfb, 0x6b, 0xeb, 0x3b, 0xeb,
	0xad, 0xef, 0xfc, 0xde, 0xfa, 0xce, 0xf7, 0x77, 0x71, 0x42, 0x3f, 0xaa, 0x45, 0xb0, 0x94, 0x79,
	0x28, 0x1a, 0xc8, 0x84, 0x2a, 0x80, 0x7e, 0x4a, 0x95, 0x9a, 0xbf, 0xb7, 0x4b, 0xa9, 0x20, 0x6c,
	0xc2, 0xdd, 0x66, 0xe9, 0x45, 0x5a, 0x8c, 0xf4, 0x26, 0x5d, 0x3c, 0x04, 0x00, 0x00, 0xff, 0xff,
	0x62, 0x47, 0x56, 0x76, 0xde, 0x02, 0x00, 0x00,
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ValidatorStatuses) > 0 {
		for iNdEx := len(m.ValidatorStatuses) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ValidatorStatuses[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x3a
		}
	}
	if len(m.Signatures) > 0 {
		for iNdEx := len(m.Signatures) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Signatures[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x32
		}
	}
	if len(m.ExternalKeys) > 0 {
		for iNdEx := len(m.ExternalKeys) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ExternalKeys[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x2a
		}
	}
	if len(m.MultisigInfos) > 0 {
		for iNdEx := len(m.MultisigInfos) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.MultisigInfos[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.Keys) > 0 {
		for iNdEx := len(m.Keys) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Keys[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.KeyRecoveryInfos) > 0 {
		for iNdEx := len(m.KeyRecoveryInfos) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.KeyRecoveryInfos[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if len(m.KeyRecoveryInfos) > 0 {
		for _, e := range m.KeyRecoveryInfos {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.Keys) > 0 {
		for _, e := range m.Keys {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.MultisigInfos) > 0 {
		for _, e := range m.MultisigInfos {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.ExternalKeys) > 0 {
		for _, e := range m.ExternalKeys {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.Signatures) > 0 {
		for _, e := range m.Signatures {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.ValidatorStatuses) > 0 {
		for _, e := range m.ValidatorStatuses {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field KeyRecoveryInfos", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.KeyRecoveryInfos = append(m.KeyRecoveryInfos, KeyRecoveryInfo{})
			if err := m.KeyRecoveryInfos[len(m.KeyRecoveryInfos)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Keys", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Keys = append(m.Keys, exported.Key{})
			if err := m.Keys[len(m.Keys)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MultisigInfos", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MultisigInfos = append(m.MultisigInfos, MultisigInfo{})
			if err := m.MultisigInfos[len(m.MultisigInfos)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExternalKeys", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ExternalKeys = append(m.ExternalKeys, ExternalKeys{})
			if err := m.ExternalKeys[len(m.ExternalKeys)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Signatures", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Signatures = append(m.Signatures, exported.Signature{})
			if err := m.Signatures[len(m.Signatures)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ValidatorStatuses", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ValidatorStatuses = append(m.ValidatorStatuses, ValidatorStatus{})
			if err := m.ValidatorStatuses[len(m.ValidatorStatuses)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)
