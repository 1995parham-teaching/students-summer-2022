package store

import (
	"context"

	"github.com/1995parham-teaching/students/internal/model"
	"go.uber.org/zap"
)

type StudentInMemory struct {
	students map[uint64]model.Student
	logger   *zap.Logger
}

func NewStudentInMemory(logger *zap.Logger) *StudentInMemory {
	return &StudentInMemory{
		students: make(map[uint64]model.Student),
		logger:   logger,
	}
}

func (m *StudentInMemory) Save(_ context.Context, s model.Student) error {
	if _, ok := m.students[s.ID]; ok {
		return DuplicateStudentError{
			ID: s.ID,
		}
	}

	m.students[s.ID] = s

	m.logger.Debug("current students", zap.Any("students", m.students))

	return nil
}

func (m *StudentInMemory) Get(_ context.Context, id uint64) (model.Student, error) {
	s, ok := m.students[id]
	if ok {
		return s, nil
	}

	return s, StudentNotFoundError{
		ID: id,
	}
}

func (m *StudentInMemory) GetAll(_ context.Context) ([]model.Student, error) {
	ss := make([]model.Student, 0)

	for _, s := range m.students {
		ss = append(ss, s)
	}

	return ss, nil
}
