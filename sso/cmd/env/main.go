package main

import (
	"fmt"
	"os"
)

func main() {
	os.Setenv("CONFIG_PATH", "../../config/local.yaml")

	// Get the USERNAME environment variable
	configPath := os.Getenv("CONFIG_PATH")

	fmt.Printf("%s", configPath)
}
