package chatdao

import (
	"log"
	config "we.care/they-chat/configs"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Chat struct {
	ID          bson.ObjectId 	`bson:"_id" json:"id"`
	IDPatient   bson.ObjectId 	`bson:"idpatient" json:"idpatient"`
	IDDoctor   	bson.ObjectId 	`bson:"iddoctor" json:"iddoctor"`
	Active 		bool 			`bson:"active" json:"active"`
	Start 		string 			`bson:"start" json:"start"`
	End 		string 			`bson:"end" json:"end"`
}

type ChatDAO struct {
	Server   string
	Database string
}

var db *mgo.Database

const (
	COLLECTION = "chat"
)

func (m *ChatDAO) Connect() {
	c := config.Config{}
	c.Read()
	session, err := mgo.Dial(c.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(c.Database)
}

func (m *ChatDAO) GetAll() ([]Chat, error) {
	var Chats []Chat
	err := db.C(COLLECTION).Find(bson.M{}).All(&Chats)
	return Chats, err
}

func (m *ChatDAO) GetByID(id string) (Chat, error) {
	var Chat Chat
	err := db.C(COLLECTION).FindId(bson.ObjectIdHex(id)).One(&Chat)
	return Chat, err
}

func (m *ChatDAO) GetActiveByDoctor(idDoctor string) (Chat, error) {
	var Chat Chat
	err := db.C(COLLECTION).Find(bson.M{"iddoctor" : bson.ObjectIdHex(idDoctor), "active" : true}).One(&Chat)
	return Chat, err
}

func (m *ChatDAO) GetActiveByPatient(idPatient string) (Chat, error) {
	var Chat Chat
	err := db.C(COLLECTION).Find(bson.M{"idpatient" : bson.ObjectIdHex(idPatient), "active" : true}).One(&Chat)
	return Chat, err
}

func (m *ChatDAO) Create(Chat Chat) error {
	err := db.C(COLLECTION).Insert(&Chat)
	return err
}

func (m *ChatDAO) Delete(id string) error {
	err := db.C(COLLECTION).RemoveId(bson.ObjectIdHex(id))
	return err
}

func (m *ChatDAO) Update(id string, Chat Chat) error {
	err := db.C(COLLECTION).UpdateId(bson.ObjectIdHex(id), &Chat)
	return err
}
