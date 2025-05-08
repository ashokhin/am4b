package utils

import (
	"testing"
)

func TestIntFromString(t *testing.T) {
	testCases := map[string]struct {
		testString string
		expected   int
	}{
		"test01": {"Alice got 8 pieces of pizza", 8},
		"test02": {"Alice gave $15 to Bob", 15},
		"test03": {"We got 5 files and 8 more", 58},
		"test04": {"Rank is 3,210 now", 3210},
	}

	for testName, testData := range testCases {
		t.Run(testName, func(t *testing.T) {
			result, err := intFromString(testData.testString)
			if err != nil {
				t.Errorf(`IntFromString("%+v") returned unexpected error: %+v`, testData.testString, err)
			}

			if result != testData.expected {
				t.Errorf(`IntToString("%+v") returned '%+v', expected '%+v'`, testData.testString, result, testData.expected)
			}
		})
	}
}
