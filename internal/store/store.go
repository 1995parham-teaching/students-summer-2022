package store

import (
	"context"
	"fmt"

	"github.com/1995parham-teaching/students/internal/model"
)

type StudentNotFoundError struct {
	ID uint64
}

func (err StudentNotFoundError) Error() string {
	return fmt.Sprintf("student %d doesn't exist", err.ID)
}

type DuplicateStudentError struct {
	ID uint64
}

func (err DuplicateStudentError) Error() string {
	return fmt.Sprintf("student %d already exists", err.ID)
}

type Student interface {
	Save(context.Context, model.Student) error
	Get(context.Context, uint64) (model.Student, error)
	GetAll(context.Context) ([]model.Student, error)
}
