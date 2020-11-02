package example

import (
	nested "github.com/antonfisher/nested-logrus-formatter"
	log "github.com/sirupsen/logrus"
)

func init() {

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

//
// SaySomething for example function
//
func SaySomething(words string) {
	log.WithFields(log.Fields{
		"words": words,
	}).Infof("Say: %s", words)
}
