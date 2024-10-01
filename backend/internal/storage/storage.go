package storage

import (
	"context"
)

type Storage interface {
	Close (context.Context) error

}
