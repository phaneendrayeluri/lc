package main

import (
	"fmt"
	"testing"
)

func TestPathSum(t *testing.T) {

	testScenarios := []struct {
		input     []int
		targetSum int
		expected  int
	}{
		{input: []int{10, 5, -3, 3, 2, 0, 11, 3, -2, 0, 1}, targetSum: 8, expected: 3},
		{input: []int{5, 4, 8, 11, 0, 13, 4, 7, 2, 0, 0, 5, 1}, targetSum: 22, expected: 3},
	}

	for idx, test := range testScenarios {
		t.Run(fmt.Sprintf("test at index %d", idx), func(t *testing.T) {
			if test.expected != PathSum(test.input, test.targetSum) {
				t.Fail()
			}
		})
	}

}
