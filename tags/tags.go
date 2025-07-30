package tags

import (
	"reflect"
	"strings"
)

const tagName = "json"

type Tag struct {
	Name      string
	OmitEmpty bool
	String    bool
}

func FieldMetadata(field reflect.StructField) Tag {
	tag, ok := field.Tag.Lookup(tagName)
	if !ok {
		return Tag{Name: field.Name}
	}
	name, rest, hasOptions := strings.Cut(tag, ",")
	if name == "" {
		name = field.Name
	}
	if !hasOptions {
		return Tag{Name: name}
	}
	return Tag{
		Name:      name,
		OmitEmpty: strings.Contains(rest, "omitempty"),
		String:    strings.Contains(rest, "string"),
	}
}
