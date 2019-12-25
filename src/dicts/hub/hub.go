package hub

import (
	"2019_2_Shtoby_shto/src/customType"
	"2019_2_Shtoby_shto/src/dicts/boardUsers"
	сardUsers "2019_2_Shtoby_shto/src/dicts/cardUsers"
	"2019_2_Shtoby_shto/src/dicts/models"
)

// Hub maintains the set of active clients and broadcasts messages to the
// clients.
type Hub struct {
	// Registered clients.
	Clients map[customType.StringUUID]*Client

	// Inbound messages from the clients.
	Broadcast chan []byte

	// Register requests from the clients.
	Register chan *Client

	// Unregister requests from clients.
	Unregister chan *Client

	cardUsersService  сardUsers.HandlerCardUsersService
	boardUsersService boardUsers.HandlerBoardUsersService
}

func NewHub(cardUsersService сardUsers.HandlerCardUsersService,
	boardUsersService boardUsers.HandlerBoardUsersService) *Hub {
	return &Hub{
		Broadcast:         make(chan []byte),
		Register:          make(chan *Client),
		Unregister:        make(chan *Client),
		Clients:           make(map[customType.StringUUID]*Client),
		cardUsersService:  cardUsersService,
		boardUsersService: boardUsersService,
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
			h.Clients[user.UserID] = client
			client.Send <- []byte(`"message":"Status Ok"`)
		case client := <-h.Unregister:
			for userID, cli := range h.Clients {
				if cli == client {
					delete(h.Clients, userID)
					close(client.Send)
				}
			}
		case message := <-h.Broadcast:
			attachBoardRequest := &models.BoardsUserAttachRequest{}

			if err := attachBoardRequest.UnmarshalJSON(message); err == nil {
				println(err)
			}

			userIDs, err := h.boardUsersService.FetchUserIDsByBoardID(attachBoardRequest.BoardID)
			if err != nil {
				println(err)
			}

			h.sendUserMessage(userIDs, message)

			attachCardRequest := &models.CardsUserAttachRequest{}
			err = attachCardRequest.UnmarshalJSON(message)
			if err != nil {
				println(err)
			}

			userIDs, err = h.cardUsersService.FetchUserIDsByCardID(attachCardRequest.CardID)
			if err != nil {
				println(err)
			}

			h.sendUserMessage(userIDs, message)
		}
	}
}

func (h *Hub) sendUserMessage(userIDs map[string]struct{}, message []byte) {
	for userID, client := range h.Clients {
		println(userID)
		if _, ok := userIDs[userID.String()]; ok {
			select {
			case client.Send <- message:

			default:
				close(client.Send)
				delete(h.Clients, userID)
			}
		}
	}
}
