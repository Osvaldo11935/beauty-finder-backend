package usecase

import (
	"fmt"
	"src/internal/domain/entities"
	"sync"

	"github.com/google/uuid"
)

type Pool struct {
	Register   chan *entities.User
	Unregister chan *entities.User
	Clients    map[*entities.User]bool
	Broadcast  chan entities.Message
	mu         sync.Mutex
}

func NewPool() *Pool {
	return &Pool{
		Register:   make(chan *entities.User),
		Unregister: make(chan *entities.User),
		Clients:    make(map[*entities.User]bool),
		Broadcast:  make(chan entities.Message),
	}
}

func (pool *Pool) Start() {
	for {
		select {
		case user := <-pool.Register:
			pool.mu.Lock()
			pool.Clients[user] = true
			fmt.Printf("Usuário conectado: %s\n", user.UserName)
			pool.mu.Unlock()

		case user := <-pool.Unregister:
			pool.mu.Lock()
			if _, ok := pool.Clients[user]; ok {
				delete(pool.Clients, user)
				fmt.Printf("Usuário desconectado: %s\n", user.UserName)
			}
			pool.mu.Unlock()

		case message := <-pool.Broadcast:
			pool.mu.Lock()

			if message.ReceiverId == uuid.Nil {
				for client := range pool.Clients {
					if client.ID != message.Sender.ID {
						err := client.SendMessage(&message)
						if err != nil {
							fmt.Printf("rro ao enviar mensagem para %s: %v\n", client.UserName, err)
							client.CloseConnection()
							delete(pool.Clients, client)
						}
					}
				}
			} else {
				var receiver *entities.User
				for client := range pool.Clients {
					if client.ID == message.ReceiverId {
						receiver = client
						break
					}
				}
				if receiver != nil {
					err := receiver.SendMessage(&message)
					if err != nil {
						fmt.Printf("Erro ao enviar mensagem para %s: %v\n", receiver.UserName, err)
						receiver.CloseConnection()
						delete(pool.Clients, receiver)
					}
				} else {
					fmt.Printf("Receptor com ID %s não encontrado\n", message.ReceiverId)
				}
			}

			pool.mu.Unlock()
		}
	}
}
