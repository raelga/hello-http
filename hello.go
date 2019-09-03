package main

import (
	"fmt"
	"os"
)

const envKeyHelloName = "HELLO_NAME"
const defaultHelloName = "World"

func helloName(envKey string) string {
	name := os.Getenv(envKey)
	if name == "" {
		return defaultHelloName
	}
	return name
}

func main() {
	fmt.Println("Hello " + helloName(envKeyHelloName) + "!")
}
