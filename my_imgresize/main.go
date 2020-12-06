package main

import (
	"fmt"
	"log"
	"os"
)

func showHelp() {
	// -v convert .

	fmt.Println("show help")

}
func main() {
	log.Println("@@@ start")

	argsWithProg := os.Args
	argsWithoutProg := os.Args[1:]

	log.Println(argsWithProg)
	log.Println(argsWithoutProg)
	log.Println(len(os.Args))

	if len(os.Args) > 1 {
		fmt.Println("go further")
		for i := 0; i < len(os.Args); i++ {
			fmt.Println("@@@", os.Args[i])
		}
	} else {
		showHelp()
	}

	//arg := os.Args[3]

	//log.Println(arg)

	// log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	// app := &cli.App{
	// 	Name:    "go_ImageResizer",
	// 	Version: "1.0.0",
	// 	Commands: []*cli.Command{
	// 		{
	// 			Name:    "convert",
	// 			Aliases: []string{"c"},
	// 			Usage:   "convert a directory of images",
	// 			Action: func(c *cli.Context) error {
	// 				log.Println("@@@ start")
	// 				myjobscheduler.RunConvert(c.Args().First())
	// 				return nil
	// 			},
	// 		},
	// 	},
	// }

	// err := app.Run(os.Args)
	// if err != nil {
	// 	log.Fatal(err)
	// }
}
