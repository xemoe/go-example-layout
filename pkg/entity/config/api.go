package config

//
// APIConfig for api cmd
//
type APIConfig struct {
	Port int `validate:"required,min=1,max=65535"`
}
