package chat

type WsServer struct {
	Clients    map[*Client]bool
	Register   chan *Client
	UnRegister chan *Client
}

// NewWebsocketServer creates a new WsServer type
func NewWebsocketServer() *WsServer {
	return &WsServer{
		Clients:    make(map[*Client]bool),
		Register:   make(chan *Client),
		UnRegister: make(chan *Client),
	}
}

// Run our websocket server, accepting various requests
func (server *WsServer) Run() {

	for {
		select {
		case client := <-server.Register:
			server.RegisterClient(client)

		case client := <-server.UnRegister:
			server.UnRegisterClient(client)
		}
	}
}

func (server *WsServer) RegisterClient(client *Client) {
	server.Clients[client] = true
}

func (server *WsServer) UnRegisterClient(client *Client) {
	if _, ok := server.Clients[client]; ok {
		delete(server.Clients, client)
	}
}
