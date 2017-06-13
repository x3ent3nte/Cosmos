package main

import (
	"runtime"
	"math/rand"
	"time"
	"log"
	"agent"
	"server"
)

func main() {

	runtime.GOMAXPROCS(8)
	rand.Seed(time.Now().Unix())

	server := server.CreateServer()
	go server.StartServer()

	odin := agent.CreateOdin(200, 40000)

	var last = time.Now().UnixNano()

	for i := 0; i < 200000; i++ {
		start := time.Now().UnixNano()
		time_delta := float64(start - last) / 1000000000.0

		log.Println("time taken: " , time_delta, "  frames/sec: ", 1.0 / time_delta)
		client_data := server.GetClientsData()
		odin.UpdatePlayerData(client_data)
		odin.Simulate(time_delta)
		server.ServeData(odin.GetEntityJSONData())

		last = start

		time.Sleep(time.Millisecond * 1)
	}
}







