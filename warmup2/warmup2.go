// Exercises from codingbat.com ported to Golang.
package warmup2

import "strings"

// Given an slice of ints, return true if the sequence of numbers 1, 2, 3
// appears in the slice somewhere.
func Array123(s []int) bool {
	for a := range s {
		if len(s) >= a+3 &&
			1 == s[a] &&
			2 == s[a+1] &&
			3 == s[a+2] {
			return true
		}
	}
	return false
}

// Given a slice of ints, return the number of 9's in the slice.
func ArrayCount9(s []int) int {
	ret := 0
	for a := range s {
		if 9 == s[a] {
			ret++
		}
	}
	return ret
}

// Given a slice of ints, return true if one of the first 4 elements in the
// slice is a 9. The slice length may be less than 4.
func ArrayFront(s []int) bool {
	for a := range s {
		if a > 3 {
			break
		}
		if 9 == s[a] {
			return true
		}
	}
	return false
}

// Given a string and a non-negative int n, we'll say that the front of the
// string is the first 3 chars, or whatever is there if the string is less
// than length 3. Return n copies of the front
func FrontTimes(s string, n int) string {
	rpt := ""
	if len(s) <= 3 {
		rpt = s
	} else {
		rpt = s[:3]
	}
	return strings.Repeat(rpt, n)
}

// Given a string, return the count of the number of times that a substring length
// 2 appears in the string and also as the last 2 chars of the string, so
// "hixxxhi" yields 1 (we won't count the end substring).
func Last2(s string) int {
	if len(s) <= 2 {
		return 0
	}
	substr := s[len(s)-2:]
	cnt := 0
	for a := range s {
		if a < len(s)-1 {
			if substr[0] == s[a] && substr[1] == s[a+1] {
				cnt++
			}
		}
	}
	return cnt - 1
}

// Given a string, return a new string made of every other char starting with
// the first, so "Hello" yields "Hlo".
func StringBits(s string) string {
	var ret []byte
	for a := range s {
		if 1 == (a+1)%2 {
			ret = append(ret, s[a])
		}
	}
	return string(ret)
}

// Given 2 strings, a and b, return the number of the positions where they contain
// the same length 2 substring. So "xxcaazz" and "xxbaaz" yields 3, since the
// "xx", "aa", and "az" substrings appear in the same place in both strings.
func StringMatch(a, b string) int {
	cnt := 0
	shorter := a
	longer := b
	if len(a) > len(b) {
		shorter = b
		longer = a
	}
	for c := range shorter {
		if c+1 < len(shorter) && shorter[c] == longer[c] && shorter[c+1] == longer[c+1] {
			cnt++
		}
	}
	return cnt
}

// Given a non-empty string like "Code" return a string like "CCoCodCode".
func StringSplosion(s string) string {
	var ret []byte
	for a := range s {
		rpt := a + 1
		for b := 0; b < rpt; b++ {
			ret = append(ret, s[b])
		}
	}
	return string(ret)
}

// Given a string and a non-negative int n, return a larger string that is n
// copies of the original string.
func StringTimes(s string, n int) string {
	ret := ""
	for a := 0; a < n; a++ {
		ret += s
	}
	return ret
}
