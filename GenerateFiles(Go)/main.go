package main

import (
	"os"
	"fmt"
	"math/rand"
	"math"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func randStr(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyz")
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func generateFile(name string, done chan string) {
	f, err := os.Create(name)
	check(err)

	defer f.Close()

	for i := 0; i < 5 * int(math.Pow(10, 6)); i++ {
		d1 := []byte(randStr(10) + "\r\n")
		_ , err := f.Write(d1)
		check(err)
	}
	f.Sync()

	done <- "Compleate " + name
}

func main() {
	done := make(chan string, 1)
	fileCount := 5

	for i := 0; i < fileCount; i++ {
		go generateFile("data" + strconv.Itoa(i) + ".txt", done)
	}

	for i := 0; i < fileCount; i++ {
		msg := <- done
		fmt.Println(msg)
	}
}
