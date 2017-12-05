package main

import (
	"os"
	"github.com/urfave/cli"
	"log"
	"path/filepath"
	"encoding/json"
	"github.com/eugeis/fit"
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
			},
			Action: func(c *cli.Context) (err error) {
				replacer := &fit.Replacer{
					Name:       c.String("name"),
					Expression: c.String("expr"), Replacement: c.String("repl"),
					Nop:        c.GlobalBool("nop")}
				json, _ := json.Marshal(replacer)
				println(string(json))
				err = filepath.Walk(c.String("path"), replacer.Replace)
				if err != nil {
					panic(err)
				}
				return nil
			},
		}, {
			Name:    "ansiToUtf8",
			Aliases: []string{"ansi"},
			Usage:   "encode ansi to ansi",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "path",
					Usage: "path of the source file",
				},
				cli.StringFlag{
					Name:  "target",
					Usage: "path of the target file",
				},
			},
			Action: func(c *cli.Context) (err error) {
				return fit.AnsiToUtf8(
					c.String("path"), c.String("target"), c.GlobalBool("nop"))
			},
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Println(err)
	}
}
