package example

import (
	"encoding/json"
	"fmt"
	"net/http"

	log "github.com/sirupsen/logrus"
)

//
// PingResponseMessage json response
//
type PingResponseMessage struct {
	StatusCode int
	StatusText string
	Message    string
}

func pingHandler(w http.ResponseWriter, r *http.Request) {

	log.WithFields(log.Fields{
		"req": r.Method + " ping",
	}).Debug("Receive ping")

	//
	// Prepare response message
	//
	message := PingResponseMessage{200, "ok", "Ok!"}

	js, err := json.Marshal(message)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

//
// Serve for start api server
//
func Serve(port int) {
	http.HandleFunc("/ping", pingHandler)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
