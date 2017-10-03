package lib

import (
	"github.com/go-mangos/mangos"
	"github.com/go-mangos/mangos/transport/ipc"
	"github.com/go-mangos/mangos/transport/tcp"

	"encoding/json"
	"fmt"
	"github.com/go-mangos/mangos/protocol/req"
)

// eventSenderSocketsPool - socket Pool for Send Reqeists
var eventSenderSocketsPool = make(map[string]mangos.Socket, 0)

// SendEvent - Send Message Event - Request
// REQ/REP message queue protocol
func SendEvent(evp EventProxy) (msg EventData, err error) {
	var sock mangos.Socket
	// TBD: evp
	if sock, err = req.NewSocket(); err != nil {
		err = fmt.Errorf("can't get new REQ socket: %s", err)
		return
	}
	sock.AddTransport(ipc.NewTransport())
	sock.AddTransport(tcp.NewTransport())
	if err = sock.Dial(evp.Url); err != nil {
		err = fmt.Errorf("Can't dial on REQ socket: %s", err.Error())
		return
	}
	msgData, err := json.Marshal(evp.Event)
	fmt.Printf("EventData: %#v\n", evp.Event)
	if err != nil {
		err = fmt.Errorf("Failed parse JSON: %s", evp.Url, evp.Name, err.Error())
		return
	}
	fmt.Printf("Msg: %s\n", string(msgData))
	// Send data to REQ server
	if err = sock.Send(msgData); err != nil {
		err = fmt.Errorf("Can't send message to socket for [%s] %s Error:", evp.Url, evp.Name, err.Error())
		return
	}

	// Recieve data as REP client
	if msgData, err = sock.Recv(); err != nil {
		err = fmt.Errorf("Can't receive date from Event Proxy: %s", err.Error())
		return
	} else {
		err = json.Unmarshal(msgData, &msg)
		if err != nil {
			err = fmt.Errorf("Failed to parse Event Message: %s", err)
			return
		}
	}
	return msg, nil
}
