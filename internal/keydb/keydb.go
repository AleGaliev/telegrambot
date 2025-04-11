package keydb

import (
	"botinfotime/internal/loging"
	"context"
	"encoding/json"
	"fmt"
	"github.com/redis/go-redis/v9"
)

type SurveyList struct {
	Surveys []Survey
}

type Survey struct {
	ChatId    int64
	FirstName string
}

type KeyDb struct {
	ClientKeyDB *redis.Client
}

func (sl *SurveyList) AddByChatId(Survey Survey) {
	check := true
	for _, survey := range sl.Surveys {
		if survey.ChatId == Survey.ChatId {
			check = false
		}
	}
	if check {
		sl.Surveys = append(sl.Surveys, Survey)
	}
}

func (sl *SurveyList) RemoveByChatId(chatId int64) {
	for i, survey := range sl.Surveys {
		if survey.ChatId == chatId {
			sl.Surveys = append(sl.Surveys[:i], sl.Surveys[i+1:]...)
			return
		}
	}
}

func InitKeyDb(keyDBAdrr string) *KeyDb {
	return &KeyDb{
		ClientKeyDB: redis.NewClient(&redis.Options{
			Addr:     keyDBAdrr,
			Password: "",
			DB:       0,
		}),
	}
}

func (kdb KeyDb) PushValueKeyDb(ctx context.Context, key string, value SurveyList) error {

	jsonData, err := json.Marshal(value)
	if err != nil {
		return err
	}

	// Запись в Redis
	if err = kdb.ClientKeyDB.Set(ctx, key, jsonData, 0).Err(); err != nil {
		return err
	}
	loging.LogMessage(0, "admin", fmt.Sprintf("Значение запискано в БД в ключ %s", key))

	return nil
}

func (kdb KeyDb) GetValueKeyDb(ctx context.Context, key string) (SurveyList, error) {
	val, err := kdb.ClientKeyDB.Get(ctx, key).Bytes()
	if err == redis.Nil { // Специальная ошибка "ключ не найден" в Redis
		return SurveyList{}, nil // или можно вернуть SurveyList{} и nil, если это допустимо
	}
	if err != nil {
		return SurveyList{}, err
	}

	var resultChatId SurveyList

	if err = json.Unmarshal(val, &resultChatId); err != nil {
		return SurveyList{}, err
	}

	return resultChatId, nil
}
