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

func TestAny(test *testing.T) {
	strs := []interface{}{"a", "B", "b"}

	f := func(o interface{}) bool {
		s := o.(string)
		return s == "B"
	}

	if !Any(strs, f) {
		test.Errorf("Expected Any to return true if the given func returns true")
	}

	f = func(o interface{}) bool {
		return false
	}

	if Any(strs, f) {
		test.Errorf("Expected Any to return false if the given func returns false")
	}
}
