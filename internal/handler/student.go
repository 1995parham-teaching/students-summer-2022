package handler

import (
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/1995parham-teaching/students/internal/store"
	"github.com/labstack/echo/v4"
)

type Student struct {
	Store store.Student
}

func (s Student) GetAll(c echo.Context) error {
	ss, err := s.Store.GatAll()
	if err != nil {
		return echo.ErrInternalServerError
	}

	return c.JSON(http.StatusOK, ss)
}

func (s Student) Get(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return echo.ErrBadRequest
	}

	st, err := s.Store.Get(id)
	if err != nil {
		var errNotFound store.StudentNotFoundError
		if ok := errors.As(err, &errNotFound); ok {
			log.Println(err)

			return echo.ErrNotFound
		}

		return echo.ErrInternalServerError
	}

	return c.JSON(http.StatusOK, st)
}

func (s Student) Register(g *echo.Group) {
	g.GET("", s.GetAll)
	g.GET("/:id", s.Get)
}
