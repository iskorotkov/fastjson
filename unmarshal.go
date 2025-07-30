package fastjson

import (
	"fmt"
	"io"
	"reflect"

	"github.com/iskorotkov/fastjson/decoder"
	"github.com/iskorotkov/fastjson/tokenizer"
)

func NewDecoder[T any]() Decoder[T] {
	var v T
	return Decoder[T]{
		dec: decoder.New(reflect.TypeOf(v)),
	}
}

type Decoder[T any] struct {
	dec decoder.Decoder
}

func (d Decoder[T]) Unmarshal(data []byte, v *T) (err error) {
	defer func() {
		if r := recover(); r != nil {
			switch e := r.(type) {
			case error:
				err = e
			default:
				err = fmt.Errorf("unexpected panic: %v", r)
			}
		}
	}()

	val := reflect.Indirect(reflect.ValueOf(v))
	if !val.CanAddr() {
		return &NotAddressableError{
			Value: val,
		}
	}

	tokens := tokenizer.NewFromBytes(data)
	d.dec(val, &tokens)
	return nil
}

func (d Decoder[T]) UnmarshalString(s string, v *T) (err error) {
	defer func() {
		if r := recover(); r != nil {
			switch e := r.(type) {
			case error:
				err = e
			default:
				err = fmt.Errorf("unexpected panic: %v", r)
			}
		}
	}()

	val := reflect.Indirect(reflect.ValueOf(v))
	if !val.CanAddr() {
		return &NotAddressableError{
			Value: val,
		}
	}

	tokens := tokenizer.NewFromString(s)
	d.dec(val, &tokens)
	return nil
}

func (d Decoder[T]) UnmarshalReader(r io.Reader, v *T) (err error) {
	defer func() {
		if r := recover(); r != nil {
			switch e := r.(type) {
			case error:
				err = e
			default:
				err = fmt.Errorf("unexpected panic: %v", r)
			}
		}
	}()

	val := reflect.Indirect(reflect.ValueOf(v))
	if !val.CanAddr() {
		return &NotAddressableError{
			Value: val,
		}
	}

	tokens := tokenizer.NewFromReader(r)
	d.dec(val, &tokens)
	return nil
}
