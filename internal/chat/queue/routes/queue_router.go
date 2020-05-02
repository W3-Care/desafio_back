package ChatQueuerouter

import (
	// "fmt"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	. "we.care/they-chat/internal/chat/queue"
	"gopkg.in/mgo.v2/bson"
)

var dao = ChatQueueDAO{}

func init(){
	dao.Connect()
}

func Load(r *mux.Router){
	r.HandleFunc("/api/v1/queues", GetAll).Methods("GET")
	r.HandleFunc("/api/v1/queues/{id}", GetByID).Methods("GET")
	r.HandleFunc("/api/v1/queues", Create).Methods("POST")
	r.HandleFunc("/api/v1/queues/{id}", Delete).Methods("DELETE")
	r.HandleFunc("/api/v1/queues/first", PopFirst).Methods("GET")
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
	ChatQueues, err := dao.GetAll()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, ChatQueues)
}

func GetByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	ChatQueue, err := dao.GetByID(params["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid ChatQueue ID")
		return
	}
	respondWithJson(w, http.StatusOK, ChatQueue)
}

func PopFirst(w http.ResponseWriter, r *http.Request) {
	ChatQueue, err := dao.GetFirst()
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid ChatQueue ID")
		return
	}
	respondWithJson(w, http.StatusOK, ChatQueue)
}

func Create(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var ChatQueue ChatQueue
	if err := json.NewDecoder(r.Body).Decode(&ChatQueue); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	ChatQueue.ID = bson.NewObjectId()
	if err := dao.Create(ChatQueue); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusCreated, ChatQueue)
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