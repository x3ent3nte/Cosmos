package control

import (
	"agent"
	"vec"
	"physics"
	"sync"
)

type Odin struct {
	ents []agent.Entity
}

func (odin *Odin) Simulate(time_delta float64) {
	for i := 0; i < len(odin.ents); i++ {
		ent := odin.ents[i]
		if ent.Alive() {
			ent.Act(time_delta)
			//ent.Move(time_delta / 1000.0)
		} else {
			odin.ents[i] = odin.ents[len(odin.ents) - 1]
			odin.ents[len(odin.ents) - 1] = nil
			odin.ents = odin.ents[:len(odin.ents) - 1]
			i--
		}
	}
	//physics.DetectCollisions(odin.ents)
	physics.HandleMovement(odin.ents, time_delta)
}

func (odin *Odin) GetEntities() []agent.Entity {
	return odin.ents
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
		go EntityJSONDataWorker(odin.ents[start: end], data[start:], &wg)
		
		wg.Wait()
	}
	return data
}

func EntityJSONDataWorker(ents []agent.Entity, data []string, wg *sync.WaitGroup) {
	for i, ent := range ents {
		data[i] = ent.GetJSON()
	}
	wg.Done()
}

func CreateOdin(initial_pop int, scope float64) Odin {
	ents := make([]agent.Entity, initial_pop)
	for i := 0; i < len(ents); i++ {
		pos := vec.Vec3Random(scope)
		var ent agent.Entity = agent.SpawnAnimal(int64(i), pos)
		ents[i] = ent
	}
	return Odin{ents}
}




