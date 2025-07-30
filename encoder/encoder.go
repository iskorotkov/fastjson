package encoder

import (
	"encoding"
	"encoding/json"
	"math"
	"reflect"
	"sync"
	"time"

	"github.com/iskorotkov/fastjson/buffer"
	"github.com/iskorotkov/fastjson/cast"
	"github.com/iskorotkov/fastjson/fields"
	"github.com/iskorotkov/fastjson/view"
)

var (
	encodersByKind [26]func(typ reflect.Type) Encoder

	encodersByType [4]CustomEncoder

	encoderCache sync.Map // map[reflect.Type]*cachedEncoder
)

type cachedEncoder struct {
	enc Encoder
}

func init() {
	encodersByKind = [...]func(typ reflect.Type) Encoder{
		reflect.Bool:      boolEncoder,
		reflect.Int:       intEncoder,
		reflect.Int8:      intEncoder,
		reflect.Int16:     intEncoder,
		reflect.Int32:     intEncoder,
		reflect.Int64:     intEncoder,
		reflect.Uint:      uintEncoder,
		reflect.Uint8:     uintEncoder,
		reflect.Uint16:    uintEncoder,
		reflect.Uint32:    uintEncoder,
		reflect.Uint64:    uintEncoder,
		reflect.Uintptr:   uintEncoder,
		reflect.Float32:   float32Encoder,
		reflect.Float64:   float64Encoder,
		reflect.String:    stringEncoder,
		reflect.Array:     arrayEncoder,
		reflect.Slice:     sliceEncoder,
		reflect.Map:       mapEncoder,
		reflect.Struct:    structEncoder,
		reflect.Pointer:   pointerEncoder,
		reflect.Interface: interfaceEncoder,
	}

	encodersByType = [...]CustomEncoder{
		{
			Type:    reflect.TypeFor[time.Duration](),
			Encoder: encodeTimeDuration,
		},
		{
			Type:    reflect.TypeFor[json.Marshaler](),
			Encoder: encodeJsonMarshaler,
		},
		{
			Type:    reflect.TypeFor[encoding.TextMarshaler](),
			Encoder: encodeTextMarshaler,
		},
		{
			Type:    reflect.TypeFor[encoding.BinaryUnmarshaler](),
			Encoder: encodeBinaryUnmarshaler,
		},
	}
}

func New(typ reflect.Type) Encoder {
	if typ == nil {
		return encodeNil
	}

	if cached, ok := encoderCache.Load(typ); ok {
		return cached.(*cachedEncoder).enc
	}

	entry := &cachedEncoder{}
	actual, loaded := encoderCache.LoadOrStore(typ, entry)
	if loaded {
		return actual.(*cachedEncoder).enc
	}

	enc := buildEncoder(typ)
	entry.enc = enc
	return enc
}

func buildEncoder(typ reflect.Type) Encoder {
	for _, enc := range encodersByType {
		if typ.AssignableTo(enc.Type) || reflect.PointerTo(typ).AssignableTo(enc.Type) {
			return enc.Encoder
		}
	}

	if typ.Kind() == reflect.Slice && typ.Elem().Kind() == reflect.Uint8 {
		return encodeByteSlice
	}

	kind := typ.Kind()
	if kind >= reflect.Kind(len(encodersByKind)) {
		panic(&UnsupportedTypeError{
			Type: typ,
		})
	}

	f := encodersByKind[kind]
	if f == nil {
		panic(&UnsupportedTypeError{
			Type: typ,
		})
	}

	return f(typ)
}

type Encoder func(value reflect.Value, b *buffer.Buffer)

type CustomEncoder struct {
	Type    reflect.Type
	Encoder Encoder
}

func encodeNil(value reflect.Value, b *buffer.Buffer) {
	b.PutNull()
}

func boolEncoder(typ reflect.Type) Encoder {
	return encodeBool
}

func encodeBool(value reflect.Value, b *buffer.Buffer) {
	b.PutBool(value.Bool())
}

func intEncoder(typ reflect.Type) Encoder {
	return encodeInt
}

func encodeInt(value reflect.Value, b *buffer.Buffer) {
	b.PutInt(value.Int())
}

func uintEncoder(typ reflect.Type) Encoder {
	return encodeUint
}

func encodeUint(value reflect.Value, b *buffer.Buffer) {
	b.PutUint(value.Uint())
}

func float32Encoder(typ reflect.Type) Encoder {
	return encodeFloat32
}

func encodeFloat32(value reflect.Value, b *buffer.Buffer) {
	f := float32(value.Float())
	if math.IsNaN(float64(f)) || math.IsInf(float64(f), 0) {
		panic(&UnsupportedValueError{
			Value: value,
			Str:   "NaN and Inf float values are not supported",
		})
	}
	b.PutFloat32(f)
}

func float64Encoder(typ reflect.Type) Encoder {
	return encodeFloat64
}

func encodeFloat64(value reflect.Value, b *buffer.Buffer) {
	f := value.Float()
	if math.IsNaN(f) || math.IsInf(f, 0) {
		panic(&UnsupportedValueError{
			Value: value,
			Str:   "NaN and Inf float values are not supported",
		})
	}
	b.PutFloat(f)
}

func stringEncoder(typ reflect.Type) Encoder {
	return encodeString
}

func encodeString(value reflect.Value, b *buffer.Buffer) {
	b.PutQuotedString(value.String())
}

func arrayEncoder(typ reflect.Type) Encoder {
	elemType := typ.Elem()
	itemsEncoder := New(elemType)
	length := typ.Len()
	return func(value reflect.Value, b *buffer.Buffer) {
		b.PutArrayStart()
		for i := range length {
			if i > 0 {
				b.PutComma()
			}
			itemsEncoder(value.Index(i), b)
		}
		b.PutArrayEnd()
	}
}

