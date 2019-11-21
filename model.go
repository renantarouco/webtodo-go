package main

import "time"

type Item struct {
	ID       int
	Title    string
	Deadline time.Time
}

func (i Item) FormatedDate() string {
	return i.Deadline.Format("2006-01-02")
}
