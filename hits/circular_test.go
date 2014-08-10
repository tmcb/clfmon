package hits

import (
	"testing"
)

func TestCircularBuffer(t *testing.T) {
	const size = uint(10)
	c := (&CircularHitCounter{}).Init(size)
	for i := uint64(0); i < uint64(size); i++ {
		if c.Hits() != i {
			t.Fatal()
		}
		c.Hit()
		if c.FrameHits() != 1 {
			t.Fatal()
		}
		c.Rotate()
	}
	for i := uint64(size - 1); i > 0; i-- {
		if c.Hits() != i {
			t.Fatal()
		}
		if c.FrameHits() != 0 {
			t.Fatal()
		}
		c.Rotate()
	}
	if c.Hits() != 0 {
		t.Fatal()
	}
}
