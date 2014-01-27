package ut

import (
	"testing"
)

func TestContains(test *testing.T) {
	strs := []interface{}{"a", "B", "b"}

	if !Contains(strs, "b") {
		test.Errorf("Expected Contains to return true if it contains a given string")
	}
	if Contains(strs, "c") {
		test.Errorf("Expected Contains to return false if it does not contain a given string")
	}

	mixed := []interface{}{1, "B", false}

	if !Contains(mixed, "B") {
		test.Errorf("Expected Contains to return true if it contains a given string")
	}
	if Contains(mixed, "c") {
		test.Errorf("Expected Contains to return false if it does not contain a given string")
	}
}
