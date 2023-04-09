package api

import (
	"github.com/RB-PRO/norsi-trans/internal/rest/bd"
	"github.com/gin-gonic/gin"
)

type Api struct {
	R  *gin.Engine // объект Gin для работы Api
	DB *bd.Base    // Объект базы данных

}

// Создать экземпляр API сервиса
func New() (*Api, error) {

	// Создаём экземпляр базы данных
	DB, ErrorDB := bd.New()
	if ErrorDB != nil {
		return nil, ErrorDB
	}

	return &Api{
		R:  gin.Default(),
		DB: DB,
	}, nil
}
