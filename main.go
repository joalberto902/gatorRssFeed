package main

import (
	"os"
	"fmt"
	"github.com/joalberto902/gator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	s := state{Config: &cfg}

	cmds := commands{CommandMap: make(map[string]func(*state, command) error)}
	cmds.register("login", handlerLogin)

	if len(os.Args) < 2 {
		fmt.Println("Not enough arguments were provided")
		os.Exit(1)
	}
	
	var cmd command
	cmd.Name = os.Args[1]
	if len(os.Args) == 2 {
		cmd.Args = []string{}
	} else {
		cmd.Args = os.Args[2:]
	}


	if err = cmds.run(&s, cmd); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
