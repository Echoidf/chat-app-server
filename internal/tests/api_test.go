package tests

import (
	"chat-app/internal/api"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"testing"
)

func TestInsertOne(t *testing.T) {
	var user = &api.User{
		FullName:   "zuoql3",
		Username:   "zql3",
		Password:   "12345",
		Gender:     "12345",
		ProfilePic: "https://avatar.iran.liara.run/public/boy?username=zql",
	}
	err := api.InsertOneUser(user)
	if err != nil {
		t.Error(err.Error())
		return
	}
	fmt.Println(user.ID.Hex())
}

func TestGetUserList(t *testing.T) {
	userList, err := api.GetUserList()
	if err != nil {
		t.Error(err.Error())
		return
	}

	for _, user := range *userList {
		fmt.Println(user)
	}
}

func TestFindByUsername(t *testing.T) {
	user, err := api.FindByUsername("zql")
	if err != nil {
		t.Error(err.Error())
		return
	}

	fmt.Println(user.ID.String())
}

func TestInsertOneConversation(t *testing.T) {
	conv := &api.Conversation{
		Participants: [2]string{"65e5ab2b6c91dde1ec908dd5", "65e5ab95b36bf566db1699dd"},
	}
	err := api.InsertOneConversation(conv)
	if err != nil {
		t.Error(err.Error())
		return
	}
	fmt.Println("conversationId:" + conv.ID.Hex())
}

func TestUpdateConversation(t *testing.T) {
	id, _ := primitive.ObjectIDFromHex("65e5d8e72a819aa42a484b83")
	conv := &api.Conversation{
		ID:           id,
		Participants: [2]string{"65e5ab2b6c91dde1ec908dd5", "65e5ab95b36bf566db1699dd"},
		Messages:     []string{"65e5cfe8263d52476d2832e4"},
	}
	err := api.UpdateConversationMessages(conv)
	if err != nil {
		t.Error(err.Error())
		return
	}
}

func TestInsertOneMessage(t *testing.T) {
	msg := &api.Message{
		ReceiverId: "65e5ab2b6c91dde1ec908dd5",
		SenderId:   "65e5ab95b36bf566db1699dd",
		Message:    "hello",
	}
	err := api.InsertOneMessage(msg)
	if err != nil {
		t.Error(err.Error())
		return
	}

	fmt.Println("messageId:" + msg.ID.Hex())
}
