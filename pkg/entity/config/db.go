package config

//
// DBConfig for api cmd
//
type DBConfig struct {
	DBFileName int `validate:"required"`
}
