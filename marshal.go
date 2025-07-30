package fastjson

import (
	"reflect"

	"github.com/iskorotkov/fastjson/encoder"
	"github.com/iskorotkov/fastjson/tiler"
	"github.com/iskorotkov/fastjson/xstrconv"
)

func NewEncoder[T any]() Encoder[T] {
	var v T
	tiler := tiler.New()
	return Encoder[T]{
		enc:   encoder.New(reflect.TypeOf(v)),
		tiler: &tiler,
	}
}

type Encoder[T any] struct {
	enc   encoder.Encoder
	tiler *tiler.Tiler
}

func (e Encoder[T]) Marshal(v T) (b []byte, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = r.(error)
		}
	}()

	e.enc(reflect.ValueOf(v), e.tiler)

	res := e.tiler.Clone()
	e.tiler.Reset()

	return res, nil
}

func (e Encoder[T]) MarshalString(v T) (s string, err error) {
	b, err := e.Marshal(v)
	return xstrconv.BytesToString(b), err
}
