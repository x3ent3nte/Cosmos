package agent

import(
	"vec"
	"mysort"
)

type Entity interface {
	Act(float64)
	Move(float64)
	Alive() bool
	GetPos() vec.Vec3
	SetPos(vec.Vec3)
	Velocity() vec.Vec3
	AddVelocity(vec.Vec3)
	Mass() float64
	Radius() float64
	GetJSON() string
}

type SortEntityXPos []Entity

func (elems SortEntityXPos) Len() int {
	return len(elems)
}

func (elems SortEntityXPos) Less(i, j int) bool {
	return elems[i].GetPos().X < elems[j].GetPos().X
}

func (elems SortEntityXPos) Swap(i, j int) {
	temp := elems[i]
	elems[i] = elems[j]
	elems[j] = temp
}

func (elems SortEntityXPos) Reslice(start, end int) mysort.Sortable {
	return SortEntityXPos(elems[start:end])
}





