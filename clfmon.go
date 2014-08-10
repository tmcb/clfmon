package main

import (
	"fmt"
	"github.com/tmcb/clfmon/control"
	"github.com/tmcb/clfmon/routines"
	"os"
	"strconv"
	"sync"
	"time"
)

const (
	alarmPeriod time.Duration = 2 * time.Minute
	infoPeriod  time.Duration = 10 * time.Second
)

func usage() {
	fmt.Fprintf(os.Stderr, "usage: %s [logpath] [hitlimit]\n", os.Args[0])
	os.Exit(1)
}

func main() {
	if len(os.Args) < 3 {
		usage()
	}
	logPath := os.Args[1]
	hitLimit, err := strconv.ParseUint(os.Args[2], 10, 0)
	if err != nil {
		usage()
	}
	ctl := (&control.Control{}).Init(uint(alarmPeriod/infoPeriod),
		hitLimit)
	wg := &sync.WaitGroup{}
	wg.Add(3)
	defer wg.Wait()
	go routines.Alert(infoPeriod, ctl, wg)
	go routines.Consume(logPath, ctl, wg)
	go routines.Info(infoPeriod, ctl, wg)
}
