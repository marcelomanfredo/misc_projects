package main

import (
	"fmt"
	"os"

	"marcelomanfredo/misc/Go/cmd/Validator/args"
)

func main() {
	if err := args.Execute(); err != nil {
		fmt.Println("Error executing command:", err)
		os.Exit(1)
	}
}
