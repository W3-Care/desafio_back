package chatbox

import (
	"log"
	"net/http"
	"gopkg.in/mgo.v2/bson"
	"github.com/gorilla/websocket"
	utils "we.care/they-chat/internal/utils"
	. "we.care/they-chat/internal/chat/dao"
)

var clients = make(map[string]map[*websocket.Conn]bool)
var waitClients = make(map[string]*websocket.Conn) // connected clients
// var clients = make(map[string]map) // connected clients
var broadcast = make(chan Message)           // broadcast channel

var chatDAO = ChatDAO{}
// Configure the upgrader
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// Define our message object
type Message struct {
	ID    string `json:"id"`
	Username string `json:"username"`
	Message  string `json:"message"`
	ChatID  string `json:"chatId"`
}

func HandleConnections(w http.ResponseWriter, r *http.Request) {
	// Upgrade initial GET request to a websocket
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}
	// Make sure we close the connection when the function returns
	defer ws.Close()

	// Register our new client
	

	for {
		var msg Message
		// Read in a new message as JSON and map it to a Message object
		err := ws.ReadJSON(&msg)
		if msg.ChatID == "" {
			token, _ := utils.ExtractTokenID(r)
			chatObj, chaterr := chatDAO.GetActiveByPatient(token)
			if chaterr == nil {
				msg.ChatID = chatObj.ID.Hex()
			}
		}
		if msg.ChatID != "" {
			if clients[msg.ChatID] == nil{
				clients[msg.ChatID] = createMessageChatBox(msg,ws)
			}
			if clients[msg.ChatID][ws] != true {
				if !checkBeforeStartWS(msg, r) {
					// TODO Add error msg to reponse and close WS
					log.Fatal("Without permission for this room")
				}
				clients[msg.ChatID][ws] = true
			}
		} else {
			token, _ := utils.ExtractTokenID(r)
			if waitClients[token] == nil {
				waitClients[token]=ws
			}
		}
		if err != nil {
			log.Printf("error: %v", err)
			delete(clients[msg.ChatID], ws)
			if len(clients[msg.ChatID]) == 0 {
				delete(clients, msg.ChatID)
			}
			break
		}
		// Send the newly received message to the broadcast channel
		broadcast <- msg
	}
}

func HandleMessages() {
	for {
		// Grab the next message from the broadcast channel
		msg := <-broadcast
		// Send it out to every client that is currently connected
		for client := range clients[msg.ChatID] {
			err := client.WriteJSON(msg)
			if err != nil {
				log.Printf("error: %v", err)
				client.Close()
				delete(clients[msg.ChatID], client)
				if len(clients[msg.ChatID]) == 0 {
					delete(clients, msg.ChatID)
				}
			}
		}
	}
}

func createMessageChatBox(msg Message, ws *websocket.Conn) (map[*websocket.Conn]bool){
	chatObj, err := (chatDAO.GetByID(msg.ChatID))
	if err != nil {
		return nil
	}
	m := make(map[*websocket.Conn]bool)
	m[ws] = true
	cws := waitClients[chatObj.IDPatient.Hex()]
	if cws != nil {
		m[cws] = true
	}
	return m
}
func checkBeforeStartWS(msg Message, r *http.Request) (bool) {
	chatObj, err := (chatDAO.GetByID(msg.ChatID))
	if err != nil {
		return false
	}
	// TODO change when auth enabled
	token, _ := utils.ExtractTokenID(r)
	userId := bson.ObjectIdHex(token)
	if userId != chatObj.IDPatient && userId != chatObj.IDDoctor {
		return false
	}
	return true
}
