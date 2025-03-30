package loging

import (
	"encoding/json"
	"fmt"
)

type Logger struct {
	ChatId   int64  `json:"chatId"`
	UserName string `json:"userName"`
	Message  string `json:"message"`
}

func LogMessage(chatId int64, userName, message string) {
	messageJson, _ := json.Marshal(&Logger{ChatId: chatId, UserName: userName, Message: message})
	fmt.Println(string(messageJson))
}
