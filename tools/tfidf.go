package tools

func ComputeNormalizeTFIDF(tf int, maxTf int, df int) float64{
	return float64(tf)/float64(maxTf) * (1.0/float64(df))
}
