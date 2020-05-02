package userdao

import (
	"log"
	. "we.care/they-chat/internal/user/model"
	config "we.care/they-chat/configs"
	security "we.care/they-chat/internal/security"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type UsersDAO struct {
	Server   string
	Database string
}

var db *mgo.Database

const (
	COLLECTION = "users"
)

func (m *UsersDAO) Connect() {
	c := config.Config{}
	c.Read()
	session, err := mgo.Dial(c.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(c.Database)
}

func (m *UsersDAO) GetAll() ([]User, error) {
	var Users []User
	err := db.C(COLLECTION).Find(bson.M{}).All(&Users)
	return Users, err
}

func (m *UsersDAO) GetByID(id string) (User, error) {
	var User User
	err := db.C(COLLECTION).FindId(bson.ObjectIdHex(id)).One(&User)
	return User, err
}

func (m *UsersDAO) GetByLogin(login string) (User, error) {
	var User User
	err := db.C(COLLECTION).Find(bson.M{"login" : login}).One(&User)
	return User, err
}

func (m *UsersDAO) Create(User User) error {
	err := db.C(COLLECTION).Insert(&User)
	return err
}

func (m *UsersDAO) Delete(id string) error {
	err := db.C(COLLECTION).RemoveId(bson.ObjectIdHex(id))
	return err
}

func (m *UsersDAO) Update(id string, User User) error {
	err := db.C(COLLECTION).UpdateId(bson.ObjectIdHex(id), &User)
	return err
}

func (m *UsersDAO) UpdatePassword(id string, UserAuth UserAuth) error {
	hash, _ := security.HashPassword(UserAuth.Password)
	err := db.C(COLLECTION).UpdateId(bson.ObjectIdHex(id), bson.M{"$set": bson.M{"hash": hash}})
	return err
}
