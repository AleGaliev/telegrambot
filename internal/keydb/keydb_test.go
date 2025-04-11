package keydb

import (
	"context"
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	keyDateList   string   = "DateList"
	keyTimeList   string   = "TimeList"
	keyListChatId string   = "ListChat"
	dateList      []string = []string{
		"2025-03-25",
		"2021-04-05",
		"2021-04-06",
		"2021-04-07",
	}
	timeList SurveyList = SurveyList{}

	listChatId SurveyList = SurveyList{
		Surveys: []Survey{
			Survey{
				ChatId:    224268678,
				FirstName: "Aleksandr",
			},
		},
	}
	//listChatId telegrambot.SurveyList = map[int64]string{
	//	224268678: "Aleksandr",
	//}
	errout error = errors.New("redis: nil")
)

type caseTestData struct {
	dataList []string
	timeList map[string][]string
}

func TestPushValueKeyDb(t *testing.T) {
	KeyDbConfig := InitKeyDb("localhost:6379")
	if err := KeyDbConfig.PushValueKeyDb(context.Background(), keyListChatId, listChatId); err != nil {
		t.Error(err)
	}

}

func TestGetValueKeyDb(t *testing.T) {
	KeyDbConfig := InitKeyDb("localhost:6379")
	listChatIdOut, err := KeyDbConfig.GetValueKeyDb(context.Background(), keyListChatId)

	if err != nil {
		t.Fatal(err)
	}
	timeListOut, err := KeyDbConfig.GetValueKeyDb(context.Background(), keyTimeList)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(listChatIdOut)
	assert.Equal(t, listChatId, listChatIdOut)
	assert.Equal(t, timeList, timeListOut)

}
