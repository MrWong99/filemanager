package fman

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"time"
)

type File struct {
	Path       string
	LastUpdate time.Time
}

func (file *File) Read() ([]string, error) {
	f, err := os.Open(file.Path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	info, err := f.Stat()
	if err != nil {
		return nil, err
	}
	file.LastUpdate = info.ModTime()

	scanner := bufio.NewScanner(f)
	var result []string

	for scanner.Scan() {
		result = append(result[:], scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return result[:], nil
}

func (file *File) Write(content []string) error {
	if exists(file.Path) {
		fmt.Printf("File '%q' already exists, trying to delete...\n", file.Path)
		err := os.Remove(file.Path)
		if err != nil {
			return err
		}
	}
	f, err := os.Create(file.Path)
	if err != nil {
		return err
	}
	defer f.Close()
	file.LastUpdate = time.Now()
	writer := bufio.NewWriter(f)
	for _, line := range content {
		_, err := writer.WriteString(line + "\n")
		if err != nil {
			return err
		}
	}
	err = writer.Flush()
	if err == nil {
		file.LastUpdate = time.Now()
	}
	return err
}

func exists(path string) bool {
	_, err := os.Stat(path)
	return !errors.Is(err, os.ErrNotExist)
}
