// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: envoy/config/filter/http/transcoder/v2/transcoder.proto

/*
	Package v2 is a generated protocol buffer package.

	It is generated from these files:
		envoy/config/filter/http/transcoder/v2/transcoder.proto

	It has these top-level messages:
		GrpcJsonTranscoder
*/
package v2

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/lyft/protoc-gen-validate/validate"

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

type GrpcJsonTranscoder struct {
	// Types that are valid to be assigned to DescriptorSet:
	//	*GrpcJsonTranscoder_ProtoDescriptor
	//	*GrpcJsonTranscoder_ProtoDescriptorBin
	DescriptorSet isGrpcJsonTranscoder_DescriptorSet `protobuf_oneof:"descriptor_set"`
	// A list of strings that supplies the service names that the
	// transcoder will translate. If the service name doesn't exist in ``proto_descriptor``, Envoy
	// will fail at startup. The ``proto_descriptor`` may contain more services than the service names
	// specified here, but they won't be translated.
	Services []string `protobuf:"bytes,2,rep,name=services" json:"services,omitempty"`
	// Control options for response JSON. These options are passed directly to
	// `JsonPrintOptions <https://developers.google.com/protocol-buffers/docs/reference/cpp/
	// google.protobuf.util.json_util#JsonPrintOptions>`_.
	PrintOptions *GrpcJsonTranscoder_PrintOptions `protobuf:"bytes,3,opt,name=print_options,json=printOptions" json:"print_options,omitempty"`
	// Whether to skip recalculating the route after the
	// outgoing headers have been transformed to the match the upstream gRPC service
	// If set to true, Envoy will determine the cluster to
	// route to based on the route for the incoming request
	// Note: this will require routes to be added to match
	// the incoming HTTP requests rather than the outgoing
	// requests (after transcoding). This means routes that
	// point directly to gRPC services cannot be reused for
	// transcoded requests. Defaults to false.
	MatchIncomingRequestRoute bool `protobuf:"varint,5,opt,name=match_incoming_request_route,json=matchIncomingRequestRoute,proto3" json:"match_incoming_request_route,omitempty"`
}

func (m *GrpcJsonTranscoder) Reset()                    { *m = GrpcJsonTranscoder{} }
func (m *GrpcJsonTranscoder) String() string            { return proto.CompactTextString(m) }
func (*GrpcJsonTranscoder) ProtoMessage()               {}
func (*GrpcJsonTranscoder) Descriptor() ([]byte, []int) { return fileDescriptorTranscoder, []int{0} }

type isGrpcJsonTranscoder_DescriptorSet interface {
	isGrpcJsonTranscoder_DescriptorSet()
	MarshalTo([]byte) (int, error)
	Size() int
}

type GrpcJsonTranscoder_ProtoDescriptor struct {
	ProtoDescriptor string `protobuf:"bytes,1,opt,name=proto_descriptor,json=protoDescriptor,proto3,oneof"`
}
type GrpcJsonTranscoder_ProtoDescriptorBin struct {
	ProtoDescriptorBin []byte `protobuf:"bytes,4,opt,name=proto_descriptor_bin,json=protoDescriptorBin,proto3,oneof"`
}

func (*GrpcJsonTranscoder_ProtoDescriptor) isGrpcJsonTranscoder_DescriptorSet()    {}
func (*GrpcJsonTranscoder_ProtoDescriptorBin) isGrpcJsonTranscoder_DescriptorSet() {}

func (m *GrpcJsonTranscoder) GetDescriptorSet() isGrpcJsonTranscoder_DescriptorSet {
	if m != nil {
		return m.DescriptorSet
	}
	return nil
}

func (m *GrpcJsonTranscoder) GetProtoDescriptor() string {
	if x, ok := m.GetDescriptorSet().(*GrpcJsonTranscoder_ProtoDescriptor); ok {
		return x.ProtoDescriptor
	}
	return ""
}

func (m *GrpcJsonTranscoder) GetProtoDescriptorBin() []byte {
	if x, ok := m.GetDescriptorSet().(*GrpcJsonTranscoder_ProtoDescriptorBin); ok {
		return x.ProtoDescriptorBin
	}
	return nil
}

func (m *GrpcJsonTranscoder) GetServices() []string {
	if m != nil {
		return m.Services
	}
	return nil
}

func (m *GrpcJsonTranscoder) GetPrintOptions() *GrpcJsonTranscoder_PrintOptions {
	if m != nil {
		return m.PrintOptions
	}
	return nil
}

