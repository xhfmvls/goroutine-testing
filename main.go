package main

import (
	"fmt"
	"testing"
	"time"
)

func SayHelloWorld() {
	fmt.Println("HELLO WORLD")
}

func PrintNumber(num int) {
	fmt.Println("[+]", num)
}

func TestCreateGoroutine(t *testing.T) {
	go SayHelloWorld()

	fmt.Println("RUN")
	time.Sleep(1 * time.Second)
	fmt.Println("RUN 2")
}

func TestGoroutines(t *testing.T) {
	for i := 0; i < 100000; i++ {
		go PrintNumber(i)
	}
	time.Sleep(5 * time.Second)
}

func main() {
	// TestCreateGoroutine(&testing.T{})
	TestGoroutines(&testing.T{})
}