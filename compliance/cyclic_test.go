package compliance

import (
	"testing"
)

func TestMarshal_Cyclic(t *testing.T) {
	a := &cyclicA{}
	b := &cyclicB{A: a}
	a.B = b

	runMarshalErrorTest(t, "simple_cycle", a)

	type selfRef struct {
		Self *selfRef `json:"self"`
	}
	s := &selfRef{}
	s.Self = s
	runMarshalErrorTest(t, "self_referential", s)
}

func TestMarshal_DeepNesting(t *testing.T) {
	type nested struct {
		Inner *nested `json:"inner"`
	}

	root := &nested{}
	current := root
	for i := 0; i < 100; i++ {
		current.Inner = &nested{}
		current = current.Inner
	}

	runMarshalTest(t, "deep_100", root, "")
}

func TestMarshal_RecursiveMap(t *testing.T) {
	runMarshalTest(t, "map/empty", recursiveMap{}, "{}")
	runMarshalTest(t, "map/nested", recursiveMap{"a": recursiveMap{"b": nil}}, `{"a":{"b":null}}`)
}

func TestMarshal_RecursiveSlice(t *testing.T) {
	runMarshalTest(t, "slice/empty", recursiveSlice{}, "[]")
	runMarshalTest(t, "slice/nested", recursiveSlice{recursiveSlice{nil}}, "[[null]]")
}

func TestMarshal_RecursivePointer(t *testing.T) {
	runMarshalTest(t, "pointer/nil", recursivePointer{}, `{"p":null}`)
	runMarshalTest(t, "pointer/nested", recursivePointer{P: &recursivePointer{}}, `{"p":{"p":null}}`)
}
