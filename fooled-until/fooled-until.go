package main

import (
	"fmt"
	"time"

	"k8s.io/apimachinery/pkg/util/wait"
)

func main() {
	stop := make(chan struct{})
	wait.Until(func() { fmt.Println("stuck in a loop") }, time.Second, stop)
}
