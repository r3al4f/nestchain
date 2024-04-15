// Code generated by protoc-gen-go-pulsar. DO NOT EDIT.
package hotels

import (
	v1beta1 "cosmossdk.io/api/cosmos/base/v1beta1"
	fmt "fmt"
	runtime "github.com/cosmos/cosmos-proto/runtime"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoiface "google.golang.org/protobuf/runtime/protoiface"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	io "io"
	reflect "reflect"
	sync "sync"
)

var (
	md_Booking             protoreflect.MessageDescriptor
	fd_Booking_booking_id  protoreflect.FieldDescriptor
	fd_Booking_hotel_id    protoreflect.FieldDescriptor
	fd_Booking_customer_id protoreflect.FieldDescriptor
	fd_Booking_start_date  protoreflect.FieldDescriptor
	fd_Booking_end_date    protoreflect.FieldDescriptor
	fd_Booking_room_type   protoreflect.FieldDescriptor
	fd_Booking_price       protoreflect.FieldDescriptor
)

func init() {
	file_nestchain_hotels_booking_proto_init()
	md_Booking = File_nestchain_hotels_booking_proto.Messages().ByName("Booking")
	fd_Booking_booking_id = md_Booking.Fields().ByName("booking_id")
	fd_Booking_hotel_id = md_Booking.Fields().ByName("hotel_id")
	fd_Booking_customer_id = md_Booking.Fields().ByName("customer_id")
	fd_Booking_start_date = md_Booking.Fields().ByName("start_date")
	fd_Booking_end_date = md_Booking.Fields().ByName("end_date")
	fd_Booking_room_type = md_Booking.Fields().ByName("room_type")
	fd_Booking_price = md_Booking.Fields().ByName("price")
}

var _ protoreflect.Message = (*fastReflection_Booking)(nil)

type fastReflection_Booking Booking

func (x *Booking) ProtoReflect() protoreflect.Message {
	return (*fastReflection_Booking)(x)
}

