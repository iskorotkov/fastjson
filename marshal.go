package fastjson

import (
	"fmt"
	"reflect"

	"github.com/iskorotkov/fastjson/buffer"
	"github.com/iskorotkov/fastjson/encoder"
	"github.com/iskorotkov/fastjson/view"
)

func NewEncoder[T any]() Encoder[T] {
	var v T
	buf := buffer.New()
	return Encoder[T]{
		enc: encoder.New(reflect.TypeOf(v)),
		buf: &buf,
	}
}

type Encoder[T any] struct {
	enc encoder.Encoder
	buf *buffer.Buffer
}

func (e Encoder[T]) Marshal(v T) (b []byte, err error) {
	defer func() {
		if r := recover(); r != nil {
			switch re := r.(type) {
			case error:
				err = re
			default:
				err = fmt.Errorf("unexpected panic: %v", r)
			}
		}
	}()

	e.enc(reflect.ValueOf(v), e.buf)

	res := e.buf.Clone()
	e.buf.Reset()

	return res, nil
}

func (e Encoder[T]) MarshalString(v T) (s string, err error) {
	b, err := e.Marshal(v)
	return view.BytesAsStr(b), err
}
