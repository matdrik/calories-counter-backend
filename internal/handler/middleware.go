package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

const (
	authHeader = "Authorization"
	userCtxKey = "userId"
)

func (h *Handler) AuthMiddleware(c *gin.Context) {
	header := c.GetHeader(authHeader)
	if header == "" {
		newErrorResponse(c, http.StatusUnauthorized, "Missing auth header")
		return
	}

	headerParts := strings.SplitN(header, " ", 2)
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		newErrorResponse(c, http.StatusUnauthorized, "Invalid auth header")
		return
	}

	userId, err := h.services.Authorization.ParseToken(headerParts[1])
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	c.Set(userCtxKey, userId)
	// todo - сделать тут сохранение роли пользователя
}

func getUserId(c *gin.Context) (int, error) {
	id, ok := c.Get(userCtxKey)
	if !ok {
		newErrorResponse(c, http.StatusUnauthorized, "Missing user id")
		return 0, errors.New("missing user id")
	}

	idInt, ok := id.(int)
	if !ok {
		newErrorResponse(c, http.StatusUnauthorized, "Invalid user id")
		return 0, errors.New("invalid user id")
	}

	return idInt, nil
}
