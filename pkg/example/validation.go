package example

import (
	"os"

	"github.com/go-playground/validator/v10"
	log "github.com/sirupsen/logrus"
)

// use a single instance of Validate, it caches struct info
var validate *validator.Validate

//
// ValidateAPIConfig for validate APIConfig struct
//
func ValidateAPIConfig(c *APIConfig) {

	validate = validator.New()

	err := validate.Struct(c)
	if err != nil {

		if _, ok := err.(*validator.InvalidValidationError); ok {
			log.Error(err)
		}

		for _, err := range err.(validator.ValidationErrors) {
			log.Error(err)
		}

		os.Exit(1)
	}
}
