package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// func (api *Api) Get_Setup(GetD  bd.FuncDef) {
func (api *Api) Get_Setup() {
	api.R.GET("/user/:name", func(c *gin.Context) {

		// Получить значение по путь - Имя
		user := c.Params.ByName("name")

		// Проверка входных данных на пустоту
		if user == "" {
			c.JSON(http.StatusNotAcceptable, gin.H{"user": user, "status": "nill 'user'"})
		}

		// получить данные из базы данных
		password, IsExist := api.DB.Get(user)

		// Возвращаем результат
		if IsExist != nil {
			c.JSON(http.StatusOK, gin.H{"user": user, "status": "no value"})
		} else {
			c.JSON(http.StatusOK, gin.H{"user": user, "pass": password})
		}
	})
}
