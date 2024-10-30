package storage

import (
	"context"

	"github.com/CamPlume1/khoury-classroom/internal/models"
)

type Storage interface {
	Close(context.Context)
	Test
}

type Test interface {
	GetTests(ctx context.Context) ([]models.Test, error)
}

