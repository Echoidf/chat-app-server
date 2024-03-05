package api

import (
	"chat-app/internal/db/mongodb"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"time"
)

type Conversation struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	Participants [2]string
	Messages     []string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func InsertOneConversation(con *Conversation) error {
	if con == nil {
		return errors.New("conversation provided must not be nil")
	}

	con.CreatedAt = time.Now()
	con.UpdatedAt = time.Now()
	res, err := mongodb.DB.Collection("conversation").InsertOne(ctx, con)
	con.ID = res.InsertedID.(primitive.ObjectID)
	return err
}

func GetConversationList() (*[]Conversation, error) {
	cursor, err := mongodb.DB.Collection("conversation").Find(ctx, bson.M{})
	if err != nil {
		log.Fatal("query users failed,err: ", err.Error())
		return nil, err
	}

	var res []Conversation
	err = cursor.All(ctx, &res)
	return &res, err
}

func FindOneConversationByParticipants(participants [2]string) (*Conversation, error) {
	var res Conversation
	err := mongodb.DB.Collection("conversation").FindOne(ctx, bson.M{"participants": bson.M{
		"$all": participants,
	}}).Decode(&res)

	return &res, err
}

func UpdateConversationMessages(con *Conversation) error {
	if con == nil {
		return errors.New("conversation provided must not be nil")
	}

	update := bson.D{{Key: "$set", Value: bson.D{{Key: "messages", Value: con.Messages}, {Key: "updatedAt", Value: time.Now()}}}}
	_, err := mongodb.DB.Collection("conversation").UpdateByID(ctx, con.ID, update)

	return err
}
