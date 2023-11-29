package main

import (
	"fmt"
	"time"
)

func run(name string) {
	for i := 0; i < 3; i++ {
		time.Sleep(1 * time.Second) // attendre 1 seconde
		fmt.Println(name, " : ", i)
	}
}

func main() {
	debut := time.Now()
	go run("A")
	go run("B")
	run("C")
	fin := time.Now()
	fmt.Println(fin.Sub(debut))

}
