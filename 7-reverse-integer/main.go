package main

import "math"

/*
https://leetcode.com/problems/reverse-integer/

Given a 32-bit signed integer, reverse digits of an integer.

Examples:
Input: 123
Output: 321

Input: -123
Output: -321

Input: 120
Output: 21

Note: Assume we are dealing with an environment which could only store integers within the 32-bit signed integer range: [−231,  231 − 1].
For the purpose of this problem, assume that your function returns 0 when the reversed integer overflows.
*/
func reverse(x int) int {
	reversedValue := 0

	for x != 0 {
		reversedValue = reversedValue*10 + x%10
		x /= 10
	}

	if reversedValue < math.MinInt32 || reversedValue > math.MaxInt32 {
		reversedValue = 0
	}

	return reversedValue
}