func (x *Booking) slowProtoReflect() protoreflect.Message {
	mi := &file_nestchain_hotels_booking_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_Booking_messageType fastReflection_Booking_messageType
var _ protoreflect.MessageType = fastReflection_Booking_messageType{}

type fastReflection_Booking_messageType struct{}

func (x fastReflection_Booking_messageType) Zero() protoreflect.Message {
	return (*fastReflection_Booking)(nil)
}
func (x fastReflection_Booking_messageType) New() protoreflect.Message {
	return new(fastReflection_Booking)
}
func (x fastReflection_Booking_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_Booking
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_Booking) Descriptor() protoreflect.MessageDescriptor {
	return md_Booking
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_Booking) Type() protoreflect.MessageType {
	return _fastReflection_Booking_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_Booking) New() protoreflect.Message {
	return new(fastReflection_Booking)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_Booking) Interface() protoreflect.ProtoMessage {
	return (*Booking)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_Booking) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.BookingId != "" {
		value := protoreflect.ValueOfString(x.BookingId)
		if !f(fd_Booking_booking_id, value) {
			return
		}
	}
	if x.HotelId != "" {
		value := protoreflect.ValueOfString(x.HotelId)
		if !f(fd_Booking_hotel_id, value) {
			return
		}
	}
	if x.CustomerId != "" {
		value := protoreflect.ValueOfString(x.CustomerId)
		if !f(fd_Booking_customer_id, value) {
			return
		}
	}
	if x.StartDate != "" {
		value := protoreflect.ValueOfString(x.StartDate)
		if !f(fd_Booking_start_date, value) {
			return
		}
	}
	if x.EndDate != "" {
		value := protoreflect.ValueOfString(x.EndDate)
		if !f(fd_Booking_end_date, value) {
			return
		}
	}
	if x.RoomType != "" {
		value := protoreflect.ValueOfString(x.RoomType)
		if !f(fd_Booking_room_type, value) {
			return
		}
	}
	if x.Price != nil {
		value := protoreflect.ValueOfMessage(x.Price.ProtoReflect())
		if !f(fd_Booking_price, value) {
			return
		}
	}
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_Booking) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "nestchain.hotels.Booking.booking_id":
		return x.BookingId != ""
	case "nestchain.hotels.Booking.hotel_id":
		return x.HotelId != ""
	case "nestchain.hotels.Booking.customer_id":
		return x.CustomerId != ""
	case "nestchain.hotels.Booking.start_date":
		return x.StartDate != ""
	case "nestchain.hotels.Booking.end_date":
		return x.EndDate != ""
	case "nestchain.hotels.Booking.room_type":
		return x.RoomType != ""
	case "nestchain.hotels.Booking.price":
		return x.Price != nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: nestchain.hotels.Booking"))
		}
		panic(fmt.Errorf("message nestchain.hotels.Booking does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Booking) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "nestchain.hotels.Booking.booking_id":
		x.BookingId = ""
	case "nestchain.hotels.Booking.hotel_id":
		x.HotelId = ""
	case "nestchain.hotels.Booking.customer_id":
		x.CustomerId = ""
	case "nestchain.hotels.Booking.start_date":
		x.StartDate = ""
	case "nestchain.hotels.Booking.end_date":
		x.EndDate = ""
	case "nestchain.hotels.Booking.room_type":
		x.RoomType = ""
	case "nestchain.hotels.Booking.price":
		x.Price = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: nestchain.hotels.Booking"))
		}
		panic(fmt.Errorf("message nestchain.hotels.Booking does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_Booking) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "nestchain.hotels.Booking.booking_id":
		value := x.BookingId
		return protoreflect.ValueOfString(value)
	case "nestchain.hotels.Booking.hotel_id":
		value := x.HotelId
		return protoreflect.ValueOfString(value)
	case "nestchain.hotels.Booking.customer_id":
		value := x.CustomerId
		return protoreflect.ValueOfString(value)
	case "nestchain.hotels.Booking.start_date":
		value := x.StartDate
		return protoreflect.ValueOfString(value)
	case "nestchain.hotels.Booking.end_date":
		value := x.EndDate
		return protoreflect.ValueOfString(value)
	case "nestchain.hotels.Booking.room_type":
		value := x.RoomType
		return protoreflect.ValueOfString(value)
	case "nestchain.hotels.Booking.price":
		value := x.Price
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: nestchain.hotels.Booking"))
		}
		panic(fmt.Errorf("message nestchain.hotels.Booking does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Booking) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "nestchain.hotels.Booking.booking_id":
		x.BookingId = value.Interface().(string)
	case "nestchain.hotels.Booking.hotel_id":
		x.HotelId = value.Interface().(string)
	case "nestchain.hotels.Booking.customer_id":
		x.CustomerId = value.Interface().(string)
	case "nestchain.hotels.Booking.start_date":
		x.StartDate = value.Interface().(string)
	case "nestchain.hotels.Booking.end_date":
		x.EndDate = value.Interface().(string)
	case "nestchain.hotels.Booking.room_type":
		x.RoomType = value.Interface().(string)
	case "nestchain.hotels.Booking.price":
		x.Price = value.Message().Interface().(*v1beta1.Coin)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: nestchain.hotels.Booking"))
		}
		panic(fmt.Errorf("message nestchain.hotels.Booking does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Booking) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "nestchain.hotels.Booking.price":
		if x.Price == nil {
			x.Price = new(v1beta1.Coin)
		}
		return protoreflect.ValueOfMessage(x.Price.ProtoReflect())
	case "nestchain.hotels.Booking.booking_id":
		panic(fmt.Errorf("field booking_id of message nestchain.hotels.Booking is not mutable"))
	case "nestchain.hotels.Booking.hotel_id":
		panic(fmt.Errorf("field hotel_id of message nestchain.hotels.Booking is not mutable"))
	case "nestchain.hotels.Booking.customer_id":
		panic(fmt.Errorf("field customer_id of message nestchain.hotels.Booking is not mutable"))
	case "nestchain.hotels.Booking.start_date":
		panic(fmt.Errorf("field start_date of message nestchain.hotels.Booking is not mutable"))
	case "nestchain.hotels.Booking.end_date":
		panic(fmt.Errorf("field end_date of message nestchain.hotels.Booking is not mutable"))
	case "nestchain.hotels.Booking.room_type":
		panic(fmt.Errorf("field room_type of message nestchain.hotels.Booking is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: nestchain.hotels.Booking"))
		}
		panic(fmt.Errorf("message nestchain.hotels.Booking does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_Booking) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "nestchain.hotels.Booking.booking_id":
		return protoreflect.ValueOfString("")
	case "nestchain.hotels.Booking.hotel_id":
		return protoreflect.ValueOfString("")
	case "nestchain.hotels.Booking.customer_id":
		return protoreflect.ValueOfString("")
	case "nestchain.hotels.Booking.start_date":
		return protoreflect.ValueOfString("")
	case "nestchain.hotels.Booking.end_date":
		return protoreflect.ValueOfString("")
	case "nestchain.hotels.Booking.room_type":
		return protoreflect.ValueOfString("")
	case "nestchain.hotels.Booking.price":
		m := new(v1beta1.Coin)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: nestchain.hotels.Booking"))
		}
		panic(fmt.Errorf("message nestchain.hotels.Booking does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_Booking) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in nestchain.hotels.Booking", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_Booking) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Booking) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_Booking) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_Booking) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*Booking)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		l = len(x.BookingId)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.HotelId)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.CustomerId)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.StartDate)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.EndDate)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.RoomType)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.Price != nil {
			l = options.Size(x.Price)
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*Booking)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if x.Price != nil {
			encoded, err := options.Marshal(x.Price)
			if err != nil {
				return protoiface.MarshalOutput{
					NoUnkeyedLiterals: input.NoUnkeyedLiterals,
					Buf:               input.Buf,
				}, err
			}
			i -= len(encoded)
			copy(dAtA[i:], encoded)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
			i--
			dAtA[i] = 0x3a
		}
		if len(x.RoomType) > 0 {
			i -= len(x.RoomType)
			copy(dAtA[i:], x.RoomType)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.RoomType)))
			i--
			dAtA[i] = 0x32
		}
		if len(x.EndDate) > 0 {
			i -= len(x.EndDate)
			copy(dAtA[i:], x.EndDate)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.EndDate)))
			i--
			dAtA[i] = 0x2a
		}
		if len(x.StartDate) > 0 {
			i -= len(x.StartDate)
			copy(dAtA[i:], x.StartDate)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.StartDate)))
			i--
			dAtA[i] = 0x22
		}
		if len(x.CustomerId) > 0 {
			i -= len(x.CustomerId)
			copy(dAtA[i:], x.CustomerId)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.CustomerId)))
			i--
			dAtA[i] = 0x1a
		}
		if len(x.HotelId) > 0 {
			i -= len(x.HotelId)
			copy(dAtA[i:], x.HotelId)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.HotelId)))
			i--
			dAtA[i] = 0x12
		}
		if len(x.BookingId) > 0 {
			i -= len(x.BookingId)
			copy(dAtA[i:], x.BookingId)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.BookingId)))
			i--
			dAtA[i] = 0xa
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*Booking)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: Booking: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: Booking: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field BookingId", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
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
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.BookingId = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field HotelId", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
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
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.HotelId = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 3:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field CustomerId", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
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
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.CustomerId = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 4:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field StartDate", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
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
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.StartDate = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 5:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field EndDate", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
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
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.EndDate = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 6:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field RoomType", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
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
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.RoomType = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 7:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Price", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if x.Price == nil {
					x.Price = &v1beta1.Coin{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.Price); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.0
