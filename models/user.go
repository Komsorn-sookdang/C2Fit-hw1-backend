package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID            primitive.ObjectID `bson:"_id"`
	Firstname     string             `bson:"Firstname"`
	Lastname      string             `bson:"Lastname"`
	Displayname   string             `bson:"Displayname"`
	Email         string             `bson:"Email"`
	Password      string             `bson:"-"`
	ConfirmPasswd string             `bson:"-"`
	HashPasswd    string             `bson:"HashPasswd" json:"-"`
	Phone         string             `bson:"Phone"`
	Address       string             `bson:"Address"`
}

type UserLogIn struct {
	Email    string `bson:"Email"`
	Password string `bson:"Password"`
}

func (user User) Validate() []string {
	var errors []string
	if user.Firstname == "" || user.Lastname == "" {
		errors = append(errors, "Firstname or Lastname must be filled")
	}
	if user.Email == "" {
		errors = append(errors, "Email must be filled")
	}
	if len(user.Password) < 8 {
		errors = append(errors, "Password must be longer than 8 characters")
	}
	if user.Password != user.ConfirmPasswd {
		errors = append(errors, "Password and ConfirmPasswd not match")
	}
	return errors
}
