package main

import (
	"os"

	"github.com/hlatimer266/passmanage/internal/parse"
	"github.com/pterm/pterm"
)

func main() {

	err := parse.CmdArgs(os.Args)
	if err != nil {
		pterm.FgRed.Println(err.Error())
		return
	}

}
