package main

import (
	"go-br-task/internal"
	"go-br-task/internal/task"
	"go-br-task/internal/user"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
)

func main() {
	storage_task := task.NewMapStorageTask()
	storage_user := user.NewMapStorageUser()
	service_task := task.NewTaskService(storage_task)
	service_user := user.NewUserService(storage_user)
	handler_user := user.NewHandlerUser(service_user)
	handler_task := task.NewHandler(service_task)

	r := gin.Default()
	internal.Init(r, handler_task, handler_user)
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
