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
func New() Api {

	// Создаём экземпляр базы данных
	DB := bd.New()

	return Api{
		R:  gin.Default(),
		DB: DB,
	}
}
