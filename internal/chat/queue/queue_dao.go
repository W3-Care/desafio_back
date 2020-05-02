package queuedao

import (
	"log"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	config "we.care/they-chat/configs"
)

type ChatQueue struct {
	ID          bson.ObjectId `bson:"_id" json:"id"`
	IDPatient   bson.ObjectId `bson:"idpatient" json:"idpatient"`
}

type ChatQueueDAO struct {
	Server   string
	Database string
}

var db *mgo.Database

const (
	COLLECTION = "chat-queue"
)

func (m *ChatQueueDAO) Connect() {
	c := config.Config{}
	c.Read()
	session, err := mgo.Dial(c.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(c.Database)
}

func (m *ChatQueueDAO) GetAll() ([]ChatQueue, error) {
	var ChatQueues []ChatQueue
	err := db.C(COLLECTION).Find(bson.M{}).All(&ChatQueues)
	return ChatQueues, err
}

func (m *ChatQueueDAO) GetByID(id string) (ChatQueue, error) {
	var ChatQueue ChatQueue
	err := db.C(COLLECTION).FindId(bson.ObjectIdHex(id)).One(&ChatQueue)
	return ChatQueue, err
}

func (m *ChatQueueDAO) GetByIDPatient(idPatient bson.ObjectId) (ChatQueue, error)  {
	var ChatQueue ChatQueue
	err := db.C(COLLECTION).Find(bson.M{"idpatient" : idPatient}).One(&ChatQueue)
	return ChatQueue, err
}

func (m *ChatQueueDAO) GetFirst() (ChatQueue, error) {
	var ChatQueue ChatQueue
	err := db.C(COLLECTION).Find(bson.M{}).One(&ChatQueue)
	return ChatQueue, err
}

func (m *ChatQueueDAO) Create(ChatQueue ChatQueue) error {
	err := db.C(COLLECTION).Insert(&ChatQueue)
	return err
}

func (m *ChatQueueDAO) Delete(id string) error {
	err := db.C(COLLECTION).RemoveId(bson.ObjectIdHex(id))
	return err
}

func (m *ChatQueueDAO) Update(id string, ChatQueue ChatQueue) error {
	err := db.C(COLLECTION).UpdateId(bson.ObjectIdHex(id), &ChatQueue)
	return err
}
