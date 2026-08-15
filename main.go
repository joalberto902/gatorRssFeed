package main

import (
	"log"
	"fmt"
	"github.com/joalberto902/gator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatal(err)
	}

	if err = cfg.SetUser("joalberto902"); err != nil {
		log.Fatal(err)
	}

	newCfg, err := config.Read()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", newCfg)	
}
