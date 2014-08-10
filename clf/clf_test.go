package clf

import (
	"testing"
)

const (
	remoteHost = "1.1.1.1"
	rfc931     = "-"
	authUser   = "-"
	dateStamp  = "[01/Jan/2000:00:00:00"
	timeStamp  = "-0000]"
	method     = "GET"
	resource   = "/section"
	protocol   = "HTTP/1.1"
	status     = "200"
	bytes      = "1024"
)

func TestIncompleteEntries(t *testing.T) {
	line := ""
	if _, err := Parse(line); err == nil {
		t.Error()
	}
	line += remoteHost
	if _, err := Parse(line); err == nil {
		t.Error()
	}
	line += " " + rfc931
	if _, err := Parse(line); err == nil {
		t.Error()
	}
	line += " " + authUser
	if _, err := Parse(line); err == nil {
		t.Error()
	}
	line += " " + dateStamp
	if _, err := Parse(line); err == nil {
		t.Error()
	}
	line += " " + timeStamp
	if _, err := Parse(line); err == nil {
		t.Error()
	}
	line += " \"" + method + " " + resource + " " + protocol + "\""
	if _, err := Parse(line); err == nil {
		t.Error()
	}
	line += " " + status
	if _, err := Parse(line); err == nil {
		t.Error()
	}
	line += " " + bytes
	// Now the line is complete; the comparison between err and nil
	// changed
	if _, err := Parse(line); err != nil {
		t.Error(err)
	}
	line += " garbage"
	// Line should be correctly parsed; garbage should be overlooked
	if _, err := Parse(line); err != nil {
		t.Error(err)
	}
}

func TestRequests(t *testing.T) {
	linePrefix := remoteHost + " " +
		rfc931 + " " +
		authUser + " " +
		dateStamp + " " +
		timeStamp
	lineSuffix := status + " " +
		bytes
	var line string
	// Empty request
	line = linePrefix + " \"\" " +
		lineSuffix
	if entry, err := Parse(line); err != nil ||
		entry.Method != "" ||
		entry.Resource != "" ||
		entry.Protocol != "" {
		t.Error(err)
	}
	// Method-only request
	line = linePrefix + " \"" +
		method + "\" " +
		lineSuffix
	if entry, err := Parse(line); err != nil ||
		entry.Method != method ||
		entry.Resource != "" ||
		entry.Protocol != "" {
		t.Error(err)
	}
	// Method + resource request
	line = linePrefix + " \"" +
		method + " " +
		resource + "\" " +
		lineSuffix
	if entry, err := Parse(line); err != nil ||
		entry.Method != method ||
		entry.Resource != resource ||
		entry.Protocol != "" {
		t.Error(err)
	}
	// Method + resource + protocol request
	line = linePrefix + " \"" +
		method + " " +
		resource + " " +
		protocol + "\" " +
		lineSuffix
	if entry, err := Parse(line); err != nil ||
		entry.Method != method ||
		entry.Resource != resource ||
		entry.Protocol != protocol {
		t.Error(err)
	}
	// Method + resource + protocol + garbage request; line should be
	// correctly parsed and garbage should be overlooked
	line = linePrefix + " \"" +
		method + " " +
		resource + " " +
		protocol + " garbage" + "\" " +
		lineSuffix
	if entry, err := Parse(line); err != nil ||
		entry.Method != method ||
		entry.Resource != resource ||
		entry.Protocol != protocol {
		t.Error(err)
	}
	// Open request
	line = linePrefix + " \"" +
		method + " " +
		resource + " " +
		protocol + " " +
		lineSuffix
	if _, err := Parse(line); err == nil {
		t.Error()
	}
}

func TestBadDates(t *testing.T) {
	linePrefix := remoteHost + " " +
		rfc931 + " " +
		authUser
	lineSuffix := " \"" + method + " " +
		resource + " " +
		protocol + "\" " +
		status + " " +
		bytes
	var line string
	// Empty dates
	line = linePrefix + " [] " + lineSuffix
	if _, err := Parse(line); err == nil {
		t.Error()
	}
	// Garbage instead of date, once
	line = linePrefix + " [garbage] " + lineSuffix
	if _, err := Parse(line); err == nil {
		t.Error()
	}
	// Garbage instead of date, twice
	line = linePrefix + " [garbage garbage] " + lineSuffix
	if _, err := Parse(line); err == nil {
		t.Error()
	}
	// Garbage instead of date, twice; tests the 5-char limit at the time
	// place
	line = linePrefix + " [garbage garbg] " + lineSuffix
	if _, err := Parse(line); err == nil {
		t.Error()
	}
	// Garbage instead of date, three times
	line = linePrefix + " [garbage garbage garbage] " + lineSuffix
	if _, err := Parse(line); err == nil {
		t.Error()
	}
}

func TestBadStatuses(t *testing.T) {
	linePrefix := remoteHost + " " +
		rfc931 + " " +
		authUser + " " +
		dateStamp + " " +
		timeStamp + " \"" +
		method + " " +
		resource + " " +
		protocol + "\""
	lineSuffix := bytes
	var line string
	// Negative status
	line = linePrefix + " -1 " + lineSuffix
	if _, err := Parse(line); err == nil {
		t.Error()
	}
	// String-like status
	line = linePrefix + " four-oh-four " + lineSuffix
	if _, err := Parse(line); err == nil {
		t.Error()
	}
	// Mixed-type status (an error should be returned because there are more
	// fields to read)
	line = linePrefix + " 200garbage " + lineSuffix
	if _, err := Parse(line); err == nil {
		t.Error(err)
	}
}

func TestBadBytes(t *testing.T) {
	linePrefix := remoteHost + " " +
		rfc931 + " " +
		authUser + " " +
		dateStamp + " " +
		timeStamp + " \"" +
		method + " " +
		resource + " " +
		protocol + "\" " +
		status
	var line string
	// Negative bytes
	line = linePrefix + " -1"
	if _, err := Parse(line); err == nil {
		t.Error(err)
	}
	// String-like byte count
	line = linePrefix + " one-megabyte"
	if _, err := Parse(line); err == nil {
		t.Error()
	}
}
