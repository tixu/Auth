package handlers_test

import (
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/tixu/Auth/handlers"
)

func TestVersionCheckHandler(t *testing.T) {
	version := "1.0.0"
	responseexpected := `{"version":"1.0.0"}`
	log.Printf("hello")
	// Create a request to pass to our handler. We don't have any query parameters for now, so we'll
	// pass 'nil' as the third parameter.
	req, err := http.NewRequest("GET", "/version", nil)
	if err != nil {
		t.Fatal(err)
	}
	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	handler := handlers.VersionHandler(version)

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.

	if strings.TrimSpace(rr.Body.String()) != responseexpected {
		t.Errorf("handler returned unexpected body: got -%v- want -%v-",
			rr.Body.String(), responseexpected)
	}
}
