package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println(getGreeting("Alan"))
}

func getGreeting(name string) (string, error) {
	if name == "" {
		return "", errors.New("name is empty")
	}
	return fmt.Sprintf("Hello, %s!", name), nil
}
