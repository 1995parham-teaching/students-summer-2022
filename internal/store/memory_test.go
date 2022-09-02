package store_test

import (
	"context"
	"testing"

	"github.com/1995parham-teaching/students/internal/model"
	"github.com/1995parham-teaching/students/internal/store"
	"go.uber.org/zap"
)

func TestInMemorySave(t *testing.T) {
	st := store.NewStudentInMemory(zap.NewNop())
	ctx := context.Background()

	st.Save(ctx, model.Student{
		ID:        9231058,
		FirstName: "Parham",
		LastName:  "Alvani",
	})

	m, err := st.Get(ctx, 9231058)
	if err != nil {
		t.Fatal(err)
	}

	if m.FirstName != "Parham" {
		t.Fatal("first name should be Parham")
	}

	if m.LastName != "Alvani" {
		t.Fatal("last name should be Alvani")
	}
}
