package mock

import (
	_ "embed"
	"io"
	"log"
	"os"
	"path/filepath"
)

var (
	mockPath string
)

func Configure(path string) {
	if _, err := os.Stat(path); err != nil {
		panic(err)
	}
	mockPath = path
}

func Gamelist() ([]byte, error) {
	return readFile("gamelist.json")
}

func Game(gameId string) ([]byte, error) {
	return readFile(filepath.Join("game", gameId+".json"))
}

func Song(songId string) ([]byte, error) {
	return readFile(filepath.Join("songs", songId+".json"))
}

func SongBlob(songId string) ([]byte, error) {
	return readFile(filepath.Join("blobs", songId+".brstm"))
}

func readFile(relpath string) ([]byte, error) {
	pwd, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	path := filepath.Join(pwd, mockPath, relpath)
	log.Default().Printf("reading: %s", path)
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	return io.ReadAll(f)
}
