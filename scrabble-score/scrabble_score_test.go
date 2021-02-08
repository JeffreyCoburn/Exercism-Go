package scrabble

import "testing"

func TestScore(t *testing.T) {
	for _, test := range scrabbleScoreTests {
		if actual := Score(test.input); actual != test.expected {
			t.Errorf("Score(%q) expected %d, Actual %d", test.input, test.expected, actual)
		}
	}
}

func BenchmarkScore(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range scrabbleScoreTests {
			Score(test.input)
		}
	}
}

func TestExtendedScore(t *testing.T) {
	for _, test := range scrabbleExtendedScoreTests {
		actual, err := ExtendedScore(test.word, test.letterModifiers, test.wordModifier)
		if err != test.err {
			t.Errorf("Score(%v) with letter modifiers(%v) and word modifier(%v)\n"+
				"Expected error: %v\n"+
				"Actual error: %v",
				test.word, test.letterModifiers, test.wordModifier, test.err, err)
		}
		if actual != test.expected {
			t.Errorf("Score(%v) with letter modifiers(%v) and word modifier(%v) expected %v, Actual %v",
				test.word, test.letterModifiers, test.wordModifier, test.expected, actual)
		}
	}
}
