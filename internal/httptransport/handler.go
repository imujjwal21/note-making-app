package httptransport

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/imujjwal21/note-making-app/notes"
)

type BaseHandler struct {
	db *sql.DB
}

func NewBaseHandler(db *sql.DB) *BaseHandler {
	return &BaseHandler{
		db: db,
	}
}

func NewHandler(storage notes.Storage) http.Handler {
	router := mux.NewRouter()
	router.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("content-type", "application/json")
			next.ServeHTTP(w, r)
		})
	})

	router.Path("/note/{id}").Methods(http.MethodGet).HandlerFunc(readByIDHandler(storage)) // router.HandleFunc("/{id}", readByIDHandler(storage)).Methods("GET")
	router.Path("/note/{id}").Methods(http.MethodDelete).HandlerFunc(deleteByIDHandler(storage))
	router.Path("/note/name/{name}/content/{content}").Methods(http.MethodDelete).HandlerFunc(createNotesHandler(storage))
	router.Path("/note").Methods(http.MethodDelete).HandlerFunc(GetAllNotesHandler(storage))

	// router.HandleFunc("/note/", CreateNote()).Methods("POST")
	// router.HandleFunc("/note/", GetNote()).Methods("GET")
	// router.HandleFunc("/delete/{noteId}", deleteNotesById()).Methods("DELETE")
	// router.HandleFunc("/notes/{noteId}", getNotesById()).Methods("GET")

	return router
}

func createNotesHandler(storage notes.Storage) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		name := mux.Vars(r)["name"]
		content := mux.Vars(r)["content"]

		_, err := storage.Create(r.Context(), name, content)
		if err != nil {
			log.Printf("cannot Created note: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}

func GetAllNotesHandler(storage notes.Storage) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		// var note []notes.Note

		note, err := storage.GetAll(r.Context())
		if err != nil {
			log.Printf("cannot read note: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(note)
	}
}

func readByIDHandler(storage notes.Storage) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		id := mux.Vars(r)["id"]
		note, err := storage.Read(r.Context(), id)
		if err != nil {
			log.Printf("cannot read note: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(note)
	}
}

func deleteByIDHandler(storage notes.Storage) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		id := mux.Vars(r)["id"]
		err := storage.Delete(r.Context(), id)
		if err != nil {
			log.Printf("cannot able Delete notes : %v", err)
			return
		}

	}
}

// func GetNote(w http.ResponseWriter, r *http.Request) {
// 	NewNote := notes.GetAllNotes()
// 	rs, _ := json.Marshal(NewNote)

// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(rs)
// }
// func getNotesById(w http.ResponseWriter, r *http.Request) {
// 	vars := mun
// }
