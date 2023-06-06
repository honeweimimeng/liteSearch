package store

import (
	"encoding/json"
	"fmt"
	"liteSearch/directory"
	"liteSearch/schema"
)

const BlockSize = 16_384

type StoreWriter struct {
	doc    int
	docs   []*schema.Document
	writer directory.WriteCloseSyncer
}

func NewStoreWriter(writer directory.WriteCloseSyncer) *StoreWriter {
	return &StoreWriter{
		writer: writer,
	}
}

// TODO: implement
func (s *StoreWriter) Store(document *schema.Document) error {
	s.docs = append(s.docs, document)
	s.doc++
	return nil
}

func (s *StoreWriter) Serialize() error {
	docJSON, err := json.Marshal(s.docs)
	if err != nil {
		return fmt.Errorf("marshal document: %w", err)
	}
	if _, err := s.writer.Write(docJSON); err != nil {
		return fmt.Errorf("write document: %w", err)
	}
	return nil
}

func (s *StoreWriter) Close() error {
	return s.writer.Close()
}
