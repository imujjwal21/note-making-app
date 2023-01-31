package httptransport

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/imujjwal21/note-making-app/notes"
)

func NewHandler(storage notes.Storage) http.Handler {
	router := mux.NewRouter()
	router.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("content-type", "application/json")
			next.ServeHTTP(w, r)
		})
	})

	router.Path("/{id}").Methods(http.MethodGet).HandlerFunc(readByIDHandler(storage))

	return router
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
