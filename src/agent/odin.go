package agent

import (
	"vec"
	"sync"
	"math/rand"
	"concurrent"
)

type Odin struct {
	sync.RWMutex
	ids concurrent.IdHandler
	players map[int64]*Player
	ents []Entity
	ents_spatial SpatialMap
}

func (odin *Odin) UpdatePlayerData(data map[int64]int) {
	for id, code := range data {
		if player, ok := odin.players[id]; ok {
			player.UpdateKeyCode(code)
		} else {
			player_new := SpawnPlayer(odin, id, odin.ids.NextId(), vec.Vec3Random(100))
			odin.players[id] = player_new
			player_new.UpdateKeyCode(code)
			odin.AddEntity(player_new)
		}
	}
}

func (odin *Odin) Simulate(time_delta float64) {
	odin.ents_spatial.SpatialReset()
	for _, ent := range odin.ents {
		odin.ents_spatial.SpatialAdd(ent)
	}

	num_workers := 4
	var wg sync.WaitGroup
	wg.Add(num_workers)

	interval := len(odin.ents) / num_workers
	start := 0
	end := interval

	for i := 0; i < num_workers - 1; i++ {
		go SimulateWorker(odin.ents[start:end], time_delta, &wg)
		start = end
		end += interval
	}
	go SimulateWorker(odin.ents[start:], time_delta, &wg)
	wg.Wait()

	DetectCollisionsParallel(odin.ents)
	//HandleMovement(odin.ents, time_delta)
}

func SimulateWorker(ents []Entity, time_delta float64, wg *sync.WaitGroup) {
	for _, ent := range ents {
		ent.Act(time_delta)
		//ent.Orientate()
		ent.Move(time_delta)
	}
	wg.Done()
}

func (odin *Odin) AddEntity(ent Entity) {
	odin.Lock()
	odin.ents = append(odin.ents, ent)
	odin.Unlock()
}

func (odin *Odin) GetEntityJSONData() []string {
	data := make([]string, len(odin.ents)) 
	if len(odin.ents) < 10 {
		for i, ent := range odin.ents {
			data[i] = ent.GetJSON()
		}
	} else {
		num_workers := 4
		var wg sync.WaitGroup
		wg.Add(num_workers)
		
		interval := len(odin.ents) / num_workers
		start := 0
		end := interval
		for i := 0; i < num_workers - 1; i++ {
			go EntityJSONDataWorker(odin.ents[start: end], data[start: end], &wg)
			start = end
			end += interval
		}
		go EntityJSONDataWorker(odin.ents[start:], data[start:], &wg)
		
		wg.Wait()
	}
	return data
}

func EntityJSONDataWorker(ents []Entity, data []string, wg *sync.WaitGroup) {
	for i, ent := range ents {
		data[i] = ent.GetJSON()
	}
	wg.Done()
}

func CreateOdin(initial_pop int, scope float64) Odin {
	players := make(map[int64]*Player)
	ents := make([]Entity, initial_pop)
	ents_spatial := CreateSpatialMap()
	ids := concurrent.CreateIdHandler()

	odin := Odin{sync.RWMutex{}, ids, players, ents, ents_spatial}

	for i := 0; i < len(ents); i++ {
		pos := vec.Vec3Random(scope)
		gen := rand.Intn(100)
		id := odin.ids.NextId()

		var ent Entity
		if gen < 10 {
			ent = SpawnAnimal(&odin, id, pos)
		} else {
			if gen < 30 {
				ent = SpawnAnimal(&odin, id, pos)
			} else {
				ent = SpawnAnimal(&odin, id, pos)
			}
		}
		ents[i] = ent
	}
	return odin
}










