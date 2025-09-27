package main

import (
	"fmt"
	"log"

	"github.com/nick-ccc/CLIborg/internal/git"
)

func main() {
	branch, err := git.CurrentBranch()
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	fmt.Println("Default branch:", branch)
}
