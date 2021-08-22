package main

import (
	"flag"

	"github.com/sainu/profile-api/config"
	"github.com/sainu/profile-api/server"
)

func main() {
	env := flag.String("e", "development", "")
	flag.Parse()

	config.Init(*env)

	if err := server.Init(); err != nil {
		panic(err)
	}
}
