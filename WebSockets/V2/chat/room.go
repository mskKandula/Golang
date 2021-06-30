package chat

type Room struct {
	Name       string
	Clients    map[*Client]bool
	Register   chan *Client
	UnRegister chan *Client
	Broadcast  chan *Message
}

func NewRoom(name string) *Room {
	return &Room{
		Name:       name,
		Clients:    make(map[*Client]bool),
		Register:   make(chan *Client),
		UnRegister: make(chan *Client),
		Broadcast:  make(chan []byte),
	}
}

func (room *Room) RunRoom() {
	for {
		select {
		case client := <-room.Register:
			room.RegisterClientInRoom(client)

		case msg := <-room.Broadcast:
			room.SendMsg(msg)

		case client := <-room.UnRegister:
			room.UnRegisterClientInRoom(client)
		}
	}
}

func (room *Room) RegisterClientInRoom(client *Client) {
	room.Clients[client] = true
}

func (room *Room) SendMsg(msg []byte) {
	for client := range room.Clients {
		client.send <- msg
	}
}

func (room *Room) UnRegisterClientInRoom(client *Client) {
	if _, ok := room.Clients[client]; ok {
		delete(room.Clients, client)
	}
}
