package process

import (
	"github.com/Mindgamesnl/wnal/utils"
	"github.com/Mindgamesnl/wnal/wio"
	"io"
	"log"
	"os/exec"
)

var Command *exec.Cmd
var CommandWriter io.WriteCloser

func WrapCommand(command string, args []string, onWrite func(a []byte), onError func(a []byte)) {
	Command = exec.Command(command, args...)

	Command.Stdout = wio.WrappedWriter{
		OnWrite: onWrite,
		Replaces: wio.RealHandlerSet.Out,
	}

	Command.Stderr = wio.WrappedWriter{
		OnWrite: onError,
		Replaces: wio.RealHandlerSet.Error,
	}

	epi, er := Command.StdinPipe()
	if er != nil {
		println(er)
	}
	CommandWriter = epi

	Command.Start()
	error := Command.Wait()
	if error != nil {
		log.Println("Could not execute command " + utils.FindCommand())
	}
}
