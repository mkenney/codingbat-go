package list2

import "sort"

// Return the number of even ints in the given array. Note: the % "mod" operator
// computes the remainder, e.g. 5 % 2 is 1.
func CountEvens(a []int) int {
	var cnt int
	for _, b := range a {
		if 0 == b%2 {
			cnt++
		}
	}
	return cnt
}

// Return the sum of the numbers in the array, returning 0 for an empty array.
// Except the number 13 is very unlucky, so it does not count and numbers that
// come immediately after a 13 also do not count.
func Sum13(a []int) int {
	var sum int
	var lastWas13 bool
	for _, b := range a {
		if 13 != b && !lastWas13 {
			sum += b
		}
		if 13 == b {
			lastWas13 = true
		} else {
			lastWas13 = false
		}
	}
	return sum
}

// Given an array length 1 or more of ints, return the difference between the
// largest and smallest values in the array.
func BigDiff(a []int) int {
	var biggest, smallest int
	for _, b := range a {
		if b > biggest {
			biggest = b
		}
	}
	smallest = biggest
	for _, b := range a {
		if b < smallest {
			smallest = b
		}
	}
	return biggest - smallest
}

// Return the sum of the numbers in the array, except ignore sections of numbers
//starting with a 6 and extending to the next 7 (every 6 will be followed by at
// least one 7). Return 0 for no numbers.
func Sum67(a []int) int {
	var sum int
	var skip bool
	for _, b := range a {
		if 6 == b {
			skip = true
		}
		if !skip {
			sum += b
		}
		if 7 == b {
			skip = false
		}
	}
	return sum
}

// Return the "centered" average of an array of ints, which we'll say is the mean
// average of the values, except ignoring the largest and smallest values in the array.
// If there are multiple copies of the smallest value, ignore just one copy, and
// likewise for the largest value. Use int division to produce the final average.
// You may assume that the array is length 3 or more.
func CenteredAverage(a []int) int {
	var sum int
	sort.Ints(a)
	for b := 1; b < len(a)-1; b++ {
		sum += a[b]
	}
	return sum / (len(a) - 2)
}

// Given an array of ints, return true if the array contains a 2 next to a 2 somewhere.
func Has22(a []int) bool {
	var had2 bool
	for _, b := range a {
		if 2 == b && had2 {
			return true
		}
		had2 = 2 == b
	}
	return false
}
