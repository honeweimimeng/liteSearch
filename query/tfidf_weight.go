package query

import "math"

type TfIDFWeight struct {
	idf float64
}

func NewTFIDFWeight(totalDocNum uint64, documentFrequency int) *TfIDFWeight {
	return &TfIDFWeight{
		idf: 1.0 + math.Log(float64(totalDocNum)/float64(1+documentFrequency)),
	}
}

func (t *TfIDFWeight) Score(termFrequency float64) float64 {
	return termFrequency * t.idf
}
