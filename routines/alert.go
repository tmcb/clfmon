package routines

import (
	"github.com/tmcb/clfmon/control"
	"log"
	"sync"
	"time"
)

func checkAlert(alright bool, nHits uint64, hitLimit uint64) bool {
	if alright && nHits > hitLimit {
		alright = false
		log.Printf(
			"[ALERT] High traffic generated an alert (total hits: %d)\n",
			nHits)
	} else if !alright && nHits < hitLimit {
		alright = true
		log.Println("[ALERT] System recovered from the last alert")
	}
	return alright
}

/*
Alert executes the alert routine of the monitor. It checks the number of hits of
the control structure ctl according to the provided period, printing log
messages if the hit limit is crossed in any direction.

A wait group wg can be informed if desired.
*/
func Alert(period time.Duration, ctl *control.Control, wg ...*sync.WaitGroup) {
	if ctl == nil {
		log.Println("[ERROR] Alert routine has no control structure.")
		return
	}
	if wg != nil {
		defer wg[0].Done()
	}
	alright := true
	for {
		time.Sleep(period)
		alright = checkAlert(alright, ctl.RoundHits().Hits(),
			ctl.HitLimit())
	}
}
