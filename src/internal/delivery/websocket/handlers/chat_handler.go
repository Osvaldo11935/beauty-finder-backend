package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"src/internal/domain/entities"
	"src/internal/usecase"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func WebSocketHandler(pool *usecase.Pool, userUseCase *usecase.UserUseCase, w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Erro ao conectar WebSocket:", err)
		return
	}

	userId := uuid.MustParse(r.URL.Query().Get("userId"))

	user, _ := userUseCase.FindUserById(userId)

	user.Conn = conn

	pool.Register <- user

	go read(user, userUseCase, pool)
}

func read(user *entities.User, userUseCase *usecase.UserUseCase, pool *usecase.Pool) {
	if user == nil {
		return
	}

	defer func() {
		pool.Unregister <- user
		user.CloseConnection()
	}()

	for {
		// Ler a mensagem bruta
		_, msg, err := user.Conn.ReadMessage()
		if err != nil {
			fmt.Println("Erro ao ler mensagem:", err)
			return
		}

		fmt.Printf("Mensagem recebida: %s\n", msg)

		var tempMessage struct {
			Type       int       `json:"type"`
			Body       string    `json:"body"`
			ReceiverId uuid.UUID `json:"receiverId"`
		}
		err = json.Unmarshal(msg, &tempMessage)

		if err != nil {
			fmt.Println("Erro ao decodificar mensagem JSON:", err)
			return
		}

		userReceiver, _ := userUseCase.FindUserById(tempMessage.ReceiverId)

		message := entities.Message{
			Type:     tempMessage.Type,
			Body:     tempMessage.Body,
			Receiver: userReceiver,
			Sender:   user,
			SenderId: user.ID,
		}

		if userReceiver != nil {
			message.ReceiverId = userReceiver.ID
		}

		pool.Broadcast <- message
	}
}
