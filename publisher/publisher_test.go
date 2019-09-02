package publisher

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"testing"
	"time"

	"../node"
	"../rcl"
	"../types"
)

func TestPublisherStringMsg(t *testing.T) {

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	msg := make(chan string, 1)
	go func() {
		// Receive input in a loop.
		for {
			var s string
			fmt.Scan(&s)
			// Send what we read over the channel.
			msg <- s
		}
	}()

	// Initialization.
	var contextObject rcl.RCLContext
	// Get initialized context object.
	contextObject.GetInitializedContext()

	// Initialize contextObject and rcl. This function has to be executed always.
	result := rcl.RCLInit(contextObject.ContextKey)
	if result != 0 {
		log.Fatal(fmt.Errorf("Error initializing the context object!"))
	}

	// Obtain initialized node struct and options.
	myNode := node.GetZeroInitializedNode()
	myNodeOpts := node.GetNodeDefaultOptions()

	fmt.Printf("Creating the node! \n")
	// Initialize node object and associate context object to it.
	result = node.NodeInit(myNode, "GoPublisher", "", myNodeOpts, contextObject.ContextKey)
	if result != 0 {
		log.Fatal(fmt.Errorf("Error initializing the node object!"))
	}

	// Create the publisher.
	myPub := GetZeroInitializedPublisher()
	myPubOpts := GetPublisherDefaultOptions()

	// Create the msg type.
	var myMsg types.StdMsgsString
	myMsg.InitMessage()

	fmt.Printf("Creating the publisher! \n")
	// Initializing the publisher.
	PublisherInit(myPub, myPubOpts, myNode, "/myGoTopic", myMsg.GetMessage())

	index := 0
loop:
	for {
		//Update my msg.
		myMsg.SetText("Greetings from GO! #" + strconv.Itoa(index))
		//Publish the message.
		retRCL := Publish(myPub, myMsg.GetMessage(), myMsg.GetData())

		if retRCL == types.RCL_RET_OK {
			fmt.Printf("(Publisher) Published: %s\n", myMsg.GetDataAsString())
		}
		time.Sleep(500 * time.Millisecond)
		index++

		//Loop breaker.
		select {
		case <-sigs:
			fmt.Println("Got shutdown, exiting")
			// Break out of the outer for statement and end the program.
			break loop
		case <-msg:

		}
	}

	fmt.Printf("Shutting down!! \n")

	myMsg.DestroyMessage()
	PublisherFini(myPub, myNode)
	// Shutdown node object.
	node.NodeFini(myNode)
	// Shutdown context object.
	rcl.RCLShutdown(contextObject.ContextKey)
}

/*
func TestPublisherInt8Msg(t *testing.T) {

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	msg := make(chan string, 1)
	go func() {
		// Receive input in a loop
		for {
			var s string
			fmt.Scan(&s)
			// Send what we read over the channel
			msg <- s
		}
	}()

	// Initialization
	// Get context object.
	contextObject := rcl.GetZeroInitializedContext()
	rcl.Init(&contextObject)
	myNode := node.GetZeroInitializedNode()
	myNodeOpts := node.GetNodeDefaultOptions()

	fmt.Printf("Creating the node! \n")
	node.NodeInit(myNode, "GoPublisher", "", myNodeOpts)

	//Create the publisher
	myPub := GetZeroInitializedPublisher()
	myPubOpts := GetPublisherDefaultOptions()

	//Create the msg type
	var myMsg types.StdMsgsInt8
	myMsg.InitMessage()
	myMsg.SetInt8(32)

	fmt.Printf("Creating the publisher! \n")
	//Initializing the publisher
	PublisherInit(myPub, myPubOpts, myNode, "/myGoTopic", myMsg.GetMessage())

	index := 0
loop:
	for {
		//Update my msg
		myMsg.SetInt8(myMsg.GetInt8() + 1)

		//Publish the message
		retRCL := Publish(myPub, myMsg.GetMessage(), myMsg.GetData())

		if retRCL == types.RCL_RET_OK {
			var value string
			fmt.Sprintf("%v", value)

			fmt.Printf("(Publisher) Published: %v\n", myMsg.GetInt8())
		}
		time.Sleep(500 * time.Millisecond)
		index++

		//Loop breaker
		select {
		case <-sigs:
			fmt.Println("Got shutdown, exiting")
			// Break out of the outer for statement and end the program
			break loop
		case <-msg:

		}
	}

	fmt.Printf("Shutting down!! \n")

	myMsg.DestroyMessage()
	PublisherFini(myPub, myNode)
	node.NodeFini(myNode)
	rcl.Shutdown(&contextObject)

}

func TestPublisherFloat64Msg(t *testing.T) {

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	msg := make(chan string, 1)
	go func() {
		// Receive input in a loop
		for {
			var s string
			fmt.Scan(&s)
			// Send what we read over the channel
			msg <- s
		}
	}()

	// Initialization
	// Get context object.
	contextObject := rcl.GetZeroInitializedContext()
	rcl.Init(&contextObject)
	myNode := node.GetZeroInitializedNode()
	myNodeOpts := node.GetNodeDefaultOptions()

	fmt.Printf("Creating the node! \n")
	node.NodeInit(myNode, "GoPublisher", "", myNodeOpts)

	//Create the publisher
	myPub := GetZeroInitializedPublisher()
	myPubOpts := GetPublisherDefaultOptions()

	//Create the msg type
	var myMsg types.StdMsgsFloat64
	myMsg.InitMessage()
	myMsg.SetFloat64(32.0)

	fmt.Printf("Creating the publisher! \n")
	//Initializing the publisher
	PublisherInit(myPub, myPubOpts, myNode, "/myGoTopic", myMsg.GetMessage())

	index := 0
loop:
	for {
		//Update my msg
		myMsg.SetFloat64(myMsg.GetFloat64() + 0.25)

		//Publish the message
		retRCL := Publish(myPub, myMsg.GetMessage(), myMsg.GetData())

		if retRCL == types.RCL_RET_OK {
			var value string
			fmt.Sprintf("%v", value)

			fmt.Printf("(Publisher) Published: %v\n", myMsg.GetFloat64())
		}
		time.Sleep(500 * time.Millisecond)
		index++

		//Loop breaker
		select {
		case <-sigs:
			fmt.Println("Got shutdown, exiting")
			// Break out of the outer for statement and end the program
			break loop
		case <-msg:

		}
	}

	fmt.Printf("Shutting down!! \n")

	myMsg.DestroyMessage()
	PublisherFini(myPub, myNode)
	node.NodeFini(myNode)
	rcl.Shutdown(&contextObject)

}
*/
