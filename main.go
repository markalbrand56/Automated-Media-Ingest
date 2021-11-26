package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	var sourceImages string
	var images []string
	var sourceVideos string
	var videos []string
	const localDestination = "C:\\Users\\marka\\Coding\\Proyectos\\Automated Media Ingest\\tests\\"
	var letterSD string

	dataTypes := []string{".MP4", ".ARW", ".JPG"}

	fmt.Println("Enter the path letter for the SD card: ")
	fmt.Scanf("%s", &letterSD)
	letterSD = strings.ToUpper(letterSD) // Always on uppercase

	sourceImages = pathImages(letterSD)           // Sony file structure
	images = searchMedia(dataTypes, sourceImages) // Files to be copied

	sourceVideos = pathVideos(letterSD)
	videos = searchMedia(dataTypes, sourceVideos)

	for _, file := range images {
		fileOrigin := sourceImages + "\\" + file // Complete pathImages to the original file

		sourceFileStat, err := os.Stat(fileOrigin) // Information about the file

		if err != nil { // If the file exists
			continue
		} else if !sourceFileStat.Mode().IsRegular() { // If it's a regular file
			continue
		}

		date := sourceFileStat.ModTime().Format("2006-01-02") // Modification date.
		newDestiny := destiny(date, localDestination)         // Complete pathImages to new destination

		bytes, err := copy(fileOrigin, newDestiny, file)

		if err == nil {
			fmt.Printf("Copied '%s' correctly (%d bytes)\n", file, bytes)
		} else {
			fmt.Printf("Failed to copy '%s'\n", file)
			fmt.Println(err)
		}
	}

	for _, file := range videos {
		fileOrigin := sourceVideos + "\\" + file // Complete pathImages to the original file

		sourceFileStat, err := os.Stat(fileOrigin) // Information about the file

		if err != nil { // If the file exists
			continue
		} else if !sourceFileStat.Mode().IsRegular() { // If it's a regular file
			continue
		}

		date := sourceFileStat.ModTime().Format("2006-01-02") // Modification date.
		newDestiny := destiny(date, localDestination)         // Complete pathImages to new destination

		bytes, err := copy(fileOrigin, newDestiny, file)

		if err == nil {
			fmt.Printf("Copied '%s' correctly (%d bytes)\n", file, bytes)
		} else {
			fmt.Printf("Failed to copy '%s'\n", file)
			fmt.Println(err)
		}
	}

}
