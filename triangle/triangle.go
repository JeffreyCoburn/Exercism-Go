// Package triangle determines the type of a triangle given its three sides.
package triangle

import (
	"errors"
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
	t, err := newTriangle(a, b, c)
	if err != nil {
		return NaT
	}
	if isEquilateral(t) {
		return Equ
	}
	if isIsosceles(t) {
		return Iso
	}
	return Sca
}

func isEquilateral(t triangle) bool {
	if t.a == t.b && t.b == t.c {
		return true
	}
	return false
}

func isIsosceles(t triangle) bool {
	if !isEquilateral(t) && t.a == t.b || t.b == t.c || t.a == t.c {
		return true
	}
	return false
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

type triangle struct {
	a, b, c float64
}

func newTriangle(a, b, c float64) (triangle, error) {
	if !isTriangle(a, b, c) {
		return triangle{0, 0, 0}, errInvalidSides
	}
	return triangle{a, b, c}, nil
}

var errInvalidSides = errors.New("Provided sides are not a triangle")
