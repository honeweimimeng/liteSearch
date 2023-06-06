package indexer

import "liteSearch/index"

type SegmentEntry struct {
	meta *index.SegmentMeta
}

func NewSegmentEntry(segmentMeta *index.SegmentMeta) *SegmentEntry {
	return &SegmentEntry{meta: segmentMeta}
}

func (s SegmentEntry) SegmentID() index.SegmentID {
	return s.meta.SegmentID
}
