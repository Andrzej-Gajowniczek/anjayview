package main

import (
	"log"
	"os"
	"path/filepath"
	"sort"
)

func (konsola *termboxRes) dirInit() {
	var err error
	konsola.dirPath, err = os.Getwd()
	checkError(err, "Get current dirr")
	// List all files in the current directory
	files, err := os.ReadDir(konsola.dirPath)
	if err != nil {
		log.Println("Error:", err)
		return
	}
	for _, file := range files {
		if !file.IsDir() && isGraphicsFile(file.Name()) {
			konsola.filelist = append(konsola.filelist, filepath.Join(konsola.dirPath, file.Name()))
		}
	}
	if len(konsola.filelist) == 0 {
		log.Println("There is no gfx files in this directory:", konsola.dirPath)
	} else {
		//log.Printf("len(konsola.filelist)=%d\n", len(konsola.filelist))
	}
	// Sort the list by creation date
	sort.Slice(konsola.filelist, func(i, j int) bool {
		time1, _ := getFileCreationTime(konsola.filelist[i])
		time2, _ := getFileCreationTime(konsola.filelist[j])
		return time1.Before(time2)
	})
}
