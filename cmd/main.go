package main

import (
	"go-br-task/internal/handlers"
	"go-br-task/internal/services"
	"go-br-task/internal/storages"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
)

func main() {
	storage_task := storages.NewMapStorageTask()
	storage_user := storages.NewMapStorageUser()
	service_task := services.NewTaskService(storage_task)
	service_user := services.NewUserService(storage_user)
	handler_user := handlers.NewHandlerUser(service_user)
	handler_task := handlers.NewHandler(service_task)

	r := gin.Default()
	handlers.Init(r, handler_task, handler_user)
	go r.Run(":8080")

	Shutdown()
}

// Прекращение процесса gin
func Shutdown() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	slog.Info("shut down successfully")
}
