package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (api *Api) Ping_Setup() {
	api.R.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
}
