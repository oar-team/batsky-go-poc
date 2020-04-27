package pkg2

import (
	"fmt"

	time "github.com/oar-team/batsky-go"
)

func ReqTime() {
	now := time.Now()
	// t := fmt.Sprintf("%f",float64(now.UnixNano())/1e9)
	//time.Sleep(1 * time.Second)
	fmt.Println("Hello from package 2")
	for i := 0; i < 4; i++ {
		fmt.Println("Now: ", now)
	}
	//fmt.Println("now from epoch: ", t)
}
