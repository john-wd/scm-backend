package server

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/john-wd/scm-backend/mock"
)

type server struct {
}

func New() *server {
	return &server{}
}

func (s *server) RegisterRoutes(mux *chi.Mux) *chi.Mux {
	mux.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedHeaders: []string{"*"},
		MaxAge:         300, // Maximum value not ignored by any of major browsers
	}))

	mux.Use(func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			log.Default().Printf("got requestr for %s", r.URL)
			h.ServeHTTP(w, r)
		})
	})
	mux.Get("/json/gamelist/", s.GetGamelist)
	mux.Get("/json/game/{gameId}", s.GetGame)
	mux.Get("/json/song/{songId}", s.GetSongById)
	mux.Get("/brstm/{songId}", s.DownloadSongById)
	return mux
}

func (s *server) GetGamelist(w http.ResponseWriter, r *http.Request) {
	list, err := mock.Gamelist()
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("could not get gamelist: %v", err)))
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(list)
}
func (s *server) GetGame(w http.ResponseWriter, r *http.Request) {
	gameId := chi.URLParam(r, "gameId")
	game, err := mock.Game(gameId)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("could not get game %s: %v", gameId, err)))
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(game)
}
func (s *server) GetSongById(w http.ResponseWriter, r *http.Request) {
	songId := chi.URLParam(r, "songId")
	song, err := mock.Song(songId)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("could not get song %s: %v", songId, err)))
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(song)
}
func (s *server) DownloadSongById(w http.ResponseWriter, r *http.Request) {
	songId := chi.URLParam(r, "songId")
	blob, err := mock.SongBlob(songId)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("could not get song %s: %v", songId, err)))
		return
	}
	w.Header().Add("Content-Type", "application/octet-stream")
	w.Header().Add("content-length", fmt.Sprint(len(blob)))
	w.WriteHeader(http.StatusOK)
	w.Write(blob)
}
