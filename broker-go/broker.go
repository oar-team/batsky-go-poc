package main

import (
	"fmt"
	"time"

	zmq "github.com/pebbe/zmq4"
)

func main() {
	responder, _ := zmq.NewSocket(zmq.REP)
	defer responder.Close()
	responder.Bind("tcp://127.0.0.1:27000")

	for {
		//  Wait for next request from client
		msg, _ := responder.RecvMessage(0)
		if len(msg) > 0 {
			//fake_time := fmt.Sprintf("%f", float64(time.Now().UnixNano())/1e9-1587117000)
			//fake_time := fmt.Sprintf("%f",float64(time.Now().UnixNano())/1e9)
			fake_time := fmt.Sprintf("%f", float64(0))
			fmt.Println("fake time:", fake_time)
			msg[0] = fake_time //"1587117099.5642722"
			responder.SendMessage(msg)
			fmt.Println("Sent ", msg)
			time.Sleep(500 * time.Millisecond)
		}
	}
}
