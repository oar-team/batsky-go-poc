package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/oar-team/batsky-go/requester"
	zmq "github.com/pebbe/zmq4"
)

func main() {
	responder, _ := zmq.NewSocket(zmq.REP)
	defer responder.Close()
	responder.Bind("tcp://127.0.0.1:27000")

	for {
		//  Wait for next request from client
		msg, err := responder.RecvBytes(0)
		var messages []requester.Message
		if err != nil {
			panic("Error receiving message:" + err.Error())
		}
		if err = json.Unmarshal(msg, &messages); err != nil {
			panic("Could not unmarshal data:" + err.Error())
		}
		//fake_time := fmt.Sprintf("%f", float64(time.Now().UnixNano())/1e9-1587117000)
		fake_time := fmt.Sprintf("%f", float64(time.Now().UnixNano())/1e9)
		//fake_time := fmt.Sprintf("%f", float64(1000))
		fmt.Println("fake time:", fake_time)

		var reply []requester.Message
		for _, m := range messages {
			reply = append(reply, requester.Message{
				RequestType: m.RequestType,
				Data:        fake_time,
			})
		}

		msg, err = json.Marshal(reply)
		if err != nil {
			panic(fmt.Sprintf("Error marshaling message %v:", reply) + err.Error())
		}
		_, err = responder.SendBytes(msg, 0)
		if err != nil {
			panic("Error sending message: " + err.Error())
		}
		fmt.Println("Sent ", reply)
		time.Sleep(500 * time.Millisecond)
	}
}
