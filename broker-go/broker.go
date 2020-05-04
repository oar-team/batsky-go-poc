package main

import (
	"encoding/binary"
	"encoding/json"
	"fmt"

	zmq "github.com/pebbe/zmq4"
)

func main() {
	requester, _ := zmq.NewSocket(zmq.REQ)
	defer requester.Close()
	requester.Bind("tcp://127.0.0.1:27000")

	now := float64(0)

	for {
		// Tell time we're ready to receive
		_, err := requester.SendBytes([]byte("ready"), 0)

		//  Wait for next request from client
		msg, err := requester.RecvBytes(0)
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
				//fmt.Printf("call me later %f\n", now+float64(d/1e6)/1e3)
			}
		}

		reply := uint64(now * 1e9)
		b := make([]byte, 8)
		binary.LittleEndian.PutUint64(b, reply)
		_, err = requester.SendBytes(b, 0)
		if err != nil {
			panic("Error sending message: " + err.Error())
		}

		//if int64(now*1e3)%1e3 == 0 {
		//	fmt.Println("now:", now)
		//}
		fmt.Println("now", now)
		// now is in seconds and its precision is in milliseconds
		// Minimum resolution is 2ms. 1ms introduces conversion problems.
		deltams := int64(2)
		//fmt.Println("now*1000", now*1000)
		//fmt.Println("int64", int64(now*1000))
		//fmt.Println("+delta", int64(now*1000)+deltams)
		//fmt.Println("float64", float64(int64(now*1000)+deltams))
		//fmt.Println("/1000", float64(int64(now*1000)+deltams)/1000)
		now = float64((int64(now*1000) + deltams)) / 1000

		msg, err = requester.RecvBytes(0)
		if err != nil {
			panic(err)
		}
		done := string(msg)
		if done != "done" {
			panic(fmt.Sprintf("Failed exchange: Expected %s, got %s", "done", done))
		}
	}
}
