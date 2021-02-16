// Package twofer writes the message, "One for <name>, one for me."
package twofer

// ShareWith takes a name and returns a string with the message, "One for <name>, one for me."
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}

	return "One for " + name + ", one for me."
}
