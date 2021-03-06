package vtf

import (
	"os"
	"testing"
)

func TestReadFromFile(t *testing.T) {
	_, err := ReadFromFile("samples/read/test.vtf")
	if err != nil {
		t.Error(err)
	}
}

func TestReadFromStream(t *testing.T) {
	f, err := os.Open("samples/read/test.vtf")
	if err != nil {
		t.Error(err)
	}
	defer f.Close()

	_, err = ReadFromStream(f)
	if err != nil {
		t.Error(err)
	}
}
