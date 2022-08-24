package request

import (
	"regexp"

	"github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type Student struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

var StudentIDRegex = regexp.MustCompile("[8-9][0-9][0-9]{2}[0-9]{3}")

func (req Student) Validate() error {
	return validation.ValidateStruct(&req,
		validation.Field(&req.FirstName, validation.Required, validation.Length(1, 255), is.UTFLetterNumeric),
		validation.Field(&req.LastName, validation.Required, validation.Length(1, 255), is.UTFLetterNumeric),
		validation.Field(&req.ID, validation.Required, validation.Length(7, 7), is.Int, validation.Match(StudentIDRegex)),
	)
}
