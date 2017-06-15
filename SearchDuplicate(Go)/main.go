package main

import (
	"fmt"
	"strconv"
	"os"
	"log"
	"bufio"
	"time"
)

func processingFile(values chan<- string, sync chan<- string, fileName string) {
	timeStart := time.Now()
	file, err := os.Open(fileName)
	if err != nil {
		sync <- fileName + " finished with error"
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		values <- scanner.Text()
	}

	if err := scanner.Err(); err != nil {
		sync <- fileName + " finished with error"
		log.Fatal(err)
	}

	fmt.Printf(fileName + " processing elapsed: %v\n\r", time.Since(timeStart))
	sync <- fileName + " finished"
}

func main() {
	timeStart := time.Now()
	filesCount := 5
	queue := make(chan string)
	syncChanel := make(chan string)
	for i := 0; i < filesCount; i++ {
		go processingFile(queue, syncChanel, "data" + strconv.Itoa(i) + ".txt")
	}
	lines := make(map[string]int)
	finished := 0
	for finished < filesCount {
		select {
		case key := <- queue:
			lines[key]++
		case syncMsg := <- syncChanel:
			finished++
			fmt.Println(syncMsg)
		}
	}

	repeatedLine := 0
	for k, v := range lines {
		if v > 1 {
			fmt.Println(k, ": ", v)
			repeatedLine += v
		}
	}

	fmt.Println("Repeated lines: ", repeatedLine)
	fmt.Printf("Elapsed programm time: %v\n\r", time.Since(timeStart))
	fmt.Println("End programm")
}