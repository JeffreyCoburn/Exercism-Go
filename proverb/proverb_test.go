package proverb

import (
	"fmt"
	"testing"
)

func TestProverb(t *testing.T) {
	for _, test := range stringTestCases {
		actual := Proverb(test.input)
		if fmt.Sprintf("%q", actual) != fmt.Sprintf("%q", test.expected) {
			t.Fatalf("FAIL %s - Proverb test [%s]\n\texpected: [%s],\n\tactual:   [%s]",
				test.description, test.input, test.expected, actual)
		}
		t.Logf("PASS %s", test.description)
	}
}

// func TestProverb(t *testing.T) {
// 	t.Run("one line", func(t *testing.T) {
// 		got := Proverb([]string{"nail"})
// 		want := []string{"And all for the want of a nail."}

// 		if !reflect.DeepEqual(got, want) {
// 			t.Errorf("got %v want %v ", got, want)
// 		}
// 	})
// 	t.Run("empty", func(t *testing.T) {
// 		got := Proverb([]string{})
// 		want := []string{}

// 		if !reflect.DeepEqual(got, want) {
// 			t.Errorf("got %v want %v ", got, want)
// 		}
// 	})
// 	t.Run("multiple lines", func (t *testing.T) {
// 		got
// 	})
// }

func BenchmarkProverb(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range stringTestCases {
			Proverb(test.input)
		}
	}
}
