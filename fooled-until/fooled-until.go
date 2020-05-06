package main

import (
	"fmt"

	"github.com/oar-team/batsky-go/time"

	"github.com/oar-team/batsky-apimachinery/pkg/util/wait"
)

func main() {
	stop := make(chan struct{})
	wait.Until(func() { fmt.Println("stuck in a loop") }, time.Second, stop)
}
