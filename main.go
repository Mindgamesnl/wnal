package main

import (
	"github.com/Mindgamesnl/wnal/process"
	"github.com/Mindgamesnl/wnal/socket"
	"github.com/Mindgamesnl/wnal/utils"
)

func main() {
	port := utils.FindPortOrFail()

	go socket.StartSocket(port)

	process.WrapCommand("jshell", func(a []byte) {
		socket.Broadcast(socket.MakeOutNormal(string(a)), socket.BroadcasterCh)
	}, func(a []byte) {
		socket.Broadcast(socket.MakeOutError(string(a)), socket.BroadcasterCh)
	})
}
