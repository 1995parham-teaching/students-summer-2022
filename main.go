package main

import (
	"log"

	"github.com/1995parham-teaching/students/internal/handler"
	"github.com/1995parham-teaching/students/internal/model"
	"github.com/1995parham-teaching/students/internal/store"
	"github.com/labstack/echo/v4"
)

func main() {
	app := echo.New()

	var studentStore store.Student = store.NewStudentInMemory()

	studentStore.Save(model.Student{
		FirstName: "Parham",
		LastName:  "Alvani",
		ID:        9231058,
		Average:   18,
	})

	h := handler.Student{
		Store: studentStore,
	}

	h.Register(app.Group("/api/students"))

	if err := app.Start(":1234"); err != nil {
		log.Println(err)
	}
}
