package indexer

import (
	"liteSearch/internal/opstamp"
	"liteSearch/schema"
)

type AddOperation struct {
	opstamp  opstamp.OpStamp
	document *schema.Document
	result   func(error)
}
