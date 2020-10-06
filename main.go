package main

import (
	"github.com/Mindgamesnl/wnal/process"
	"github.com/Mindgamesnl/wnal/utils"
)

func main()  {
	_ = utils.FindPortOrFail()
	process.WrapCommand("jshell")
}