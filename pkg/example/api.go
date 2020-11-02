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

//
// PingHandler for example http handle
//
func PingHandler(w http.ResponseWriter, r *http.Request) {

	log.WithFields(log.Fields{
		"req": r.Method + " ping",
	}).Debug("Receive ping")

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(PingResponseMessage{200, "ok", "Ok!"})
}

//
// Serve for start api server
//
func Serve(port int) {
	http.HandleFunc("/ping", PingHandler)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
