package notes

import (
	"context"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type Note struct {
	ID        string    `json:"ID"`
	Name      string    `json:"Name"`
	Content   string    `json:"Content"`
	CreatedAt time.Time `json:"CreatedAt"`
}

type Storage interface {
	Read(ctx context.Context, id string) (Note, error)
	Create(ctx context.Context, name, content string) (Note, error)
	Delete(ctx context.Context, id string) error
	GetAll(ctx context.Context) ([]Note, error)
}

// func GetNotesById(Id int64) Note {
// 	row, err := db.Query(`select ID,Name,Content,CreatedAt from note where ID=?`, Id)

// 	var notes Note

// 	if err != nil {
// 		panic(err)
// 		return notes
// 	}

// 	defer row.Close()

// 	for row.Next() {
// 		err := row.Scan(&notes.ID, &notes.Name, &notes.Content, &notes.CreatedAt)

// 		if err != nil {
// 			panic(err)
// 			return notes
// 		}
// 	}

// 	if err := row.Err(); err != nil {
// 		panic(err)
// 		return notes
// 	}

// 	return notes
// }

// func DeleteNote(Id int64) error {
// 	_, err := db.Exec(`delete from note where ID=?`, Id)

// 	if err != nil {
// 		return errors.New("Can't able to Delete")
// 	}
// 	return nil
// }

// func (n *Note) CreateNote() error {
// 	createdAt := time.Now()Read

// 	_, err := db.Exec(`insert into notes (name,content,created_at) values (?,?,?)`, n.Name, n.Content, createdAt)
// 	if err != nil {
// 		return errors.New("Can't able to Insert")
// 	}
// 	return nil
// }

// func GetAllNotes() []Note {
// 	row, err := db.Query(`select id,name,content,created_at from notes`)

// 	var notes []Note

// 	if err != nil {
// 		panic(err)
// 		return notes
// 	}
// 	defer row.Close()

// 	for row.Next() {
// 		var tempNote Note
// 		err := row.Scan(&tempNote.ID, &tempNote.Name, &tempNote.Content, &tempNote.CreatedAt)

// 		if err != nil {
// 			panic(err)
// 			return notes
// 		}
// 		notes = append(notes, tempNote)
// 	}

// 	if err := row.Err(); err != nil {
// 		panic(err)
// 		return notes
// 	}

// 	return notes
// }
