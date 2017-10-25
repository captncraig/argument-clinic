package models

import "time"

type CreateCommentRequest struct {
	Name  string
	Email string
	Text  string
}

type Comment struct {
	ID   uint64
	Name string
	Text string
	Date time.Time
}
