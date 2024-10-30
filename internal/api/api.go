package api

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"server/internal/middleware"
	"server/internal/repository"
)

// todo - вынести следующие константы и переменные в конфиг
var cookieSecret = []byte("secret")

type api struct {
	r  *gin.Engine
	db *repository.PGRepo
}

func New(router *gin.Engine, db *repository.PGRepo) *api {
	return &api{r: router, db: db}
}

func (api *api) Handle() {
	store := cookie.NewStore(cookieSecret)
	store.Options(sessions.Options{MaxAge: 60 * 10})
	api.r.Use(sessions.Sessions("auth_session", store))

	err := api.r.SetTrustedProxies([]string{"127.0.0.1"})
	if err != nil {
		log.Fatal(err)
	}
	api.r.Use(gin.Logger())

	//api.r.Use(middleware.AdminRoleRequired)
	//api.r.Handle(http.MethodPost, "/test", func(context *gin.Context) {
	//	context.JSON(http.StatusOK, gin.H{"message": "this is test"})
	//})

	api.r.Use(cors.Default())

	superGroup := api.r.Group("/api")
	{
		superGroup.Handle(http.MethodPost, "/login", api.LoginHandler)
		superGroup.Handle(http.MethodPost, "/register", api.RegisterHandler)

		private := superGroup.Group("/private")
		{
			private.Use(middleware.AuthRequired)
			private.Handle(http.MethodGet, "/me", GetUserHandler)
			private.Handle(http.MethodGet, "/status", GetStatusHandler)
			private.Handle(http.MethodGet, "/logout", LogoutHandler)
			private.Handle(http.MethodGet, "/food", api.GetFoodHandler)

			admin := private.Group("/admin")
			{
				admin.Use(middleware.AdminRoleRequired)
				admin.Handle(http.MethodPost, "/food", api.CreateFoodHandler)
			}
		}
	}
}

func (api *api) ListenAndServe(host string, port string) {
	if err := api.r.Run(fmt.Sprintf("%v:%v", host, port)); err != nil {
		log.Fatal("Не удалось запустить сервер: ", err)
	}
}
