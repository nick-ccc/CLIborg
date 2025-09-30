package main

import (
	"log"

	"github.com/nick-ccc/CLIborg/internal/git"
	"github.com/nick-ccc/CLIborg/internal/repository"
)

func main() {
	_, err := git.LogChanges()
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	// fmt.Println(m)

	// _ = repository.CreateChangelog(
	// 	"changelogs/CHANGELOG-v0.1.0.md",
	// 	"v0.1.0",
	// 	"",
	// 	"https://go.dev/blog/go-brand/Go-Logo/SVG/Go-Logo_Aqua.svg",
	// )
	err = repository.ConsolidateChangelog(
		"/git/CLIborg/changelogs/CHANGELOG-v0.1.0.md",
	)

	if err != nil {
		log.Fatalf("Error: %v", err)
	}
}
