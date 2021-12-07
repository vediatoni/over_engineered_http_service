package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
	"time"
)

func TestRun(t *testing.T) {
	svc := new(getPort())

	serviceRunning := make(chan struct{})
	serviceDone := make(chan struct{})
	go func() {
		close(serviceRunning)
		if err := svc.run(); err != http.ErrServerClosed && err != nil {
			t.Errorf("failed to run service: %v", err)
			return
		}
		defer close(serviceDone)
	}()

	<-serviceRunning

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err := svc.httpServer.Shutdown(ctx)
	if err != nil {
		t.Fatalf("failed to shutdown server: %v", err)
	}

	<-serviceDone
}

func TestCustomPort(t *testing.T) {
	os.Setenv("PORT", "8086")
	svc := new(getPort())

	serviceRunning := make(chan struct{})
	serviceDone := make(chan struct{})
	go func() {
		close(serviceRunning)
		if err := svc.run(); err != http.ErrServerClosed && err != nil {
			t.Errorf("failed to run service: %v", err)
			return
		}
		defer close(serviceDone)
	}()

	<-serviceRunning

	if svc.httpServer.Addr != ":8086" {
		t.Errorf("failed to set custom port")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err := svc.httpServer.Shutdown(ctx)
	if err != nil {
		t.Errorf("failed to shutdown server: %v", err)
	}

	<-serviceDone
}

func TestService_Custom(t *testing.T) {
	svc := new(getPort())

	type response struct {
		res        string
		statusCode int
	}
	testCases := []struct {
		name    string
		request string
		expect  *response
	}{
		{
			name:    "Test 1",
			request: "/NOOOO/data",
			expect: &response{
				res:        FailedToParseAccountID,
				statusCode: http.StatusBadRequest,
			},
		},
		{
			name:    "Test 2",
			request: "/",
			expect: &response{
				res:        "Hello",
				statusCode: http.StatusOK,
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			r := httptest.NewRequest("GET", testCase.request, nil)
			w := httptest.NewRecorder()
			svc.accountId(w, r)

			resp := w.Result()
			defer resp.Body.Close()
			out, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				t.Fatalf("failed to decode response: %v", err)
			}

			if resp.StatusCode != testCase.expect.statusCode {
				t.Fatalf("expected %v, got %v", testCase.expect.statusCode, resp.StatusCode)
			}

			result := strings.TrimRight(string(out), "\n")
			if result != testCase.expect.res {
				t.Fatalf("expected %v, got %v", testCase.expect.res, result)
			}
		})
	}
}

func TestService_NumHandler(t *testing.T) {
	svc := new(getPort())

	type response struct {
		res        *responsePayload
		statusCode int
	}
	testCases := []struct {
		name    string
		request string
		expect  *response
	}{
		{
			name:    "Test 1",
			request: "/1/data",
			expect: &response{
				res: &responsePayload{
					Data:      RandomText,
					AccountID: 1,
				},
				statusCode: http.StatusOK,
			},
		},
		{
			name:    "Test 2",
			request: "/55/data",
			expect: &response{
				res: &responsePayload{
					Data:      RandomText,
					AccountID: 55,
				},
				statusCode: http.StatusOK,
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			r := httptest.NewRequest("GET", testCase.request, nil)
			w := httptest.NewRecorder()
			svc.accountId(w, r)

			resp := w.Result()
			defer resp.Body.Close()
			out, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				t.Fatalf("failed to decode response: %v", err)
			}

			if resp.StatusCode != testCase.expect.statusCode {
				t.Fatalf("expected %v, got %v", testCase.expect.statusCode, resp.StatusCode)
			}

			response := &responsePayload{}
			err = json.Unmarshal(out, response)
			if err != nil {
				t.Fatalf("failed to decode response: %v", err)
			}

			if response.AccountID != testCase.expect.res.AccountID {
				t.Fatalf("expected %v, got %v", testCase.expect.res.AccountID, response.AccountID)
			}

			if response.Data != testCase.expect.res.Data {
				t.Fatalf("expected %v, got %v", testCase.expect.res.Data, response.Data)
			}
		})
	}
}

func TestService_HealthCheck(t *testing.T) {
	r := httptest.NewRequest("GET", "/healtz", nil)
	w := httptest.NewRecorder()

	svc := new(getPort())

	svc.healthCheck(w, r)
	resp := w.Result()
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK{
		t.Fatalf("Expected %v status code but got %v", http.StatusOK, resp.StatusCode)
	}
}
