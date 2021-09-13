package lasagna

// Returns how many minutes the lasagna should be in the oven
func OvenTime() int {
	return 40
}

// Takes the actual minutes the lasagna has been in the oven as a parameter and returns how many minutes the
// lasagna still has to remain in the oven
func RemainingOvenTime(elapsedOvenTime int) int {
	return OvenTime() - elapsedOvenTime
}

// Takes the number of layers you added to the lasagna as a parameter and returns how many minutes you spent
// preparing the lasagna, assuming each layer takes you 2 minutes to prepare.
func PreparationTime(layers int) int {
	return layers * 2
}

// Takes the number of layers you added to the lasagna, and the number of minutes the lasagna has been in the oven.
// The function returns how many minutes in total you've worked on cooking the lasagna, which is the sum of the
// preparation time in minutes, and the time in minutes the lasagna has spent in the oven at the moment.
func ElapsedTime(layers, elapsedOvenTime int) int {
	return PreparationTime(layers) + elapsedOvenTime
}
