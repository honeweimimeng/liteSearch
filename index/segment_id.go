package index

import "liteSearch/pkg/uuid"

type SegmentID string

func NewSegmentID() SegmentID {
	return SegmentID(uuid.Generate())
}
