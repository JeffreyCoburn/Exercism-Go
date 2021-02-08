package scrabble

// Source: exercism/problem-specifications
// Commit: 0d882ed scrabble-score: Apply new "input" policy
// Problem Specifications Version: 1.1.0

type scrabbleTest struct {
	input    string
	expected int
}

var scrabbleScoreTests = []scrabbleTest{
	{"a", 1},                           // lowercase letter
	{"A", 1},                           // uppercase letter
	{"f", 4},                           // valuable letter
	{"at", 2},                          // short word
	{"zoo", 12},                        // short, valuable word
	{"street", 6},                      // medium word
	{"quirky", 22},                     // medium, valuable word
	{"OxyphenButazone", 41},            // long, mixed-case word
	{"pinata", 8},                      // english-like word
	{"", 0},                            // empty input
	{"abcdefghijklmnopqrstuvwxyz", 87}, // entire alphabet available
	{"pinata0", 8},                     // contains digits
	{"pinata\n\t", 8},                  // contains special characters
}

type scrabbleExtendedTest struct {
	word            string
	letterModifiers []int
	wordModifier    int
	expected        int
	err             error
}

var scrabbleExtendedScoreTests = []scrabbleExtendedTest{
	{"", []int{}, single, 0, nil},                                 // empty input
	{"pinata", []int{2, 1, 1, 1, 1}, single, 0, ErrInvalidLength}, // different number of letters vs letter modifiers
	{"a", []int{4}, single, 0, ErrInvalidModifier},                // letter modifier that is too large
	{"a", []int{0}, single, 0, ErrInvalidModifier},                // letter modifier that is too small
	{"a", []int{1}, 4, 0, ErrInvalidModifier},                     // word modifier that is too large
	{"a", []int{1}, 0, 0, ErrInvalidModifier},                     // word modifier that is too small
	{"pinata", []int{2, 1, 1, 1, 1, 1}, single, 11, nil},          // lower case word with one double letter modifier and normal word modifier
	{"zOo", []int{1, 1, 1}, double, 24, nil},                      // mixed case double word
	{
		"abcdefghijklmnopqrstuvwxyz",
		[]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
		triple,
		261,
		nil,
	}, // entire alphabet available with triple word
}
