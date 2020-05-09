package redis

import "github.com/tidwall/redcon"



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
	return
}

func (i *Instance) Get(conn redcon.Conn, cmd redcon.Command) {
	if len(cmd.Args) != 2 {
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
	return
}

func (i *Instance) Set(conn redcon.Conn, cmd redcon.Command) {
	if len(cmd.Args) != 3 {
		conn.WriteError("ERR wrong number of arguments for '" + string(cmd.Args[0]) + "' command")
		return
	}
	i.Lock()
	i.data[string(cmd.Args[1])] = cmd.Args[2]
	i.Unlock()
	conn.WriteString("OK")
	return
}
