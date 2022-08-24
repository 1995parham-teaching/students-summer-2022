package main

import (
	"log"

	"github.com/1995parham-teaching/students/internal/handler"
	"github.com/1995parham-teaching/students/internal/store"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

func main() {
	app := echo.New()
	logger, err := zap.NewDevelopment()
	if err != nil {
		log.Fatal(err)
	}

	var studentStore store.Student

	{
		logger := logger.Named("store")

		studentStore = store.NewStudentInMemory(
			logger.Named("student"),
		)
	}

	{
		logger := logger.Named("http")

		h := handler.Student{
			Store:  studentStore,
			Logger: logger.Named("student"),
		}

		h.Register(app.Group("/api/students"))
	}

	if err := app.Start(":1234"); err != nil {
		log.Println(err)
	}
}
