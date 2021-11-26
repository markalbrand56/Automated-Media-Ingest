package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	var sourceImages string                                                                         // Sony's path for images
	var images []string                                                                             // Images to be copied
	var sourceVideos string                                                                         // Sony's path for videos
	var videos []string                                                                             // Videos to be copied
	var letterSD string                                                                             // Path letter ('D:', 'F:', ...)
	var dataTypes = []string{".MP4", ".ARW", ".JPG"}                                                // File formats to copy
	const localDestination = "C:\\Users\\marka\\Coding\\Proyectos\\Automated Media Ingest\\tests\\" // Main destination

	fmt.Println("Enter the path letter for the SD card: ")
	fmt.Scanf("%s", &letterSD)
	letterSD = strings.ToUpper(letterSD) // Always on uppercase

	// Media search
	sourceImages = pathImages(letterSD)
	images = searchMedia(dataTypes, sourceImages)

	sourceVideos = pathVideos(letterSD)
	videos = searchMedia(dataTypes, sourceVideos)

	// Copying images
	for _, file := range images {
		fileOrigin := sourceImages + "\\" + file // Complete pathImages to the original file

		sourceFileStat, err := os.Stat(fileOrigin) // Information about the file

		if err != nil { // If the file exists
			continue
		} else if !sourceFileStat.Mode().IsRegular() { // If it's a regular file
			continue
		}

		date := sourceFileStat.ModTime().Format("2006-01-02") // Modification date.
		newDestiny := destiny(date, localDestination)         // Complete path to new destination

		bytes, err := copy(fileOrigin, newDestiny, file)

		if err == nil {
			fmt.Printf("Copied '%s' correctly (%d bytes)\n", file, bytes)
		} else {
			fmt.Printf("Failed to copy '%s'\n", file)
			fmt.Println(err)
		}
	}

	// Copying videos
	for _, file := range videos {
		fileOrigin := sourceVideos + "\\" + file // Complete path to the original file

		sourceFileStat, err := os.Stat(fileOrigin) // Information about the file

		if err != nil { // If the file exists
			continue
		} else if !sourceFileStat.Mode().IsRegular() { // If it's a regular file
			continue
		}

		date := sourceFileStat.ModTime().Format("2006-01-02") // Modification date.
		newDestiny := destiny(date, localDestination)         // Complete path to new destination

		bytes, err := copy(fileOrigin, newDestiny, file)

		if err == nil {
			fmt.Printf("Copied '%s' correctly (%d bytes)\n", file, bytes)
		} else {
			fmt.Printf("Failed to copy '%s'\n", file)
			fmt.Println(err)
		}
	}

}
