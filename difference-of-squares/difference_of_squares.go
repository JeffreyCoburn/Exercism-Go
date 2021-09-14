package diffsquares

// SumOfSquares returns the sum of the squares of the first N natural numbers: 1^2 + 2^2 + ... + n^2
func SumOfSquares(n int) int {
	// 1^2 + 2^2 + 3^2 ... + n^2 = n(n+1)(2n+1)/6
	return n * (n + 1) * (2*n + 1) / 6
}

// SquareOfSum returns the square of the sums first N natural numbers: (1+2+...+n)^2
func SquareOfSum(n int) int {
	sum := (n * (n + 1)) / 2
	return sum * sum
}

// Difference calculates the difference between the square of the sum and the sum of the squares of the first N natural numbers
func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
