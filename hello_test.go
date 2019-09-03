package main

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

type testEnv struct {
	set   bool
	key   string
	value string
}

func Test_helloName(t *testing.T) {

	tests := []struct {
		name string
		want string
		env  testEnv
	}{
		{name: "unsetEnvVar", want: "World", env: testEnv{set: false, key: "HELLO_NAME", value: ""}},
		{name: "emptyEnvVar", want: "World", env: testEnv{set: true, key: "HELLO_NAME", value: ""}},
		{name: "setEnvVar", want: "Rael", env: testEnv{set: true, key: "HELLO_NAME", value: "Rael"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.env.set {
				os.Setenv(tt.env.key, tt.env.value)
				defer os.Unsetenv(tt.env.key)
			} else {
				os.Unsetenv(tt.env.key)
			}
			if got := helloName(tt.env.key); got != tt.want {
				t.Errorf("helloName(\"%v\") = %v, want %v", tt.env.key, got, tt.want)
			}
		})
	}
}

func Test_helloHandler(t *testing.T) {
	type testReq struct {
		path     string
		httpCode int
	}
	tests := []struct {
		name string
		want string
		req  testReq
	}{
		{name: "rootWorldRequest", want: "Hello World!", req: testReq{path: "/", httpCode: 200}},
		{name: "fakePathRequest", want: "Hello World!", req: testReq{path: "/fakePathRequest", httpCode: 200}},
	}

	for _, tt := range tests {

		// Create the test HTTP request
		req, err := http.NewRequest("GET", tt.req.path, nil)
		if err != nil {
			t.Errorf("main() NewRequest failed: %v", err)
		}

		// Execute the HTTP request
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(helloHandler)
		handler.ServeHTTP(rr, req)

		t.Run(tt.name, func(t *testing.T) {
			if got := rr.Code; got != tt.req.httpCode {
				t.Errorf("helloHandler() Status = %v, want %v", got, tt.want)
			}
		})

		t.Run(tt.name, func(t *testing.T) {
			if got := rr.Body.String(); got != tt.want {
				t.Errorf("helloHandler() Body = %v, want %v", got, tt.want)
			}
		})
	}
}
