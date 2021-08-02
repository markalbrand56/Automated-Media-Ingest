package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	/*
		source := path()
		dataTypes := []string{".MP4", ".ARW", ".JPG"}
		var files = searchMedia(dataTypes, source)

		currentTime := time.Now().Format("2006-01-02")

		source = source + "\\M4K04804.JPG" // Example

		localDest := "C:\\Users\\marka\\Coding\\Proyectos\\Automated Media Ingest\\DD\\"
		destination := destiny(currentTime, localDest)

		fmt.Println(source)
		fmt.Println(destination)
		//fmt.Print(copy(source, destination))

		fmt.Println(searchMedia(dataTypes, "C:\\Users\\marka\\Pictures\\Sony Alpha\\29-07-2021"))
	*/
	var source string
	var files []string
	var localDestination = "C:\\Users\\marka\\Coding\\Proyectos\\Automated Media Ingest\\DD\\"
	var date string = time.Now().Format("2006-01-02")
	var destination string = destiny(date, localDestination)
	var letterSD string

	dataTypes := []string{".MP4", ".ARW", ".JPG"}

	fmt.Println("Ingrese la letra de la tarjeta SD: ") // Temporal
	fmt.Scanf("%s", &letterSD)

	source = path(letterSD)

	files = searchMedia(dataTypes, source)

	fmt.Println(destination)

	os.Create(destination)

	for _, file := range files {
		fileOrigin := source + "\\" + file
		fmt.Println("File: " + fileOrigin)
		bytes, err := copy(fileOrigin, destination)
		fmt.Println(bytes)
		fmt.Println(err)
	}

}
