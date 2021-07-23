package main

import (
	"fmt"
	"time"
)

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