// 	protoc        (unknown)
// source: nestchain/hotels/booking.proto

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Booking struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BookingId  string        `protobuf:"bytes,1,opt,name=booking_id,json=bookingId,proto3" json:"booking_id,omitempty"`    // Unique identifier for the booking
	HotelId    string        `protobuf:"bytes,2,opt,name=hotel_id,json=hotelId,proto3" json:"hotel_id,omitempty"`          // Identifier for the hotel
	CustomerId string        `protobuf:"bytes,3,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"` // Identifier for the customer
	StartDate  string        `protobuf:"bytes,4,opt,name=start_date,json=startDate,proto3" json:"start_date,omitempty"`    // Start date of the booking
	EndDate    string        `protobuf:"bytes,5,opt,name=end_date,json=endDate,proto3" json:"end_date,omitempty"`          // End date of the booking
	RoomType   string        `protobuf:"bytes,6,opt,name=room_type,json=roomType,proto3" json:"room_type,omitempty"`       // Type of room booked
	Price      *v1beta1.Coin `protobuf:"bytes,7,opt,name=price,proto3" json:"price,omitempty"`                             // Price in coins
}

func (x *Booking) Reset() {
	*x = Booking{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nestchain_hotels_booking_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Booking) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Booking) ProtoMessage() {}

// Deprecated: Use Booking.ProtoReflect.Descriptor instead.
func (*Booking) Descriptor() ([]byte, []int) {
	return file_nestchain_hotels_booking_proto_rawDescGZIP(), []int{0}
}

func (x *Booking) GetBookingId() string {
	if x != nil {
		return x.BookingId
	}
	return ""
}

func (x *Booking) GetHotelId() string {
	if x != nil {
		return x.HotelId
	}
	return ""
}

func (x *Booking) GetCustomerId() string {
	if x != nil {
		return x.CustomerId
	}
	return ""
}

func (x *Booking) GetStartDate() string {
	if x != nil {
		return x.StartDate
	}
	return ""
}

func (x *Booking) GetEndDate() string {
	if x != nil {
		return x.EndDate
	}
	return ""
}

func (x *Booking) GetRoomType() string {
	if x != nil {
		return x.RoomType
	}
	return ""
}

func (x *Booking) GetPrice() *v1beta1.Coin {
	if x != nil {
		return x.Price
	}
	return nil
}

var File_nestchain_hotels_booking_proto protoreflect.FileDescriptor

