package main

import (
	"fmt"
	"os"

	"github.com/inifomeeo/rssflow/internal/config"
)

type state struct {
	cfg *config.Config
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Println("Error reading config:", err)
		os.Exit(1)
	}

	programState := &state{
		cfg: &cfg,
	}

	cmds := commands{
		handlers: make(map[string]func(*state, command) error),
	}
	cmds.register("login", handlerLogin)

	if len(os.Args) < 2 {
		fmt.Println("Error: not enough arguments")
		fmt.Println("usage: cli <command> [args...]")
		os.Exit(1)
	}

	cmd := command{
		Name: os.Args[1],
		Arguments: os.Args[2:],
	}

	err = cmds.run(programState, cmd)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}
