package rest

import (
	"github.com/RB-PRO/norsi-trans/internal/rest/api"
)

// Структура данных, которая приходит с методов
type Query struct {
	User     string `json:"user"`
	Password string `json:"pass"`
}

func Start() {

	// Создаём экзепляр API
	API := api.New()

	// Настраиваем методы
	API.Get_Setup()
	API.Post_Setup()
	API.Delete_Setup()
	API.Ping_Setup()

	// Запускаем
	API.R.Run(":8080") // http://localhost:8080
}
