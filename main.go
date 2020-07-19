package main

import (
	"go-balls/appcontext"
	"go-balls/cli"
	"os"
)

func main() {
	c := cli.NewCli(os.Args)
	app := appcontext.NewApp()
	c.Run(app)
}
