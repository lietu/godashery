package godashery

import (
	"fmt"
	"log"
	"net/http"
	"github.com/nu7hatch/gouuid"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

func SendToConnection(uuid uuid.UUID, data []byte) {
	conn, ok := connections[uuid]

	if ok {
		conn.Send(data)
	}
}

func SendMessageToConnection(uuid uuid.UUID, m interface{}) {
	conn, ok := connections[uuid]

	if ok {
		conn.SendMessage(m)
	}
}

func websocketHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		log.Print("upgrade:", err)
		return
	}

	defer conn.Close()

	if err != nil {
		log.Fatalf("error:", err)
	}

	wrapper := NewConnectionWrapper(conn)

	log.Printf("Connection %s from %s", wrapper.uuid.String(), r.RemoteAddr)

	SendHello(wrapper.uuid)

	for {
		_, message, err := conn.ReadMessage()

		if err != nil {
			if err.Error() == "websocket: close 1001 " {
				log.Printf("Client %s disconnected", wrapper.uuid.String())
			} else {
				log.Printf("Client %s error, %s", err.Error())
			}
			wrapper.Close()
			break
		}

		HandleMessage(wrapper.uuid, message)
	}
}

func RunServer(settings Settings) {
	address := fmt.Sprintf("%s:%d", settings.ListenAddress, settings.ListenPort)

	log.Printf("Listening to %s", address)
	log.Printf("Serving static files from %s", settings.WwwPath)

	http.HandleFunc("/data", websocketHandler)
	http.Handle("/", http.FileServer(http.Dir(settings.WwwPath)))

	log.Fatal(http.ListenAndServe(address, nil))
}