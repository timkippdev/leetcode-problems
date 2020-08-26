package main

import "testing"

func TestTwoSum(t *testing.T) {
	result := twoSum([]int{2, 11, 7, 15}, 9)
	if result[0] != 0 || result[1] != 2 {
		t.Fail()
	}
}
