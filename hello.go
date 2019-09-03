package main

import (
	"fmt"
	"net/http"
	"os"
)

var defaultHTTPPort = 8080

const envKeyHelloName = "HELLO_NAME"
const defaultHelloName = "World"

func helloName(envKey string) string {
	name := os.Getenv(envKey)
	if name == "" {
		return defaultHelloName
	}
	return name
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello %s!", helloName(envKeyHelloName))
}

func main() {
	http.HandleFunc("/", helloHandler)

	if err := http.ListenAndServe(fmt.Sprintf(":%d", defaultHTTPPort), nil); err != nil {
		panic(err)
	}
}
