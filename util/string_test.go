package util

import (
	"reflect"
	"testing"
)

func TestStringToSlice(t *testing.T) {
	arg := "line1\nline2\nlin'e'3"
	expected := []string{"line1", "line2", "lin'e'3"}
	result := StringToSlice(arg)

	// 使用 reflect.DeepEqual 对比两个切片是否一致
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Test failed, expected: '%v', got:  '%v'", expected, result)
	}
}
