package concurrent

import (
	"sync"
)

type IdHandler struct {
	sync.RWMutex
	id int64
}

func (ids *IdHandler) NextId() int64 {
	ids.Lock()
	next := ids.id
	ids.id++
	ids.Unlock()
	return next
}

func CreateIdHandler() IdHandler {
	return IdHandler{sync.RWMutex{}, int64(0)}
}