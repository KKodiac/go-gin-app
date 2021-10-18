// handlers.article_test.go
// Testing to check for following conditions :
// 1. The handler responds with an HTTP status code of 200
// 2. The returned HTML contains a title tag containing the text "Home Page"
// 3. Test that the application returns a JSON list of articles when the Accept header is set to "application/json"
// 4. Test that the application returns a XML list of articles when the Accept header is set to "application/xml"
package main

// Test that a GET request to the home page returns the list of articles
// in JSON format when the Accept header is set to application/json
// func TestArticleListJSON(t *testing.T) {
// 	r := getRouter(true)

// 	// Define the route similar to its definition in the routes file
// 	r.GET("/", showIndexPage)

// 	// Create a request to send to the above route
// 	req, _ := http.NewRequest("GET", "/", nil)
// 	req.Header.Add("Accept", "application/json")

// 	testHTTPResponse(t, r, req, func(w *httptest.ResponseRecorder) bool {
// 		// Test that the http status code is 200
// 		statusOK := w.Code == http.StatusOK

// 		// Test that the response is JSON which can be converted to
// 		// an array of Article structs
// 		p, err := ioutil.ReadAll(w.Body)
// 		if err != nil {
// 			return false
// 		}
// 		var articles []article
// 		err = json.Unmarshal(p, &articles)

// 		return err == nil && len(articles) >= 2 && statusOK
// 	})
// }

// Test that a GET request to an article page returns the article in XML
// format when the Accept header is set to application/xml
// func TestArticleXML(t *testing.T) {
// 	r := getRouter(true)

// 	// Define the route similar to its definition in the routes file
// 	r.GET("/article/view/:article_id", getArticle)

// 	// Create a request to send to the above route
// 	req, _ := http.NewRequest("GET", "/article/view/1", nil)
// 	req.Header.Add("Accept", "application/xml")

// 	testHTTPResponse(t, r, req, func(w *httptest.ResponseRecorder) bool {
// 		// Test that the http status code is 200
// 		statusOK := w.Code == http.StatusOK

// 		// Test that the response is JSON which can be converted to
// 		// an array of Article structs
// 		p, err := ioutil.ReadAll(w.Body)
// 		if err != nil {
// 			return false
// 		}
// 		var a article
// 		err = xml.Unmarshal(p, &a)

// 		return err == nil && a.ID == 1 && len(a.Title) >= 0 && statusOK
// 	})
// }
