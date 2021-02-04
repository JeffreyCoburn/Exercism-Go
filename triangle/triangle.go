// Package triangle determines the type of a triangle given its three sides.
package triangle

import (
	"math"
)

// Kind is the type of triangle
type Kind int

// Types of triangles
const (
	NaT = iota // not a triangle
	Equ        // equilateral
	Iso        // isosceles
	Sca        // scalene
)

// KindFromSides returns the type of triangle given three sides
func KindFromSides(a, b, c float64) Kind {
	if !isTriangle(a, b, c) {
		return NaT
	}
	if isEquilateral(a, b, c) {
		return Equ
	}
	if isIsosceles(a, b, c) {
		return Iso
	}
	return Sca
}

func isTriangle(a, b, c float64) bool {
	// Verify all sides are numbers
	if math.IsNaN(a) || math.IsNaN(b) || math.IsNaN(c) {
		return false
	}
	// Verify no numbers are inifinity
	if math.IsInf(a, 0) || math.IsInf(b, 0) || math.IsInf(c, 0) {
		return false
	}
	// Verify all sides are greater than zero
	if a <= 0 || b <= 0 || c <= 0 {
		return false
	}
	// Verify the sum of the lengths of any two sides must be greater than or equal to the length of the third side
	if a+b < c || a+c < b || b+c < a {
		return false
	}
	return true
}

func isEquilateral(a, b, c float64) bool {
	if a == b && b == c {
		return true
	}
	return false
}

// An Equilateral triangle is also considered isosceles
func isIsosceles(a, b, c float64) bool {
	if a == b || b == c || a == c {
		return true
	}
	return false
}
