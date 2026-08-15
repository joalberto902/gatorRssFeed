package main

import (
	"fmt"
	"errors"
	"github.com/joalberto902/gator/internal/config"
)

//state struct stores the state of the application to be used by the handlers
type state struct {
	Config *config.Config
}

//command holds a name for a command and the arguments for the command
type command struct {
	Name string
	Args []string
}

//commands holds all the commands the CLI can handle
type commands struct {
	CommandMap map[string]func(*state, command) error
}
//handlerLogin is function that handles the input when the command is login
func handlerLogin(s *state, cmd command) error {
	if len(cmd.Args) == 0 {
		return errors.New("login command expects a single argument <username>")
	} 

	if err := s.Config.SetUser(cmd.Args[0]); err != nil {
		return err
	}

	fmt.Printf("New user <%s> has been set", cmd.Args[0])
	return nil
}

func (c *commands) run(s *state, cmd command) error {
	if s == nil {
		return errors.New("state does not exist")
	}
	
	if err := c.CommandMap[cmd.Name](s, cmd); err != nil {
		return err
	}

	return nil
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.CommandMap[name] = f
}


