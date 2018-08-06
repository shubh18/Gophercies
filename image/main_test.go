package main

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/CloudBroker/dash_utils/dashtest"
)

func getRequestBody(t *testing.T) *bytes.Buffer {
	file, err := os.Open("./input.png")
	if err != nil {
		t.Error("error in opening file")
	}
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("image", file.Name())
	if err != nil {
		t.Error("Expected nil got", err)
	}
	_, err = io.Copy(part, file)
	if err != nil {
		t.Error("Expected nil got", err)
	}
	err = writer.Close()
	if err != nil {
		t.Error("Expected nil got", err)
	}
	return body
}

func TestIndex(t *testing.T) {
	testServer := httptest.NewServer(getHandlers())
	defer testServer.Close()

	newreq := func(method, url string, body io.Reader) *http.Request {
		req, err := http.NewRequest(method, url, body)
		if err != nil {
			t.Fatal(err)
		}
		return req
	}

	testCases := []struct {
		name   string
		req    *http.Request
		status int
	}{
		{name: "TC1", req: newreq("GET", testServer.URL+"/", nil), status: 200},
	}
	for _, tests := range testCases {
		t.Run(tests.name, func(t *testing.T) {
			resp, err := http.DefaultClient.Do(tests.req)
			if err != nil {
				t.Fatal(err)
			}
			defer resp.Body.Close()
			if resp.StatusCode != tests.status {
				t.Error("Expected 200 got", tests.status)
			}
		})
	}
}

func TestUpload(t *testing.T) {
	testServer := httptest.NewServer(getHandlers())
	defer testServer.Close()

	newreq := func(method, url string) *http.Request {
		file, err := os.Open("./input.png")
		if err != nil {
			t.Error("Expected nil got", err)
		}
		body := &bytes.Buffer{}
		writer := multipart.NewWriter(body)
		part, err := writer.CreateFormFile("image", file.Name())
		if err != nil {
			t.Error("Expected nil got", err)
		}
		_, err = io.Copy(part, file)
		if err != nil {
			t.Error("Expected nil got", err)
		}
		err = writer.Close()
		if err != nil {
			t.Error("Expected nil got", err)
		}
		req, err := http.NewRequest(method, url, body)
		if err != nil {
			t.Fatal(err)
		}
		req.Header.Set("Content-Type", writer.FormDataContentType())
		return req
	}

	testCases := []struct {
		name   string
		req    *http.Request
		status int
	}{
		{name: "TC2", req: newreq("POST", testServer.URL+"/upload"), status: 200},
	}
	for _, tests := range testCases {
		t.Run(tests.name, func(t *testing.T) {
			resp, err := http.DefaultClient.Do(tests.req)
			if err != nil {
				t.Fatal(err)
			}
			defer resp.Body.Close()
			if resp.StatusCode != tests.status {
				t.Error("Expected 200 got", tests.status)
			}
		})
	}
}

func TestModify(t *testing.T) {
	testServer := httptest.NewServer(getHandlers())
	defer testServer.Close()

	newreq := func(method, url string, body io.Reader) *http.Request {
		req, err := http.NewRequest(method, url, body)
		if err != nil {
			t.Fatal(err)
		}
		return req
	}

	testCases := []struct {
		name   string
		req    *http.Request
		status int
	}{
		{name: "TC3", req: newreq("GET", testServer.URL+"/modify/img/690011768.png?mode=2", nil), status: 200},
		{name: "TC4", req: newreq("GET", testServer.URL+"/modify/img/011501652.png?mode=2&number=100", nil), status: 200},
	}
	for _, tests := range testCases {
		t.Run(tests.name, func(t *testing.T) {
			resp, err := http.DefaultClient.Do(tests.req)
			if err != nil {
				t.Error("Expected nil got", err)
			}
			defer resp.Body.Close()
			if resp.StatusCode != tests.status {
				t.Error("Expected 200 got", resp.StatusCode)
			}
		})
	}
}

func TestTemp(t *testing.T) {
	_, err := tempfile("test", "png")
	if err != nil {
		t.Error("Expected nil got", err)
	}
}

func TestImage(t *testing.T) {
	var image io.Reader
	image, _ = os.Open("input.png")
	testImage, err := genrateImage(image, "png", "2", "10")
	if err != nil {
		t.Error("Expected image got", err, testImage)
	}
}

func TestMain(m *testing.M) {
	dashtest.ControlCoverage(m)
	main()
}
