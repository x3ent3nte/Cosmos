package main

import (
	"math/rand"
	"time"
	"fmt"
	"control"
	"server"
)

func main() {
	rand.Seed(time.Now().Unix())

	server := server.CreateServer()
	go server.StartServer()

	odin := control.CreateOdin(750, 90000)

	var last = time.Now().UnixNano()

	for i := 0; i < 200000; i++ {
		start := time.Now().UnixNano()
		time_delta := float64(start - last) / 1000000.0

		fmt.Println("time taken: " , time_delta, "  frames/sec: ", 1000 / time_delta)
		odin.Simulate(time_delta)
		server.Serve(odin.GetEntityJSONData())

		last = start

		time.Sleep(time.Millisecond * 1)
	}
}