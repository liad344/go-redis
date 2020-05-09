package main

import (
	"github.com/tidwall/redcon"
)



func onConnectionClosed (conn redcon.Conn, err error){

}

func onNewConnection (conn redcon.Conn) bool{
	return true
}
