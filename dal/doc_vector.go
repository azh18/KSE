package dal

import "math"

type DocVector map[string]float64

// todo: not used
func (d DocVector) GetCosineSimilarity(a DocVector) float64{
	return 0.0
}

func (d DocVector) GetMagnitude() float64{
	sumSquare := 0.0
	for k := range d{
		sumSquare += d[k]
	}
	return math.Sqrt(sumSquare)
}