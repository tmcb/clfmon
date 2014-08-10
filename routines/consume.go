package routines

import (
	"errors"
	"github.com/ActiveState/tail"
	"github.com/tmcb/clfmon/bytes"
	"github.com/tmcb/clfmon/clf"
	"github.com/tmcb/clfmon/control"
	"github.com/tmcb/clfmon/hits"
	"log"
	"os"
	"strings"
	"sync"
)

func section(s string) (string, error) {
	sectionSplit := strings.Split(s, "/")
	if len(sectionSplit) < 2 {
		return "", errors.New("invalid resource")
	}
	return strings.Join(sectionSplit[:2], "/"), nil

}

func consumeSectionHit(entry clf.Entry, sectionHits *hits.SectionHitCounter) {
	if section, err := section(entry.Resource); err == nil {
		sectionHits.Hit(section)
	}
}

func consumeRoundHit(roundHits *hits.CircularHitCounter) {
	roundHits.Hit()
}

func consumeBytes(entry clf.Entry, byteCount *bytes.Bytes) {
	byteCount.Add(entry.Bytes)
}

func consume(line string,
	roundHits *hits.CircularHitCounter,
	sectionHits *hits.SectionHitCounter, byteCount *bytes.Bytes) {
	entry, err := clf.Parse(line)
	if err != nil {
		log.Println("[ERROR] Unable to parse an entry:", err.Error())
	}
	if entry.Method == "GET" && strings.HasPrefix(entry.Protocol, "HTTP") {
		consumeSectionHit(entry, sectionHits)
	}
	consumeRoundHit(roundHits)
	consumeBytes(entry, byteCount)
}

/*
Consume executes the conspumption routine of the monitor. It consumes lines from
a log file specified by logPath and then updates the control structure ctl.

A wait group wg can be informed if desired.
*/
func Consume(logPath string, ctl *control.Control, wg ...*sync.WaitGroup) {
	if ctl == nil {
		log.Println("[ERROR] Consume routine has no control structure.")
		return
	}
	if wg != nil {
		defer wg[0].Done()
	}
	stream, err := tail.TailFile(logPath, tail.Config{
		Follow:   true,
		ReOpen:   true,
		Location: &tail.SeekInfo{Offset: 0, Whence: os.SEEK_END}})
	if err != nil {
		log.Printf("[ERROR] Unable to open log file \"%s\"", logPath)
		return
	}
	for line := range stream.Lines {
		if line.Err != nil {
			log.Println("[ERROR] Unable to get last line")
		}
		consume(line.Text, ctl.RoundHits(), ctl.SectionHits(),
			ctl.ByteCount())
	}
}
