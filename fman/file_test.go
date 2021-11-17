package fman

import (
	"os"
	"reflect"
	"testing"
	"time"
)

func TestRead(t *testing.T) {
	file := File{
		Path: "../testdata/for_test.txt",
	}
	var initialTime time.Time
	if !initialTime.Equal(file.LastUpdate) {
		t.Fatalf(`File.LastUpdate = %q, but expected is %q`, file.LastUpdate, initialTime)
	}
	expected := []string{"Hello, cruel", "World!"}
	msg, err := file.Read()
	if !reflect.DeepEqual(expected, msg) || err != nil {
		t.Fatalf(`File.Read() = %q, %v, but expected were %q, nil`, msg, err, expected)
	}
	expectedTimestamp := time.Date(2021, 11, 17, 10, 00, 56, 171300800, time.Local)
	if !expectedTimestamp.Equal(file.LastUpdate) {
		t.Fatalf(`File.LastUpdate = %q, but expected is %q`, file.LastUpdate, expectedTimestamp)
	}
}

func TestWrite(t *testing.T) {
	testPath := "../testdata/temp_to_write.txt"
	testContent := []string{"Hello", "Mehrnie :)"}
	expectedTestSize := int64(len("Hello\nMehrnie :)\n"))
	t.Cleanup(deleteFileFunction(testPath))
	file := File{
		Path: testPath,
	}
	if exists(testPath) {
		t.Fatalf(`Wrong test environment! File "%q" already exists`, testPath)
	}
	file.Write(testContent)
	resFile, err := os.Stat(testPath)
	if err != nil {
		t.Fatalf(`Test file %q was not created by Write(). Error: %v`, testPath, err)
	}
	if resFile.Size() != expectedTestSize {
		t.Fatalf(`Test file %q had size %d, but expected is %d`, testPath, resFile.Size(), expectedTestSize)
	}
}

func deleteFileFunction(path string) func() {
	fnc := func() {
		if exists(path) {
			os.Remove(path)
		}
	}
	return fnc
}
