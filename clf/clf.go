package clf

import (
	"fmt"
	"strings"
	"time"
)

/*
Struct Entry holds the different fields of a CLF entry.
*/
type Entry struct {
	RemoteHost string
	RFC931     string
	AuthUser   string
	Date       time.Time
	Method     string
	Resource   string
	Protocol   string
	Status     uint16
	Bytes      uint64
}

/*
Parse fills a struct Entry from a given log line. This line is returned to the
caller, with a non-nil error message if it was not possible to parse the line.
*/
func Parse(line string) (entry Entry, err error) {
	var tmpDate, tmpRequest, tmpTime string
	// The %q mask matches the request, since it is double-quoted.
	_, err = fmt.Sscanf(line, "%s %s %s [%s %5s] %q %d %d",
		&entry.RemoteHost,
		&entry.RFC931,
		&entry.AuthUser,
		&tmpDate,
		&tmpTime,
		&tmpRequest,
		&entry.Status,
		&entry.Bytes)
	if err != nil {
		return
	}
	const dateLayout = "02/Jan/2006:15:04:05 -0700"
	entry.Date, err = time.Parse(dateLayout, tmpDate+" "+tmpTime)
	if err != nil {
		return
	}
	fmt.Sscanf(tmpRequest, "%s %s %s",
		&entry.Method,
		&entry.Resource,
		&entry.Protocol)
	entry.Method = strings.ToUpper(entry.Method)
	entry.Protocol = strings.ToUpper(entry.Protocol)
	return
}
