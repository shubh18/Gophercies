package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSourceCodeHandler(t *testing.T) {

	testTable := []struct {
		testCaseName string
		url          string
		status       int
	}{
		{
			testCaseName: "TC1",
			url:          "line=24&path=/home/gslab/go/src/recover/main.go",
			status:       200,
		}, {
			testCaseName: "TC2",
			url:          "line=ewr&path=/home/gslab/go/src/recover/main.go",
			status:       200,
		},
		{
			testCaseName: "TC3",
			url:          "line=24&path=/home/gslab/go/main.go",
			status:       500,
		},
	}
	for i := 0; i < len(testTable); i++ {
		req, err := http.NewRequest("GET", "http://localhost:8000/debug/?"+testTable[i].url, nil)
		if err != nil {
			t.Fatalf("could not create request: %v", err)
		}
		rec := httptest.NewRecorder()
		SourceCodeHandler(rec, req)
		res := rec.Result()
		if res.StatusCode != testTable[i].status {
			t.Errorf("Test case Number: %v Expected %v , Actual status %v", testTable[i].testCaseName, testTable[i].status, res.StatusCode)
		}
	}

}

func TestMiddleware(t *testing.T) {
	handler := http.HandlerFunc(PanicHandler)
	executeRequest("Get", "/panic", Middleware(handler))
}

func executeRequest(method string, url string, handler http.Handler) (*httptest.ResponseRecorder, error) {
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}
	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)
	return rr, err
}
