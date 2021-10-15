package clock

import "fmt"

const minutesInDay = 24 * 60

// Clock represents a clock that shows the time in hours and minutes
type Clock struct {
	hour,
	minute int
}

// New creates a Clock given the hours and minutes
func New(hour, minute int) Clock {
	minutes := hour*60 + minute
	minutes = minutes % minutesInDay
	if minutes < 0 {
		minutes += minutesInDay
	}
	hour = minutes / 60
	minute = minutes - hour*60
	return Clock{hour: hour, minute: minute}
}

// String returns the time the Clock is set to in the format "hh:mm"
func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.hour, c.minute)
}

// Add sets the time of a clock by adding the given number of minutes
func (c Clock) Add(minutes int) Clock {
	return New(c.hour, c.minute+minutes)
}

// Subtract sets the time of a clock by subtracting the given number of minutes
func (c Clock) Subtract(minutes int) Clock {
	return New(c.hour, c.minute-minutes)
}
