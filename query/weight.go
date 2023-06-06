package query

import (
	"liteSearch/reader"
	"liteSearch/schema"
)

func ForEach(scorer reader.Scorer, callback func(docID schema.DocID, score float64)) error {
	doc := scorer.Doc()
	for !doc.IsTerminated() {
		callback(doc, scorer.Score())
		doc = scorer.Advance()
	}
	return nil
}

func ForEachPruning(scorer reader.Scorer, threshold float64, callback func(docID schema.DocID, score float64) float64) error {
	doc := scorer.Doc()
	for doc.IsTerminated() {
		if score := scorer.Score(); score > threshold {
			threshold = callback(doc, threshold)
		}
		doc = scorer.Advance()
	}
	return nil
}
