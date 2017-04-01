package main

import (
	"os"
	"github.com/urfave/cli"
	"log"
	"path/filepath"
	"fit/core"
	"encoding/json"
)

func main() {
	app := cli.NewApp()
	app.Name = "fit"
	app.Usage = "file tools"
	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "nop",
			Usage: "Simulation mode, no changes will be done",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:    "replace",
			Aliases: []string{"r"},
			Usage:   "replace a string in files",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "path",
					Usage: "path",
				},
				cli.StringFlag{
					Name:  "name",
					Usage: "Name pattern of files",
				},
				cli.StringFlag{
					Name:  "expr",
					Usage: "Expression to search",
				},
				cli.StringFlag{
					Name:  "repl",
					Usage: "Replace with",
				},
				cli.StringFlag{
					Name:  "nop",
					Usage: "Simulation mode, no changes will be done",
				},
			},
			Action: func(c *cli.Context) (err error) {
				replacer := &core.Replacer{
					Name:       c.String("name"),
					Expression: c.String("expr"), Replacement: c.String("repl"),
					Nop:        c.GlobalBool("nop")}
				json,_ := json.Marshal(replacer)
				println(string(json))
				err = filepath.Walk(c.String("path"), replacer.Replace)
				if err != nil {
					panic(err)
				}
				return nil
			},

		},
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Println(err)
	}
}
