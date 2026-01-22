package code

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetSize(t *testing.T) {
	path := "./testdata/plain.txt"
	var expectedSize int64
	fileInfo, err := os.Lstat(path)
	if err != nil {
		log.Fatal("err read test directory")
	}

	expectedSize = fileInfo.Size()
	expectedResult := fmt.Sprintf("%dB %s", expectedSize, path)
	actualResult, _ := GetPathSize(path, false, false, false)

	require.Equal(t, expectedResult, actualResult, "File size should match")
}

func TestGetSizeHumanFormat(t *testing.T) {
	path := "./testdata/plain.txt"
	var expectedSize int64
	fileInfo, err := os.Lstat(path)
	if err != nil {
		log.Fatal("err read test directory")
	}

	expectedSize = fileInfo.Size()
	expectedHumanFormat := FormatSize(expectedSize, true)
	expectedResult := fmt.Sprintf("%s %s", expectedHumanFormat, path)
	actualResult, _ := GetPathSize(path, true, false, false)

	require.Equal(t, expectedResult, actualResult, "File size with human format should match")
}

func TestGetSizeWithHiddenFiles(t *testing.T) {
	path := "./testdata"
	var expectedSize int64
	files, err := os.ReadDir(path)
	if err != nil {
		log.Fatal("err read test directory")
	}

	for _, file := range files {
		if !file.IsDir() {
			info, _ := file.Info()
			expectedSize += info.Size()
		}
	}

	expectedHumanFormat := FormatSize(expectedSize, true)
	expectedResult := fmt.Sprintf("%s %s", expectedHumanFormat, path)
	actualResult, _ := GetPathSize(path, false, true, true)

	require.Equal(t, expectedResult, actualResult, "Size with hide files should match")
}
