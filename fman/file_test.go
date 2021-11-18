package fman

import (
	"os"
	"reflect"
	"testing"
	"time"
)

// Test the File.Read() method for the following myFunctionalities:
// - initial LastUpdate time is a not initialized time.Time struct
// - the read testdata has the expected content
// - the LastUpdate timestamp is updated correctly after the Read()
func TestRead(t *testing.T) {
	file := File{
		Path: "../testdata/for_test.txt",
	}

	// First test
	var initialTime time.Time
	if !initialTime.Equal(file.LastUpdate) {
		t.Fatalf(`File.LastUpdate = %q, but expected is %q`, file.LastUpdate, initialTime)
	}

	// Second test
	expected := []string{"Hello, cruel", "World!"}
	msg, err := file.Read()
	if !reflect.DeepEqual(expected, msg) || err != nil {
		t.Fatalf(`File.Read() = %q, %v, but expected were %q, nil`, msg, err, expected)
	}

	// Third test
	expectedTimestamp := time.Date(2021, 11, 17, 10, 00, 56, 171300800, time.Local)
	if !expectedTimestamp.Equal(file.LastUpdate) {
		t.Fatalf(`File.LastUpdate = %q, but expected is %q`, file.LastUpdate, expectedTimestamp)
	}
}

// Test that a file not found error is returned upon wrong input
func TestReadNotExists(t *testing.T) {
	file := File{
		Path: "../testdata/non_existend.file",
	}
	var initialTime time.Time
	if !initialTime.Equal(file.LastUpdate) {
		t.Fatalf(`File.LastUpdate = %q, but expected is %q`, file.LastUpdate, initialTime)
	}
	msg, err := file.Read()
	if err == nil || !os.IsNotExist(err) {
		t.Fatalf(`File.Read() = %q, %v, but expected were [], %v`, msg, err, os.ErrNotExist)
	}
}

// Test the File.Write() method for the following checks:
// - the test file does not exist before running the test
// - the file was created successfully after the Write() call
// - the file has the expected byte size
func TestWrite(t *testing.T) {
	testPath := "../testdata/temp_to_write.txt"
	testContent := []string{"Hello", "Mehrnie :)"}
	expectedTestSize := int64(len("Hello\nMehrnie :)\n"))

	t.Cleanup(deleteFileFunction(testPath)) // Cleanup any created test files after the test finishes

	file := File{
		Path: testPath,
	}

	// First test
	if exists(testPath) {
		t.Fatalf(`Wrong test environment! File "%q" already exists`, testPath)
	}

	file.Write(testContent)
	resFile, err := os.Stat(testPath)

	// Second test
	if err != nil {
		t.Fatalf(`Test file %q was not created by Write(). Error: %v`, testPath, err)
	}

	// Third test
	if resFile.Size() != expectedTestSize {
		t.Fatalf(`Test file %q had size %d, but expected is %d`, testPath, resFile.Size(), expectedTestSize)
	}
}

// This method returns a myFunction that can be called to delete the file under given path if it exists
func deleteFileFunction(path string) func() {
	myFunction := func() {
		if exists(path) {
			os.Remove(path)
		}
	}
	// with this myFunction in a variable you could write this to call it:
	// myFunction()
	return myFunction
}
