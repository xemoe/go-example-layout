package example

import (
	"os"

	"github.com/spf13/viper"

	log "github.com/sirupsen/logrus"
)

const (
	//
	// DefaultConfigPath for default config.yml file
	//
	DefaultConfigPath = "/etc/go-layout/"
)

//
// InitEnv for setup .env configuration
//
func InitEnv(v *viper.Viper) *viper.Viper {

	v.SetConfigFile(".env")
	v.AddConfigPath(".")
	if err := v.MergeInConfig(); err != nil {
		log.Error(err)
		os.Exit(1)
	}

	log.WithFields(log.Fields{
		"file.env": v.ConfigFileUsed(),
	}).Infof("Using .env config file")

	return v
}

//
// InitYaml for setup config.yaml configuration
//
func InitYaml(v *viper.Viper, cfgFile string) *viper.Viper {

	if cfgFile != "" {
		v.SetConfigFile(cfgFile)
	} else {
		v.SetConfigName("config")
		v.AddConfigPath(DefaultConfigPath)
		v.AddConfigPath(".")
	}

	if err := v.MergeInConfig(); err != nil {
		log.Error(err)
		os.Exit(1)
	}

	log.WithFields(log.Fields{
		"file.configuration": v.ConfigFileUsed(),
	}).Infof("Using yaml config file")

	return v
}