func sliceEncoder(typ reflect.Type) Encoder {
	elemType := typ.Elem()
	itemsEncoder := New(elemType)
	return func(value reflect.Value, b *buffer.Buffer) {
		if value.IsNil() {
			b.PutNull()
			return
		}
		b.PutArrayStart()
		for i := range value.Len() {
			if i > 0 {
				b.PutComma()
			}
			itemsEncoder(value.Index(i), b)
		}
		b.PutArrayEnd()
	}
}

func mapEncoder(typ reflect.Type) Encoder {
	keysEncoder := New(typ.Key())
	itemsEncoder := New(typ.Elem())
	return func(value reflect.Value, b *buffer.Buffer) {
		if value.IsNil() {
			b.PutNull()
			return
		}

		b.PutObjectStart()

		iter := value.MapRange()
		var i int
		for iter.Next() {
			if i > 0 {
				b.PutComma()
			}
			keysEncoder(iter.Key(), b)
			b.PutColon()
			itemsEncoder(iter.Value(), b)
			i++
		}

		b.PutObjectEnd()
	}
}

func structEncoder(typ reflect.Type) Encoder {
	var properties Properties
	hasOmitEmpty := false
	for f := range fields.Extract(typ) {
		properties = append(properties, Property{
			Index:           f.Index,
			Name:            f.Tag.Name,
			PrecomputedName: precomputeFieldName(f.Tag.Name),
			Encoder:         New(f.Type),
			OmitEmpty:       f.Tag.OmitEmpty,
		})
		if f.Tag.OmitEmpty {
			hasOmitEmpty = true
		}
	}
	if hasOmitEmpty {
		return structEncoderWithOmitEmpty(properties)
	}
	return structEncoderSimple(properties)
}

func structEncoderSimple(properties Properties) Encoder {
	return func(value reflect.Value, b *buffer.Buffer) {
		b.PutObjectStart()
		for i, prop := range properties {
			if i > 0 {
				b.PutComma()
			}
			b.PutBytes(prop.PrecomputedName)
			prop.Encoder(value.Field(prop.Index), b)
		}
		b.PutObjectEnd()
	}
}

func structEncoderWithOmitEmpty(properties Properties) Encoder {
	return func(value reflect.Value, b *buffer.Buffer) {
		b.PutObjectStart()
		first := true
		for _, prop := range properties {
			if prop.OmitEmpty && isEmpty(value.Field(prop.Index)) {
				continue
			}
			if !first {
				b.PutComma()
			}
			first = false
			b.PutBytes(prop.PrecomputedName)
			prop.Encoder(value.Field(prop.Index), b)
		}
		b.PutObjectEnd()
	}
}

func isEmpty(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Bool:
		return !v.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return v.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return v.Float() == 0
	case reflect.String:
		return v.String() == ""
	case reflect.Slice, reflect.Map:
		return v.IsNil() || v.Len() == 0
	case reflect.Pointer, reflect.Interface:
		return v.IsNil()
	case reflect.Array:
		return v.Len() == 0
	}
	return false
}

func pointerEncoder(typ reflect.Type) Encoder {
	dec := New(typ.Elem())
	return func(value reflect.Value, b *buffer.Buffer) {
		if !value.IsNil() {
			dec(value.Elem(), b)
			return
		}
		b.PutNull()
	}
}

func interfaceEncoder(typ reflect.Type) Encoder {
	return encodeInterface
}

func encodeInterface(value reflect.Value, b *buffer.Buffer) {
	if value.IsNil() {
		b.PutNull()
		return
	}
	New(value.Elem().Type())(value.Elem(), b)
}

func encodeTimeDuration(value reflect.Value, b *buffer.Buffer) {
	dur, _ := reflect.TypeAssert[time.Duration](value)
	b.PutDuration(dur)
}

func encodeJsonMarshaler(value reflect.Value, b *buffer.Buffer) {
	if value.Kind() == reflect.Pointer && value.IsNil() {
		b.PutNull()
		return
	}
	marshaler, _ := cast.To[json.Marshaler](value)
	bytes, err := marshaler.MarshalJSON()
	if err != nil {
		panic(&MarshalerError{
			Err:   err,
			Value: value,
		})
	}
	b.PutBytes(bytes)
}

func encodeTextMarshaler(value reflect.Value, b *buffer.Buffer) {
	if value.Kind() == reflect.Pointer && value.IsNil() ||
		value.Kind() == reflect.Slice && value.IsNil() {
		b.PutNull()
		return
	}
	marshaler, _ := cast.To[encoding.TextMarshaler](value)
	bytes, err := marshaler.MarshalText()
	if err != nil {
		panic(&MarshalerError{
			Err:   err,
			Value: value,
		})
	}
	b.PutQuotedBytes(bytes)
}

func encodeBinaryUnmarshaler(value reflect.Value, b *buffer.Buffer) {
	marshaler, _ := cast.To[encoding.BinaryMarshaler](value)
	bytes, err := marshaler.MarshalBinary()
	if err != nil {
		panic(&MarshalerError{
			Err:   err,
			Value: value,
		})
	}
	b.PutQuotedString(view.BytesAsStr(bytes))
}

func encodeByteSlice(value reflect.Value, b *buffer.Buffer) {
	bytes := value.Bytes()
	if len(bytes) == 0 {
		b.PutString(`""`)
		return
	}
	b.PutString(`"`)
	b.PutBase64(bytes)
	b.PutString(`"`)
}
