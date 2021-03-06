// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: envoy/extensions/transport_sockets/tap/v4alpha/tap.proto

package envoy_extensions_transport_sockets_tap_v4alpha

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	v4alpha1 "github.com/datawire/ambassador/pkg/api/envoy/config/core/v4alpha"
	v4alpha "github.com/datawire/ambassador/pkg/api/envoy/extensions/common/tap/v4alpha"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
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

// Configuration for tap transport socket. This wraps another transport socket, providing the
// ability to interpose and record in plain text any traffic that is surfaced to Envoy.
type Tap struct {
	// Common configuration for the tap transport socket.
	CommonConfig *v4alpha.CommonExtensionConfig `protobuf:"bytes,1,opt,name=common_config,json=commonConfig,proto3" json:"common_config,omitempty"`
	// The underlying transport socket being wrapped.
	TransportSocket      *v4alpha1.TransportSocket `protobuf:"bytes,2,opt,name=transport_socket,json=transportSocket,proto3" json:"transport_socket,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *Tap) Reset()         { *m = Tap{} }
func (m *Tap) String() string { return proto.CompactTextString(m) }
func (*Tap) ProtoMessage()    {}
func (*Tap) Descriptor() ([]byte, []int) {
	return fileDescriptor_6a4a69e98ab9f8a2, []int{0}
}
func (m *Tap) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Tap) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Tap.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Tap) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Tap.Merge(m, src)
}
func (m *Tap) XXX_Size() int {
	return m.Size()
}
func (m *Tap) XXX_DiscardUnknown() {
	xxx_messageInfo_Tap.DiscardUnknown(m)
}

var xxx_messageInfo_Tap proto.InternalMessageInfo

func (m *Tap) GetCommonConfig() *v4alpha.CommonExtensionConfig {
	if m != nil {
		return m.CommonConfig
	}
	return nil
}

func (m *Tap) GetTransportSocket() *v4alpha1.TransportSocket {
	if m != nil {
		return m.TransportSocket
	}
	return nil
}

func init() {
	proto.RegisterType((*Tap)(nil), "envoy.extensions.transport_sockets.tap.v4alpha.Tap")
}

func init() {
	proto.RegisterFile("envoy/extensions/transport_sockets/tap/v4alpha/tap.proto", fileDescriptor_6a4a69e98ab9f8a2)
}

var fileDescriptor_6a4a69e98ab9f8a2 = []byte{
	// 344 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0x86, 0xd9, 0x88, 0xa5, 0x44, 0xc5, 0xd2, 0x8b, 0xa5, 0x60, 0x50, 0xf1, 0x20, 0x82, 0xbb,
	0x62, 0x7b, 0x90, 0xe2, 0x29, 0xc5, 0x7b, 0xd1, 0x9c, 0x2d, 0xd3, 0x74, 0xad, 0xc1, 0x76, 0x67,
	0xc9, 0x6e, 0x43, 0x7b, 0xf3, 0xa6, 0xf8, 0x08, 0x3e, 0x8a, 0x77, 0xa1, 0x47, 0x7d, 0x03, 0xe9,
	0x63, 0x78, 0x92, 0xec, 0x26, 0xb5, 0x35, 0x17, 0xbd, 0x2d, 0x33, 0xf3, 0xfd, 0xff, 0xdf, 0xbf,
	0x71, 0xcf, 0xb9, 0x48, 0x70, 0xca, 0xf8, 0x44, 0x73, 0xa1, 0x22, 0x14, 0x8a, 0xe9, 0x18, 0x84,
	0x92, 0x18, 0xeb, 0xae, 0xc2, 0xf0, 0x9e, 0x6b, 0xc5, 0x34, 0x48, 0x96, 0x34, 0x61, 0x28, 0xef,
	0x20, 0x7d, 0x53, 0x19, 0xa3, 0xc6, 0x2a, 0x35, 0x24, 0xfd, 0x21, 0x69, 0x81, 0xa4, 0xe9, 0x75,
	0x46, 0xd6, 0x0f, 0xad, 0x53, 0x88, 0xe2, 0x36, 0x1a, 0xb0, 0x10, 0x63, 0xbe, 0x10, 0xed, 0x81,
	0xe2, 0x56, 0xb5, 0x7e, 0x5a, 0xc8, 0x13, 0xe2, 0x68, 0x84, 0x62, 0x25, 0x84, 0x1d, 0x65, 0xc4,
	0xee, 0xb8, 0x2f, 0x81, 0x81, 0x10, 0xa8, 0x41, 0x1b, 0x42, 0x69, 0xd0, 0x63, 0x95, 0xad, 0xf7,
	0x0b, 0xeb, 0x84, 0xc7, 0xa9, 0x72, 0x24, 0x06, 0xd9, 0xc9, 0x4e, 0x02, 0xc3, 0xa8, 0x0f, 0x9a,
	0xb3, 0xfc, 0x61, 0x17, 0x07, 0x8f, 0x8e, 0xbb, 0x16, 0x80, 0xac, 0x46, 0xee, 0x96, 0xb5, 0xec,
	0xda, 0xf4, 0x35, 0xb2, 0x47, 0x8e, 0x36, 0xce, 0x5a, 0xc5, 0x0a, 0xb2, 0x64, 0x4b, 0xbf, 0x9b,
	0xb6, 0xcd, 0xe8, 0x32, 0xbf, 0x69, 0x1b, 0x05, 0xbf, 0xfc, 0xe5, 0xaf, 0x3f, 0x13, 0xa7, 0x42,
	0xae, 0x36, 0x2d, 0x63, 0xe7, 0xd5, 0xae, 0x5b, 0xf9, 0x5d, 0x63, 0xcd, 0x31, 0x6e, 0xc7, 0x99,
	0x9b, 0x8d, 0x40, 0xd3, 0x02, 0x17, 0x1e, 0x41, 0x8e, 0x5c, 0x1b, 0x62, 0x49, 0x7d, 0x5b, 0xaf,
	0xae, 0x5a, 0xcd, 0x97, 0xb7, 0x27, 0x8f, 0xb9, 0x27, 0x7f, 0xfd, 0xf7, 0x1a, 0x34, 0x00, 0xe9,
	0xdf, 0xcc, 0xe6, 0x1e, 0x79, 0x9f, 0x7b, 0xe4, 0x73, 0xee, 0x91, 0xd7, 0x87, 0xd9, 0x47, 0xc9,
	0xa9, 0x38, 0xee, 0x45, 0x84, 0x36, 0x94, 0x8c, 0x71, 0x32, 0xfd, 0xe7, 0x07, 0xe1, 0x97, 0x03,
	0x90, 0x9d, 0xb4, 0xe7, 0x0e, 0xe9, 0x95, 0x4c, 0xe1, 0x8d, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x60, 0x83, 0xf9, 0x42, 0x8f, 0x02, 0x00, 0x00,
}

func (m *Tap) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Tap) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Tap) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.TransportSocket != nil {
		{
			size, err := m.TransportSocket.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTap(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.CommonConfig != nil {
		{
			size, err := m.CommonConfig.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTap(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintTap(dAtA []byte, offset int, v uint64) int {
	offset -= sovTap(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Tap) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.CommonConfig != nil {
		l = m.CommonConfig.Size()
		n += 1 + l + sovTap(uint64(l))
	}
	if m.TransportSocket != nil {
		l = m.TransportSocket.Size()
		n += 1 + l + sovTap(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovTap(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTap(x uint64) (n int) {
	return sovTap(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Tap) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTap
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
			return fmt.Errorf("proto: Tap: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Tap: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CommonConfig", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTap
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
				return ErrInvalidLengthTap
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTap
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.CommonConfig == nil {
				m.CommonConfig = &v4alpha.CommonExtensionConfig{}
			}
			if err := m.CommonConfig.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TransportSocket", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTap
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
				return ErrInvalidLengthTap
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTap
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.TransportSocket == nil {
				m.TransportSocket = &v4alpha1.TransportSocket{}
			}
			if err := m.TransportSocket.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTap(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTap
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTap
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipTap(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTap
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
					return 0, ErrIntOverflowTap
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
					return 0, ErrIntOverflowTap
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
				return 0, ErrInvalidLengthTap
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTap
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTap
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTap        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTap          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTap = fmt.Errorf("proto: unexpected end of group")
)
