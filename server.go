package main

import (
	"github.com/mrlsd/reqrep-events/lib"
	"fmt"
)

const MQ_URL_SERVER = "tcp://127.0.0.1:40899"

func main() {
	ev := lib.EventProxy{
		Url:  MQ_URL_SERVER,
		Name: "test-mq-server",
	}
	if err := lib.ServeEvents(ev, eventHandler); err != nil {
		fmt.Printf("Failed start server: %s", err)
	}
}

func eventHandler(ev lib.EventData) (lib.EventData, error) {
	evd := lib.EventData{}
	return evd, nil
}
