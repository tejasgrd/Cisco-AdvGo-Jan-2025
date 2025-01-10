package utils_test

import (
	"fmt"
	"testing"
	"testing-demo/utils"
)

/* func TestIsPrime_11(t *testing.T) {
	// arrange
	var no int64 = 11
	expectedResult := true

	// act
	actualResult := utils.IsPrime(no)

	// assert
	if actualResult != expectedResult {
		t.Errorf("IsPrime(%d), expected = %t but actual = %t\n", no, expectedResult, actualResult)
	}
}

func TestIsPrime_14(t *testing.T) {

	// arrange
	var no int64 = 14
	expectedResult := false

	// act
	actualResult := utils.IsPrime(no)

	// assert
	if actualResult != expectedResult {
		t.Errorf("IsPrime(%d), expected = %t but actual = %t\n", no, expectedResult, actualResult)
	}
} */

// data driven tests (by dynamically creating test functions at runtime)
func TestIsPrime(t *testing.T) {
	// prepare the data
	testData := []struct {
		no       int64
		expected bool
	}{
		{no: 11, expected: true},
		{no: 12, expected: false},
		{no: 13, expected: false},
		{no: 17, expected: true},
		{no: 19, expected: true},
	}
	for _, td := range testData {
		t.Run(fmt.Sprintf("IsPrime(%d)", td.no), func(t *testing.T) {
			actual := utils.IsPrime(td.no)
			if actual != td.expected {
				t.Errorf("IsPrime(%d), expected = %t but actual = %t\n", td.no, td.expected, actual)
			}
		})
	}
}
