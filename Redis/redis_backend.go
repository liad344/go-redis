package Redis

import (
	"sync"
)

type Instance struct {
	data map[string][]byte
	sync.Mutex
}

func NewInstance() Instance {
	return Instance{
		data:  make(map[string][]byte),
		Mutex: sync.Mutex{},
	}
}
