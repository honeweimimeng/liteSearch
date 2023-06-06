package postings

import (
	"liteSearch/directory"
	"liteSearch/index"
	"liteSearch/internal/termdict"
	"liteSearch/schema"
)

type InvertedIndexSerializer struct {
	termsWrite    *termdict.TermWriter
	postingsWrite directory.WriteCloseSyncer
	schema        *schema.Schema
}

func NewInvertedIndexSerializer(segment *index.Segment) (*InvertedIndexSerializer, error) {
	termsWrite, err := segment.OpenWrite(index.SegmentComponentTerms)
	if err != nil {
		return nil, err
	}
	postingsWrite, err := segment.OpenWrite(index.SegmentComponentPostings)
	if err != nil {
		return nil, err
	}
	return &InvertedIndexSerializer{
		termsWrite:    termdict.NewTermWriter(termsWrite),
		postingsWrite: postingsWrite,
		schema:        segment.Schema(),
	}, nil
}

func (i *InvertedIndexSerializer) Close() error {
	if err := i.termsWrite.Close(); err != nil {
		return err
	}
	if err := i.postingsWrite.Close(); err != nil {
		return err
	}
	return nil
}
