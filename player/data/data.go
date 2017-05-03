package data

import (
	"math"
)

type Data struct {
	PlayerName  string
	CodedBase   string
	CodedPoints string
	NumericBase int
	// FIXME: naming is not clear
	DecodedPoints []int
	// Base 10 actual player points
	Points float64
}

// CalcBase calculates the numeric base for the given codified string
func (d *Data) CalcBase() {
	d.NumericBase = len(d.CodedBase)
}

// Decode abstract representation to numeric, on the given numeric base
func (d *Data) Decode() {
	m := make(map[string]int)

	// Map the coded symbol to its integer
	for k, v := range d.CodedBase {
		m[string(v)] = k
	}

	// DecodedPoints represents the numeric values on the given base
	// For example, oF8 Fo
	// oF8 is a base 3 abstract representation, 012 is the numeric representation
	// so Fo is 10 in base 3
	for _, v := range d.CodedPoints {
		d.DecodedPoints = append(d.DecodedPoints, m[string(v)])
	}
}

// ToBase10 converts the decodedPoints in given base to base 10
func (d *Data) ToBase10() {
	// Convert to base 10 operations
	i := 0
	for distance := len(d.DecodedPoints) - 1; distance >= 0; distance-- {
		// Base to the power of distance times the digit "of the distance"
		d.Points += float64(d.DecodedPoints[i]) * math.Pow(float64(d.NumericBase), float64(distance))

		i++
	}
}

type SortByPoints []*Data

func (s SortByPoints) Len() int {
	return len(s)
}

func (s SortByPoints) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s SortByPoints) Less(i, j int) bool {
	return s[i].Points > s[j].Points
}
