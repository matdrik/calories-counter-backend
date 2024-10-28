package main

import (
	"context"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/jackc/pgx/v5/stdlib"
	"log"
	"net/http"
	"os"
	"strings"
)

type LoginBody struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// todo - вынести следующие константы и переменные в конфиг
const connectionString string = "postgresql://postgres:8563@localhost:5432/calories_counter"

var cookieSecret = []byte("secret")

const userSessionKey = "user"

func main() {
	//todo - должен ли следующий код быть в приложении всегда?
	//if err := migration.RunMigrations(connectionString); err != nil {
	//	log.Fatalf("Migration failed: %v", err)
	//}

	dbpool, err := pgxpool.New(context.Background(), connectionString)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Не удалось подключиться к базе данных: %v\n", err)
		os.Exit(1)
	}
	defer dbpool.Close()

	r := engine()
	r.Use(gin.Logger())
	if err := r.Run("localhost:8080"); err != nil {
		log.Fatal("Не удалось запустить сервер: ", err)
	}

	//r := gin.Default()
	//err = r.SetTrustedProxies([]string{"127.0.0.1"})
	//if err != nil {
	//	log.Fatal(err)
	//}
	//r.GET("/ping", func(c *gin.Context) {
	//	c.JSON(200, gin.H{
	//		"message": "pong",
	//	})
	//})
	//err = r.Run("localhost:8080")
	//if err != nil {
	//	log.Fatal(err)
	//}

	//var greeting string
	//err = dbpool.QueryRow(context.Background(), "select 'Hello, world!'").Scan(&greeting)
	//if err != nil {
	//	fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
	//	os.Exit(1)
	//}

	//fmt.Println(greeting)
}

func engine() *gin.Engine {
	r := gin.New()

	err := r.SetTrustedProxies([]string{"127.0.0.1"})
	if err != nil {
		log.Fatal(err)
	}

	r.Use(sessions.Sessions("auth_session", cookie.NewStore(cookieSecret)))

	r.Handle(http.MethodPost, "/login", login)
	r.Handle(http.MethodGet, "/logout", logout)

	private := r.Group("/private")
	private.Use(AuthRequired)
	{
		private.Handle(http.MethodGet, "/me", me)
		private.Handle(http.MethodGet, "/status", status)
	}
	return r
}

func AuthRequired(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get(userSessionKey)
	if user == nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	c.Next()
}

// todo - сделать нормальную проверку параметров
func checkParammetrs(login, password string) bool {
	return login == "admin" || password == "admin"
}

func login(c *gin.Context) {
	session := sessions.Default(c)
	//username := c.PostForm("name")
	//password := c.PostForm("password")

	var content LoginBody
	if err := c.ShouldBindJSON(&content); err != nil {
		c.Error(err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	log.Println(content)

	username := content.Username
	password := content.Password

	if strings.Trim(username, " ") == "" || strings.Trim(password, " ") == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Пустые параметры авторизации"})
		return
	}

	checkParammetrs(username, password)

	//todo - заменить username айдишник пользователя или что-то другое
	session.Set(userSessionKey, username)
	if err := session.Save(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Внутренняя ошибка сервера"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Пользователь успешно авторизирован"})
}

func logout(c *gin.Context) {
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

func me(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get(userSessionKey)
	c.JSON(http.StatusOK, gin.H{"user": user})
}

func status(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "Вы авторизированны"})
}
