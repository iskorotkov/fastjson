package encoder

import (
	"encoding"
	"encoding/json"
	"reflect"
	"time"

	"github.com/iskorotkov/fastjson/tiler"
	"github.com/iskorotkov/fastjson/xreflect"
	"github.com/iskorotkov/fastjson/xstrconv"
)

var encodersByKind [26]func(typ reflect.Type) Encoder

var encodersByType [4]CustomEncoder

func init() {
	encodersByKind = [...]func(typ reflect.Type) Encoder{
		reflect.Bool:    boolEncoder,
		reflect.Int:     intEncoder,
		reflect.Int8:    intEncoder,
		reflect.Int16:   intEncoder,
		reflect.Int32:   intEncoder,
		reflect.Int64:   intEncoder,
		reflect.Uint:    uintEncoder,
		reflect.Uint8:   uintEncoder,
		reflect.Uint16:  uintEncoder,
		reflect.Uint32:  uintEncoder,
		reflect.Uint64:  uintEncoder,
		reflect.Float32: floatEncoder,
		reflect.Float64: floatEncoder,
		reflect.String:  stringEncoder,
		reflect.Array:   arrayEncoder,
		reflect.Slice:   sliceEncoder,
		reflect.Map:     mapEncoder,
		reflect.Struct:  structEncoder,
		reflect.Pointer: pointerEncoder,
	}

	encodersByType = [...]CustomEncoder{
		{
			Type:    reflect.TypeFor[time.Duration](),
			Encoder: encodeTimeDuration,
		},
		{
			Type:    reflect.TypeFor[json.Unmarshaler](),
			Encoder: encodeJsonUnmarshaler,
		},
		{
			Type:    reflect.TypeFor[encoding.TextUnmarshaler](),
			Encoder: encodeTextUnmarshaler,
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

	for _, enc := range encodersByType {
		if typ.AssignableTo(enc.Type) || reflect.PointerTo(typ).AssignableTo(enc.Type) {
			return enc.Encoder
		}
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

type Encoder func(value reflect.Value, t *tiler.Tiler)

type CustomEncoder struct {
	Type    reflect.Type
	Encoder Encoder
}

func encodeNil(value reflect.Value, t *tiler.Tiler) {}

func boolEncoder(typ reflect.Type) Encoder {
	return encodeBool
}

func encodeBool(value reflect.Value, t *tiler.Tiler) {
	t.PutBool(value.Bool())
}

func intEncoder(typ reflect.Type) Encoder {
	return encodeInt
}

func encodeInt(value reflect.Value, t *tiler.Tiler) {
	t.PutInt(value.Int())
}

func uintEncoder(typ reflect.Type) Encoder {
	return encodeUint
}

func encodeUint(value reflect.Value, t *tiler.Tiler) {
	t.PutUint(value.Uint())
}

func floatEncoder(typ reflect.Type) Encoder {
	return encodeFloat
}

func encodeFloat(value reflect.Value, t *tiler.Tiler) {
	t.PutFloat(value.Float())
}

func stringEncoder(typ reflect.Type) Encoder {
	return encodeString
}

func encodeString(value reflect.Value, t *tiler.Tiler) {
	t.PutQuotedString(value.String())
}

func arrayEncoder(typ reflect.Type) Encoder {
	elemType := typ.Elem()
	itemsEncoder := New(elemType)
	length := typ.Len()
	return func(value reflect.Value, t *tiler.Tiler) {
		t.PutArrayStart()
		for i := range length {
			if i > 0 {
				t.PutComma()
			}
			itemsEncoder(value.Index(i), t)
		}
		t.PutArrayEnd()
	}
}

func sliceEncoder(typ reflect.Type) Encoder {
	elemType := typ.Elem()
	itemsEncoder := New(elemType)
	return func(value reflect.Value, t *tiler.Tiler) {
		if value.IsNil() {
			t.PutNull()
			return
		}
		t.PutArrayStart()
		for i := range value.Len() {
			if i > 0 {
				t.PutComma()
			}
			itemsEncoder(value.Index(i), t)
		}
		t.PutArrayEnd()
	}
}

func mapEncoder(typ reflect.Type) Encoder {
	keysEncoder := New(typ.Key())
	itemsEncoder := New(typ.Elem())
	return func(value reflect.Value, t *tiler.Tiler) {
		if value.IsNil() {
			t.PutNull()
			return
		}

		t.PutObjectStart()

		iter := value.MapRange()
		var i int
		for iter.Next() {
			if i > 0 {
				t.PutComma()
			}
			keysEncoder(iter.Key(), t)
			t.PutColon()
			itemsEncoder(iter.Value(), t)
			i++
		}

		t.PutObjectEnd()
	}
}

func structEncoder(typ reflect.Type) Encoder {
	var properties Properties
	for i := range typ.NumField() {
		field := typ.Field(i)
		name := xreflect.JSONTag(field)
		if name == "" || name == "-" {
			continue
		}
		enc := New(field.Type)
		properties = append(properties, Property{Index: i, Name: name, Encoder: enc})
	}
	return func(value reflect.Value, t *tiler.Tiler) {
		t.PutObjectStart()
		for i, prop := range properties {
			if i > 0 {
				t.PutComma()
			}
			t.PutQuotedString(prop.Name)
			t.PutColon()
			prop.Encoder(value.Field(prop.Index), t)
		}
		t.PutObjectEnd()
	}
}

func pointerEncoder(typ reflect.Type) Encoder {
	dec := New(typ.Elem())
	return func(value reflect.Value, t *tiler.Tiler) {
		if !value.IsNil() {
			dec(value.Elem(), t)
			return
		}
		t.PutNull()
	}
}

func encodeTimeDuration(value reflect.Value, t *tiler.Tiler) {
	dur, _ := xreflect.TypeAssert[time.Duration](value)
	t.PutDuration(dur)
}

func encodeJsonUnmarshaler(value reflect.Value, t *tiler.Tiler) {
	marshaler, ok := xreflect.TypeAssert[json.Marshaler](value)
	if !ok {
		marshaler, _ = xreflect.TypeAssert[json.Marshaler](value.Addr())
	}
	b, err := marshaler.MarshalJSON()
	if err != nil {
		panic(&MarshalerError{
			Err:   err,
			Value: value,
		})
	}
	t.PutBytes(b)
}

func encodeTextUnmarshaler(value reflect.Value, t *tiler.Tiler) {
	marshaler, ok := xreflect.TypeAssert[encoding.TextMarshaler](value)
	if !ok {
		marshaler, _ = xreflect.TypeAssert[encoding.TextMarshaler](value.Addr())
	}
	b, err := marshaler.MarshalText()
	if err != nil {
		panic(&MarshalerError{
			Err:   err,
			Value: value,
		})
	}
	t.PutQuotedBytes(b)
}

func encodeBinaryUnmarshaler(value reflect.Value, t *tiler.Tiler) {
	marshaler, ok := xreflect.TypeAssert[encoding.BinaryMarshaler](value)
	if !ok {
		marshaler, _ = xreflect.TypeAssert[encoding.BinaryMarshaler](value.Addr())
	}
	b, err := marshaler.MarshalBinary()
	if err != nil {
		panic(&MarshalerError{
			Err:   err,
			Value: value,
		})
	}
	t.PutQuotedString(xstrconv.BytesToString(b))
}
