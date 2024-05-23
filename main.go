package main

import (
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"syscall"

	"github.com/alexflint/go-arg"
)

var args struct {
	silent bool `arg:"-q,--quiet" help:"don't print anything"`
	server bool `arg:"-S,--server" help:"run a websocket server that can be connected to listen to changes"`
	port   int  `arg:"-p,--port" help:"port to start the server on"`
}

func main() {
	cmd, err := splitArgs(os.Args)

	if err != nil {
		fmt.Println(err)
		fmt.Println("Error: Unable to parse arguments")
		os.Exit(1)
	}

	runCommand(cmd, c)

	arg.MustParse(&args)

}

func runCommand(command []string, c chan bool) {
	cmd := exec.Command(command[0], command[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin


}

func splitArgs(inputArgs []string) (command []string, err error) {
	command = []string{}
	lookingForCommand := false

	for _, arg := range inputArgs {
		if lookingForCommand {
			command = append(command, arg)
		} else {
			if arg == "--" {
				lookingForCommand = true
			}
		}
	}

	if len(command) == 0 {
		err = fmt.Errorf("No command found")
		return command, err
	}

	return command, nil
}
