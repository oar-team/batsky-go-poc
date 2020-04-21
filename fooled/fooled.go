package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	t := fmt.Sprintf("%f",float64(now.UnixNano())/1e9)
	//time.Sleep(10 * time.Second)
	fmt.Println("Hello, world.")
	fmt.Println("Now: ", now)
	fmt.Println("now from epoch: ", t)
}
