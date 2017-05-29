package physics

import (
	"agent"
	"mysort"
)

type EntityPairTime struct {
	one agent.Entity
	two agent.Entity
	time float64
}

type SortEntityPairTime []EntityPairTime

func (elems SortEntityPairTime) Len() int {
	return len(elems)
}

func (elems SortEntityPairTime) Less(i, j int) bool {
	return elems[i].time < elems[j].time
}

func (elems SortEntityPairTime) Swap(i, j int) {
	temp := elems[i]
	elems[i] = elems[j]
	elems[j] = temp
}

func (elems SortEntityPairTime) Reslice(start, end int) mysort.Sortable {
	return SortEntityPairTime(elems[start:end])
}