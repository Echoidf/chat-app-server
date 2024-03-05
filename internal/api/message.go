package api

import (
	"chat-app/internal/db/mongodb"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type Message struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	ReceiverId string
	SenderId   string
	Message    string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func InsertOneMessage(message *Message) error {
	if message == nil {
		return errors.New("message provided must not be nil")
	}
	message.CreatedAt = time.Now()
	message.UpdatedAt = time.Now()
	res, err := mongodb.DB.Collection("messages").InsertOne(ctx, message)
	message.ID = res.InsertedID.(primitive.ObjectID)
	return err
}

func GetMessagesById(msgIds []string) (*[]Message, error) {
	var ids []primitive.ObjectID
	for _, id := range msgIds {
		objId, _ := primitive.ObjectIDFromHex(id)
		ids = append(ids, objId)
	}
	var messages []Message
	filter := bson.D{}
	filter = append(filter, bson.E{Key: "_id", Value: bson.M{"$in": ids}})
	cursor, err := mongodb.DB.Collection("messages").Find(ctx, filter)
	if errors.Is(err, mongo.ErrNoDocuments) {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	if err = cursor.All(ctx, &messages); err != nil {
		return nil, err
	}
	return &messages, err
}
