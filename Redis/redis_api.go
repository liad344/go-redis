package redis

import (
	log "github.com/sirupsen/logrus"
	"github.com/tidwall/redcon"
)


func (i *Instance) Ping(conn redcon.Conn, cmd redcon.Command) {
	conn.WriteString("PONG")
	log.Info("Ponged ip" , conn.RemoteAddr() )
}


func (i *Instance) Del(conn redcon.Conn, cmd redcon.Command) {
	if len(cmd.Args) != 2 {
		conn.WriteError("ERR wrong number of arguments for '" + string(cmd.Args[0]) + "' command")
		return
	}
	i.Lock()
	_, ok := i.data[string(cmd.Args[1])]
	delete(i.data, string(cmd.Args[1]))
	i.Unlock()
	if !ok {
		conn.WriteInt(0)
	} else {
		conn.WriteInt(1)
	}
	log.Info("Deleted")
	return
}

func (i *Instance) Get(conn redcon.Conn, cmd redcon.Command) {
	if len(cmd.Args) < 2 {
		conn.WriteError("ERR wrong number of arguments for '" + string(cmd.Args[0]) + "' command")
		return
	}
	i.Lock()
	val, ok := i.data[string(cmd.Args[1])]
	i.Unlock()
	if !ok {
		conn.WriteNull()
	} else {
		conn.WriteBulk(val)
	}
	log.Info("Got val " , string(val))
	return
}

func (i *Instance) Set(conn redcon.Conn, cmd redcon.Command) {
	if len(cmd.Args) < 3 {
		conn.WriteError("ERR wrong number of arguments for '" + string(cmd.Args[0]) + "' command")
		return
	}
	i.Lock()
	i.data[string(cmd.Args[1])] = cmd.Args[2]
	i.Unlock()
	conn.WriteString("OK")
	log.Info("Set " , string(cmd.Args[2]))
	return
}
