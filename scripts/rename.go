package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// Change this to the name you want to replace "MBFacto" with
const oldName = "MBFacto"
const newName = "MBFacto"

func main() {
	rootDir := "." // Start in the current directory

	err := filepath.Walk(rootDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Process only .go files and go.mod
		if !info.IsDir() && (strings.HasSuffix(info.Name(), ".go") || info.Name() == "go.mod") {
			replaceInFile(path)
		}
		return nil
	})

	if err != nil {
		fmt.Println("Error:", err)
	}
}

func replaceInFile(path string) {
	content, err := os.ReadFile(path)
	if err != nil {
		fmt.Printf("Failed to read file: %s, Error: %v\n", path, err)
		return
	}

	newContent := strings.ReplaceAll(string(content), oldName, newName)
	if string(content) == newContent {
		return // No changes, skip writing
	}

	err = os.WriteFile(path, []byte(newContent), 0644)
	if err != nil {
		fmt.Printf("Failed to write file: %s, Error: %v\n", path, err)
		return
	}

	fmt.Println("Updated:", path)
}
