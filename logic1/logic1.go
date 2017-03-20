package logic1

// When squirrels get together for a party, they like to have cigars. A squirrel
// party is successful when the number of cigars is between 40 and 60, inclusive.
// Unless it is the weekend, in which case there is no upper bound on the number
// of cigars. Return true if the party with the given values is successful, or false otherwise.
func CigarParty(cigars int, weekend bool) bool {
	return (40 <= cigars && weekend) || (40 <= cigars && 60 >= cigars)
}

// You are driving a little too fast, and a police officer stops you. Write code
// to compute the result, encoded as an int value: 0=no ticket, 1=small ticket, 2=big ticket.
// If speed is 60 or less, the result is 0. If speed is between 61 and 80 inclusive,
// the result is 1. If speed is 81 or more, the result is 2. Unless it is your
// birthday -- on that day, your speed can be 5 higher in all cases.
func CaughtSpeeding(speed int, isBirthday bool) int {
	bdayLeeway := 0
	result := 0
	if isBirthday {
		bdayLeeway = 5
	}
	if speed <= 60+bdayLeeway {
		result = 0
	} else if speed <= 80+bdayLeeway {
		result = 1
	} else {
		result = 2
	}

	return result
}

// The number 6 is a truly great number. Given two int values, a and b, return true
// if either one is 6. Or if their sum or difference is 6.
func Love6(a, b int) bool {
	return (6 == a ||
		6 == b ||
		6 == a+b ||
		6 == a-b ||
		6 == b-a)
}

// You and your date are trying to get a table at a restaurant. The parameter "you"
// is the stylishness of your clothes, in the range 0..10, and "date" is the stylishness
// of your date's clothes. The result getting the table is encoded as an int value
// with 0=no, 1=maybe, 2=yes. If either of you is very stylish, 8 or more, then the
// result is 2 (yes). With the exception that if either of you has style of 2 or less,
// then the result is 0 (no). Otherwise the result is 1 (maybe).
func DateFashion(you, date int) int {
	result := 0
	if 2 >= you || 2 >= date {
		result = 0
	} else if 8 > you && 8 > date {
		result = 1
	} else {
		result = 2
	}
	return result
}

// Given 2 ints, a and b, return their sum. However, sums in the range 10..19 inclusive,
// are forbidden, so in that case just return 20.
func SortaSum(a, b int) int {
	sum := a + b
	if 10 <= sum && 19 >= sum {
		sum = 20
	}
	return sum
}

// Given a number n, return true if n is in the range 1..10, inclusive. Unless
// "outsideMode" is true, in which case return true if the number is less or equal
// to 1, or greater or equal to 10.
func In1To10(n int, outsideMode bool) bool {
	ret := false
	if !outsideMode && 1 <= n && 10 >= n {
		ret = true
	} else if outsideMode && (1 >= n || 10 <= n) {
		ret = true
	}
	return ret
}

// The squirrels in Palo Alto spend most of the day playing. In particular, they
// play if the temperature is between 60 and 90 (inclusive). Unless it is summer,
// then the upper limit is 100 instead of 90. Given an int temperature and a boolean
// isSummer, return true if the squirrels play and false otherwise.
func SquirrelPlay(temp int, isSummer bool) bool {
	maxTemp := 90
	if isSummer {
		maxTemp = 100
	}
	return (temp >= 60 && temp <= maxTemp)
}

// Given a day of the week encoded as 0=Sun, 1=Mon, 2=Tue, ...6=Sat, and a boolean
// indicating if we are on vacation, return a string of the form "7:00" indicating
// when the alarm clock should ring. Weekdays, the alarm should be "7:00" and on the
// weekend it should be "10:00". Unless we are on vacation -- then on weekdays it
// should be "10:00" and weekends it should be "off".
func AlarmClock(day int, isVacation bool) string {
	weekend := false
	weekdayAlarm := "7:00"
	weekendAlarm := "10:00"
	if isVacation {
		weekdayAlarm = "10:00"
		weekendAlarm = "off"
	}
	if 0 == day || 6 == day {
		weekend = true
	}
	if weekend {
		return weekendAlarm
	}
	return weekdayAlarm
}

// Given a non-negative number "num", return true if num is within 2 of a multiple of 10.
// Note: (a % b) is the remainder of dividing a by b, so (7 % 5), is 2. See also:
// Introduction to Mod (http://codingbat.com/doc/practice/mod-introduction.html)
func NearTen(num int) bool {
	return (2 >= 10-(num%10) || 2 >= num%10)
}
