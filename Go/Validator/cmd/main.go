package main

import (
	"fmt"
	"os"

	"Validator/internal"
)

func main() {
	if err := internal.Execute(); err != nil {
		fmt.Println("Error executing command:", err)
		os.Exit(1)
	}
}
