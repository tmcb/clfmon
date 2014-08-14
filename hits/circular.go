package hits

import (
	"sync"
)

/*
Struct CircularHitCounter holds the computed hits between two time instants.
The difference between these two times is called a round.

Frames provide granularity to the round, storing the number of hits between the
first and the last rotation. For a round with n frames, it takes n+1 rotations
to get the same frame again. The frame to be used is restarted at every
rotation.
*/
type CircularHitCounter struct {
	sync.Mutex
	framesHits []uint
	current    uint
}

/*
Init fills a struct CircularHitCounter with basic values. It requires the
number of frames the round should have. It returns a pointer to the very
structure initialized.
*/
func (c *CircularHitCounter) Init(size uint) *CircularHitCounter {
	*c = CircularHitCounter{sync.Mutex{}, make([]uint, size), 0}
	return c
}

/*
Rotate goes to the next frame in the round. The already used frames are
restarted (set with zero again).
*/
func (c *CircularHitCounter) Rotate() {
	if len(c.framesHits) == 0 {
		return
	}
	c.Lock()
	defer c.Unlock()
	c.current += 1
	c.current %= uint(len(c.framesHits))
	c.framesHits[c.current] = 0
}

/*
FrameHits gets the number of hits in the current frame.
*/
func (c *CircularHitCounter) FrameHits() uint {
	return c.framesHits[c.current]
}

/*
Hits computes the total number of hits in the current round.
*/
func (c *CircularHitCounter) Hits() uint64 {
	hits := uint64(0)
	for _, v := range c.framesHits {
		hits += uint64(v)
	}
	return hits
}

/*
Hit stores a hit in the current frame.
*/
func (c *CircularHitCounter) Hit() {
	c.Lock()
	defer c.Unlock()
	c.framesHits[c.current] += 1
}
