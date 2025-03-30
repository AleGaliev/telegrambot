package dateinfo

import (
	"botinfotime/internal/preparation"
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/url"
	"path"
	"time"
)

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

func GetDateFree(configPostPayload preparation.ConfigPostPayload, info TimeListInfo) (Response, error) {
	urlGetDataFree, _ := url.JoinPath(info.BaseUrlAddr, info.PrefixDate)
	configPostPayload.Req.URL, _ = url.Parse(urlGetDataFree)
	jsonStr, err := json.Marshal(initPayloadDateNow(info.LocationId, info.StaffId, info.TypeId))
	if err != nil {
		return Response{}, err
	}
	configPostPayload.Req.Body = io.NopCloser(bytes.NewReader(jsonStr))
	resp, err := configPostPayload.Client.Do(configPostPayload.Req)
	if err != nil {
		return Response{}, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return Response{}, err
	}
	var out Response
	err = json.Unmarshal(body, &out)
	if err != nil {
		return Response{}, err
	}
	return out, nil
}

func GetTimeFree(configPostPayload preparation.ConfigPostPayload, info TimeListInfo, date string) (Response, error) {
	configPostPayload.Req.URL.Path = path.Join(info.BaseUrlAddr, info.PrefixDate)
	urlGetTimeFree, _ := url.JoinPath(info.BaseUrlAddr, info.PrefixTime)
	configPostPayload.Req.URL, _ = url.Parse(urlGetTimeFree)
	jsonStr, err := json.Marshal(initPayloadTime(info.LocationId, info.StaffId, info.TypeId, date))
	if err != nil {
		return Response{}, err
	}
	configPostPayload.Req.Body = io.NopCloser(bytes.NewReader(jsonStr))
	resp, err := configPostPayload.Client.Do(configPostPayload.Req)
	if err != nil {
		return Response{}, err
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	var out Response
	err = json.Unmarshal(body, &out)
	if err != nil {
		return Response{}, err
	}
	return out, nil
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
