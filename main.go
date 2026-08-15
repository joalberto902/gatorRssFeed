package main


import _ "github.com/lib/pq"
import (
	"os"
	"fmt"
	"database/sql"
	"github.com/joalberto902/gator/internal/config"
	"github.com/joalberto902/gator/internal/database"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	db, err := sql.Open("postgres", cfg.DbURL)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	dbQueries := database.New(db)
	s := state{Config: &cfg, Database: dbQueries}

	cmds := commands{CommandMap: make(map[string]func(*state, command) error)}
	cmds.register("login", handlerLogin)
	cmds.register("register", handlerRegister)
	cmds.register("reset", handlerReset)
	cmds.register("users", handlerUsers)

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
