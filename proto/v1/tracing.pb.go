// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: proto/v1/tracing.proto

package grpc_kit_api_proto_v1

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	golang_proto "github.com/golang/protobuf/proto"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = golang_proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type TracingRequest struct {
	// Id 对应 github.com/grpc-kit/cfg.HTTPHeaderRequestID 请求头内容
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (m *TracingRequest) Reset()         { *m = TracingRequest{} }
func (m *TracingRequest) String() string { return proto.CompactTextString(m) }
func (*TracingRequest) ProtoMessage()    {}
func (*TracingRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_38b784e272bf6054, []int{0}
}
func (m *TracingRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TracingRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TracingRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TracingRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TracingRequest.Merge(m, src)
}
func (m *TracingRequest) XXX_Size() int {
	return m.Size()
}
func (m *TracingRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TracingRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TracingRequest proto.InternalMessageInfo

func (m *TracingRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func init() {
	proto.RegisterType((*TracingRequest)(nil), "grpc.kit.api.proto.v1.TracingRequest")
	golang_proto.RegisterType((*TracingRequest)(nil), "grpc.kit.api.proto.v1.TracingRequest")
}

func init() { proto.RegisterFile("proto/v1/tracing.proto", fileDescriptor_38b784e272bf6054) }
func init() { golang_proto.RegisterFile("proto/v1/tracing.proto", fileDescriptor_38b784e272bf6054) }

var fileDescriptor_38b784e272bf6054 = []byte{
	// 249 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x8f, 0x31, 0x4e, 0x03, 0x31,
	0x10, 0x45, 0xd7, 0x29, 0x90, 0xd8, 0x22, 0x45, 0x24, 0x10, 0x4a, 0x31, 0x8a, 0xa8, 0x28, 0x58,
	0x5b, 0x11, 0x35, 0x0d, 0x47, 0x08, 0x5c, 0xc0, 0xeb, 0x98, 0xc1, 0x02, 0x3c, 0xc6, 0xf6, 0x26,
	0xca, 0x2d, 0x38, 0x46, 0x8e, 0x41, 0xb9, 0x65, 0x4a, 0x4a, 0x58, 0x5f, 0x04, 0x61, 0xa7, 0x00,
	0x6d, 0xf7, 0xff, 0x9f, 0xf9, 0x4f, 0x33, 0xf5, 0xb9, 0xf3, 0x14, 0x49, 0x6c, 0x96, 0x22, 0x7a,
	0xa9, 0x8c, 0x45, 0x9e, 0x83, 0xd9, 0x19, 0x7a, 0xa7, 0xf8, 0xb3, 0x89, 0x5c, 0x3a, 0x53, 0x32,
	0xbe, 0x59, 0xce, 0x1b, 0x34, 0xf1, 0xa9, 0x6b, 0xb9, 0xa2, 0x57, 0x81, 0x84, 0x24, 0xf2, 0xa4,
	0xed, 0x1e, 0xb3, 0x2b, 0xac, 0x5f, 0x55, 0x1a, 0xf3, 0xdb, 0x7f, 0xeb, 0x84, 0x2f, 0x5a, 0x3a,
	0x13, 0xc6, 0x52, 0x48, 0x67, 0x84, 0xb4, 0x96, 0xa2, 0x8c, 0x86, 0x6c, 0x38, 0xd6, 0xef, 0xff,
	0xd6, 0xbd, 0x53, 0x8d, 0x56, 0x14, 0x76, 0x21, 0xea, 0xa3, 0x45, 0x19, 0xf5, 0x56, 0xee, 0xca,
	0x11, 0xaa, 0x41, 0x6d, 0x9b, 0xb0, 0x95, 0x88, 0xda, 0x0b, 0x72, 0x19, 0x34, 0x86, 0x5e, 0x2e,
	0xea, 0xe9, 0x43, 0x79, 0x75, 0xa5, 0xdf, 0x3a, 0x1d, 0xe2, 0x6c, 0x5a, 0x4f, 0xcc, 0xfa, 0x82,
	0x2d, 0xd8, 0xd5, 0xe9, 0x6a, 0x62, 0xd6, 0x77, 0xd7, 0xfd, 0x37, 0xb0, 0xfd, 0x00, 0xac, 0x1f,
	0x80, 0x1d, 0x06, 0x60, 0x5f, 0x03, 0xb0, 0xf7, 0x04, 0xd5, 0x3e, 0x01, 0xfb, 0x48, 0xc0, 0xfa,
	0x04, 0xd5, 0x21, 0x41, 0xf5, 0x99, 0xa0, 0x6a, 0x4f, 0x32, 0xf6, 0xe6, 0x27, 0x00, 0x00, 0xff,
	0xff, 0x15, 0xf2, 0xce, 0xe1, 0x4a, 0x01, 0x00, 0x00,
}

func (this *TracingRequest) Compare(that interface{}) int {
	if that == nil {
		if this == nil {
			return 0
		}
		return 1
	}

	that1, ok := that.(*TracingRequest)
	if !ok {
		that2, ok := that.(TracingRequest)
		if ok {
			that1 = &that2
		} else {
			return 1
		}
	}
	if that1 == nil {
		if this == nil {
			return 0
		}
		return 1
	} else if this == nil {
		return -1
	}
	if this.Id != that1.Id {
		if this.Id < that1.Id {
			return -1
		}
		return 1
	}
	return 0
}
func (this *TracingRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*TracingRequest)
	if !ok {
		that2, ok := that.(TracingRequest)
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
	return true
}
func (m *TracingRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TracingRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TracingRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintTracing(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintTracing(dAtA []byte, offset int, v uint64) int {
	offset -= sovTracing(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *TracingRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovTracing(uint64(l))
	}
	return n
}

func sovTracing(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTracing(x uint64) (n int) {
	return sovTracing(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *TracingRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTracing
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
			return fmt.Errorf("proto: TracingRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TracingRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTracing
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
				return ErrInvalidLengthTracing
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTracing
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTracing(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTracing
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTracing
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
func skipTracing(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTracing
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
					return 0, ErrIntOverflowTracing
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
					return 0, ErrIntOverflowTracing
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
				return 0, ErrInvalidLengthTracing
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTracing
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTracing
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTracing        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTracing          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTracing = fmt.Errorf("proto: unexpected end of group")
)
