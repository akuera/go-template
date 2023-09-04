package requests

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type CreateDomain struct {
	Name string `json:"name"`
}

func (r CreateDomain) Validate() error {
	return validation.ValidateStruct(&r,
		validation.Field(&r.Name, validation.Required))
}

type UpdateDomain struct {
	Name string `json:"name"`
}
