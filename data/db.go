package data

import "github.com/captncraig/argument-clinic/models"

type DataAccess interface {
	// CreateComment creates a comment and returns the ID as a string
	CreateComment(*models.CreateCommentRequest) (string, error)
}
