package main

import (
	"context"
	"errors"
	"fmt"
	"time"
	"github.com/joalberto902/gator/internal/config"
	"github.com/joalberto902/gator/internal/database"
	"github.com/google/uuid"
)

//state struct stores the state of the application to be used by the handlers
type state struct {
	Config *config.Config
	Database *database.Queries
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
	
	usr, err := s.Database.GetUser(context.Background(), cmd.Args[0])
	if err != nil {
		return err
	}

	if err := s.Config.SetUser(usr.Name); err != nil {
		return err
	}

	fmt.Printf("New user <%s> has been set", cmd.Args[0])
	return nil
}

//handlerRegister registers a new user at the database and fails if user already exists
func handlerRegister(s *state, cmd command) error {
	if len(cmd.Args) == 0 {
		return errors.New("register command expects a single argument <username>")
	}
	usr, err := s.Database.CreateUser(context.Background(), database.CreateUserParams{
		ID: uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name: cmd.Args[0],
	})	
	if err != nil {
		return err
	}

	if err = s.Config.SetUser(usr.Name); err != nil {
		return err
	}
	fmt.Printf("User <%s> was registered sucessfully\nData: ID = %s\nCreatedAt = %v\nUpdatedAt = %v\n", 
		usr.Name, 
		usr.ID,
		usr.CreatedAt,
		usr.UpdatedAt)
	return nil
}

//handlerReset function reset the data in the database for ease of use
func handlerReset(s *state, cmd command) error {
	err := s.Database.ResetUsers(context.Background())
	if err != nil  {
		return err
	}

	fmt.Println("All the users of the database were cleared")
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


