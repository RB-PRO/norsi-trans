package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (api *Api) Post_Setup() {
	api.R.POST("/add", func(c *gin.Context) {

		// Получить данные с параметров
		User := c.Query("user")
		Password := c.Query("pass")

		// Проверяем заполненность входных данных
		if User == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "'user' is null"})
			return
		}
		if Password == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "'password' is null"})
			return
		}

		// Ответить положительно на ответ
		c.JSON(http.StatusOK, gin.H{
			"user": User,
			"pass": Password,
		})
	})
}