func (m *GrpcJsonTranscoder) GetMatchIncomingRequestRoute() bool {
	if m != nil {
		return m.MatchIncomingRequestRoute
	}
	return false
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*GrpcJsonTranscoder) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _GrpcJsonTranscoder_OneofMarshaler, _GrpcJsonTranscoder_OneofUnmarshaler, _GrpcJsonTranscoder_OneofSizer, []interface{}{
		(*GrpcJsonTranscoder_ProtoDescriptor)(nil),
		(*GrpcJsonTranscoder_ProtoDescriptorBin)(nil),
	}
}

func _GrpcJsonTranscoder_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*GrpcJsonTranscoder)
	// descriptor_set
	switch x := m.DescriptorSet.(type) {
	case *GrpcJsonTranscoder_ProtoDescriptor:
		_ = b.EncodeVarint(1<<3 | proto.WireBytes)
		_ = b.EncodeStringBytes(x.ProtoDescriptor)
	case *GrpcJsonTranscoder_ProtoDescriptorBin:
		_ = b.EncodeVarint(4<<3 | proto.WireBytes)
		_ = b.EncodeRawBytes(x.ProtoDescriptorBin)
	case nil:
	default:
		return fmt.Errorf("GrpcJsonTranscoder.DescriptorSet has unexpected type %T", x)
	}
	return nil
}

func _GrpcJsonTranscoder_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*GrpcJsonTranscoder)
	switch tag {
	case 1: // descriptor_set.proto_descriptor
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.DescriptorSet = &GrpcJsonTranscoder_ProtoDescriptor{x}
		return true, err
	case 4: // descriptor_set.proto_descriptor_bin
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeRawBytes(true)
		m.DescriptorSet = &GrpcJsonTranscoder_ProtoDescriptorBin{x}
		return true, err
	default:
		return false, nil
	}
}

func _GrpcJsonTranscoder_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*GrpcJsonTranscoder)
	// descriptor_set
	switch x := m.DescriptorSet.(type) {
	case *GrpcJsonTranscoder_ProtoDescriptor:
		n += proto.SizeVarint(1<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.ProtoDescriptor)))
		n += len(x.ProtoDescriptor)
	case *GrpcJsonTranscoder_ProtoDescriptorBin:
		n += proto.SizeVarint(4<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.ProtoDescriptorBin)))
		n += len(x.ProtoDescriptorBin)
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type GrpcJsonTranscoder_PrintOptions struct {
	// Whether to add spaces, line breaks and indentation to make the JSON
	// output easy to read. Defaults to false.
	AddWhitespace bool `protobuf:"varint,1,opt,name=add_whitespace,json=addWhitespace,proto3" json:"add_whitespace,omitempty"`
	// Whether to always print primitive fields. By default primitive
	// fields with default values will be omitted in JSON output. For
	// example, an int32 field set to 0 will be omitted. Setting this flag to
	// true will override the default behavior and print primitive fields
	// regardless of their values. Defaults to false.
	AlwaysPrintPrimitiveFields bool `protobuf:"varint,2,opt,name=always_print_primitive_fields,json=alwaysPrintPrimitiveFields,proto3" json:"always_print_primitive_fields,omitempty"`
	// Whether to always print enums as ints. By default they are rendered
	// as strings. Defaults to false.
	AlwaysPrintEnumsAsInts bool `protobuf:"varint,3,opt,name=always_print_enums_as_ints,json=alwaysPrintEnumsAsInts,proto3" json:"always_print_enums_as_ints,omitempty"`
	// Whether to preserve proto field names. By default protobuf will
	// generate JSON field names using the ``json_name`` option, or lower camel case,
	// in that order. Setting this flag will preserve the original field names. Defaults to false.
	PreserveProtoFieldNames bool `protobuf:"varint,4,opt,name=preserve_proto_field_names,json=preserveProtoFieldNames,proto3" json:"preserve_proto_field_names,omitempty"`
}

func (m *GrpcJsonTranscoder_PrintOptions) Reset()         { *m = GrpcJsonTranscoder_PrintOptions{} }
func (m *GrpcJsonTranscoder_PrintOptions) String() string { return proto.CompactTextString(m) }
func (*GrpcJsonTranscoder_PrintOptions) ProtoMessage()    {}
func (*GrpcJsonTranscoder_PrintOptions) Descriptor() ([]byte, []int) {
	return fileDescriptorTranscoder, []int{0, 0}
}

func (m *GrpcJsonTranscoder_PrintOptions) GetAddWhitespace() bool {
	if m != nil {
		return m.AddWhitespace
	}
	return false
}

