package main

import (
	"github.com/liad344/go-redis/Redis"
	"github.com/tidwall/redcon"
	"strings"
)

func handleClose(conn redcon.Conn, err error) {

}

func handleConnection(conn redcon.Conn) bool {

	return true
}

func handleCommands(conn redcon.Conn, cmd redcon.Command) {
	go handle(conn, cmd)
}

func handle(conn redcon.Conn, cmd redcon.Command) {
	switch strings.ToLower(string(cmd.Args[0])) {
	default:
		conn.WriteError("ERR unknown command '" + string(cmd.Args[0]) + "'")
	case "ping":
		conn.WriteString("PONG")
	case "quit":
		conn.WriteString("OK")
		conn.Close()
	case "set":
		if Redis.Set(conn, cmd) {
			return
		}
	case "get":
		if Get(conn, cmd) {
			return
		}
	case "del":
		if Del(conn, cmd) {
			return
		}
	}
}
