package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Простейший метод пинга сервера
func (api *Api) Ping_Setup() {
	api.R.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
}
