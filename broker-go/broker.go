package main

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"time"

	zmq "github.com/pebbe/zmq4"
)

func main() {
	responder, _ := zmq.NewSocket(zmq.REP)
	handshake, _ := zmq.NewSocket(zmq.REQ)
	defer responder.Close()
	defer handshake.Close()
	responder.Bind("tcp://127.0.0.1:27000")
	handshake.Bind("tcp://127.0.0.1:27001")

	now := float64(0)

	for {
		// Tell time we're ready to receive
		_, err := handshake.SendBytes([]byte("ready"), 0)
		handshakeReplyBytes, _ := handshake.RecvBytes(0)
		handshakeReply := string(handshakeReplyBytes)
		if handshakeReply != "ok" {
			panic(fmt.Sprintf("Failed handshake : Expected %s, got %s", "ok", handshakeReply))
		}

		//  Wait for next request from client
		msg, err := responder.RecvBytes(0)
		durations := make([]int64, 0)
		if err != nil {
			panic("Error receiving message:" + err.Error())
		}
		if err = json.Unmarshal(msg, &durations); err != nil {
			panic("Could not unmarshal data:" + err.Error())
		}
		//fake_time := fmt.Sprintf("%f", float64(time.Now().UnixNano())/1e9-1587117000)
		//fake_time := fmt.Sprintf("%f", float64(time.Now().UnixNano())/1e9)
		//fake_time := fmt.Sprintf("%f", float64(1000))

		for _, d := range durations {
			if d > 0 {
				fmt.Printf("call me later %f\n", now+float64(d/1e6)/1e3)
			}
		}

		reply := uint64(now * 1e9)
		b := make([]byte, 8)
		binary.LittleEndian.PutUint64(b, reply)
		_, err = responder.SendBytes(b, 0)
		if err != nil {
			panic("Error sending message: " + err.Error())
		}

		// now is in seconds and its precision is in milliseconds
		fmt.Println("now:", now)
		deltams := int64(10)
		now = float64((int64(now*1000) + deltams)) / 1000
		time.Sleep(time.Millisecond * 10)
	}
}
