package api

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"server/internal/models"
	"server/internal/utils"
	"strings"
)

// todo - вынести следующие константы и переменные в конфиг
const userSessionKey = "user"

func (api *api) RegisterHandler(c *gin.Context) {
	var body models.LimitedUser
	if err := c.ShouldBindJSON(&body); err != nil {
		if err := c.Error(err); err != nil {
		}
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Неверный формат параметров регистрации"})
		return
	}

	username := body.Username
	passwordHash, err := utils.HashPassword(body.Password)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Внутренняя ошибка сервера"})
		return
	}

	userExists, err := api.db.UsernameExists(username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Внутренняя ошибка сервера: %v", err.Error())})
		return
	}
	if userExists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Пользователь с такими именем уже существует"})
		return
	}

	err = api.db.CreateUser(username, passwordHash)
	if err != nil {
		if err := c.Error(err); err != nil {
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Успешная регистрация"})
}

func (api *api) LoginHandler(c *gin.Context) {
	session := sessions.Default(c)

	var body models.LimitedUser
	if err := c.ShouldBindJSON(&body); err != nil {
		if err := c.Error(err); err != nil {
		}
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Неверный формат параметров авторизации"})
		return
	}

	username := body.Username
	password := body.Password

	// todo - при неправильном формате body до сюда проверка не дойдет. Проверить
	if strings.Trim(username, " ") == "" || strings.Trim(password, " ") == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Пустые параметры авторизации"})
		return
	}

	userData, err := api.db.GetUserByUsername(username)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Пользователь с таким именем не найден"})
		return
	}

	err = utils.CheckParameters(username, userData.Username, password, userData.PasswordHash)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//todo - заменить username айдишник пользователя или что-то другое
	session.Set(userSessionKey, userData.ID)
	if err := session.Save(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Внутренняя ошибка сервера"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Пользователь успешно авторизирован"})
}

func LogoutHandler(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get(userSessionKey)
	if user == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный токен сессии авторизации"})
		return
	}
	session.Delete(userSessionKey)
	if err := session.Save(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Внутренняя ошибка сервера"})
		return
	}
}

func GetUserHandler(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get(userSessionKey)
	c.JSON(http.StatusOK, gin.H{"user": user})
}

func GetStatusHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "Вы авторизированны"})
}
