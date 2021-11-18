package fman

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"time"
)

// A simple file datatype
type File struct {
	Path       string
	LastUpdate time.Time
}

// Read the given file and return a slice of the contents split by line
func (file *File) Read() ([]string, error) {
	// Open the file and exit on any error
	f, err := os.Open(file.Path)
	if err != nil {
		return nil, err
	}
	defer f.Close() // Defer means execute this when this function returns at any point

	// Get the file info like last modified time or size
	info, err := f.Stat()
	if err != nil {
		return nil, err
	}
	file.LastUpdate = info.ModTime()

	// Scan the results line by line and append it to the result
	scanner := bufio.NewScanner(f)
	var result []string

	for scanner.Scan() {
		result = append(result, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return result, nil
}

// Write the given content line by line into the file
func (file *File) Write(content []string) error {
	// Delete any file that might exist in the path
	if exists(file.Path) {
		fmt.Printf("File '%q' already exists, trying to delete...\n", file.Path)
		err := os.Remove(file.Path)
		if err != nil {
			return err
		}
	}

	// Create a new file
	f, err := os.Create(file.Path)
	if err != nil {
		return err
	}
	defer f.Close()

	file.LastUpdate = time.Now() // File has just been created. Update last update time accordingly

	// Write to the file
	writer := bufio.NewWriter(f)
	for _, line := range content {
		_, err := writer.WriteString(line + "\n")
		if err != nil {
			return err
		}
	}
	err = writer.Flush()

	// Update last update time when no error occured
	if err == nil {
		file.LastUpdate = time.Now()
	}
	return err
}

// Check if given path exists
func exists(path string) bool {
	_, err := os.Stat(path)
	return !errors.Is(err, os.ErrNotExist)
}
