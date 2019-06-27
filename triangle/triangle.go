// Package triangle compare side of triangles
package triangle

import "math"

// Kind mean 4 types of triangle
type Kind string

// kinds of triangles
const (
	NaT Kind = "not a triangle"
	Equ Kind = "equilateral"
	Iso Kind = "isosceles"
	Sca Kind = "scalene"
)

// KindFromSides implements check of 3 kinds of triangles
func KindFromSides(a, b, c float64) Kind {
	if isValidTriangle(a, b, c) {
		if isEqu(a, b, c) {
			return Equ
		} else if isIso(a, b, c) {
			return Iso
		} else if isSca(a, b, c) {
			return Sca
		}
	}

	return NaT
}

func isValidTriangle(a, b, c float64) bool {
	return isValidSide(a) && isValidSide(b) && isValidSide(c) && isValidSumSides(a, b, c)
}

func isValidSumSides(a, b, c float64) bool {
	return (a <= b+c) && (b <= a+c) && (c <= a+b)
}

func isValidSide(side float64) bool {
	return side > 0 && side != math.Inf(1)
}

func isEqu(a, b, c float64) bool {
	return a == b && b == c
}

func isIso(a, b, c float64) bool {
	return (a == b || a == c || b == c) && !isEqu(a, b, c)
}

func isSca(a, b, c float64) bool {
	return a != b && a != c && b != c
}
