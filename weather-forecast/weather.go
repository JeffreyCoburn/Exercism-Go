// Package weather creates log entry strings based on a location and condition
package weather

var (
	// CurrentCondition contains the current weather conditions
	CurrentCondition string
	// CurrentLocation contains the current location
	CurrentLocation string
)

// Log returns a description based on the location and condition passed in
func Log(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
