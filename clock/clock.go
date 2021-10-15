package clock

import "fmt"

// Clock API:
//
// type Clock                      // define the clock type
// New(hour, minute int) Clock     // a "constructor"
// (Clock) String() string         // a "stringer"
// (Clock) Add(minutes int) Clock
// (Clock) Subtract(minutes int) Clock
//
// To satisfy the README requirement about clocks being equal, values of
// your Clock type need to work with the == operator. This means that if your
// New function returns a pointer rather than a value, your clocks will
// probably not work with ==.
//
// While the time.Time type in the standard library (https://golang.org/pkg/time/#Time)
// doesn't necessarily need to be used as a basis for your Clock type, it might
// help to look at how constructors there (Date and Now) return values rather
// than pointers. Note also how most time.Time methods have value receivers
// rather than pointer receivers.
//
// For some useful guidelines on when to use a value receiver or a pointer
// receiver see: https://github.com/golang/go/wiki/CodeReviewComments#receiver-type

const minutesInDay = 24 * 60

type Clock struct {
	hour,
	minute int
}

func New(hour, minute int) Clock {
	minutes := hour*60 + minute
	hour, minute = getTime(minutes)
	return Clock{hour: hour, minute: minute}
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.hour, c.minute)
}

func (c Clock) Add(minutes int) Clock {
	minutes = c.hour*60 + c.minute + minutes
	hour, minute := getTime(minutes)
	return Clock{hour: hour, minute: minute}
}

func (c Clock) Subtract(minutes int) Clock {
	minutes = c.hour*60 + c.minute - minutes
	hour, minute := getTime(minutes)
	return Clock{hour: hour, minute: minute}
}

// getTime returns the hour and minute given the minutes since the start of the day
func getTime(minutes int) (int, int) {
	minutes = minutes % minutesInDay
	if minutes < 0 {
		minutes += minutesInDay
	}
	hour := minutes / 60
	minutes = minutes - hour*60
	return hour, minutes
}
