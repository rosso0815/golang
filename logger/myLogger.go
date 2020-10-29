package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"

	//log "github.com/Sirupsen/logrus"
	"github.com/Sirupsen/logrus"
)

type Configuration struct {
	Users  []string
	Groups []string
}

func main() {

	var buf bytes.Buffer

	logger := log.New(&buf, "main: ", log.Lshortfile)

	logger.Print("Hello, log file!")
	logger.Print("test 02")
	//logger.Fatal("error")

	fmt.Print(&buf)

	customFormatter := new(logrus.TextFormatter)
	customFormatter.TimestampFormat = "2006-01-02 15:04:05"
	customFormatter.FullTimestamp = true
	logrus.SetFormatter(customFormatter)
	logrus.SetLevel(logrus.DebugLevel)

	//logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.WithFields(logrus.Fields{
		"animal": "walrus",
	}).Info("A walrus appears")
	logrus.Info("test info")
	logrus.Debug("debug info")

	file, _ := os.Open("conf.json")
	decoder := json.NewDecoder(file)
	configuration := Configuration{}
	err := decoder.Decode(&configuration)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println(configuration.Users) // output: [UserA, UserB]
}
