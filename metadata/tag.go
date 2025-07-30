package metadata

import (
	"reflect"
	"strings"
)

const tagName = "json"

func JSONFieldName(field reflect.StructField) string {
	tag, ok := field.Tag.Lookup(tagName)
	if  !ok {
		return field.Name
	}
	if !strings.Contains(tag, ",") {
		return tag
	}
	return strings.SplitN(tag, ",", 2)[0]
}
