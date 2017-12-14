package main

import (
	"fmt"
	"github.com/mrlsd/reqrep-events/lib"
)

const MQ_URL_CLIENT = "tcp://127.0.0.1:40899"

func main() {
	lib.LogInfo("Starting REQREP client")
	evd := lib.EventData{
		EventId:   100,
		EventName: "Test",
	}
	ev := lib.EventProxy{
		Url:   MQ_URL_CLIENT,
		Name:  "test-mq-client",
		Event: evd,
	}
	if msg, err := lib.SendEvent(ev); err != nil {
		fmt.Printf("EventData: %#v\n", msg)
		fmt.Printf("Error: %s\n", err)
	}
}
