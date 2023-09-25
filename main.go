package main

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"
)

func main() {
	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// List all files in the current directory
	fileList := []string{}
	files, err := os.ReadDir(currentDir)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	for _, file := range files {
		if !file.IsDir() && isGraphicsFile(file.Name()) {
			fileList = append(fileList, filepath.Join(currentDir, file.Name()))
		}
	}

	// Sort the list by creation date
	sort.Slice(fileList, func(i, j int) bool {
		time1, _ := getFileCreationTime(fileList[i])
		time2, _ := getFileCreationTime(fileList[j])
		return time1.Before(time2)
	})

	for _, file := range fileList {
		fmt.Println(file)
	}
}

// Function to check if a file has a graphic file extension
func isGraphicsFile(filename string) bool {
	extensions := []string{".jpg", ".jpeg", ".png", ".gif", ".bmp", ".svg"}
	ext := strings.ToLower(filepath.Ext(filename))
	for _, validExt := range extensions {
		if ext == validExt {
			return true
		}
	}
	return false
}

// Function to get the creation time of a file
func getFileCreationTime(filename string) (time.Time, error) {
	fileInfo, err := os.Stat(filename)
	if err != nil {
		return time.Time{}, err
	}
	return fileInfo.ModTime(), nil
}
