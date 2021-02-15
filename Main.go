package main

import (
	"fmt"
	"time"
)

func path() string { /*Path of the SD card itself*/
	sdPath := "F:" /*Temporary*/
	return sdPath
}

func destiny(folderName string, localPath string) string { /*Path to the new folder in the device*/
	destinyPath := localPath + folderName
	return destinyPath
}

func searchMedia(mediaTypes string, pathFolder string) { /*Search for all media in the SD card*/
	/*Find all media to ingest*/
}

func main() {
	dataTypes := [4]string{".mp4", ".arw", ".jpg"}

	currentTime := time.Now().Format("2006-01-02")
	fmt.Printf("%s\n", currentTime)
	fmt.Println(dataTypes)
}
