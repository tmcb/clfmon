package bytes

import (
	"sync"
)

/*
Struct Bytes holds amounts of bytes informed as transmitted.

It holds the total amount informed, as well as the amount informed during the
last time frame.

The time frame counter can be resetted.
*/
type Bytes struct {
	sync.Mutex
	frameBytes uint64
	totalBytes uint64
}

/*
ResetFrame throws away all info stored about the current time frame, so it is
restarted afresh.
*/
func (b *Bytes) ResetFrame() {
	b.Lock()
	defer b.Unlock()
	b.frameBytes = 0
}

/*
FrameBytes returns the number of bytes transmitted during the current time
frame.
*/
func (b *Bytes) FrameBytes() uint64 {
	return b.frameBytes
}

/*
TotalBytes returns the total number of bytes transmitted.
*/
func (b *Bytes) TotalBytes() uint64 {
	return b.totalBytes
}

/*
Add sums the informed number of bytes both to the time frame and to the total
counters.
*/
func (b *Bytes) Add(bytes uint64) {
	b.Lock()
	defer b.Unlock()
	b.frameBytes += bytes
	b.totalBytes += bytes
}
