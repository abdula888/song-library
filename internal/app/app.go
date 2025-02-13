package app

import (
	"song-library/internal/config"
	"song-library/internal/delivery/http"
	"song-library/internal/infrastructure/postgres/repository"
	"song-library/internal/usecase"
	"song-library/migrations"
	"song-library/pkg/db/conn"
	"song-library/pkg/log"
)

func Run(conf *config.Config) {
	log.SetUpLogger(conf.LogLevel)

	db, err := conn.InitDB(conf.DatabaseURL)
	if err != nil {
		log.Logger.Fatal("Failed to connect to the database:", err)
	}
	log.Logger.Info("Successfully connected to the database")

	if err := migrations.RunMigrations(db); err != nil {
		log.Logger.Fatal("Error applying migration: ", err)
	}
	log.Logger.Info("Migrations applied successfully")

	// Регистрация маршрутов
	r := http.NewRouter(usecase.New(
		repository.New(db),
		conf.APIConfig,
	))
	log.Logger.Info("Routes registered successfully")

	// Запуск сервера
	log.Logger.Info("Server started on port 8080")
	log.Logger.Fatal(r.Run(":8080"))
}
