package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func run(name string) {

	defer wg.Done()

	for i := 0; i < 3; i++ {
		time.Sleep(1 * time.Second) // attendre 1 seconde
		fmt.Println(name, " : ", i)
	}
}

func main() {
	debut := time.Now()
	wg.Add(1)
	go run("A")
	wg.Add(1)
	go run("B")
	wg.Add(1)
	go run("C")

	wg.Wait()

	fin := time.Now()
	fmt.Println(fin.Sub(debut))
}
