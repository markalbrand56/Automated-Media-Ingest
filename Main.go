package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"time"
)

func path() string { /*Path of the SD card itself*/
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Ingrese la letra de la tarjeta SD: ") // Temporal
	scanner.Scan()
	letterSD := scanner.Text()

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

func main() {
	dataTypes := [4]string{".mp4", ".arw", ".jpg"}

	currentTime := time.Now().Format("2006-01-02")

	fmt.Printf("%s\n", currentTime)
	fmt.Println(dataTypes)
	/*
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Println("Ingrese el nombre del archivo a copiar: ") // Temporal
		scanner.Scan()
		file := scanner.Text()*/

	source := path()
	source = source + "\\M4K04804.JPG" // Example

	localDest := "C:\\Users\\marka\\Coding\\Proyectos\\Automated Media Ingest\\DD\\"
	destination := destiny(currentTime, localDest)

	fmt.Print(copy(source, destination))

}
