package agent

import (
	"vec"
	"mysort"
	"sync"
)

func DetectCollisions(ents []Entity) {
	for i := 0; i < len(ents); i++ {
		one := ents[i]
		one_pos := one.GetPos()
		for j := i + 1; j < len(ents); j++ {
			two := ents[j]
			two_pos := two.GetPos()
			if two_pos.X - one_pos.X > one.GetRadius() + two.GetRadius() {
				break
			} else {
				if vec.Vec3DistanceBetween(one_pos, two_pos) < one.GetRadius() + two.GetRadius() {
					bodyCollide(one, two)
				}
			}
		}
	}
}

func DetectCollisionsParallel(ents_int interface{}) {
	ents := ents_int.([]Entity)
	mysort.Insertion(SortEntityXPos(ents))
	ch := make(chan []EntityPair)

	num_workers := 4
	interval := len(ents) / num_workers

	start := 0
	end := interval
	for i := 0; i < num_workers - 1; i++ {
		go detectCollisionsWorker(ents[start:], end - start, ch)
		start = end
		end += interval
	}
	go detectCollisionsWorker(ents[start:], len(ents) - start, ch)

	performCollisions := func(pairs []EntityPair, wg *sync.WaitGroup) {
		for _, pair := range pairs {
			bodyCollide(pair.one, pair.two)
		}
		wg.Done()
	}

	var wg sync.WaitGroup
	wg.Add(num_workers)

	col1 := <- ch
	go performCollisions(col1, &wg)
	col2 := <- ch
	go performCollisions(col2, &wg)
	col3 := <- ch
	go performCollisions(col3, &wg)
	col4 := <- ch
	go performCollisions(col4, &wg)
	wg.Wait()
}

func detectCollisionsWorker(ents []Entity, end int, ch chan []EntityPair) {
	pairs := make([]EntityPair, 0, 10)
	for i := 0; i < end; i++ {
		one := ents[i]
		one_pos := one.GetPos()
		for j := i + 1; j < len(ents); j++ {
			two := ents[j]
			two_pos := two.GetPos()
			if two_pos.X - one_pos.X > one.GetRadius() + two.GetRadius() {
				break
			} else {
				if vec.Vec3DistanceBetween(one_pos, two_pos) < one.GetRadius() + two.GetRadius() {
					pairs = append(pairs, EntityPair{one, two})
				}
			}
		}
	}
	ch <- pairs
}














