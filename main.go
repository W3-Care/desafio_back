package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	userrouter "we.care/they-chat/internal/user/routes"
	chatrouter "we.care/they-chat/internal/chat/routes"
	chatbox "we.care/they-chat/internal/chat/chatbox"
	securityrouter "we.care/they-chat/internal/security/routes"
	qrcoderouter "we.care/they-chat/internal/qrcode"
	jwtAuthentication "we.care/they-chat/internal/security/jwt-auth"
	
)

func main() {
	r  := mux.NewRouter()
	r.Use(jwtAuthentication.JwtAuthentication) //attach JWT auth middleware
	userrouter.Load(r)
	chatrouter.Load(r)
	securityrouter.Load(r)
	qrcoderouter.Load(r)
	// TODO chat view
	r.PathPrefix("/simplechat/").Handler(http.StripPrefix("/simplechat/", 
	http.FileServer(http.Dir("./public"))))

	// Configure websocket route
	r.HandleFunc("/chatbox", chatbox.HandleConnections)

	// Start listening for incoming chat messages
	go chatbox.HandleMessages()
	var port = ":3000"
	fmt.Println("Server running in port:", port)
	log.Fatal(http.ListenAndServe(port, r))
}

