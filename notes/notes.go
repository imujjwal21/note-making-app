package notes

import (
	"context"
	"time"
)

type Note struct {
	ID        string
	Name      string
	Content   string
	CreatedAt time.Time
}

type Storage interface {
	Read(ctx context.Context, id string) (Note, error)
	Create(ctx context.Context, name, content string) (Note, error)
	Delete(ctx context.Context, id string) error
}
