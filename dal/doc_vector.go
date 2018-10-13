package dal

import "math"

type DocVector map[string]float64

// Compute magnitude for a Document Vector
func (d DocVector) GetMagnitude() float64{
	sumSquare := 0.0
	for k := range d{
		sumSquare += d[k]*d[k]
	}
	return math.Sqrt(sumSquare)
}