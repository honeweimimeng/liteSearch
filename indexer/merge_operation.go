package indexer

import (
	"liteSearch/index"
	"liteSearch/internal/opstamp"
)

type MergeOperation struct {
	targetOpStamp opstamp.OpStamp
	segmentIDs    []index.SegmentID
}

func NewMergeOperation(targetOpStamp opstamp.OpStamp, segmentIDs []index.SegmentID) *MergeOperation {
	return &MergeOperation{
		targetOpStamp: targetOpStamp,
		segmentIDs:    segmentIDs,
	}
}
