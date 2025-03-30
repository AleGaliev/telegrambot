package app

import (
	"botinfotime/internal/checknewdate"
	"botinfotime/internal/config"
	"botinfotime/internal/dateinfo"
	"botinfotime/internal/preparation"
	"botinfotime/internal/timefilter"
	"fmt"
)

func AppRun(oldTime map[string][]string) (map[string][]string, string) {
	changeTime := make(map[string][]string)
	var message string
	newTime := getInfo()
	changeTime = checknewdate.CheckNewDate(newTime, oldTime)
	for key, value := range changeTime {
		message = fmt.Sprintf("%v\n%v", message, fmt.Sprintf("Для Стефании появилось свободное время время %v на дату %s", value, key))
	}
	return newTime, message
}

func AppRunNow() (map[string][]string, string) {
	var message string
	newTime := getInfo()
	if len(newTime) == 0 {
		return newTime, "К сожалении у Стефании нет свободного времени на запись"
	}
	for key, value := range newTime {
		message = fmt.Sprintf("%v\n%v", message, fmt.Sprintf("Для Стефании есть свободное время время %v на дату %s", value, key))
	}
	return newTime, message
}

func getInfo() map[string][]string {
	appConf := config.NewAppConfig()
	configPayload := preparation.InitConfigPayload(appConf.HeaderAuth, appConf.BaseUrl)
	initTimeList := dateinfo.InitTimeList(configPayload, appConf)
	responseDateFree, _ := dateinfo.GetDateFree(configPayload, initTimeList)
	dateList := dateinfo.DateList(responseDateFree)
	newTime := dateinfo.TimeList(dateList, initTimeList)
	newTime, _ = timefilter.TimeFilter(newTime, "", "")
	return newTime
}
