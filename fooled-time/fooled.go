package main

import (
	"fmt"
	gotime "time"

	"github.com/oar-team/batsky-go/time"
)

func main() {
	now := time.Now()
	t := gotime.Now()
	fmt.Println("Now: ", now)
	fmt.Println("now from epoch: ", t)
	fmt.Println()

	fmt.Println("now", time.Now())
	fmt.Println("sleeping for 5h")
	time.Sleep(5 * time.Hour)
	fmt.Println("now", time.Now())
	fmt.Println()

	fmt.Println("now", time.Now())
	fmt.Println("sleeping for 5 years")
	time.Sleep(5 * 8766 * time.Hour)
	fmt.Println("now", time.Now())
	fmt.Println()
}
