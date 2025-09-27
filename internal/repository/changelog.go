package repository

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
)

const changelogTemplate = `
<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8" />
  <meta name="viewport" content="width=device-width, initial-scale=1.0" />
  <title>Changelog v0.1.0</title>
  <style>

    .container {
      max-width: 700px;
      margin: auto;
      background: white;
      box-shadow: 0 4px 10px rgba(0, 0, 0, 0.1);
      border-radius: 8px;
      padding: 2rem;
    }

    .logo {
      text-align: center;
      margin-bottom: 1rem;
    }

    .logo img {
      width: 100px;
    }

    h1 {
      text-align: center;
      color: #4A90E2;
      font-size: 1.8rem;
      margin-bottom: 0.5rem;
    }

    .section {
      margin-top: 1.5rem;
    }

    .section h2 {
      font-size: 1.2rem;
      padding: 0.5rem;
      color: white;
      border-radius: 5px;
    }

    .added h2 {
    background: linear-gradient(to right, #0273a4ff, transparent);
    }
    .changed h2 {
    background: linear-gradient(to right, #00857cff, transparent);
    }
    .fixed h2 {
    background: linear-gradient(to right, #006e4fff, transparent);
    }
    .removed h2 {
    background: linear-gradient(to right, #9e1010ff, transparent);
    }
    .deprecated h2 {
    background: linear-gradient(to right, #9b59b6, transparent);
    }
    .security h2 {
    background: linear-gradient(to right, #7d22e6ff, transparent);
    }


    ul {
      margin-top: 0.5rem;
      padding-left: 1.5rem;
    }

    li {
      margin: 0.5rem 0;
    }
  </style>
</head>
<body>
  <div class="container">
    <div class="logo">
      <a href="">
        <img src="INSERT_IMAGE_SRC_HERE" alt="Changelog Image" />
      </a>
    </div>
    <h1>[INSERT_VERSION_HERE] - INSERT_DATE_SRC_HERE</h1>

<div class="section added">
  <h2>➕ Added</h2>
  <ul>
    <li></li>
  </ul>
</div>

<div class="section changed">
  <h2>🛠 Changed</h2>
  <ul>
    <li></li>
  </ul>
</div>

<div class="section fixed">
  <h2>🧰 Fixed</h2>
  <ul>
    <li></li>
  </ul>
</div>

<div class="section removed">
  <h2>🗑 Removed</h2>
  <ul>
    <li></li>
  </ul>
</div>

<div class="section deprecated">
  <h2>⚠️ Deprecated</h2>
  <ul>
    <li></li>
  </ul>
</div>

<div class="section security">
  <h2>🛡 Security</h2>
  <ul>
    <li></li>
  </ul>
</div>
  </div>
</body>
</html>

`

// Revise this
func CreateChangelog(filepath string, version string, date string) error {
	if date == "" {
		date = time.Now().Format("2001-12-01")
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

	content := fmt.Sprintf(changelogTemplate)
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
	inSection := false
	validSection := false

	for _, line := range lines {
		if strings.HasPrefix(line, "### ") {
			// flush previous section if valid
			if validSection {
				result = append(result, currentSection...)
			}
			// reset section
			currentSection = []string{line}
			inSection = true
			validSection = false
			continue
		}

		if inSection {
			currentSection = append(currentSection, line)
			if strings.HasPrefix(strings.TrimSpace(line), "-") {
				validSection = true
			}
		} else {
			// outside of section, just append
			result = append(result, line)
		}
	}

	// flush the last section if it's valid
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
