package main

import (
	"go-br-task/internal/api"
	"go-br-task/internal/task"
	"go-br-task/internal/user"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
)

func main() {
	storageTask := task.NewMapStorageTask()
	storageUser := user.NewMapStorageUser()
	serviceTask := task.NewTaskService(storageTask)
	serviceUser := user.NewUserService(storageUser)
	handlerUser := user.NewHandlerUser(serviceUser)
	handlerTask := task.NewHandler(serviceTask)

	r := gin.Default()
	api.Init(r, handlerTask, handlerUser)
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
