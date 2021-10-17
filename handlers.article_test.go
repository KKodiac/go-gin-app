// handlers.article_test.go
// Testing to check for following conditions :
// 1. The handler responds with an HTTP status code of 200
// 2. The returned HTML contains a title tag containing the text "Home Page"
// 3. Test that the application returns a JSON list of articles when the Accept header is set to "application/json"
// 4. Test that the application returns a XML list of articles when the Accept header is set to "application/xml"
package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// Tests sending GET Request to "/" page and receiving status code 200
// for unauthenticated user
func TestShowIndexPageUnauthenticated(t *testing.T) {
	// Creates router
	r := getRouter(true)

	// Sends GET request to the root home page "/"
	r.GET("/", showIndexPage)

	// response received from the previous GET request
	req, _ := http.NewRequest("GET", "/", nil)

	testHTTPResponse(t, r, req, func(w *httptest.ResponseRecorder) bool {
		// Tests if HTTP status code is 200 OK
		statusOK := w.Code == http.StatusOK

		// Reads previously received response HTML.
		// Checks if Title tag of the HTML is "Home Page"
		// A lot more testing can be done by using html parser libraries
		p, err := ioutil.ReadAll(w.Body)
		pageOK := err == nil && strings.Index(string(p), "<title>Home Page</title>") > 0

		return statusOK && pageOK
	})
}
