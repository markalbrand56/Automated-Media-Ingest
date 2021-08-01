package main

import (
	"fmt"
	"time"
)

func main() {
	dataTypes := []string{".MP4", ".ARW", ".JPG"}

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
	fmt.Println(source)
	fmt.Println(destination)
	//fmt.Print(copy(source, destination))
	fmt.Println(searchMedia(dataTypes, "C:\\Users\\marka\\Pictures\\Sony Alpha\\29-07-2021"))

}
