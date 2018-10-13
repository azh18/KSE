package tools

import "math"

// The method of computing TF-IDF, invoked by other modules
func ComputeNormalizeTFIDF(tf int, maxTf int, df int, nDoc int) float64{
	return float64(tf)/float64(maxTf) * math.Log2(float64(nDoc)/float64(df))
}
