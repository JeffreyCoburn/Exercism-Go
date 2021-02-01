// Exercism second exercise
package twofer

// Given a name, returns a string with the message, "One for <name>, one for me."
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}

	return "One for " + name + ", one for me."
}
