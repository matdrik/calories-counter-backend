package main

import (
	"context"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"os/signal"
	caloriescounterbackend "server"
	"server/internal/handler"
	"server/internal/repository"
	"server/internal/service"
	"server/internal/ws"
	"syscall"
)

func main() {
	if err := initConfig(); err != nil {
		logrus.Fatalf("Ошибка при считывании конфигурационного файла: %v", err)
	}

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("Ошибка при считывании переменных окружения: %v", err)
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   viper.GetString("db.dbName"),
		SSLMode:  viper.GetString("db.sslMode"),
	})
	if err != nil {
		logrus.Fatalf("Ошибка при инициализации БД: %v", err)
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	hub := ws.NewHub()
	go hub.Run()

	//router := gin.Default()
	//router.GET("/ws", ws.HandleWebSocket(hub))

	//fmt.Println("Gin server started at :8080")
	//if err := router.Run(":8080"); err != nil {
	//	log.Fatalf("Server failed: %v", err)
	//}

	server := new(caloriescounterbackend.Server)
	go func() {
		if err := server.Run(viper.GetString("port"), handlers.Init(hub)); err != nil {
			logrus.Fatalf("Ошибка при выполнении http-сервера: %v", err.Error())
		}

	}()

	logrus.Print("calories-counter-backend запущено")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	logrus.Print("calories-counter-backend остановлено")

	if err := server.Shutdown(context.Background()); err != nil {
		logrus.Errorf("Ошибка при остановке сервера: %v", err.Error())
	}

	if err := db.Close(context.Background()); err != nil {
		logrus.Errorf("Ошибка при закрытии соединения БД: %v", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
