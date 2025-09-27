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

	_ = repository.CreateChangelog(
		"changelogs/CHANGELOG-v0.1.0.md",
		"v0.1.0",
		"",
	)
	// err = repository.ConsolidateChangelog(
	// 	"/git/CLIborg/changelogs/CHANGELOG-v0.1.0.md",
	// )

	// if err != nil {
	// 	log.Fatalf("Error: %v", err)
	// }
}
