package example

import (
	"github.com/go-playground/validator/v10"
	log "github.com/sirupsen/logrus"

	"github.com/xemoe/go-example-layout/pkg/entity/config"
)

// use a single instance of Validate, it caches struct info
var validate *validator.Validate

//
// ValidateAPIConfig for validate APIConfig struct
//
func ValidateAPIConfig(c *config.APIConfig, errorLog bool) error {

	validate = validator.New()

	err := validate.Struct(c)
	if err != nil {

		if _, ok := err.(*validator.InvalidValidationError); ok {
			if errorLog {
				log.Error(err)
			}
			return err
		}

		for _, err := range err.(validator.ValidationErrors) {
			if errorLog {
				log.Error(err)
			}
			return err
		}
	}

	return nil
}
