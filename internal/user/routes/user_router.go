package Userrouter

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	. "we.care/they-chat/internal/user/dao"
	. "we.care/they-chat/internal/user/model"
	"gopkg.in/mgo.v2/bson"
)

var dao = UsersDAO{}

func init(){
	dao.Connect()
}

func Load(r *mux.Router){
	r.HandleFunc("/api/v1/users", GetAll).Methods("GET")
	r.HandleFunc("/api/v1/users/{id}", GetByID).Methods("GET")
	r.HandleFunc("/api/v1/users", Create).Methods("POST")
	r.HandleFunc("/api/v1/users/{id}", Update).Methods("PUT")
	r.HandleFunc("/api/v1/users/{id}", Delete).Methods("DELETE")
	// TODO remove id param and get id from user authentication
	r.HandleFunc("/api/v1/users/password/change", UpdatePassword).Methods("PUT")
	r.HandleFunc("/api/v1/users/auto/config", AutoConfig).Methods("GET")
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
	Users, err := dao.GetAll()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, Users)
}

func GetByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	User, err := dao.GetByID(params["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid User ID")
		return
	}
	respondWithJson(w, http.StatusOK, User)
}

func Create(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var User User
	if err := json.NewDecoder(r.Body).Decode(&User); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	User.ID = bson.NewObjectId()
	if err := dao.Create(User); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusCreated, User)
}

func Update(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	params := mux.Vars(r)
	var User User
	if err := json.NewDecoder(r.Body).Decode(&User); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if err := dao.Update(params["id"], User); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, map[string]string{"result": User.Name + " atualizado com sucesso!"})
}

// TODO change for not pass User full
func UpdatePassword(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var UserAuth UserAuth;
	if err := json.NewDecoder(r.Body).Decode(&UserAuth); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if err := dao.UpdatePassword(UserAuth.ID, UserAuth); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, map[string]string{"result": "password atualizado com sucesso!"})
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

func AutoConfig(w http.ResponseWriter, r *http.Request) {
	createUser("patient1","Patient","Patient 1","secret")
	createUser("patient2","Patient","Patient 2","secret")
	createUser("doctor1","Doctor","Doctor 1","secret")
	createUser("doctor2","Doctor","Doctor 2","secret")
	respondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}

func createUser(login string, usertype string, name string, pass string){
	User := User{}
	User.Name = name
	User.Type = usertype
	User.Login = login
	User.Active = true
	User.ID = bson.NewObjectId()
	dao.Create(User)

	dao.UpdatePassword(User.ID.Hex(), UserAuth{Password:pass})
}