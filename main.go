package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	var source string
	var files []string
	const localDestination = "C:\\Users\\marka\\Coding\\Proyectos\\Automated Media Ingest\\tests\\"
	var date string = time.Now().Format("2006-01-02")
	var destination string = destiny(date, localDestination)
	var letterSD string

	dataTypes := []string{".MP4", ".ARW", ".JPG"}

	fmt.Println("Enter the path letter for the SD card: ")
	fmt.Scanf("%s", &letterSD)
	letterSD = strings.ToUpper(letterSD) // Always on uppercase

	source = path(letterSD) // Sony file structure

	files = searchMedia(dataTypes, source) // Files to be copied

	for _, file := range files {
		fileOrigin := source + "\\" + file
		//fmt.Println("File: " + fileOrigin)
		bytes, err := copy(fileOrigin, destination, file)
		if err == nil {
			fmt.Printf("Copied '%s' correctly (%d bytes)\n", file, bytes)
		} else {
			fmt.Printf("Failed to copy '%s'", file)
			fmt.Println(err)
		}
	}

}
