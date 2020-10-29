package stomp

// see https://github.com/go-stomp/stomp/blob/master/examples/client_test/main.go

import (
	"flag"
	"fmt"
	"testing"

	"github.com/go-stomp/stomp"
)

const defaultPort = ":61613"

var serverAddr = flag.String("server", "localhost:61613", "STOMP server endpoint")
var queueName = flag.String("queue", "/queue/client_test", "Destination queue")
var helpFlag = flag.Bool("help", false, "Print help text")

var options []func(*stomp.Conn) error = []func(*stomp.Conn) error{
	stomp.ConnOpt.Login("guest", "guest"),
	stomp.ConnOpt.Host("/"),
}

func sendMessages(messageCount int) {
	fmt.Println("==>send start")
	// defer func() {
	// 	stop <- true
	// }()

	conn, err := stomp.Dial("tcp", *serverAddr, options...)
	if err != nil {
		println("cannot connect to server", err.Error())
		return
	}

	fmt.Println("connection OK")

	for i := 1; i <= messageCount; i++ {
		fmt.Println("MessageCount =", i)
		text := fmt.Sprintf("Message #%d", i)
		err = conn.Send(*queueName, "text/plain",
			[]byte(text), nil)
		if err != nil {
			println("failed to send to server", err)
			return
		}
	}

	fmt.Println("==> send finished")
}
func recvMessages(messageCount int) {
	// defer func() {
	// 	stop <- true
	// }()

	conn, err := stomp.Dial("tcp", *serverAddr, options...)

	if err != nil {
		println("cannot connect to server", err.Error())
		return
	}

	sub, err := conn.Subscribe(*queueName, stomp.AckAuto)
	if err != nil {
		println("cannot subscribe to", *queueName, err.Error())
		return
	}
	//close(subscribed)

	for i := 1; i <= messageCount; i++ {
		fmt.Println("messageCount=", i)
		msg := <-sub.C
		expectedText := fmt.Sprintf("Message #%d", i)
		actualText := string(msg.Body)
		if expectedText != actualText {
			fmt.Println("Expected:", expectedText)
			fmt.Println("Actual:", actualText)
		}
	}
	fmt.Println("receiver finished")

}

func Test_CallStomp(t *testing.T) {

	sendMessages(4)

	//recvMessages(10)

	fmt.Println("receiver finished")
}
