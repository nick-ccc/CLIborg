package repository

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// order is version -> date -> image
const changelogTemplate = `
<div align="center">
    <h1>[%s] - %s</h1>
  <a href="">
    <img src="%s" alt="Changelog Image" width="150" />
  </a>
</div>

## Added
- 

## Changed
- 

## Fixed
- 

## Removed
- 

## Deprecated
- 

## Security
- 

`

// Revise this
func CreateChangelog(filepath, version, date, image_src string) error {
	if date == "" {
		date = time.Now().Format("2006-01-02")
	}

	// Ensure the directory exists, create if not
	dir := filepathDir(filepath)
	err := os.MkdirAll(dir, 0755)
	if err != nil {
		return fmt.Errorf("failed to create directory %s: %w", dir, err)
	}

	file, err := os.Create(filepath)
	if err != nil {
		return fmt.Errorf("error creating changelog file: %w", err)
	}
	defer file.Close()

	content := fmt.Sprintf(changelogTemplate, version, date, image_src)
	_, err = file.WriteString(content)
	if err != nil {
		return fmt.Errorf("error writing to changelog file: %w", err)
	}

	fmt.Printf("Changelog file created: %s\n", filepath)
	return nil
}

// filepathDir safely gets directory part of a path
func filepathDir(path string) string {
	dir := filepath.Dir(path)
	if dir == "." || dir == "" {
		return "./"
	}
	return dir
}

func cleanChangelog(content string) string {
	lines := strings.Split(content, "\n")
	var result []string

	var currentSection []string
	validSection := false

	for _, line := range lines {
		if strings.HasPrefix(line, "## ") {
			// flush previous section if valid
			if validSection {
				result = append(result, currentSection...)
			}
			// start new section
			currentSection = []string{line}
			validSection = false
			continue
		}

		// if inside a section (currentSection not empty), accumulate lines
		if len(currentSection) > 0 {
			currentSection = append(currentSection, line)
			// check if line is a bullet with non-empty content
			trimmed := strings.TrimSpace(line)
			if strings.HasPrefix(trimmed, "-") && len(trimmed) > 1 {
				// check that after "-" there is at least one non-whitespace character
				contentAfterDash := strings.TrimSpace(trimmed[1:])
				if len(contentAfterDash) > 0 {
					validSection = true
				}
			}
		} else {
			// outside section, keep line
			result = append(result, line)
		}
	}

	// flush the last section if valid
	if validSection {
		result = append(result, currentSection...)
	}

	return strings.TrimSpace(strings.Join(result, "\n"))
}

func ConsolidateChangelog(filepath string) error {
	_, err := os.Stat(filepath)
	if err != nil {
		return fmt.Errorf("changelog file does not exists")
	}

	changelog, err := os.ReadFile(filepath)
	if err != nil {
		return fmt.Errorf("error reading file: %s", filepath)
	}

	cleaned := cleanChangelog(string(changelog))
	err = os.WriteFile(filepath, []byte(cleaned), 0644)
	if err != nil {
		return fmt.Errorf("failed to write cleaned content to %s: %v", filepath, err)
	}

	fmt.Printf("Changelog file Consolidated: %s\n", filepath)
	return nil
}

func CreateHTMLChangelog() int {
	return 0
}
