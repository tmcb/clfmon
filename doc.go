/*
clfmon monitors Common Log Format (CLF) files from HTTP servers.

It receives a log file path and a hit limit as entries. It analyzes the log file
and informs the section of the web site with the most hits at every 10 seconds.

If the total number of hits in the last 2 minutes surpasses the hit limit, it
emits an alert informing the date and time of the event.

When the number of hits recover and get below the limit, it will emit a warning
telling that the hit load is already normalized.

Usage:
	clfmon [logpath] [hitlimit]
*/
package main
