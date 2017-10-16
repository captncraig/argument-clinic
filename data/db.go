package data

import (
	"github.com/captncraig/arguments/models"
)

type DataAccess interface {
	CreateComment(*models.Comment) (uint64, error)
}
