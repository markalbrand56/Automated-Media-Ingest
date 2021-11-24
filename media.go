package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

// Builds the Sony file System using the specified letter of the SD card
func path(letter string) string {

	SonyFiles := "\\DCIM\\100MSDCF"
	sdPath := letter + ":" + SonyFiles

	return sdPath
}

// Builds the path to the new destination
func destiny(folderName string, localPath string) string { /*Path to the new folder in the device*/
	destinyPath := localPath + folderName + "\\"
	return destinyPath
}

// Searches for all media that will be copied from the source directory
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

// Copies the specified file from its source directory to the new directory
func copy(src, dst, file string) (int64, error) {
	/* src: Complete path of original file
	* dst: Path to new folder
	* file: Name of the file to copy
	 */
	// https://opensource.com/article/18/6/copying-files-go
	// https://golangbyexample.com/copy-file-go/

	var newFile string = dst + "\\" + file // The complete path for the new file

	sourceFileStat, err := os.Stat(src) // If the original file exists
	if err != nil {
		return 0, err
	}

	if _, err := os.Stat(newFile); err == nil { // If the file to be copied already exists in dst
		return 0, err // It's not an error. 0 bytes copied
	}

	_, err = os.Stat(dst)

	if os.IsNotExist(err) { // If the new directory exists
		errDir := os.MkdirAll(dst, 0755)
		if errDir != nil {
			log.Fatal(err)
		}

	}

	if !sourceFileStat.Mode().IsRegular() { // If the file is regular
		return 0, fmt.Errorf("%s is not a regular file", src)
	}

	source, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	defer source.Close()

	destination, err := os.Create(newFile)

	if err != nil {
		return 0, err
	}

	defer destination.Close()
	nBytes, err := io.Copy(destination, source)
	return nBytes, err
}
