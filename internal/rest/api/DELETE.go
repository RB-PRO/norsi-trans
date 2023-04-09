package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Удалить запись из базы данных
func (api *Api) Delete_Setup() error {
	api.R.DELETE("/remove/:name", func(c *gin.Context) {

		// Получить значение по путь - Имя
		user := c.Params.ByName("name")

		// Проверка входных данных на пустоту
		if user == "" {
			c.JSON(http.StatusNotAcceptable, gin.H{"user": user, "status": "nill 'user'"})
		}

		// получить данные из базы данных
		IsDelete := api.DB.Delete(user)

		// Возвращаем результат
		if IsDelete != nil {
			c.JSON(http.StatusOK, gin.H{"user": user, "status": IsDelete.Error()})
		} else {
			c.JSON(http.StatusOK, gin.H{"user": user})
		}
	})

	return nil
}
