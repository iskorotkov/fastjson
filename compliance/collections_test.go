package compliance

import "testing"

func TestMarshal_Slices(t *testing.T) {
	runMarshalTest(t, "nil", ([]int)(nil), "null")
	runMarshalTest(t, "empty", []int{}, "[]")
	runMarshalTest(t, "single", []int{1}, "[1]")
	runMarshalTest(t, "multiple", []int{1, 2, 3}, "[1,2,3]")

	runMarshalTest(t, "strings/empty", []string{}, "[]")
	runMarshalTest(t, "strings/single", []string{"hello"}, `["hello"]`)
	runMarshalTest(t, "strings/multiple", []string{"a", "b", "c"}, `["a","b","c"]`)

	runMarshalTest(t, "bools/multiple", []bool{true, false, true}, "[true,false,true]")

	runMarshalTest(t, "nested/empty", [][]int{}, "[]")
	runMarshalTest(t, "nested/single", [][]int{{1, 2}}, "[[1,2]]")
	runMarshalTest(t, "nested/multiple", [][]int{{1, 2}, {3, 4}}, "[[1,2],[3,4]]")

	runMarshalTest(t, "pointers/nil_element", []*int{nil}, "[null]")
	runMarshalTest(t, "pointers/mixed", []*int{addr(1), nil, addr(3)}, "[1,null,3]")

	runMarshalTest(t, "pointer/nil", (*[]int)(nil), "null")
	runMarshalTest(t, "pointer/empty", addr([]int{}), "[]")
	runMarshalTest(t, "pointer/single", addr([]int{1}), "[1]")
}

func TestMarshal_Arrays(t *testing.T) {
	runMarshalTest(t, "empty", [0]int{}, "[]")
	runMarshalTest(t, "single", [1]int{1}, "[1]")
	runMarshalTest(t, "multiple", [3]int{1, 2, 3}, "[1,2,3]")

	runMarshalTest(t, "strings/single", [1]string{"hello"}, `["hello"]`)
	runMarshalTest(t, "strings/multiple", [3]string{"a", "b", "c"}, `["a","b","c"]`)

	runMarshalTest(t, "bools/multiple", [3]bool{true, false, true}, "[true,false,true]")

	runMarshalTest(t, "nested", [2][2]int{{1, 2}, {3, 4}}, "[[1,2],[3,4]]")

	runMarshalTest(t, "pointer/single", addr([1]int{1}), "[1]")
	runMarshalTest(t, "pointer/nil", (*[1]int)(nil), "null")
}

func TestMarshal_Maps(t *testing.T) {
	runMarshalTest(t, "nil", (map[string]int)(nil), "null")
	runMarshalTest(t, "empty", map[string]int{}, "{}")

	runMarshalTest(t, "string_keys/single", map[string]int{"a": 1}, `{"a":1}`)

	runMarshalTest(t, "string_values/single", map[string]string{"key": "value"}, `{"key":"value"}`)

	runMarshalTest(t, "bool_values/single", map[string]bool{"flag": true}, `{"flag":true}`)

	runMarshalTest(t, "nested/single", map[string]map[string]int{"outer": {"inner": 1}}, `{"outer":{"inner":1}}`)

	runMarshalTest(t, "slice_values/single", map[string][]int{"nums": {1, 2, 3}}, `{"nums":[1,2,3]}`)

	runMarshalTest(t, "pointer/nil", (*map[string]int)(nil), "null")
	runMarshalTest(t, "pointer/empty", addr(map[string]int{}), "{}")
}

func TestUnmarshal_Slices(t *testing.T) {
	runUnmarshalTest(t, "null", "null", ([]int)(nil))
	runUnmarshalTest(t, "empty", "[]", []int{})
	runUnmarshalTest(t, "single", "[1]", []int{1})
	runUnmarshalTest(t, "multiple", "[1,2,3]", []int{1, 2, 3})
	runUnmarshalTest(t, "with_spaces", "[ 1 , 2 , 3 ]", []int{1, 2, 3})

	runUnmarshalTest(t, "strings/empty", "[]", []string{})
	runUnmarshalTest(t, "strings/single", `["hello"]`, []string{"hello"})
	runUnmarshalTest(t, "strings/multiple", `["a","b","c"]`, []string{"a", "b", "c"})

	runUnmarshalTest(t, "bools/multiple", "[true,false,true]", []bool{true, false, true})

	runUnmarshalTest(t, "nested/empty", "[]", [][]int{})
	runUnmarshalTest(t, "nested/single", "[[1,2]]", [][]int{{1, 2}})
	runUnmarshalTest(t, "nested/multiple", "[[1,2],[3,4]]", [][]int{{1, 2}, {3, 4}})

	runUnmarshalErrorTest[[]int](t, "invalid/object", "{}")
	runUnmarshalErrorTest[[]int](t, "invalid/string", `"[]"`)
	runUnmarshalErrorTest[[]int](t, "invalid/element", `[1,"two",3]`)
}

func TestUnmarshal_Arrays(t *testing.T) {
	runUnmarshalTest(t, "empty", "[]", [0]int{})
	runUnmarshalTest(t, "single", "[1]", [1]int{1})
	runUnmarshalTest(t, "multiple", "[1,2,3]", [3]int{1, 2, 3})

	runUnmarshalTest(t, "strings/single", `["hello"]`, [1]string{"hello"})
	runUnmarshalTest(t, "strings/multiple", `["a","b","c"]`, [3]string{"a", "b", "c"})

	runUnmarshalTest(t, "bools/multiple", "[true,false,true]", [3]bool{true, false, true})

	runUnmarshalTest(t, "nested", "[[1,2],[3,4]]", [2][2]int{{1, 2}, {3, 4}})

	runUnmarshalErrorTest[[3]int](t, "wrong_length/short", "[1,2]")
	runUnmarshalErrorTest[[3]int](t, "wrong_length/long", "[1,2,3,4]")
}

func TestUnmarshal_Maps(t *testing.T) {
	runUnmarshalTest(t, "null", "null", (map[string]int)(nil))
	runUnmarshalTest(t, "empty", "{}", map[string]int{})
	runUnmarshalTest(t, "single", `{"a":1}`, map[string]int{"a": 1})
	runUnmarshalTest(t, "multiple", `{"a":1,"b":2,"c":3}`, map[string]int{"a": 1, "b": 2, "c": 3})
	runUnmarshalTest(t, "with_spaces", `{ "a" : 1 , "b" : 2 }`, map[string]int{"a": 1, "b": 2})

	runUnmarshalTest(t, "string_values/single", `{"key":"value"}`, map[string]string{"key": "value"})
	runUnmarshalTest(t, "bool_values/single", `{"flag":true}`, map[string]bool{"flag": true})

	runUnmarshalTest(t, "nested/single", `{"outer":{"inner":1}}`, map[string]map[string]int{"outer": {"inner": 1}})

	runUnmarshalTest(t, "slice_values/single", `{"nums":[1,2,3]}`, map[string][]int{"nums": {1, 2, 3}})

	runUnmarshalErrorTest[map[string]int](t, "invalid/array", "[]")
	runUnmarshalErrorTest[map[string]int](t, "invalid/string", `"{}"`)
	runUnmarshalErrorTest[map[string]int](t, "invalid/value", `{"a":"one"}`)
}
