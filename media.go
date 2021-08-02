package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func path(letter string) string { /*Path of the SD card itself*/

	SonyFiles := "\\DCIM\\100MSDCF"
	sdPath := letter + ":" + SonyFiles

	return sdPath
}

func destiny(folderName string, localPath string) string { /*Path to the new folder in the device*/
	destinyPath := localPath + folderName + "\\"
	return destinyPath
}

func searchMedia(mediaTypes []string, pathFolder string) []string { /*Search for all media in the SD card*/
	/*Find all media to ingest*/
	var filesToCopy []string

	files, err := ioutil.ReadDir(pathFolder)
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		for _, fileType := range mediaTypes {
			if strings.HasSuffix(file.Name(), fileType) {
				filesToCopy = append(filesToCopy, file.Name())
			}
		}
	}
	return filesToCopy
}

func copy(src, dst string) (int64, error) {
	// https://opensource.com/article/18/6/copying-files-go
	// Needs absolute path

	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return 0, err
	}

	if !sourceFileStat.Mode().IsRegular() {
		return 0, fmt.Errorf("%s is not a regular file", src)
	}

	source, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
		return 0, err
	}
	defer destination.Close()
	nBytes, err := io.Copy(destination, source)
	return nBytes, err
}
