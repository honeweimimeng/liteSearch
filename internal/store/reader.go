package store

import (
	"encoding/json"
	"fmt"
	"liteSearch/directory"
	"liteSearch/schema"
)

func ReadDocs(storeFile *directory.FileSlice) ([]*schema.Document, error) {
	var res []*schema.Document
	bytes := make([]byte, storeFile.Len())
	_, err := storeFile.Reader().Read(bytes)
	err = json.Unmarshal(bytes, &res)
	if err != nil {
		return res, fmt.Errorf("marshal document: %w", err)
	}
	return res, nil
}
