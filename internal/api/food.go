package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"server/internal/models"
)

func (api *api) GetFoodHandler(c *gin.Context) {
	var data []models.Food
	//todo - если параметры (id еды) не переданы, то запросить все строки
	data, err := api.db.GetFood()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Внутренняя ошибка сервера: %v\n", err.Error())})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": data})
}

func (api *api) CreateFoodHandler(c *gin.Context) {
	var body models.Food
	if err := c.ShouldBindJSON(&body); err != nil {
		if err := c.Error(err); err != nil {
		}
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Неверный формат параметров добавляемого продукта"})
		return
	}

	if err := api.db.CreateFood(body); err != nil {
		if err := c.Error(err); err != nil {
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": nil})
}