func (m *GrpcJsonTranscoder_PrintOptions) GetAlwaysPrintPrimitiveFields() bool {
	if m != nil {
		return m.AlwaysPrintPrimitiveFields
	}
	return false
}

func (m *GrpcJsonTranscoder_PrintOptions) GetAlwaysPrintEnumsAsInts() bool {
	if m != nil {
		return m.AlwaysPrintEnumsAsInts
	}
	return false
}

func (m *GrpcJsonTranscoder_PrintOptions) GetPreserveProtoFieldNames() bool {
	if m != nil {
		return m.PreserveProtoFieldNames
	}
	return false
}

func init() {
	proto.RegisterType((*GrpcJsonTranscoder)(nil), "envoy.config.filter.http.transcoder.v2.GrpcJsonTranscoder")
	proto.RegisterType((*GrpcJsonTranscoder_PrintOptions)(nil), "envoy.config.filter.http.transcoder.v2.GrpcJsonTranscoder.PrintOptions")
}
func (m *GrpcJsonTranscoder) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GrpcJsonTranscoder) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.DescriptorSet != nil {
		nn1, err := m.DescriptorSet.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += nn1
	}
	if len(m.Services) > 0 {
		for _, s := range m.Services {
			dAtA[i] = 0x12
			i++
			l = len(s)
			for l >= 1<<7 {
				dAtA[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			dAtA[i] = uint8(l)
			i++
			i += copy(dAtA[i:], s)
		}
	}
	if m.PrintOptions != nil {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintTranscoder(dAtA, i, uint64(m.PrintOptions.Size()))
		n2, err := m.PrintOptions.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	if m.MatchIncomingRequestRoute {
		dAtA[i] = 0x28
		i++
		if m.MatchIncomingRequestRoute {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	return i, nil
}

func (m *GrpcJsonTranscoder_ProtoDescriptor) MarshalTo(dAtA []byte) (int, error) {
	i := 0
	dAtA[i] = 0xa
	i++
	i = encodeVarintTranscoder(dAtA, i, uint64(len(m.ProtoDescriptor)))
	i += copy(dAtA[i:], m.ProtoDescriptor)
	return i, nil
}
func (m *GrpcJsonTranscoder_ProtoDescriptorBin) MarshalTo(dAtA []byte) (int, error) {
	i := 0
	if m.ProtoDescriptorBin != nil {
		dAtA[i] = 0x22
		i++
		i = encodeVarintTranscoder(dAtA, i, uint64(len(m.ProtoDescriptorBin)))
		i += copy(dAtA[i:], m.ProtoDescriptorBin)
	}
	return i, nil
}
func (m *GrpcJsonTranscoder_PrintOptions) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GrpcJsonTranscoder_PrintOptions) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.AddWhitespace {
		dAtA[i] = 0x8
		i++
		if m.AddWhitespace {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.AlwaysPrintPrimitiveFields {
		dAtA[i] = 0x10
		i++
		if m.AlwaysPrintPrimitiveFields {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.AlwaysPrintEnumsAsInts {
		dAtA[i] = 0x18
		i++
		if m.AlwaysPrintEnumsAsInts {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.PreserveProtoFieldNames {
		dAtA[i] = 0x20
		i++
		if m.PreserveProtoFieldNames {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	return i, nil
}

func encodeVarintTranscoder(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *GrpcJsonTranscoder) Size() (n int) {
	var l int
	_ = l
	if m.DescriptorSet != nil {
		n += m.DescriptorSet.Size()
	}
	if len(m.Services) > 0 {
		for _, s := range m.Services {
			l = len(s)
			n += 1 + l + sovTranscoder(uint64(l))
		}
	}
	if m.PrintOptions != nil {
		l = m.PrintOptions.Size()
		n += 1 + l + sovTranscoder(uint64(l))
	}
	if m.MatchIncomingRequestRoute {
		n += 2
	}
	return n
}

func (m *GrpcJsonTranscoder_ProtoDescriptor) Size() (n int) {
	var l int
	_ = l
	l = len(m.ProtoDescriptor)
	n += 1 + l + sovTranscoder(uint64(l))
	return n
}
func (m *GrpcJsonTranscoder_ProtoDescriptorBin) Size() (n int) {
	var l int
	_ = l
	if m.ProtoDescriptorBin != nil {
		l = len(m.ProtoDescriptorBin)
		n += 1 + l + sovTranscoder(uint64(l))
	}
	return n
}
func (m *GrpcJsonTranscoder_PrintOptions) Size() (n int) {
	var l int
	_ = l
	if m.AddWhitespace {
		n += 2
	}
	if m.AlwaysPrintPrimitiveFields {
		n += 2
	}
	if m.AlwaysPrintEnumsAsInts {
		n += 2
	}
	if m.PreserveProtoFieldNames {
		n += 2
	}
	return n
}

func sovTranscoder(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozTranscoder(x uint64) (n int) {
	return sovTranscoder(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GrpcJsonTranscoder) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTranscoder
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
			return fmt.Errorf("proto: GrpcJsonTranscoder: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GrpcJsonTranscoder: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProtoDescriptor", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTranscoder
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
				return ErrInvalidLengthTranscoder
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DescriptorSet = &GrpcJsonTranscoder_ProtoDescriptor{string(dAtA[iNdEx:postIndex])}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Services", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTranscoder
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
				return ErrInvalidLengthTranscoder
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Services = append(m.Services, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PrintOptions", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTranscoder
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
				return ErrInvalidLengthTranscoder
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.PrintOptions == nil {
				m.PrintOptions = &GrpcJsonTranscoder_PrintOptions{}
			}
			if err := m.PrintOptions.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProtoDescriptorBin", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTranscoder
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTranscoder
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := make([]byte, postIndex-iNdEx)
			copy(v, dAtA[iNdEx:postIndex])
			m.DescriptorSet = &GrpcJsonTranscoder_ProtoDescriptorBin{v}
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MatchIncomingRequestRoute", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTranscoder
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.MatchIncomingRequestRoute = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipTranscoder(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTranscoder
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
func (m *GrpcJsonTranscoder_PrintOptions) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTranscoder
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
			return fmt.Errorf("proto: PrintOptions: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PrintOptions: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AddWhitespace", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTranscoder
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.AddWhitespace = bool(v != 0)
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AlwaysPrintPrimitiveFields", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTranscoder
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.AlwaysPrintPrimitiveFields = bool(v != 0)
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AlwaysPrintEnumsAsInts", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTranscoder
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.AlwaysPrintEnumsAsInts = bool(v != 0)
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PreserveProtoFieldNames", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTranscoder
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.PreserveProtoFieldNames = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipTranscoder(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTranscoder
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
func skipTranscoder(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTranscoder
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
					return 0, ErrIntOverflowTranscoder
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
					return 0, ErrIntOverflowTranscoder
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
				return 0, ErrInvalidLengthTranscoder
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowTranscoder
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
				next, err := skipTranscoder(dAtA[start:])
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
	ErrInvalidLengthTranscoder = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTranscoder   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("envoy/config/filter/http/transcoder/v2/transcoder.proto", fileDescriptorTranscoder)
}

var fileDescriptorTranscoder = []byte{
	// 468 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0x41, 0x6b, 0xd4, 0x40,
	0x1c, 0xc5, 0x3b, 0xbb, 0x5d, 0x49, 0xc7, 0x6d, 0x2d, 0x83, 0xb8, 0x6b, 0xd0, 0x25, 0x08, 0x96,
	0x80, 0x90, 0x40, 0x3c, 0x08, 0x7a, 0x90, 0x2e, 0x6a, 0x5b, 0x0f, 0xba, 0x0c, 0x82, 0xe0, 0x65,
	0x98, 0x26, 0xff, 0xed, 0x0e, 0x24, 0x33, 0xe3, 0xcc, 0x6c, 0x4a, 0xbf, 0x86, 0x9f, 0x46, 0x3c,
	0xf5, 0xe8, 0xd1, 0xab, 0x37, 0xd9, 0x5b, 0x8f, 0x7e, 0x03, 0xc9, 0xc4, 0x5d, 0x62, 0xbd, 0xf4,
	0x96, 0xc9, 0xfb, 0xbd, 0xf7, 0x9f, 0xff, 0x1b, 0xfc, 0x0c, 0x64, 0xad, 0x2e, 0xd2, 0x5c, 0xc9,
	0xb9, 0x38, 0x4b, 0xe7, 0xa2, 0x74, 0x60, 0xd2, 0x85, 0x73, 0x3a, 0x75, 0x86, 0x4b, 0x9b, 0xab,
	0x02, 0x4c, 0x5a, 0x67, 0x9d, 0x53, 0xa2, 0x8d, 0x72, 0x8a, 0x1c, 0x78, 0x63, 0xd2, 0x1a, 0x93,
	0xd6, 0x98, 0x34, 0xc6, 0xa4, 0x83, 0xd6, 0x59, 0x38, 0xaa, 0x79, 0x29, 0x0a, 0xee, 0x20, 0x5d,
	0x7f, 0xb4, 0x01, 0x8f, 0x7e, 0x6e, 0x63, 0x72, 0x64, 0x74, 0xfe, 0xd6, 0x2a, 0xf9, 0x61, 0x63,
	0x21, 0x4f, 0xf0, 0xbe, 0xd7, 0x59, 0x01, 0x36, 0x37, 0x42, 0x3b, 0x65, 0xc6, 0x28, 0x42, 0xf1,
	0xce, 0xf1, 0x16, 0xbd, 0xe3, 0x95, 0x57, 0x1b, 0x81, 0x1c, 0xe0, 0xc0, 0x82, 0xa9, 0x45, 0x0e,
	0x76, 0xdc, 0x8b, 0xfa, 0xf1, 0xce, 0x14, 0x7f, 0xbb, 0xba, 0xec, 0x0f, 0xbe, 0xa0, 0x5e, 0x80,
	0xe8, 0x46, 0x23, 0x25, 0xde, 0xd5, 0x46, 0x48, 0xc7, 0x94, 0x76, 0x42, 0x49, 0x3b, 0xee, 0x47,
	0x28, 0xbe, 0x9d, 0x1d, 0x25, 0x37, 0x5b, 0x22, 0xf9, 0xff, 0x9e, 0xc9, 0xac, 0xc9, 0x7b, 0xdf,
	0xc6, 0xd1, 0xa1, 0xee, 0x9c, 0x48, 0x86, 0xef, 0x5e, 0x5f, 0x81, 0x9d, 0x0a, 0x39, 0xde, 0x8e,
	0x50, 0x3c, 0x3c, 0xde, 0xa2, 0xe4, 0xda, 0x1a, 0x53, 0x21, 0xc9, 0x4b, 0xfc, 0xa0, 0xe2, 0x2e,
	0x5f, 0x30, 0x21, 0x73, 0x55, 0x09, 0x79, 0xc6, 0x0c, 0x7c, 0x5e, 0x82, 0x75, 0xcc, 0xa8, 0xa5,
	0x83, 0xf1, 0x20, 0x42, 0x71, 0x40, 0xef, 0x7b, 0xe6, 0xe4, 0x2f, 0x42, 0x5b, 0x82, 0x36, 0x40,
	0xf8, 0x1b, 0xe1, 0x61, 0xf7, 0x4e, 0xe4, 0x31, 0xde, 0xe3, 0x45, 0xc1, 0xce, 0x17, 0xc2, 0x81,
	0xd5, 0x3c, 0x07, 0x5f, 0x63, 0x40, 0x77, 0x79, 0x51, 0x7c, 0xdc, 0xfc, 0x24, 0x87, 0xf8, 0x21,
	0x2f, 0xcf, 0xf9, 0x85, 0x65, 0x6d, 0x43, 0xda, 0x88, 0x4a, 0x38, 0x51, 0x03, 0x9b, 0x0b, 0x28,
	0x8b, 0xa6, 0xd7, 0xc6, 0x15, 0xb6, 0x90, 0x9f, 0x30, 0x5b, 0x23, 0x6f, 0x3c, 0x41, 0x9e, 0xe3,
	0xf0, 0x9f, 0x08, 0x90, 0xcb, 0xca, 0x32, 0x6e, 0x99, 0x90, 0xae, 0xad, 0x3a, 0xa0, 0xf7, 0x3a,
	0xfe, 0xd7, 0x8d, 0x7e, 0x68, 0x4f, 0xa4, 0xb3, 0xe4, 0x05, 0x0e, 0xb5, 0x81, 0xe6, 0xa1, 0x80,
	0xb5, 0xa5, 0xf9, 0xb1, 0x4c, 0xf2, 0x0a, 0xac, 0x6f, 0x2c, 0xa0, 0xa3, 0x35, 0x31, 0x6b, 0x00,
	0x3f, 0xf4, 0x5d, 0x23, 0x4f, 0x47, 0x78, 0xaf, 0x53, 0xb1, 0x05, 0x47, 0x06, 0x5f, 0xaf, 0x2e,
	0xfb, 0x68, 0xba, 0xff, 0x7d, 0x35, 0x41, 0x3f, 0x56, 0x13, 0xf4, 0x6b, 0x35, 0x41, 0x9f, 0x7a,
	0x75, 0x76, 0x7a, 0xcb, 0x87, 0x3f, 0xfd, 0x13, 0x00, 0x00, 0xff, 0xff, 0x58, 0xf4, 0x4a, 0x34,
	0xf0, 0x02, 0x00, 0x00,
}
