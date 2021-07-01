package handler

import (
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAdd(t *testing.T) {
	tests := []struct {
		name       string
		params     map[string]string
		statusCode int
	}{
		{
			name: "Addition good test",
			params: map[string]string{
				"a": "5",
				"b": "5",
			},
			statusCode: 200,
		},
		{
			name: "Addition bad test",
			params: map[string]string{
				"a": "5",
				"b": "",
			},
			statusCode: 400,
		},
	}

	for _, tt := range tests {
		req, _ := http.NewRequest("GET", "/api/add", nil)
		q := req.URL.Query()

		for k, v := range tt.params {
			q.Add(k, v)
		}

		req.URL.RawQuery = q.Encode()
		rec := httptest.NewRecorder()
		Add(rec, req)
		res := rec.Result()

		if res.StatusCode != tt.statusCode {
			t.Errorf("Got response status code: %v, expected response status code: %v", res.StatusCode, tt.statusCode)
		}
	}
}

func TestSub(t *testing.T) {
	tests := []struct {
		name       string
		params     map[string]string
		statusCode int
	}{
		{
			name: "Subtraction good test",
			params: map[string]string{
				"a": "5",
				"b": "5",
			},
			statusCode: 200,
		},
		{
			name: "Subtraction bad test",
			params: map[string]string{
				"a": "5",
				"b": "",
			},
			statusCode: 400,
		},
	}

	for _, tt := range tests {
		req, _ := http.NewRequest("GET", "/api/sub", nil)
		q := req.URL.Query()

		for k, v := range tt.params {
			q.Add(k, v)
		}

		req.URL.RawQuery = q.Encode()
		rec := httptest.NewRecorder()
		Sub(rec, req)
		res := rec.Result()

		if res.StatusCode != tt.statusCode {
			t.Errorf("Got response status code: %v, expected response status code: %v", res.StatusCode, tt.statusCode)
		}
	}
}

func TestMul(t *testing.T) {
	tests := []struct {
		name       string
		params     map[string]string
		statusCode int
	}{
		{
			name: "Multiplication good test",
			params: map[string]string{
				"a": "5",
				"b": "5",
			},
			statusCode: 200,
		},
		{
			name: "Multiplication bad test",
			params: map[string]string{
				"a": "5",
				"b": "",
			},
			statusCode: 400,
		},
	}

	for _, tt := range tests {
		req, _ := http.NewRequest("GET", "/api/mul", nil)
		q := req.URL.Query()

		for k, v := range tt.params {
			q.Add(k, v)
		}

		req.URL.RawQuery = q.Encode()
		rec := httptest.NewRecorder()
		Mul(rec, req)
		res := rec.Result()

		if res.StatusCode != tt.statusCode {
			t.Errorf("Got response status code: %v, expected response status code: %v", res.StatusCode, tt.statusCode)
		}
	}
}

func TestDiv(t *testing.T) {
	tests := []struct {
		name       string
		params     map[string]string
		statusCode int
	}{
		{
			name: "Division good test",
			params: map[string]string{
				"a": "5",
				"b": "5",
			},
			statusCode: 200,
		},
		{
			name: "Subtraction bad test",
			params: map[string]string{
				"a": "5",
				"b": "",
			},
			statusCode: 400,
		},
		{
			name: "Subtraction bad test - zero division",
			params: map[string]string{
				"a": "5",
				"b": "0",
			},
			statusCode: 400,
		},
	}

	for _, tt := range tests {
		log.Println("Test with params: ", tt.params)
		req, _ := http.NewRequest("GET", "/api/div", nil)
		q := req.URL.Query()

		for k, v := range tt.params {
			q.Add(k, v)
		}

		req.URL.RawQuery = q.Encode()
		rec := httptest.NewRecorder()
		Div(rec, req)
		res := rec.Result()

		if res.StatusCode != tt.statusCode {
			t.Errorf("Got response status code: %v, expected response status code: %v", res.StatusCode, tt.statusCode)
		}
	}
}
