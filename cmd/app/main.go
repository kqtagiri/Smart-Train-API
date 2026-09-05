package main

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"smarttrain/internal/database"
	"smarttrain/internal/handler"
	"smarttrain/internal/repository"
	"smarttrain/internal/service"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {

	ctx := context.Context(context.Background())
	db, err := database.NewDB(ctx)
	if err != nil {
		slog.Error("Get	next error when connect to database:", err)
		return
	}

	userRepo := repository.NewUserRepo(db)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "HELLO FROM GIN")
	})

	users := r.Group("/users")
	users.GET("/all", userHandler.AllUsersInfo)

	server := http.Server{
		Addr:    ":9111",
		Handler: r,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			slog.Error(err.Error())
			return
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	context, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	if err := server.Shutdown(context); err != nil {
		slog.Error(err.Error())
		return
	}

	if err := db.Close(); err != nil {
		slog.Error(err.Error())
		return
	}

}
