package fields

import (
	"iter"
	"reflect"

	"github.com/iskorotkov/fastjson/tags"
)

type Field struct {
	Index int
	Tag   tags.Tag
	Type  reflect.Type
}

func Extract(typ reflect.Type) iter.Seq[Field] {
	return func(yield func(Field) bool) {
		for i := range typ.NumField() {
			field := typ.Field(i)
			if !field.IsExported() {
				continue
			}
			tag := tags.FieldMetadata(field)
			if tag.Name == "" || tag.Name == "-" {
				continue
			}
			if !yield(Field{Index: i, Tag: tag, Type: field.Type}) {
				return
			}
		}
	}
}
