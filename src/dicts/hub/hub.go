package hub

import (
	"2019_2_Shtoby_shto/src/customType"
	сardUsers "2019_2_Shtoby_shto/src/dicts/cardUsers"
	"2019_2_Shtoby_shto/src/dicts/models"
)

// Hub maintains the set of active clients and broadcasts messages to the
// clients.
type Hub struct {
	// Registered clients.
	Clients map[*Client]customType.StringUUID

	// Inbound messages from the clients.
	Broadcast chan []byte

	// Register requests from the clients.
	Register chan *Client

	// Unregister requests from clients.
	Unregister chan *Client

	cardUsersService сardUsers.HandlerCardUsersService
}

func NewHub(cardUsersService сardUsers.HandlerCardUsersService) *Hub {
	return &Hub{
		Broadcast:        make(chan []byte),
		Register:         make(chan *Client),
		Unregister:       make(chan *Client),
		Clients:          make(map[*Client]customType.StringUUID),
		cardUsersService: cardUsersService,
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.Register:
			message := <-h.Broadcast
			user := &models.RegUser{}
			err := user.UnmarshalJSON(message)
			if err != nil {
				println(err)
			}
			h.Clients[client] = user.UserID
			client.Send <- []byte(`"message":"Status Ok""`)
		case client := <-h.Unregister:
			if _, ok := h.Clients[client]; ok {
				delete(h.Clients, client)
				close(client.Send)
			}
		case message := <-h.Broadcast:
			attachRequest := &models.CardsUserAttachRequest{}
			err := attachRequest.UnmarshalJSON(message)
			if err != nil {
				println(err)
			}
			userIDs, err := h.cardUsersService.FetchUserIDsByCardID(attachRequest.CardID)
			if err != nil {
				println(err)
			}
			for client, userID := range h.Clients {
				println(userID)
				if _, ok := userIDs[userID.String()]; ok {
					select {
					case client.Send <- message:
					default:
						close(client.Send)
						delete(h.Clients, client)
					}
				}
			}
		}
	}
}
