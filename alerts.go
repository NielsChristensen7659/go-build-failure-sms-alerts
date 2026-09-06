package main

import "fmt"

type BuildEvent struct{ Project, Commit, Status string }

func AlertText(e BuildEvent) (string, bool) {
	if e.Status != "failed" {
		return "", false
	}
	return fmt.Sprintf("%s build failed at %s", e.Project, e.Commit), true
}

func HandleBuild(c *SMSClient, to string, e BuildEvent) (string, error) {
	text, send := AlertText(e)
	if !send {
		return "", nil
	}
	return c.Send(to, text, e.Project+":"+e.Commit)
}
