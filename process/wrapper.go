package process

import (
	"github.com/Mindgamesnl/wnal/wio"
	"os"
	"os/exec"
)

func WrapCommand(command string, onWrite func(a []byte), onError func(a []byte)) {
	cmd := exec.Command(command)

	cmd.Stdout = wio.WrappedWriter{
		OnWrite: onWrite,
		Replaces: wio.RealHandlerSet.Out,
	}

	cmd.Stderr = wio.WrappedWriter{
		OnWrite: onError,
		Replaces: wio.RealHandlerSet.Error,
	}

	cmd.Stdin = os.Stdin

	cmd.Run()
	cmd.Wait()
}
