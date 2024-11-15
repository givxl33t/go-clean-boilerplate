package infrastructure

import (
	"github.com/spf13/viper"

	"github.com/go-playground/validator/v10"
)

func NewValidator(config *viper.Viper) *validator.Validate {
	return validator.New()
}
