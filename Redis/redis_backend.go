package Redis

import (
	"github.com/spf13/afero"
	"sync"
)

type RedisInstance struct {
	data map[string][]byte
	sync.Mutex
}

var fs = afero.NewMemMapFs()

func init() {
	r := NewRedisInstance()
	file, _ := fs.Create("/redis/")
	file.Write(r.data["y"])

}

func NewRedisInstance() RedisInstance {
	return RedisInstance{
		data:  make(map[string][]byte),
		Mutex: sync.Mutex{},
	}
}
