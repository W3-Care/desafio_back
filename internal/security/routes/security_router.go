package security

import (
	// "fmt"
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
	. "we.care/they-chat/internal/user/dao"
	. "we.care/they-chat/internal/user/model"
	queuedao "we.care/they-chat/internal/chat/queue"
	chatdao "we.care/they-chat/internal/chat/dao"
	security "we.care/they-chat/internal/security"
	jwt "github.com/dgrijalva/jwt-go"
	"os"
	"log"
	"gopkg.in/mgo.v2/bson"

)

var dao = UsersDAO{}
var queueDao = queuedao.ChatQueueDAO{}
var chatDao = chatdao.ChatDAO{}

func init(){
	dao.Connect()
	queueDao.Connect()
	chatDao.Connect()
}

func Load(r *mux.Router){
	r.HandleFunc("/api/v1/authentication", Authenticate).Methods("POST")
}

// TODO abstract this method
func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJson(w, code, map[string]string{"error": msg})
}

// TODO abstract this method
func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func Authenticate(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var UserAuth UserAuth
	if err := json.NewDecoder(r.Body).Decode(&UserAuth); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	User, err := dao.GetByLogin(UserAuth.Login)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid User Login")
		return
	}
	if !security.CheckPasswordHash(UserAuth.Password, User.Hash) {
		respondWithError(w, http.StatusBadRequest, "Invalid User Login or Password")
		return
	}

	//Create JWT token
	tk := &Token{UserId: User.ID}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(os.Getenv("token_password")))
	tk.Token = tokenString
	tk.Name = User.Name
	tk.Type = User.Type
	if User.Type == "Patient" {
		_, err := queueDao.GetByIDPatient(User.ID)
		if err != nil {
			chatObj, chaterr := chatDao.GetActiveByPatient(User.ID.Hex())
			if chaterr != nil {
				chatQueue := queuedao.ChatQueue{}
				chatQueue.ID = bson.NewObjectId()
				chatQueue.IDPatient = User.ID
				if err := queueDao.Create(chatQueue); err != nil {
					log.Fatal("Fail when try to queue patient")
				}
			} else {
				tk.ChatId = chatObj.ID
			}
		}
	} else {
		chatObj, chaterr := chatDao.GetActiveByDoctor(User.ID.Hex())
		if chaterr == nil {
			tk.ChatId = chatObj.ID
		}
	}
	respondWithJson(w, http.StatusOK, tk)
}
