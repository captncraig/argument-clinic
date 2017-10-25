package data

import (
	"context"

	"github.com/captncraig/argument-clinic/models"
)

type DataAccess interface {
	// CreateComment creates a comment and returns the ID as a string
	CreateComment(context.Context, *models.CreateCommentRequest) (uint64, error)
}
