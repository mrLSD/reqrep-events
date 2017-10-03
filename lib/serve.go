package lib

import (
	"encoding/json"
	"fmt"
	"github.com/go-mangos/mangos"
	"github.com/go-mangos/mangos/protocol/rep"
	"github.com/go-mangos/mangos/transport/ipc"
	"github.com/go-mangos/mangos/transport/tcp"
)

// eventSenrveSocketsPool - socket Pool for Receive
var eventSenrveSocketsPool = make(map[string]mangos.Socket, 0)

// ServeEvents - serve Message Events - Responses
// REQ/REP message queue protocol
func ServeEvents(evp EventProxy, evh EventHandler) (err error) {
	var sock mangos.Socket
	var msg EventData
	// TBD: eventSenrveSocketsPool
	if sock, err = rep.NewSocket(); err != nil {
		err = fmt.Errorf("Can't get new REP socket: %s", err)
		return
	}
	sock.AddTransport(ipc.NewTransport())
	sock.AddTransport(tcp.NewTransport())
	if err = sock.Listen(evp.Url); err != nil {
		err = fmt.Errorf("Can't listen on REP socket: %s", err.Error())
		return
	}
	for {
		// Could also use sock.RecvMsg to get header
		msgData, err := sock.Recv()
		println("GET MSG")
		if err != nil {
			// TBD: error actions
		}
		err = json.Unmarshal(msgData, &msg)
		if err != nil {
			// TBD: error actions
		}
		evData, err := evh(msg)
		if err != nil {
			// TBD: error actions
		}
		msgData, err = json.Marshal(evData)
		if err != nil {
			// TBD: error actions
		}
		err = sock.Send(msgData)
		if err != nil {
			// TBD: err actions
		}
	}
	return nil
}
