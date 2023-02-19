package utils

import "math"

func TfIdf(tokenNum, docNum, totalToken, totalDoc int64) float64 {
	TF := float64(tokenNum) / float64(totalToken)
	IDF := math.Log2(float64(totalDoc) / float64(docNum+1))
	return TF * IDF
}
