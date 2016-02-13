package godashery

import (
	"github.com/nu7hatch/gouuid"
)

var subscriptions = map[string][]uuid.UUID{}

func SendHello(uuid uuid.UUID) {
	SendMessageToConnection(uuid, NewHello())

	for _, w := range widgets {
		SendToConnection(uuid, w.GetLastValue())
	}
}

func HandleMessage(uuid uuid.UUID, messageContent []byte) {

	/*
	message := MessageIn{}

	json.Unmarshal(messageContent, &message)

	switch {
	default:
		log.Printf("Unknown message received of type %s from %s", message.Type, uuid.String())

	case message.Type == "subscribe":
		log.Printf("Connection %s subscribing to %s", uuid.String(), message.SteamId64)
		onSubscribe(uuid, message)

	}
	*/

}
