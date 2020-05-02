package Chatrouter

import (
	"time"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	chatqueueroutes "we.care/they-chat/internal/chat/queue/routes"
	utils "we.care/they-chat/internal/utils"
	queue "we.care/they-chat/internal/chat/queue"
	. "we.care/they-chat/internal/chat/dao"
	"gopkg.in/mgo.v2/bson"
)

var dao = ChatDAO{}
var queueDao = &queue.ChatQueueDAO{}

func init(){
	dao.Connect()
	// queueDao.Connect()
}

func Load(r *mux.Router){
	r.HandleFunc("/api/v1/chats", GetAll).Methods("GET")
	r.HandleFunc("/api/v1/chats/{id}", GetByID).Methods("GET")
	r.HandleFunc("/api/v1/chats", Create).Methods("POST")
	r.HandleFunc("/api/v1/chats/{id}", Delete).Methods("DELETE")
	r.HandleFunc("/api/v1/chats/start", Start).Methods("POST")
	r.HandleFunc("/api/v1/chats/terminate/{id}", TerminateChat).Methods("POST")
	
	
	chatqueueroutes.Load(r)
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJson(w, code, map[string]string{"error": msg})
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func GetAll(w http.ResponseWriter, r *http.Request) {
	Chats, err := dao.GetAll()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, Chats)
}

func GetByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	Chat, err := dao.GetByID(params["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Chat ID")
		return
	}
	respondWithJson(w, http.StatusOK, Chat)
}

func Create(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var Chat Chat
	if err := json.NewDecoder(r.Body).Decode(&Chat); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	Chat.ID = bson.NewObjectId()
	if err := dao.Create(Chat); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusCreated, Chat)
}

func Start(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	firstQueue, err := queueDao.GetFirst()
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Without queue")
		return
	}
	chat := Chat{}
	chat.IDPatient = firstQueue.IDPatient
	str, _ := utils.ExtractTokenID(r)
	chat.IDDoctor = bson.ObjectIdHex(str)
	chat.Start = time.Now().String()
	chat.Active = true
	chat.ID = bson.NewObjectId()
	if err := dao.Create(chat); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	queueDao.Delete(firstQueue.ID.Hex())
	
	respondWithJson(w, http.StatusCreated, chat)
}

func Update(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	params := mux.Vars(r)
	var Chat Chat
	if err := json.NewDecoder(r.Body).Decode(&Chat); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if err := dao.Update(params["id"], Chat); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, map[string]string{"result": " atualizado com sucesso!"})
}

func Delete(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	params := mux.Vars(r)
	if err := dao.Delete(params["id"]); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}

func TerminateChat(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	Chat, err := dao.GetByID(params["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Chat ID")
		return
	}
	str, _ := utils.ExtractTokenID(r)

	if bson.ObjectIdHex(str) != Chat.IDDoctor {
		respondWithError(w, http.StatusBadRequest, "Only current doctor can terminate a chat")
		return
	}

	Chat.Active = false
	Chat.End = time.Now().String()
	if err := dao.Update(params["id"], Chat); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, Chat)
}
