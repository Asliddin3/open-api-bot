package api

import (
	"github.com/Asliddin3/open-api-bot/api/repo"
)

type ApiI interface {
	Pixabay() repo.PixabayRepo
}

type ApiStorage struct {
	// db          *sql.Conn
	pixabayRepo repo.PixabayRepo
}

func RegisterApi() ApiI {
	return &ApiStorage{}
}

func (a *ApiStorage) Pixabay() repo.PixabayRepo {
	return a.pixabayRepo
}
