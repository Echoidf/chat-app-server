package controllers

import (
	"chat-app/internal/api"
	"chat-app/pkg/response"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type sendMsgReq struct {
	Message string `json:"message"`
}

func SendMessage(c *gin.Context) {
	senderId, _ := c.Get("userId")
	receiverId := c.Param("id")

	var msg sendMsgReq
	if err := c.Bind(&msg); err != nil || msg.Message == "" {
		response.JSON(c, 400, "message can not be empty")
		return
	}

	// query conversation here
	var conversation *api.Conversation
	participants := [2]string{senderId.(string), receiverId}
	if conversation, _ = api.FindOneConversationByParticipants(participants); conversation == nil || conversation.ID == primitive.NilObjectID {
		// if conversation is nil, create one new conversation then
		conversation = &api.Conversation{
			Participants: participants,
		}
		if err := api.InsertOneConversation(conversation); err != nil {
			response.JSON(c, 500, err.Error())
			return
		}
	}

	// create one message
	message := &api.Message{
		ReceiverId: receiverId,
		SenderId:   senderId.(string),
		Message:    msg.Message,
	}
	if err := api.InsertOneMessage(message); err != nil {
		response.JSON(c, 500, err.Error())
		return
	}

	// update the related conversation
	conversation.Messages = append(conversation.Messages, message.ID.Hex())
	if err := api.UpdateConversationMessages(conversation); err != nil {
		response.JSON(c, 500, err.Error())
		return
	}

	response.JSON(c, 200, message)
}

func GetMessages(c *gin.Context) {
	userToChatId := c.Param("id")
	senderId, _ := c.Get("userId")

	var conversation *api.Conversation
	participants := [2]string{senderId.(string), userToChatId}
	if conversation, _ = api.FindOneConversationByParticipants(participants); conversation == nil {
		response.JSON(c, 404, "conversation not found")
		return
	}

	messages, err := api.GetMessagesById(conversation.Messages)
	if err != nil {
		response.Error(c, 500, err)
		return
	}

	response.JSON(c, 200, messages)
}
