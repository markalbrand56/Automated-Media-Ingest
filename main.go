package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

type Configuration struct {
	Destiny string
	Pattern string
}

func main() {
	var sourceImages string                          // Sony's path for images
	var images []string                              // Images to be copied
	var sourceVideos string                          // Sony's path for videos
	var videos []string                              // Videos to be copied
	var letterSD string                              // Path letter ('D:', 'F:', ...)
	var dataTypes = []string{".MP4", ".ARW", ".JPG"} // File formats to copy
	var config Configuration                         // Configuration for the program

	// Checking fot the config file
	found := false
	for found != true {
		_, err := os.Stat(".\\config.json")

		if err != nil {
			fmt.Println("File not found")

			newConfig := Configuration{}
			var newDestiny string
			var newPattern string
			consoleReader := bufio.NewReader(os.Stdin) // Complex strings, allows spaces in the address

			fmt.Println("Enter the destination")
			newDestiny, err = consoleReader.ReadString('\n')
			newDestiny = strings.TrimSuffix(newDestiny, "\n")
			if err != nil {
				fmt.Println(err.Error())
				return
			}

			fmt.Println("Enter the pattern")
			_, err := fmt.Scanln(&newPattern)
			if err != nil {
				fmt.Println(err.Error())
				return
			}

			newConfig.Destiny = newDestiny
			newConfig.Pattern = newPattern

			// Creation of the config file
			os.Create(".\\config.json")
			file, _ := json.MarshalIndent(newConfig, "", " ")
			err = ioutil.WriteFile(".\\config.json", file, 0644)

			if err == nil {
				found = true
			}

		} else {
			fmt.Println("File found")
			found = true
		}
	}

	// Configuration load
	content, err := ioutil.ReadFile("./config.json")
	if err != nil {
		fmt.Println("Error when opening file: ", err)
	}
	err = json.Unmarshal(content, &config)
	if err != nil {
		fmt.Println("Error during Unmarshal(): ", err)
	}

	fmt.Println("\nEnter the path letter for the SD card: ")
	fmt.Scanln(&letterSD)
	letterSD = strings.ToUpper(letterSD) // Always on uppercase

	// Media search
	var sourceErrors int = 0 // if the value is 2, both sources were not found
	sourceImages = pathImages(letterSD)
	images, err = searchMedia(dataTypes, sourceImages)
	if err != nil {
		fmt.Println("\nError locating the source folder for the images")
		sourceErrors++
	}

	sourceVideos = pathVideos(letterSD)
	videos, err = searchMedia(dataTypes, sourceVideos)
	if err != nil {
		fmt.Println("\nError locating the source folder for the videos")
		sourceErrors++
	}

	// Copying images
	filesCopied := 0
	start := time.Now()
	transferErrors := 0
	for _, file := range images {
		fileOrigin := sourceImages + "\\" + file // Complete pathImages to the original file

		sourceFileStat, err := os.Stat(fileOrigin) // Information about the file

		if err != nil { // If the file doesn't exist
			transferErrors++
			continue
		} else if !sourceFileStat.Mode().IsRegular() { // If it's not a regular file
			transferErrors++
			continue
		}

		date := sourceFileStat.ModTime().Format(config.Pattern) // Modification date.
		newDestiny := destiny(date, config.Destiny)             // Complete path to new destination

		bytes, err := copy(fileOrigin, newDestiny, file) // File transfer

		if err == nil { // No errors in the file transfer
			if bytes > 0 { // If it actually copied the file
				fmt.Printf("Copied '%s' correctly (%d bytes)\n", file, bytes)
				filesCopied++
			}
		} else {
			fmt.Printf("Failed to copy '%s'\n", file)
			transferErrors++
		}
	}

	// Copying videos
	for _, file := range videos {
		fileOrigin := sourceVideos + "\\" + file // Complete path to the original file

		sourceFileStat, err := os.Stat(fileOrigin) // Information about the file

		if err != nil { // If the file doesn't exist
			transferErrors++
			continue
		} else if !sourceFileStat.Mode().IsRegular() { // If it's not a regular file
			transferErrors++
			continue
		}

		date := sourceFileStat.ModTime().Format("2006-01-02") // Modification date.
		newDestiny := destiny(date, config.Destiny)           // Complete path to new destination

		bytes, err := copy(fileOrigin, newDestiny, file) // File transfer

		if err == nil { // No errors in the file transfer
			if bytes > 0 { // If it actually copied the file
				fmt.Printf("Copied '%s' correctly (%d bytes)\n", file, bytes)
				filesCopied++
			}
		} else {
			fmt.Printf("Failed to copy '%s'\n", file)
			transferErrors++
		}
	}

	end := time.Now()

	// Results
	if sourceErrors >= 2 { // No source folder found
		fmt.Println("\nThe program wasn't able to locate the desired media")
	} else if filesCopied == 0 && transferErrors == 0 { // All files already in destination
		fmt.Println("\nThere were no new files to be copied")
	} else if filesCopied > 0 && sourceErrors == 0 { // New files copied to destination, every source folder found
		fmt.Printf("\nCopied correctly %d files in %.2f seconds \n%d error(s) while copying", filesCopied, end.Sub(start).Seconds(), transferErrors)
	} else if filesCopied > 0 { // New files copied to destination, at least one source folder not found
		fmt.Printf("\nCopied correctly %d files in %.2f seconds \n%d error(s) while copying \n%d media source folder not found", filesCopied, end.Sub(start).Seconds(), transferErrors, sourceErrors)

	}
	fmt.Println("\n\nPress enter to exit")
	fmt.Scanln()
}
