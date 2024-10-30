package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"os"
	"server/internal/api"
	"server/internal/config"
	"server/internal/repository"
)

import (
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/jackc/pgx/v5/stdlib"
)

func init() {
	if err := godotenv.Load("../../.env"); err != nil {
		log.Print("Не найден файл конфигурации (.env)")
	}
}

func main() {
	//todo - должен ли следующий код быть в приложении всегда?
	//err := migrate.RunMigrations()
	//if err != nil {
	//	log.Fatal(err.Error())
	//	return
	//}

	conf := config.New()

	dbpool, err := repository.New(conf.DBConnectionString)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Не удалось подключиться к базе данных: %v\n", err)
		os.Exit(1)
	}

	apiObj := api.New(gin.New(), dbpool)
	apiObj.Handle()
	apiObj.ListenAndServe(conf.Host, conf.Port)
}
