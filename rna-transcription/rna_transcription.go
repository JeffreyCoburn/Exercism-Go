// Package strand returns an RNA strand given its DNA complement
package strand

var dnaToRna = map[rune]string{
	'G': "C", 'g': "C",
	'C': "G", 'c': "G",
	'T': "A", 't': "A",
	'A': "U", 'a': "U",
}

/*
ToRNA transcribes a given DNA strand to an RNA strand by replacing each nucleotide with its complement:
G -> C
C -> G
T -> A
A -> U
*/
func ToRNA(dna string) string {
	result := ""
	for _, char := range dna {
		result += dnaToRna[char]
	}
	return result
}
