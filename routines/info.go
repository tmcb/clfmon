package routines

import (
	"github.com/tmcb/clfmon/bytes"
	"github.com/tmcb/clfmon/control"
	"github.com/tmcb/clfmon/hits"
	"log"
	"sync"
	"time"
)

func sectionInfo(sectionHits *hits.SectionHitCounter) {
	greatestHit, frameHits, err := sectionHits.FrameGreatestHit()
	if err != nil {
		return
	}
	totalHits, _ := sectionHits.TotalHits(greatestHit)
	log.Printf(
		"[INFO] Section \"%s\" had %d hits in the last time frame (%d total)\n",
		greatestHit, frameHits, totalHits)
}

func totalInfo(roundHits *hits.CircularHitCounter) {
	log.Printf("[INFO] Overall hits: %d last frame, %d last round\n",
		roundHits.FrameHits(), roundHits.Hits())
}

func bytesInfo(byteCount *bytes.Bytes) {
	log.Printf("[INFO] %d KiB last frame, %d KiB last round\n",
		byteCount.FrameBytes()/1024, byteCount.TotalBytes()/1024)
}

/*
Info executes the information routine of the monitor for a control structure
ctl. From period to period it prints a log about:

- The section with most hits in the period;

- The overall hits in the last frame and the last round (see
hits.CircularHitFrame);

- The total of bytes transmitted in the last frame and the last round (see
bytes.Bytes).

A wait group wg can be informed if desired.
*/
func Info(period time.Duration, ctl *control.Control, wg ...*sync.WaitGroup) {
	if ctl == nil {
		log.Println("[ERROR] Info routine needs a control structure.")
		return
	}
	if wg != nil {
		defer wg[0].Done()
	}
	for {
		time.Sleep(period)
		sectionInfo(ctl.SectionHits())
		totalInfo(ctl.RoundHits())
		bytesInfo(ctl.ByteCount())
		ctl.SectionHits().ResetFrame()
		ctl.RoundHits().Rotate()
		ctl.ByteCount().ResetFrame()
	}
}
