package reader

import (
	"fmt"
	"liteSearch/internal/postings"
	store2 "liteSearch/internal/store"
	"liteSearch/internal/termdict"

	"liteSearch/directory"
	"liteSearch/index"
	"liteSearch/schema"
)

type SegmentReader struct {
	SegmentID index.SegmentID
	MaxDoc    schema.DocID

	perFieldTermDict map[schema.FieldID]termdict.TermDict
	stores           []*schema.Document
	storeFile        *directory.FileSlice
	postingsFile     *directory.FileSlice

	schema *schema.Schema
}

func NewSegmentReader(segment *index.Segment) (*SegmentReader, error) {
	termdictFile, err := segment.OpenRead(index.SegmentComponentTerms)
	if err != nil {
		return nil, fmt.Errorf("open termdict file: %w", err)
	}
	perFieldTermDict, err := termdict.ReadPerFieldTermDict(termdictFile)
	if err != nil {
		return nil, fmt.Errorf("read termdict: %w", err)
	}
	// TODO: make it possible to read document
	storeFile, err := segment.OpenRead(index.SegmentComponentStore)
	if err != nil {
		return nil, fmt.Errorf("open store file: %w", err)
	}
	store, err := store2.ReadDocs(storeFile)
	if err != nil {
		return nil, fmt.Errorf("read store: %w", err)
	}
	postingsFile, err := segment.OpenRead(index.SegmentComponentPostings)
	if err != nil {
		return nil, fmt.Errorf("open positings file: %w", err)
	}

	return &SegmentReader{
		SegmentID:        segment.Meta().SegmentID,
		MaxDoc:           segment.Meta().MaxDoc,
		perFieldTermDict: perFieldTermDict,
		stores:           store,
		storeFile:        storeFile,
		postingsFile:     postingsFile,
		schema:           segment.Schema(),
	}, nil
}

func (s *SegmentReader) InvertedIndex(fieldID schema.FieldID) *postings.InvertedIndexReader {
	return postings.NewInvertedIndexReader(s.perFieldTermDict[fieldID], s.postingsFile)
}

func (s *SegmentReader) GetDoc(id schema.DocID) *schema.Document {
	return s.stores[id]
}

func (s *SegmentReader) DocNum() uint32 {
	return uint32(s.MaxDoc)
}
