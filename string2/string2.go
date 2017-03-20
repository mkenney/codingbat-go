package string2

import "strings"

// Given a string, return a string where for every char in the original, there are two chars.
func DoubleChar(s string) string {
	var bytes []byte
	for a := range s {
		bytes = append(bytes, s[a], s[a])
	}
	return string(bytes)
}

// Return the number of times that the string "code" appears anywhere in the
// given string, except we'll accept any letter for the "d", so "cope" and "cooe" count.
// Try not to use regex
func CountCode(s string) int {
	var cnt int
	for a := range s {
		if len(s) >= a+4 &&
			'c' == s[a] &&
			'o' == s[a+1] &&
			'e' == s[a+3] {
			cnt++
		}
	}
	return cnt
}

// Return the number of times that the string "hi" appears anywhere in the given string.
func CountHi(s string) int {
	var cnt int
	for a := range s {
		if len(s) >= a+2 &&
			'h' == s[a] &&
			'i' == s[a+1] {
			cnt++
		}
	}
	return cnt
}

// Given two strings, return true if either of the strings appears at the very end
// of the other string, ignoring upper/lower case differences (in other words, the
// computation should not be "case sensitive"). Note: s.lower() returns the
// lowercase version of a string.
func EndOther(a, b string) bool {
	a = strings.ToLower(a)
	b = strings.ToLower(b)
	if len(a) >= len(b) {
		return a[len(a)-len(b):] == b[:]
	}
	return b[len(b)-len(a):] == a[:]
}

// Return true if the string "cat" and "dog" appear the same number of times
// in the given string.
func CatDog(s string) bool {
	var cat, dog int
	for a := range s {
		if len(s) >= a+3 {
			if 'c' == s[a] &&
				'a' == s[a+1] &&
				't' == s[a+2] {
				cat++
			}
			if 'd' == s[a] &&
				'o' == s[a+1] &&
				'g' == s[a+2] {
				dog++
			}
		}
	}
	return cat == dog
}

// Return true if the given string contains an appearance of "xyz" where the xyz
// is not directly preceeded by a period (.). So "xxyz" counts but "x.xyz" does not.
func XyzThere(s string) bool {
	for a := range s {
		if 0 == a &&
			len(s) >= a+3 &&
			'x' == s[a] &&
			'y' == s[a+1] &&
			'z' == s[a+2] {
			return true
		}
		if len(s) >= a+4 &&
			'.' != s[a] &&
			'x' == s[a+1] &&
			'y' == s[a+2] &&
			'z' == s[a+3] {
			return true
		}
	}
	return false
}
