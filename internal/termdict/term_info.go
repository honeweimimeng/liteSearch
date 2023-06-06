package termdict

import "liteSearch/pkg/xrange"

type TermInfo struct {
	Term          string
	DocFreq       int
	PostingsRange xrange.Range
}
