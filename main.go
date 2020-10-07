package main

import (
	"fmt"
	"github.com/Mindgamesnl/wnal/process"
	"github.com/Mindgamesnl/wnal/queue"
	"github.com/Mindgamesnl/wnal/socket"
	"github.com/Mindgamesnl/wnal/utils"
)

func main() {
	port := utils.FindPortOrFail()

	go socket.StartSocket(port)

	fmt.Println("Use the web session at http://wnal.craftmend.com/ and use host localhost:" + port)

	process.WrapCommand(utils.FindCommand(), utils.FindArgs(), func(a []byte) {
		queue.LogLines.AddImport(queue.Import{Text: string(a)})
		socket.Broadcast(socket.MakeOutNormal(string(a)), socket.BroadcasterCh)
	}, func(a []byte) {
		queue.LogLines.AddImport(queue.Import{Text: string(a)})
		socket.Broadcast(socket.MakeOutError(string(a)), socket.BroadcasterCh)
	})
}
