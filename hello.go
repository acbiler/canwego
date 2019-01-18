package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Printf("hello, world\n")
	rand.Seed(time.Now().UTC().UnixNano())
	fmt.Println("random number: ", rand.Intn(7))
	fmt.Println("Result: ", divide(6,3))
}

func divide (x , y int) int{
	return x / y
}