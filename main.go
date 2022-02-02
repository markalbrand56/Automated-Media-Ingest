package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
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
	var config Configuration

	// Checking fot the config file
	found := false
	for found != true {
		_, err := os.Stat(".\\config.json")

		if err != nil {
			fmt.Println("File not found")

			newConfig := Configuration{}
			var newDestiny string
			var newPattern string
			consoleReader := bufio.NewReader(os.Stdin)  // Complex strings, allows spaces in the address

			fmt.Println("Enter the destination")
			newDestiny, err = consoleReader.ReadString('\n')
			newDestiny = strings.TrimSuffix(newDestiny, "\n")
			if err != nil{
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
			_ = ioutil.WriteFile(".\\config.json", file, 0644)

			found = true

		} else {
			fmt.Println("File found")
			found = true
		}
	}

	// Configuration load
	content, err := ioutil.ReadFile("./config.json")
	if err != nil {
		fmt.Println("Error when opening file: ", err)
		// TODO Error handling, non-existing file
	}
	err = json.Unmarshal(content, &config)
	if err != nil {
		fmt.Println("Error during Unmarshal(): ", err)
	}

	fmt.Println("\nEnter the path letter for the SD card: ")
	fmt.Scanln(&letterSD)
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

		date := sourceFileStat.ModTime().Format(config.Pattern) // Modification date.
		newDestiny := destiny(date, config.Destiny)             // Complete path to new destination

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
		newDestiny := destiny(date, config.Destiny)           // Complete path to new destination

		bytes, err := copy(fileOrigin, newDestiny, file)

		if err == nil {
			fmt.Printf("Copied '%s' correctly (%d bytes)\n", file, bytes)
		} else {
			fmt.Printf("Failed to copy '%s'\n", file)
			fmt.Println(err)
		}
	}

}
