package stringutils

import (
	"reflect"
	"testing"
)

func TestExtractStringToArrayStr(t *testing.T) {
	type addTest struct {
		param    string
		expected []string
	}

	var addTests = []addTest{
		{"in(a,b,c)", []string{"a", "b", "c"}},
		{"in(1,2,3)", []string{"1", "2", "3"}},
		{"in(B00000112,B00000142,B00000123)", []string{"B00000112", "B00000142", "B00000123"}},
	}

	for _, value := range addTests {
		expectation := value.expected
		actual := ExtractStringToArrayStr(value.param)
		if !reflect.DeepEqual(actual, expectation) {
			t.Errorf("Expected %v but got %v", expectation, actual)
		}
	}

}

func TestGenerateStringInFromArrayStr(t *testing.T) {
	type addTest struct {
		param    []string
		expected string
	}

	var addTests = []addTest{
		{[]string{"1", "2", "3"}, "in(1,2,3)"},
		{[]string{"a", "b", "c"}, "in(a,b,c)"},
		{[]string{"100", "200", "300"}, "in(100,200,300)"},
	}

	for _, value := range addTests {
		expectation := value.expected
		actual := GenerateStringInFromArrayStr(value.param)
		if actual != expectation {
			t.Errorf("Expected %v but got %v", expectation, actual)
		}
	}
}

func TestGenerateStringInFromArrayInt64(t *testing.T) {
	type addTest struct {
		param    []int64
		expected string
	}

	var addTests = []addTest{
		{[]int64{1, 2, 3}, "in(1,2,3)"},
		{[]int64{100, 200, 300}, "in(100,200,300)"},
		{[]int64{1000, 2000, 3000}, "in(1000,2000,3000)"},
		{[]int64{10000, 20000, 30000}, "in(10000,20000,30000)"},
	}

	for _, value := range addTests {
		expectation := value.expected
		actual := GenerateStringInFromArrayInt64(value.param)
		if actual != expectation {
			t.Errorf("Expected %v but got %v", expectation, actual)
		}
	}
}
