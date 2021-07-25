package main

import (
	"fmt"
	"io"
	"os"
)

func path() string { /*Path of the SD card itself*/
	var letterSD string
	fmt.Println("Ingrese la letra de la tarjeta SD: ") // Temporal
	fmt.Scanf("%s", &letterSD)

	SonyFiles := "\\DCIM\\100MSDCF"
	sdPath := letterSD + ":" + SonyFiles

	return sdPath
}

func destiny(folderName string, localPath string) string { /*Path to the new folder in the device*/
	destinyPath := localPath + folderName
	return destinyPath
}

func searchMedia(mediaTypes string, pathFolder string) { /*Search for all media in the SD card*/
	/*Find all media to ingest*/
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
