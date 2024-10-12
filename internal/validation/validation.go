package validation

import (
	"slices"

	"github.com/go-playground/validator/v10"
)

func New() *validator.Validate {
	v := validator.New()
	v.RegisterValidation("valid_env", ValidateEnv)
	v.RegisterValidation("valid_owner_type", ValidateOwnerType)

	return v
}

func ValidateEnv(fl validator.FieldLevel) bool {
	return slices.Contains(
		[]string{
			"ENVIRONMENT_STAGING",
			"ENVIRONMENT_PRODUCTION",
		},
		fl.Field().String(),
	)
}

func ValidateOwnerType(fl validator.FieldLevel) bool {
	return slices.Contains(
		[]string{
			"OWNER_TYPE_INDIVIDUAL",
			"OWNER_TYPE_ORGANIZATION",
		},
		fl.Field().String(),
	)
}
