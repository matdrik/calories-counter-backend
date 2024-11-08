package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	calories_counter_backend "server"
)

type getAllFoodResponse struct {
	Data []calories_counter_backend.FoodResponse `json:"data"`
}

func (h *Handler) getAllFood(c *gin.Context) {
	_, err := getUserId(c)
	if err != nil {
		return
	}

	data, err := h.services.Food.GetAll()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllFoodResponse{Data: data})
}
