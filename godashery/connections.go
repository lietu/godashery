package godashery

import (
	"log"
	"encoding/json"
	"github.com/nu7hatch/gouuid"
	"github.com/gorilla/websocket"
)

var connections = map[uuid.UUID]ConnectionWrapper{}

type ConnectionWrapper struct {
	uuid          uuid.UUID
	connection    *websocket.Conn
	subscriptions []string
}

func (c *ConnectionWrapper) Close() {
	delete(connections, c.uuid)
}

func (c *ConnectionWrapper) Send(data []byte) {
	c.connection.WriteMessage(websocket.TextMessage, data)
}

func (c *ConnectionWrapper) SendMessage(m interface{}) {

	result, err := json.Marshal(&m)

	if err != nil {
		log.Fatalf("error: %v", err)
	}

	c.Send(result)
}

func NewConnectionWrapper(connection *websocket.Conn) *ConnectionWrapper {
	uuidPtr, err := uuid.NewV4()

	if err != nil {
		log.Fatalf("Error creating uuid: %s", err)
	}

	uuid := *uuidPtr

	wrapper := ConnectionWrapper{
		uuid,
		connection,
		[]string{},
	}

	connections[uuid] = wrapper

	return &wrapper
}