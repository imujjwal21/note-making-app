package notes

import (
	"context"
	"database/sql"
	"errors"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type inMemoryStorage struct {
	db *sql.DB
}

func NewInMemoryStorage(db *sql.DB) Storage {
	return &inMemoryStorage{db}
}

// func NewInMemoryStorage() Storage {
// 	return &inMemoryStorage{notes: map[string]Note{}}
// }

func (i *inMemoryStorage) Create(ctx context.Context, name string, content string) (Note, error) {

	createdAt := time.Now()
	result, err := i.db.Exec(`insert into notes (name,content,created_at) values (?,?,?)`, name, content, createdAt)
	if err != nil {
		return Note{}, errors.New("can't able to Insert")
	}
	id, _ := result.LastInsertId()

	s1 := strconv.FormatInt(int64(id), 10) //s1 := strconv.Itoa(id)
	createdAt = time.Now()

	n1 := Note{s1, name, content, createdAt}
	// i.notes[s1] = n1

	return n1, nil
}

func (i *inMemoryStorage) GetAll(ctx context.Context) ([]Note, error) {

	row, err := i.db.Query(`select id,name,content,created_at from notes`)

	var note []Note

	if err != nil {
		return note, errors.New("can't connect to database")
	}
	defer row.Close()

	for row.Next() {
		var tempNote Note
		err := row.Scan(&tempNote.ID, &tempNote.Name, &tempNote.Content, &tempNote.CreatedAt)

		if err != nil {
			return note, errors.New("can't able to fetch data")
		}
		note = append(note, tempNote)
	}

	if err := row.Err(); err != nil {
		return note, errors.New("end error while get data")
	}

	return note, nil
	// var note []Note
	// for key, _ := range i.notes {
	// 	temp, ok := i.notes[key]
	// 	if !ok {
	// 		return note, errors.New("not not found from map")
	// 	}
	// 	note = append(note, temp)
	// }

	// return note, nil
}

func (i *inMemoryStorage) Read(ctx context.Context, id string) (Note, error) {

	row, err := i.db.Query(`select ID,Name,Content,CreatedAt from note where ID=?`, id)

	var note Note

	if err != nil {
		return note, errors.New("can't connect to database with id")
	}

	defer row.Close()

	for row.Next() {
		err := row.Scan(&note.ID, &note.Name, &note.Content, &note.CreatedAt)

		if err != nil {
			return note, errors.New("can't able to fetch data with id")
		}
	}

	if err := row.Err(); err != nil {
		return note, errors.New("end error while get data with id")
	}

	return note, nil
	// note, ok := i.notes[id]
	// if !ok {
	// 	return Note{}, errors.New("not not found")
	// }

	// return note, nil
}

func (i *inMemoryStorage) Delete(ctx context.Context, id string) error {

	_, err := i.db.Exec(`delete from note where ID=?`, id)

	if err != nil {
		return errors.New("can't able to delete")
	}
	return nil

	// _, err := db.Exec(`delete from notes where id=?`, id)

	// if err != nil {
	// 	return errors.New("Can't able to Delete")
	// }

	// delete(i.notes, id)
	// return nil
}

// func (i *inMemoryStorage) Create(ctx context.Context, name string, content string) ([]Note, error)
// {

// }

// func (i *inMemoryStorage) Read(ctx context.Context, id string) (Note, error) {
// 	note, ok := i.notes[id]
// 	if !ok {
// 		return Note{}, errors.New("not not found")
// 	}

// 	return note, nil
// }

// func (i *inMemoryStorage) Create(ctx context.Context, name string, content string) (Note, error) {
// 	panic("not implemented") // TODO: Implement
// }

// func (i *inMemoryStorage) Delete(ctx context.Context, id string) error {
// 	panic("not implemented") // TODO: Implement
// }
