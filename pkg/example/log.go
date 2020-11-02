package example

import (
	nested "github.com/antonfisher/nested-logrus-formatter"
	log "github.com/sirupsen/logrus"
)

//
// InitLogger for setup logging format
//
func InitLogger() {

	log.SetFormatter(&nested.Formatter{
		HideKeys:       false,
		NoColors:       false,
		NoFieldsColors: false,
		ShowFullLevel:  true,
		FieldsOrder:    []string{"process", "status"},
	})

	log.SetLevel(log.DebugLevel)

	//
	// Enable to show caller
	//
	// log.SetReportCaller(true)
}

func init() {
	InitLogger()
}
