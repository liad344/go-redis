package Redis

import "github.com/tidwall/redcon"

func Del(conn redcon.Conn, cmd redcon.Command) bool {
	if len(cmd.Args) != 2 {
		conn.WriteError("ERR wrong number of arguments for '" + string(cmd.Args[0]) + "' command")
		return true
	}
	mu.Lock()
	_, ok := items[string(cmd.Args[1])]
	delete(items, string(cmd.Args[1]))
	mu.Unlock()
	if !ok {
		conn.WriteInt(0)
	} else {
		conn.WriteInt(1)
	}
	return false
}

func Get(conn redcon.Conn, cmd redcon.Command) bool {
	if len(cmd.Args) != 2 {
		conn.WriteError("ERR wrong number of arguments for '" + string(cmd.Args[0]) + "' command")
		return true
	}
	mu.RLock()
	val, ok := items[string(cmd.Args[1])]
	mu.RUnlock()
	if !ok {
		conn.WriteNull()
	} else {
		conn.WriteBulk(val)
	}
	return false
}

func Set(conn redcon.Conn, cmd redcon.Command) bool {
	if len(cmd.Args) != 3 {
		conn.WriteError("ERR wrong number of arguments for '" + string(cmd.Args[0]) + "' command")
		return true
	}
	mu.Lock()
	items[string(cmd.Args[1])] = cmd.Args[2]
	mu.Unlock()
	conn.WriteString("OK")
	return false
}
