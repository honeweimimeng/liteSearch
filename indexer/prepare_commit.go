package indexer

import "liteSearch/internal/opstamp"

type PrepareCommit struct {
	indexWriter *IndexWriter
	payload     string
	opStamp     opstamp.OpStamp
}

func NewPrepareCommit(indexWriter *IndexWriter, opStamp opstamp.OpStamp) *PrepareCommit {
	return &PrepareCommit{
		indexWriter: indexWriter,
		payload:     "",
		opStamp:     opStamp,
	}
}
