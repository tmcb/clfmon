package control

import (
	"github.com/tmcb/clfmon/bytes"
	"github.com/tmcb/clfmon/hits"
)

/*
Control holds the control variables of a monitoring session.
*/
type Control struct {
	sectionHits *hits.SectionHitCounter
	roundHits   *hits.CircularHitCounter
	byteCount   *bytes.Bytes
	hitLimit    uint64
}

/*
Init fills a struct Control with basic values. It requires a round size, which
is the number of time frames a hit evaluation round should have, and a hit
limit. It returns a pointer to the very control structure initialized.
*/
func (c *Control) Init(roundSize uint, hitLimit uint64) *Control {
	*c = Control{
		(&hits.SectionHitCounter{}).Init(),
		(&hits.CircularHitCounter{}).Init(roundSize),
		&bytes.Bytes{},
		hitLimit}
	return c
}

/*
SectionHits provides the strucutre that controls hits by sections.
*/
func (c *Control) SectionHits() *hits.SectionHitCounter {
	return c.sectionHits
}

/*
RoundHits provides the structure that controls hits in a given time round.
*/
func (c *Control) RoundHits() *hits.CircularHitCounter {
	return c.roundHits
}

/*
ByteCount provides the structure that controls transferred bytes.
*/
func (c *Control) ByteCount() *bytes.Bytes {
	return c.byteCount
}

/*
HitLimit gives the limit over which alerts about excessive hits should be
emitted in the log.
*/
func (c *Control) HitLimit() uint64 {
	return c.hitLimit
}
