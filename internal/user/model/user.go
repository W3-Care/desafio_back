package models

import (
	"gopkg.in/mgo.v2/bson"
	"github.com/dgrijalva/jwt-go"
)

type User struct {
	ID          bson.ObjectId `bson:"_id" json:"id"`
	Name        string        `bson:"name" json:"name"`
	Login 		string        `bson:"login" json:"login"`
	Type 		string        `bson:"type" json:"type"`
	Hash 		string        `bson:"hash" json:"hash"`
	Active      bool          `bson:"active" json:"active"`
}

type UserAuth struct {
	ID 			string        `bson:"id" json:"id"`
	Login 		string        `bson:"login" json:"login"`
	Password 	string        `bson:"pass" json:"pass"`
}

type Token struct {
	UserId bson.ObjectId
	jwt.StandardClaims
	Token string
	Type string
	Name string
	ChatId bson.ObjectId
}
//TODO turn password transient