package main

import (
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
