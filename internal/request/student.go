package request

import (
	"regexp"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

const (
	StudentIDLen    = 7
	FirstNameMaxLen = 255
	FirstNameMinLen = 1
	LastNameMaxLen  = 255
	LastNameMinLen  = 1
)

type Student struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

var StudentIDRegex = regexp.MustCompile("[8-9][0-9][0-9]{2}[0-9]{3}")

// nolint: wrapcheck
func (req Student) Validate() error {
	return validation.ValidateStruct(&req,
		validation.Field(&req.FirstName,
			validation.Required, validation.Length(FirstNameMinLen, FirstNameMaxLen), is.UTFLetterNumeric),
		validation.Field(&req.LastName,
			validation.Required, validation.Length(LastNameMinLen, LastNameMaxLen), is.UTFLetterNumeric),
		validation.Field(&req.ID, validation.Required,
			validation.Length(StudentIDLen, StudentIDLen), is.Int, validation.Match(StudentIDRegex)),
	)
}
