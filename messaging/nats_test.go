package stomp

// see https://github.com/go-stomp/stomp/blob/master/examples/client_test/main.go

import (
	"fmt"
	"os"
	"testing"

	nats "github.com/nats-io/nats.go"
)

func Test_CallNats(t *testing.T) {

	fmt.Println("defaulltURL =", nats.DefaultURL)
	// Connect to a server
	nc, err := nats.Connect(nats.DefaultURL)

	if err != nil {
		fmt.Println("[ERROR]", err)
		os.Exit(1)
	}
	// Simple Publisher
	nc.Publish("foo", []byte("Hello World"))
	//fmt.Println("receiver finished")

	fmt.Println("@@@ done")
}

func Test_ReceiveFromNats(t *testing.T) {

	fmt.Println("defaulltURL =", nats.DefaultURL)
	// Connect to a server
	nc, err := nats.Connect(nats.DefaultURL)

	if err != nil {
		fmt.Println("[ERROR]", err)
		os.Exit(1)
	}
	// Simple Publisher
	nc.Publish("foo", []byte("Hello World"))
	//fmt.Println("receiver finished")

	fmt.Println("@@@ done")
}
