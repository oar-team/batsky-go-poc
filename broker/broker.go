//
//  Hello World server.
//  Binds REP socket to tcp://*:5555
//  Expects "Hello" from client, replies with "World"
//

package main

import (
	zmq "github.com/pebbe/zmq4"
	"fmt"
	"time"
)

func main() {
	responder, _ := zmq.NewSocket(zmq.STREAM) // zmq.STREAM need to exchange with raw TCP connection
	defer responder.Close()
	responder.Bind("tcp://*:27000")

	for {
		//  Wait for next request from client
		msg, _ := responder.RecvMessage(0)
		client_id := msg[0] // client_id ideally should be converted to int
		l := len(msg)
		if (l == 2) {
			if (msg[1] != "") {
				fmt.Println("Client_id, real_time", msg)
				fake_time := fmt.Sprintf("%f",float64(time.Now().UnixNano())/1e9 - 1587117000)
				fmt.Println("fake time:", fake_time)
				msg[1] = fake_time //"1587117099.5642722"
				responder.SendMessage(msg)
				fmt.Println("Sent ", msg)
				time.Sleep(500 * time.Millisecond)
			} else {
				fmt.Println("Client id (de)connet: ", client_id)
			}
		}
	}
}
