// Package leap reports if a given year is a leap year
package leap

// IsLeapYear returns true if the given year is a leap year and false otherwise
func IsLeapYear(year int) bool {
	// if year is divisible by 4, but not by 100 unless it's also divisible by 400
	if (year%4 == 0) && (year%100 != 0 || year%400 == 0) {
		return true
	}
	return false
}
