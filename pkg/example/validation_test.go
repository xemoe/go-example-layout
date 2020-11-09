package example_test

import (
	"testing"

	"github.com/xemoe/go-example-layout/pkg/entity/config"
	"github.com/xemoe/go-example-layout/pkg/example"
)

func TestValidateAPIConfig(t *testing.T) {

	err1 := example.ValidateAPIConfig(
		&config.APIConfig{
			Port: 0,
		}, false)
	if err1 == nil {
		t.Errorf("should error when invalid port number")
	}

	err2 := example.ValidateAPIConfig(
		&config.APIConfig{
			Port: 1,
		}, false)
	if err2 != nil {
		t.Errorf("should not error when valid port number")
	}
}
