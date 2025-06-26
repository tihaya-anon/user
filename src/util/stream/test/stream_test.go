package test

import (
	"MVC_DI/util/stream"
	"reflect"
	"testing"
)

func Test_ListStream(t *testing.T) {
	// Case 1: Map
	// Given
	source := []int{1, 2, 3}
	// When
	doubled := stream.NewListStream(source).Map(func(item int) any { return item * 2 }).ToList()
	// Then
	expectedDoubled := []any{2, 4, 6}
	if !reflect.DeepEqual(doubled, expectedDoubled) {
		t.Errorf("case `Map` failed:\nexpected: %v\ngot: %v", expectedDoubled, doubled)
	}

	// Case 2: Filter
	// Given
	// source above
	// When
	even := stream.NewListStream(source).Filter(func(item int) bool { return item%2 == 0 }).ToList()
	// Then
	expectedEven := []int{2}
	if !reflect.DeepEqual(even, expectedEven) {
		t.Errorf("case `Filter` failed:\nexpected: %v\ngot: %v", expectedEven, even)
	}
}

func Test_MapStream(t *testing.T) {
	// Case 1: Map
	// Given
	source := map[string]int{"a": 1, "b": 2, "c": 3}

	// When
	doubled := stream.NewMapStream(source).Map(func(key string, val int) (string, any) {
		return key + "_double", val * 2
	}).ToMap()

	// Then
	expectedDoubled := map[string]any{
		"a_double": 2,
		"b_double": 4,
		"c_double": 6,
	}
	if !reflect.DeepEqual(doubled, expectedDoubled) {
		t.Errorf("case `Map` failed:\nexpected: %v\ngot: %v", expectedDoubled, doubled)
	}

	// Case 2: Filter
	// Given
	// source above
	// When
	even := stream.NewMapStream(source).Filter(func(key string, val int) bool {
		return val%2 == 0
	}).ToMap()

	// Then
	expectedEven := map[string]int{"b": 2}
	if !reflect.DeepEqual(even, expectedEven) {
		t.Errorf("case `Filter` failed:\nexpected: %v\ngot: %v", expectedEven, even)
	}
}
