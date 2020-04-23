package main

import (
	"fmt"
	"time"
)

func log(format string, a ...interface{}) {
	trueFormat := fmt.Sprintf("%s\n", format)
	fmt.Printf(trueFormat, a...)
}

func spin(stop *bool, c chan string) {
	for !(*stop) {
		for _, ch := range "-\\|/" {
			fmt.Printf("\r \r %c", ch)
			time.Sleep(time.Second)
		}
	}
	log("Stopped")
	c <- "done"
}

func main() {
	log("Started")

	stop := false
	c := make(chan string)
	go spin(&stop, c)
	time.Sleep(10 * time.Second)
	stop = true
	log("Bool stopped")
	<-c
	log("Channel returned")

	log("Done")
}
