package logic2

import "math"

// We want to make a row of bricks that is goal inches long. We have a number of
// small bricks (1 inch each) and big bricks (5 inches each). Return true if it is
// possible to make the goal by choosing from the given bricks. This is a little harder
// than it looks and can be done without any loops.
// http://codingbat.com/doc/practice/makebricks-introduction.html
func MakeBricks(brick1, brick5, goal int) bool {
	var brick1N int
	var brick5N int

	brick5N = goal / 5
	if brick5 < brick5N {
		brick5N = brick5
	}

	brick1N = goal - (brick5N * 5)
	if brick1 >= brick1N {
		return true
	}
	return false
}

// Given 3 int values, a b c, return their sum. However, if any of the values is
// a teen -- in the range 13..19 inclusive -- then that value counts as 0, except 15 and 16
// do not count as a teens. Write a separate helper "def fixteen(n):"that takes in an
// int value and returns that value fixed for the teen rule. In this way, you avoid
// repeating the teen code 3 times (i.e. "decomposition"). Define the helper below
// and at the same indent level as the main NoTeenSum().
func FixTeen(n int) int {
	if n >= 13 && n <= 19 && n != 15 && n != 16 {
		n = 0
	}
	return n
}
func NoTeenSum(a, b, c int) int {
	return FixTeen(a) + FixTeen(b) + FixTeen(c)
}

// We want make a package of goal kilos of chocolate. We have small bars (1 kilo each)
// and big bars (5 kilos each). Return the number of small bars to use, assuming we
// always use big bars before small bars. Return -1 if it can't be done.
func MakeChocolate(bar1, bar5, goal int) int {
	var bar1N int
	var bar5N int

	bar5N = goal / 5
	if bar5 < bar5N {
		bar5N = bar5
	}

	bar1N = goal - (bar5N * 5)
	if bar1 >= bar1N {
		return bar1N
	}
	return -1
}

// Given 3 int values, a b c, return their sum. However, if one of the values is
// the same as another of the values, it does not count towards the sum.
func LoneSum(a, b, c int) int {
	sum := 0
	if a != b && a != c {
		sum += a
	}
	if b != a && b != c {
		sum += b
	}
	if a != c && b != c {
		sum += c
	}
	return sum
}

// For this problem, we'll round an int value up to the next multiple of 10 if its
// rightmost digit is 5 or more, so 15 rounds up to 20. Alternately, round down to
// the previous multiple of 10 if its rightmost digit is less than 5, so 12 rounds
// down to 10. Given 3 ints, a b c, return the sum of their rounded values. To avoid
// code repetition, write a separate helper "def round10(num):" and call it 3 times.
// Write the helper entirely below and at the same indent level as RoundSum().
func RoundSum(a, b, c int) int {
	return roundSumHelper(a) + roundSumHelper(b) + roundSumHelper(c)
}
func roundSumHelper(n int) int {
	if n%10 >= 5 {
		return n + (10 - (n % 10))
	}
	return n - (n % 10)
}

// Given 3 int values, a b c, return their sum. However, if one of the values is 13
// then it does not count towards the sum and values to its right do not count. So
// for example, if b is 13, then both b and c do not count.
func LuckySum(a, b, c int) int {
	var sum int

	if 13 == a {
		return sum
	}
	sum += a

	if 13 == b {
		return sum
	}
	sum += b

	if 13 == c {
		return sum
	}
	sum += c

	return sum
}

// Given three ints, a b c, return true if one of b or c is "close" (differing from
// a by at most 1), while the other is "far", differing from BOTH other values by 2
// or more.
func CloseFar(a, b, c int) bool {
	a2b := math.Abs(float64(a - b))
	a2c := math.Abs(float64(a - c))
	b2c := math.Abs(float64(b - c))

	if 1 >= a2b {
		if 2 <= a2c && 2 <= b2c {
			return true
		}
	} else if 1 >= a2c {
		if 2 <= a2b && 2 <= b2c {
			return true
		}
	}
	return false
}
