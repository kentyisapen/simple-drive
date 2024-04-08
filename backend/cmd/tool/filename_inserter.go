// cmd/tool/filename_inserter.go
package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	err := filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && strings.HasSuffix(path, ".go") && !strings.HasSuffix(path, ".pb.go") {
			wg.Add(1)
			go func(path string) {
				defer wg.Done()
				if err := insertOrUpdateComment(path); err != nil {
					fmt.Printf("Error processing file %s: %v\n", path, err)
				}
			}(path)
		}
		return nil
	})

	if err != nil {
		fmt.Printf("Error walking through the directory: %v\n", err)
		return
	}

	wg.Wait()
}

func insertOrUpdateComment(path string) error {
	file, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("error reading file %s: %w", path, err)
	}

	lines := strings.Split(string(file), "\n")
	comment := fmt.Sprintf("// %s", path)

	// Check if the first line is a comment, and replace or prepend accordingly.
	if len(lines) > 0 && strings.HasPrefix(strings.TrimSpace(lines[0]), "//") {
		lines[0] = comment
	} else {
		lines = append([]string{comment}, lines...)
	}

	output, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("error creating file %s: %w", path, err)
	}
	defer output.Close()

	_, err = output.WriteString(strings.Join(lines, "\n"))
	if err != nil {
		return fmt.Errorf("error writing to file %s: %w", path, err)
	}

	return nil
}
