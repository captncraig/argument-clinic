package data

import (
	"context"
	"crypto/rand"

	"github.com/captncraig/argument-clinic/models"
)

type DataAccess interface {
	SiteFromHost(ctx context.Context, host string) (*models.Site, error)

	CreateComment(ctx context.Context, comment *models.Comment) (uint64, error)
}

// GenerateSalt gives 160 bytes of random data. Intended to be a site "key" for cookie encryption, hash salts, etc.
// 32 for hash key, 32 for block key for secure cookie
// 32 more for our own general-purpose hash key
func GenerateSalt() []byte {
	//128 bytes for secure cookie + 32 more for a general purpose hash salt
	dat := make([]byte, 128+32)
	rand.Read(dat)
	return dat
}
