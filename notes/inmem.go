package notes

import (
	"context"
	"errors"
	"time"
)

type inMemoryStorage struct {
	notes map[string]Note
}

func NewInMemoryStorage() Storage {
	return &inMemoryStorage{notes: map[string]Note{
		"12": {
			ID:        "12",
			Name:      "Test note",
			Content:   "Learn go",
			CreatedAt: time.Now(),
		},
	}}
}

func (i *inMemoryStorage) Read(ctx context.Context, id string) (Note, error) {
	note, ok := i.notes[id]
	if !ok {
		return Note{}, errors.New("not not found")
	}

	return note, nil
}

func (i *inMemoryStorage) Create(ctx context.Context, name string, content string) (Note, error) {
	panic("not implemented") // TODO: Implement
}

func (i *inMemoryStorage) Delete(ctx context.Context, id string) error {
	panic("not implemented") // TODO: Implement
}
