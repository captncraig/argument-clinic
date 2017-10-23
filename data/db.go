package data

import "github.com/captncraig/argument-clinic/models"

type DataAccess interface {
	CreateComment(*models.CreateCommentRequest) (string, error)
}
