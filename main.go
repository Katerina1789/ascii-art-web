package main

import (
	ascii "asciiWeb/internal/utilities"
	"fmt"
)

func main() {
	if ascii.EnsureFontFiles() != nil {
		fmt.Printf("Error retriving the files")
		return
	}
}
