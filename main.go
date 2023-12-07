package main

import (
	"fmt"
	"go_script/scripts"
	"os"
	"time"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "okkgit's scripts",
		Usage: "okkgit's scripts, it's easy to work",
	}
	app.Commands = []*cli.Command{
		{
			Name:  "tx-dns-update",
			Usage: "tx clould dns update ip to \"okkgit.top\"",
			Action: func(ctx *cli.Context) error {
				fmt.Println(">> run script tx-dns-update at: ", time.Now())
				scripts.UpdateIPToDNS()
				return nil
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		panic(err)
	}
}
