package int_utils

import (
	"reflect"
	"testing"
)

func TestIntInSlice(t *testing.T) {
	type addTest struct {
		number     int
		listNumber []int
		expected   bool
	}

	var addTests = []addTest{
		{3, []int{1, 2, 3}, true},
		{4, []int{1, 2, 3}, false},
		{5, []int{1, 2, 3, 4, 5}, true},
	}

	for _, value := range addTests {
		expectation := value.expected
		actual := intInSlice(value.number, value.listNumber)
		if !reflect.DeepEqual(actual, expectation) {
			t.Errorf("Expected %v but got %v", expectation, actual)
		}
	}

}
