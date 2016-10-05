package handlers_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/tixu/Auth/handlers"
	"github.com/tixu/Auth/mocks"
	"github.com/tixu/Auth/users"
)

func TestListAllhHandler(t *testing.T) {

	// Create a request to pass to our handler. We don't have any query parameters for now, so we'll
	// pass 'nil' as the third parameter.
	req, err := http.NewRequest("GET", "/admin/user", nil)
	if err != nil {
		t.Fatal(err)
	}
	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	handler := handlers.GetAdmin("secret", mocks.GetAdminMockService())

	// Our handler returns a funcitons that is able to process httpRequest & httpResponse

	funtotest := handler.ListAll()
	funtotest(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	var us users.Users
	if err = json.Unmarshal(rr.Body.Bytes(), &us); err != nil {
		t.Errorf("unable to unmarshall received body %s",
			string(rr.Body.Bytes()))
	}

	if len(us) != len(mocks.DB) {
		t.Errorf("received users number from mock %d not equals to  %d from Body",
			len(mocks.DB), len(us))
	}
}