var file_nestchain_hotels_booking_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x6e, 0x65, 0x73, 0x74, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2f, 0x68, 0x6f, 0x74, 0x65,
	0x6c, 0x73, 0x2f, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x10, 0x6e, 0x65, 0x73, 0x74, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x68, 0x6f, 0x74, 0x65,
	0x6c, 0x73, 0x1a, 0x1e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2f,
	0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x63, 0x6f, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xec, 0x01, 0x0a, 0x07, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x12, 0x1d,
	0x0a, 0x0a, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x49, 0x64, 0x12, 0x19, 0x0a,
	0x08, 0x68, 0x6f, 0x74, 0x65, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x68, 0x6f, 0x74, 0x65, 0x6c, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61,
	0x72, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73,
	0x74, 0x61, 0x72, 0x74, 0x44, 0x61, 0x74, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x5f,
	0x64, 0x61, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x44,
	0x61, 0x74, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x72, 0x6f, 0x6f, 0x6d, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x6f, 0x6f, 0x6d, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x2f, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x19, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x76, 0x31,
	0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x43, 0x6f, 0x69, 0x6e, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63,
	0x65, 0x42, 0xa8, 0x01, 0x0a, 0x14, 0x63, 0x6f, 0x6d, 0x2e, 0x6e, 0x65, 0x73, 0x74, 0x63, 0x68,
	0x61, 0x69, 0x6e, 0x2e, 0x68, 0x6f, 0x74, 0x65, 0x6c, 0x73, 0x42, 0x0c, 0x42, 0x6f, 0x6f, 0x6b,
	0x69, 0x6e, 0x67, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x21, 0x63, 0x6f, 0x73, 0x6d,
	0x6f, 0x73, 0x73, 0x64, 0x6b, 0x2e, 0x69, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6e, 0x65, 0x73,
	0x74, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2f, 0x68, 0x6f, 0x74, 0x65, 0x6c, 0x73, 0xa2, 0x02, 0x03,
	0x4e, 0x48, 0x58, 0xaa, 0x02, 0x10, 0x4e, 0x65, 0x73, 0x74, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e,
	0x48, 0x6f, 0x74, 0x65, 0x6c, 0x73, 0xca, 0x02, 0x10, 0x4e, 0x65, 0x73, 0x74, 0x63, 0x68, 0x61,
	0x69, 0x6e, 0x5c, 0x48, 0x6f, 0x74, 0x65, 0x6c, 0x73, 0xe2, 0x02, 0x1c, 0x4e, 0x65, 0x73, 0x74,
	0x63, 0x68, 0x61, 0x69, 0x6e, 0x5c, 0x48, 0x6f, 0x74, 0x65, 0x6c, 0x73, 0x5c, 0x47, 0x50, 0x42,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x11, 0x4e, 0x65, 0x73, 0x74, 0x63,
	0x68, 0x61, 0x69, 0x6e, 0x3a, 0x3a, 0x48, 0x6f, 0x74, 0x65, 0x6c, 0x73, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_nestchain_hotels_booking_proto_rawDescOnce sync.Once
	file_nestchain_hotels_booking_proto_rawDescData = file_nestchain_hotels_booking_proto_rawDesc
)

func file_nestchain_hotels_booking_proto_rawDescGZIP() []byte {
	file_nestchain_hotels_booking_proto_rawDescOnce.Do(func() {
		file_nestchain_hotels_booking_proto_rawDescData = protoimpl.X.CompressGZIP(file_nestchain_hotels_booking_proto_rawDescData)
	})
	return file_nestchain_hotels_booking_proto_rawDescData
}

var file_nestchain_hotels_booking_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_nestchain_hotels_booking_proto_goTypes = []interface{}{
	(*Booking)(nil),      // 0: nestchain.hotels.Booking
	(*v1beta1.Coin)(nil), // 1: cosmos.base.v1beta1.Coin
}
var file_nestchain_hotels_booking_proto_depIdxs = []int32{
	1, // 0: nestchain.hotels.Booking.price:type_name -> cosmos.base.v1beta1.Coin
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_nestchain_hotels_booking_proto_init() }
func file_nestchain_hotels_booking_proto_init() {
	if File_nestchain_hotels_booking_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_nestchain_hotels_booking_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Booking); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_nestchain_hotels_booking_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_nestchain_hotels_booking_proto_goTypes,
		DependencyIndexes: file_nestchain_hotels_booking_proto_depIdxs,
		MessageInfos:      file_nestchain_hotels_booking_proto_msgTypes,
	}.Build()
	File_nestchain_hotels_booking_proto = out.File
	file_nestchain_hotels_booking_proto_rawDesc = nil
	file_nestchain_hotels_booking_proto_goTypes = nil
	file_nestchain_hotels_booking_proto_depIdxs = nil
}