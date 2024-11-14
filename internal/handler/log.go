package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	calories_counter_backend "server"
	"strconv"
)

func (h *Handler) createLog(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	var input calories_counter_backend.Log
	if err := c.ShouldBind(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Create(userId, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": id})
}

type getAllListResponse struct {
	Data []calories_counter_backend.LogResponse `json:"data"`
}
type getAllListRequest struct {
	Date string `json:"date"`
}

func (h *Handler) getAllLogs(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	date := c.Param("date")
	if date == "" {
		newErrorResponse(c, http.StatusBadRequest, "Invalid DATE param")
		return
	}

	//var input getAllListRequest
	//if err := c.ShouldBind(&input); err != nil {
	//	newErrorResponse(c, http.StatusBadRequest, err.Error())
	//	return
	//}

	data, err := h.services.Log.GetAll(userId, date)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllListResponse{Data: data})
}

func (h *Handler) getLogById(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid ID param")
		return
	}

	data, err := h.services.GetById(userId, id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, data)
}
func (h *Handler) updateLog(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid ID param")
		return
	}

	var input calories_counter_backend.UpdateLogInput
	if err := c.ShouldBind(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.Update(userId, id, input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, StatusResponse{Status: "ok"})
}
func (h *Handler) deleteLog(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid ID param")
		return
	}

	err = h.services.Delete(userId, id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, StatusResponse{
		Status: "ok",
	})
}
