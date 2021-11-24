package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	var source string
	var files []string
	var localDestination = "C:\\Users\\marka\\Coding\\Proyectos\\Automated Media Ingest\\DD\\"
	var date string = time.Now().Format("2006-01-02")
	var destination string = destiny(date, localDestination)
	var letterSD string

	dataTypes := []string{".MP4", ".ARW", ".JPG"}

	fmt.Println("Enter the path letter for the SD card: ") // Temporal
	fmt.Scanf("%s", &letterSD)

	source = path(letterSD)

	files = searchMedia(dataTypes, source)

	fmt.Println(destination)

	os.Create(destination)

	for _, file := range files {
		fileOrigin := source + "\\" + file
		//fmt.Println("File: " + fileOrigin)
		bytes, err := copy(fileOrigin, destination, file)
		if err == nil {
			fmt.Printf("Copied '%s' correctly (%d)\n", file, bytes)
		} else {
			fmt.Printf("Failed to copy '%s'", file)
			fmt.Println(err)
		}
	}

}
