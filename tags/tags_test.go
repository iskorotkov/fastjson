package tags

import (
	"reflect"
	"testing"
)

func TestFieldMetadata(t *testing.T) {
	type testStruct struct {
		NoTag          string
		BasicName      string `json:"basic_name"`
		WithOmitEmpty  string `json:"with_omit,omitempty"`
		WithString     string `json:"with_string,string"`
		Skip           string `json:"-"`
		Combined       string `json:"combined,omitempty,string"`
		EmptyName      string `json:",omitempty"`
		OnlyOmitEmpty  string `json:",omitempty"`
		OnlyString     string `json:",string"`
		EmptyTag       string `json:""`
		NameWithHyphen string `json:"name-with-hyphen"`
		NameOnly       string `json:"name_only"`
		OmitEmptyFirst string `json:"field,omitempty,string"`
		StringFirst    string `json:"field2,string,omitempty"`
	}

	typ := reflect.TypeFor[testStruct]()

	cases := []struct {
		fieldName string
		expected  Tag
	}{
		{
			fieldName: "NoTag",
			expected:  Tag{Name: "NoTag", OmitEmpty: false, String: false},
		},
		{
			fieldName: "BasicName",
			expected:  Tag{Name: "basic_name", OmitEmpty: false, String: false},
		},
		{
			fieldName: "WithOmitEmpty",
			expected:  Tag{Name: "with_omit", OmitEmpty: true, String: false},
		},
		{
			fieldName: "WithString",
			expected:  Tag{Name: "with_string", OmitEmpty: false, String: true},
		},
		{
			fieldName: "Skip",
			expected:  Tag{Name: "-", OmitEmpty: false, String: false},
		},
		{
			fieldName: "Combined",
			expected:  Tag{Name: "combined", OmitEmpty: true, String: true},
		},
		{
			fieldName: "EmptyName",
			expected:  Tag{Name: "EmptyName", OmitEmpty: true, String: false},
		},
		{
			fieldName: "OnlyOmitEmpty",
			expected:  Tag{Name: "OnlyOmitEmpty", OmitEmpty: true, String: false},
		},
		{
			fieldName: "OnlyString",
			expected:  Tag{Name: "OnlyString", OmitEmpty: false, String: true},
		},
		{
			fieldName: "EmptyTag",
			expected:  Tag{Name: "EmptyTag", OmitEmpty: false, String: false},
		},
		{
			fieldName: "NameWithHyphen",
			expected:  Tag{Name: "name-with-hyphen", OmitEmpty: false, String: false},
		},
		{
			fieldName: "NameOnly",
			expected:  Tag{Name: "name_only", OmitEmpty: false, String: false},
		},
		{
			fieldName: "OmitEmptyFirst",
			expected:  Tag{Name: "field", OmitEmpty: true, String: true},
		},
		{
			fieldName: "StringFirst",
			expected:  Tag{Name: "field2", OmitEmpty: true, String: true},
		},
	}

	for _, c := range cases {
		t.Run(c.fieldName, func(t *testing.T) {
			field, ok := typ.FieldByName(c.fieldName)
			if !ok {
				t.Fatalf("field %s not found", c.fieldName)
			}
			got := FieldMetadata(field)
			if got != c.expected {
				t.Errorf("Parse(%s) = %+v, want %+v", c.fieldName, got, c.expected)
			}
		})
	}
}

