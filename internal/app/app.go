package app

import (
	"botinfotime/internal/config"
	"botinfotime/internal/poolinformer/checknewdate"
	"botinfotime/internal/poolinformer/datetimeinfo"
	"fmt"
)

var (
	OldTime map[string][]string = make(map[string][]string)
)

func RunGetTimeNow(appConfig config.AppConfig) (string, map[string][]string, error) {
	dataInfo := datetimeinfo.InitDateTimeInfo(appConfig)
	dataList, err := dataInfo.GetDateFree()
	if err != nil {
		return "", nil, err
	}
	dataTime, err := dataInfo.GetTimeFree(dataList)
	if err != nil {
		return "", nil, err
	}
	var message string
	for data, time := range dataTime {
		message = fmt.Sprintf("%v\n%v", message, fmt.Sprintf("Для Стефании появилось свободное время время %v на дату %s", time, data))
	}
	if len(message) == 0 {
		return "К сожалении у Стефании нет свободного времени на запись", nil, nil
	}
	return message, nil, nil
}

func RunGetChangeTime(appConfig config.AppConfig) (string, map[string][]string, error) {
	dataInfo := datetimeinfo.InitDateTimeInfo(appConfig)
	dataList, err := dataInfo.GetDateFree()
	if err != nil {
		return "", nil, err
	}
	dataTime, err := dataInfo.GetTimeFree(dataList)
	if err != nil {
		return "", nil, err
	}
	changeTime := checknewdate.CheckNewDate(dataTime, OldTime)
	OldTime = dataTime
	var message string
	for data, time := range changeTime {
		message = fmt.Sprintf("%v\n%v", message, fmt.Sprintf("Для Стефании появилось свободное время время %v на дату %s", time, data))
	}
	return message, changeTime, nil
}
