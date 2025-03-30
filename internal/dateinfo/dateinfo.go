package dateinfo

import (
	"botinfotime/internal/config"
	"botinfotime/internal/preparation"
)

type TimeListInfo struct {
	СonfigPostPayload preparation.ConfigPostPayload
	BaseUrlAddr       string
	PrefixDate        string
	PrefixTime        string
	LocationId        int
	StaffId           int
	TypeId            int
}

func InitTimeList(payload preparation.ConfigPostPayload, config config.AppConfig) TimeListInfo {
	return TimeListInfo{
		payload,
		config.BaseUrl,
		config.PreficsDates,
		config.PreficsTimeslots,
		config.LocationId,
		config.StaffId,
		config.ServiceId,
	}
}

func DateList(response Response) []string {
	var dateList []string
	for _, value := range response.Data {
		if value.Attributes.IsBookable == true {
			dateList = append(dateList, value.Attributes.Date)
		}
	}
	return dateList
}

func TimeList(DateList []string, timeListInfo TimeListInfo) map[string][]string {
	timeList := make(map[string][]string)
	for _, date := range DateList {

		responseTime, err := GetTimeFree(timeListInfo.СonfigPostPayload, timeListInfo, date)
		if err != nil {
			return make(map[string][]string)
		}
		for _, value := range responseTime.Data {
			timeList[date] = append(timeList[date], value.Attributes.Time)

		}
	}
	return timeList
}
