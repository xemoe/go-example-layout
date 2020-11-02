package example

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func AreEqualJSON(s1, s2 string) (bool, error) {

	var o1 interface{}
	var o2 interface{}

	var err error
	err = json.Unmarshal([]byte(s1), &o1)
	if err != nil {
		return false, fmt.Errorf("Error mashalling string 1 :: %s", err.Error())
	}
	err = json.Unmarshal([]byte(s2), &o2)
	if err != nil {
		return false, fmt.Errorf("Error mashalling string 2 :: %s", err.Error())
	}

	return reflect.DeepEqual(o1, o2), nil
}

//
// ref:
// - https://blog.questionable.services/article/testing-http-handlers-go/
// - https://bl.ocks.org/turtlemonvh/e4f7404e28387fadb8ad275a99596f67
//
func TestHTTPPing(t *testing.T) {
	t.Run("it should return httpCode 200", func(t *testing.T) {

		// Create a request to pass to our handler. We don't have any query parameters for now, so we'll
		// pass 'nil' as the third parameter.
		req, err := http.NewRequest(http.MethodGet, "/ping", nil)
		if err != nil {
			t.Error(err)
		}

		// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
		resp := httptest.NewRecorder()
		handler := http.HandlerFunc(PingHandler)

		// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
		// directly and pass in our Request and ResponseRecorder.
		handler.ServeHTTP(resp, req)

		// Check the status code is what we expect.
		if status := resp.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
		}

		// Check the response body is what we expect.
		expected := `{"StatusCode":200,"StatusText":"ok","Message":"Ok!"}`
		areEqual, err := AreEqualJSON(resp.Body.String(), expected)
		if err != nil {
			t.Error("Error mashalling strings", err.Error())
		}

		if !areEqual {
			t.Errorf("handler returned unexpected body: got %v want %v",
				resp.Body.String(), expected)
		}
	})
}
