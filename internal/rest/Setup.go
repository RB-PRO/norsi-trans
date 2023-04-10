package rest

import (
	"github.com/RB-PRO/norsi-trans/internal/rest/api"
)

// Структура данных, которая приходит с методов
type Query struct {
	User     string `json:"user"`
	Password string `json:"pass"`
}

func setup() (*api.Api, error) {
	// Создаём экзепляр API
	API, ErrorAPI := api.New()
	if ErrorAPI != nil {
		return nil, ErrorAPI
	}

	// Настраиваем метды
	API.Get_Setup()
	API.Post_Setup()
	API.Delete_Setup()
	API.Ping_Setup()

	return API, nil
}
