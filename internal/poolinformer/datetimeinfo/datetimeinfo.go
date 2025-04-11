package datetimeinfo

import (
	"botinfotime/internal/config"
	"botinfotime/internal/loging"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

var (
	preficsTimeslots string = "timeslots"
	preficsDates     string = "dates"
	//baseUrl          string = "https://b611880.yclients.com/api/v1/booking/search"
	baseUrl url.URL = url.URL{
		Scheme: "https",
		Host:   "b611880.yclients.com",
		Path:   "/api/v1/booking/search",
	}
)

type DateTimeInfo struct {
	PayloadTime     Payload
	PayloadTimeList []Payload
	PayloadDateNow  Payload
	Request         http.Request
	HttpClient      *http.Client
}

type Payload struct {
	Contect Context `json:"context"`
	Filter  Filter  `json:"filter"`
}

type Context struct {
	LocationId int `json:"location_id"`
}

type Filter struct {
	Date     string    `json:"date,omitempty"`
	DateFrom string    `json:"date_from,omitempty"`
	DateTo   string    `json:"date_to,omitempty"`
	Records  []Records `json:"records"`
}

type Records struct {
	StaffId                int                      `json:"staff_id"`
	AttendanceServiceItems []AttendanceServiceItems `json:"attendance_service_items"`
}

type AttendanceServiceItems struct {
	Type string `json:"type"`
	Id   int    `json:"id"`
}

type Response struct {
	Data []Data `json:"data"`
}

type Data struct {
	Attributes Attributes `json:"attributes"`
}

type Attributes struct {
	Time       string `json:"time,omitempty"`
	Date       string `json:"date,omitempty"`
	IsBookable bool   `json:"is_bookable,omitempty"`
}

func InitDateTimeInfo(AppConfig config.AppConfig) DateTimeInfo {
	req, err := http.NewRequest(http.MethodPost, baseUrl.String(), nil)
	if err != nil {
		return DateTimeInfo{}
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", AppConfig.HeaderAuth)
	PayloadDateNow := initPayloadDateNow(AppConfig.LocationId, AppConfig.StaffId, AppConfig.ServiceId)
	PayloadTime := initPayloadTime(AppConfig.LocationId, AppConfig.StaffId, AppConfig.ServiceId, "")
	return DateTimeInfo{
		Request:         *req,
		HttpClient:      &http.Client{},
		PayloadDateNow:  PayloadDateNow,
		PayloadTime:     PayloadTime,
		PayloadTimeList: []Payload{},
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

func initPayloadTime(locationId, staffId, typeId int, date string) Payload {
	payload := Payload{
		Contect: Context{
			LocationId: locationId,
		},
		Filter: Filter{
			Date: date,
			Records: []Records{
				{
					StaffId: staffId,
					AttendanceServiceItems: []AttendanceServiceItems{
						{
							Type: "service",
							Id:   typeId,
						},
					},
				},
			},
		},
	}
	return payload
}

func initPayloadDateNow(locationId, staffId, typeId int) Payload {
	nowDate := time.Now().Format("2006-01-02")
	payload := Payload{
		Contect: Context{
			LocationId: locationId,
		},
		Filter: Filter{
			DateFrom: nowDate,
			DateTo:   "9999-01-01",
			Records: []Records{
				{
					StaffId: staffId,
					AttendanceServiceItems: []AttendanceServiceItems{
						{
							Type: "service",
							Id:   typeId,
						},
					},
				},
			},
		},
	}
	return payload
}

//func (app config.ConfigPostPayload) configDataFree(urlAddr string, payload Payload) {
//	jsonStr, _ := json.Marshal(payload)
//	app.Req.URL = app.Req.URL.Parse(urlAddr)
//	app.Req.Body = bytes.NewBuffer(jsonStr)
//}

func (dti DateTimeInfo) GetDateFree() ([]string, error) {
	pathGetGateFree, err := url.JoinPath(baseUrl.Path, preficsDates)
	if err != nil {
		return nil, err
	}

	newReq := dti.Request.Clone(context.Background())

	newReq.URL.Path = pathGetGateFree

	jsonStr, err := json.Marshal(dti.PayloadDateNow)
	if err != nil {
		return nil, err
	}
	newReq.Body = io.NopCloser(bytes.NewReader(jsonStr))

	resp, err := dti.HttpClient.Do(newReq)
	if err != nil {
		return nil, fmt.Errorf("Не смог сделать http запрос %ы", err)
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("get date free: status code %d", resp.StatusCode)
	} else {
		loging.LogMessage(0000000000, "admin", "Информация о свободых датах получена")
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var out Response
	err = json.Unmarshal(body, &out)
	if err != nil {
		return nil, err
	}

	dateList := DateList(out)

	return dateList, nil
}

func (dti DateTimeInfo) GetTimeFree(dates []string) (map[string][]string, error) {
	pathGetTimeFree, err := url.JoinPath(baseUrl.Path, preficsTimeslots)
	if err != nil {
		return nil, err
	}

	newReq := dti.Request.Clone(context.Background())
	newReq.URL.Path = pathGetTimeFree

	timeList := make(map[string][]string)
	for _, date := range dates {
		newJson := dti.PayloadTime
		newJson.Filter.Date = date
		jsonStr, err := json.Marshal(newJson)
		if err != nil {
			return nil, err
		}

		newReq.Body = io.NopCloser(bytes.NewReader(jsonStr))

		resp, err := dti.HttpClient.Do(newReq)
		if err != nil {
			return nil, err
		}
		if resp.StatusCode != http.StatusOK {
			return nil, fmt.Errorf("get date free: status code %d", resp.StatusCode)
		} else {
			loging.LogMessage(0000000000, "admin", "Информация о свободном времени получена")
		}
		defer resp.Body.Close()
		body, _ := ioutil.ReadAll(resp.Body)
		var out Response
		err = json.Unmarshal(body, &out)
		if err != nil {
			return nil, err
		}
		for _, time := range out.Data {
			timeList[date] = append(timeList[date], time.Attributes.Time)
		}
	}

	return timeList, nil
}
