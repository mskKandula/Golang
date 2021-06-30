package chat

type WsServer struct {
	Clients    map[*Client]bool
	Rooms      map[*Room]bool
	Register   chan *Client
	UnRegister chan *Client
	Broadcast  chan []byte
}

// NewWebsocketServer creates a new WsServer type
func NewWebsocketServer() *WsServer {
	return &WsServer{
		Clients:    make(map[*Client]bool),
		Rooms:      make(map[*Room]bool),
		Register:   make(chan *Client),
		UnRegister: make(chan *Client),
		Broadcast:  make(chan []byte),
	}
}

// Run our websocket server, accepting various requests
func (server *WsServer) Run() {

	for {
		select {
		case client := <-server.Register:
			server.RegisterClient(client)

		case msg := <-server.Broadcast:
			server.BroadcastMsg(msg)

		case client := <-server.UnRegister:
			server.UnRegisterClient(client)
		}
	}
}

func (server *WsServer) RegisterClient(client *Client) {
	server.Clients[client] = true
}

func (server *WsServer) BroadcastMsg(msg []byte) {

	for client := range server.Clients {
		client.conn.WriteJSON(msg)
	}

}

func (server *WsServer) UnRegisterClient(client *Client) {
	if _, ok := server.Clients[client]; ok {
		delete(server.Clients, client)
	}
}

func (server *WsServer) FindRoomByName(name string) *Room {
	var foundRoom *Room
	for room := range server.Rooms {
		if room.Name == name {
			foundRoom = room
			break
		}
	}
	return foundRoom
}

func (server *WsServer) CreateNewRoom(name string) *Room {
	room := NewRoom(name)

	go room.RunRoom()

	server.Rooms[room] = true

	return room

}
