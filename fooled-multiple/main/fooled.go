package main

import (
	"fmt"
	"fooled-multiple/pkg1"
	"fooled-multiple/pkg2"

	time "github.com/oar-team/batsky-go"
)

func main() {
	now := time.Now()
	// t := fmt.Sprintf("%f",float64(now.UnixNano())/1e9)
	//time.Sleep(1 * time.Second)
	fmt.Println("Hello, world.")
	for i := 0; i < 4; i++ {
		fmt.Println("Now: ", now)
	}
	pkg1.ReqTime()
	pkg2.ReqTime()
	//fmt.Println("now from epoch: ", t)
}
