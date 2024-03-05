package api

import (
	"chat-app/internal/db/mongodb"
	"context"
	"github.com/quangdangfit/gocommon/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
)

type User struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	FullName   string
	Username   string
	Password   string
	Gender     string
	ProfilePic string
}

var ctx = context.Background()

func InsertOneUser(user *User) error {
	if user == nil {
		return errors.New("user provided must not be nil")
	}
	res, err := mongodb.DB.Collection("users").InsertOne(ctx, user)
	user.ID = res.InsertedID.(primitive.ObjectID)
	return err
}

func GetUserList() (*[]User, error) {
	cursor, err := mongodb.DB.Collection("users").Find(ctx, bson.M{})
	if err != nil {
		log.Fatal("query users failed,err: ", err.Error())
		return nil, err
	}

	var res []User
	err = cursor.All(ctx, &res)
	return &res, err
}

func FindByUsername(username string) (*User, error) {
	var user User
	err := mongodb.DB.Collection("users").FindOne(ctx, bson.M{"username": username}).Decode(&user)

	return &user, err
}
