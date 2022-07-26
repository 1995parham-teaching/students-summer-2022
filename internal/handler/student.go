package handler

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/1995parham-teaching/students/internal/model"
	"github.com/1995parham-teaching/students/internal/request"
	"github.com/1995parham-teaching/students/internal/store"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

const (
	IDBase = 10
	IDBits = 64
)

type Student struct {
	Store  store.Student
	Logger *zap.Logger
}

func (s Student) GetAll(c echo.Context) error {
	ctx := c.Request().Context()

	ss, err := s.Store.GetAll(ctx)
	if err != nil {
		return echo.ErrInternalServerError
	}

	return c.JSON(http.StatusOK, ss)
}

func (s Student) Get(c echo.Context) error {
	ctx := c.Request().Context()

	id, err := strconv.ParseUint(c.Param("id"), IDBase, IDBits)
	if err != nil {
		return echo.ErrBadRequest
	}

	st, err := s.Store.Get(ctx, id)
	if err != nil {
		var errNotFound store.StudentNotFoundError
		if ok := errors.As(err, &errNotFound); ok {
			s.Logger.Error("student not found", zap.Error(err))

			return echo.ErrNotFound
		}

		return echo.ErrInternalServerError
	}

	return c.JSON(http.StatusOK, st)
}

func (s Student) Create(c echo.Context) error {
	ctx := c.Request().Context()

	var req request.Student

	if err := c.Bind(&req); err != nil {
		s.Logger.Error("cannot bind request to student",
			zap.Error(err),
		)

		return echo.ErrBadRequest
	}

	if err := req.Validate(); err != nil {
		s.Logger.Error("request validation failed",
			zap.Error(err),
			zap.Any("request", req),
		)

		return echo.ErrBadRequest
	}

	id, _ := strconv.ParseUint(req.ID, IDBase, IDBits)

	m := model.Student{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		ID:        id,
		Average:   0,
	}

	if err := s.Store.Save(ctx, m); err != nil {
		var errDuplicate store.DuplicateStudentError
		if ok := errors.As(err, &errDuplicate); ok {
			s.Logger.Error("duplicate student",
				zap.Error(err),
				zap.Uint64("id", m.ID),
			)

			return echo.ErrBadRequest
		}

		return echo.ErrInternalServerError
	}

	s.Logger.Info("student creation success")

	return c.JSON(http.StatusCreated, m)
}

func (s Student) Register(g *echo.Group) {
	g.GET("", s.GetAll)
	g.GET("/:id", s.Get)
	g.POST("", s.Create)
}
