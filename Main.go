package main

import (
	"fmt"
	"time"
)

func main() {
	currentTime := time.Now().Format("2006-01-02")
	fmt.Printf("%s\n", currentTime)
}
