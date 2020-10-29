package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
	//"github.com/rosso0815/golang/myjobscheduler"
)

func main() {
	log.Println("@@@ start")

	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	app := &cli.App{
		Name:    "go_ImageResizer",
		Version: "1.0.0",
		Commands: []*cli.Command{
			{
				Name:    "convert",
				Aliases: []string{"c"},
				Usage:   "convert a directory of images",
				Action: func(c *cli.Context) error {
					log.Println("@@@ start")
					//myjobscheduler.RunConvert(c.Args().First())
					return nil
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
