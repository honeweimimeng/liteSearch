package index

import "liteSearch/schema"

type Document struct {
	// NOTE: commented out to pass lint, but will be used when we store original document
	// fields map[string]interface{}
}

type DocAddress struct {
	SegmentOrd int          `json:"SegmentOrd"`
	DocID      schema.DocID `json:"DocID"`
}

type DocSet interface {
	Advance() schema.DocID
	Doc() schema.DocID
	Seek(target schema.DocID) schema.DocID
	SizeHint() uint32
}

func SeekDocSet(docSet DocSet, target schema.DocID) schema.DocID {
	doc := docSet.Doc()
	for doc < target {
		doc = docSet.Advance()
	}
	return doc
}
