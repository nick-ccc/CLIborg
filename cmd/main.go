package main

import (
	"log"

	"github.com/nick-ccc/CLIborg/internal/git"
)

func main() {
	_, err := git.StageAndCommitTracked("Test staging and commit")
	_, err = git.Push("origin", "main")
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	// fmt.Println("Default branch:", branch)
}
