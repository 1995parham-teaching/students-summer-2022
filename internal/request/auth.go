package request

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type Login struct {
	Username string
	Password string
}

// nolint: wrapcheck
func (req Login) Validate() error {
	return validation.ValidateStruct(&req,
		validation.Field(&req.Username,
			validation.Required, is.Alphanumeric),
		validation.Field(&req.Password,
			validation.Required, is.Alphanumeric),
	)
}
