package hits

import (
	"errors"
	"sync"
)

/*
SectionHitCounter holds the number of hits a section had in two different ways:
its total since the start of monitoring, and the amount of hits in the last time
frame.

The time frame counter can be resetted.
*/
type SectionHitCounter struct {
	sync.Mutex
	frameHits map[string]uint
	totalHits map[string]uint
}

func hit(hitsMap map[string]uint, section string) {
	if _, ok := hitsMap[section]; ok {
		hitsMap[section] += 1
	} else {
		hitsMap[section] = 1
	}
}

/*
Init fills a struct SectionHitCounter with basic values. It returns a pointer
to the very structure initialized.
*/
func (c *SectionHitCounter) Init() *SectionHitCounter {
	*c = SectionHitCounter{sync.Mutex{}, make(map[string]uint),
		make(map[string]uint)}
	return c
}

/*
ResetFrame throws away all info stored about the current frame, so it is
restarted afresh.
*/
func (c *SectionHitCounter) ResetFrame() {
	c.Lock()
	defer c.Unlock()
	c.frameHits = make(map[string]uint)
}

/*
TotalHits returns the total number of hits for a given section.
*/
func (c *SectionHitCounter) TotalHits(section string) (hits uint, ok bool) {
	hits, ok = c.totalHits[section]
	return
}

/*
FrameGreatestHit returns the section with the most hits in the current time
frame, together with the number of hits. It also returns an error indicator,
valued nil in case of success, and non-nil otherwise.
*/
func (c *SectionHitCounter) FrameGreatestHit() (string, uint, error) {
	if len(c.frameHits) == 0 {
		return "", 0, errors.New("no hits")
	}
	var maxHits = uint(0)
	var greatestHit string
	for k, v := range c.frameHits {
		if v > maxHits {
			greatestHit = k
			maxHits = v
		}
	}
	return greatestHit, c.frameHits[greatestHit], nil
}

/*
Hit adds a hit for the given section both in the frame and in total counters.
*/
func (c *SectionHitCounter) Hit(section string) {
	c.Lock()
	defer c.Unlock()
	hit(c.frameHits, section)
	hit(c.totalHits, section)
}
