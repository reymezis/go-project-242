package code

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func GetFolderSize(path string, all, recursive bool) (int64, error) {
	var size int64
	fileInfo, err := os.Lstat(path)
	if err != nil {
		return 0, err
	}

	if fileInfo.IsDir() {
		entries, err := os.ReadDir(path)
		if err != nil {
			log.Fatal("err read directory")
			return 0, err
		}

		for _, entry := range entries {
			if !entry.IsDir() {
				info, _ := entry.Info()
				if strings.HasPrefix(info.Name(), ".") && !all {
					continue
				}
				size += info.Size()
			} else if recursive {
				innerSize, err := GetFolderSize(filepath.Join(path, entry.Name()), all, recursive)
				if err != nil {
					return 0, err
				}
				size += innerSize
			}
		}
	} else {
		size = fileInfo.Size()
	}

	return size, nil
}

func GetPathSize(path string, recursive, human, all bool) (string, error) {
	size, err := GetFolderSize(path, all, recursive)
	if err != nil {
		return "", err
	}
	formatedSize := FormatSize(size, human)
	return formatedSize, nil
}

func FormatSize(size int64, human bool) string {
	if !human {
		return fmt.Sprintf("%dB", size)
	}

	const unit = 1024
	if size < unit {
		return fmt.Sprintf("%dB", size)
	}

	units := []string{"B", "KB", "MB", "GB", "TB", "PB", "EB"}
	val := float64(size)
	i := 0

	for val >= 1024 && i < len(units)-1 {
		val /= 1024
		i++
	}

	return fmt.Sprintf("%.1f%s", val, units[i])
}
