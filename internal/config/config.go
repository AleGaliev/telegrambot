package config

import "os"

const (
	locationId int = 578233
	staffId    int = 2583227
	//staffId          int    = 3299357
	serviceId        int    = 8441801
	preficsTimeslots string = "timeslots"
	preficsDates     string = "dates"
	baseUrl          string = "https://b611880.yclients.com/api/v1/booking/search"
)

type AppConfig struct {
	LocationId       int
	StaffId          int
	ServiceId        int
	BaseUrl          string
	PreficsTimeslots string
	PreficsDates     string
	HeaderAuth       string
}

func NewAppConfig() AppConfig {
	return AppConfig{
		LocationId:       locationId,
		StaffId:          staffId,
		ServiceId:        serviceId,
		BaseUrl:          baseUrl,
		PreficsTimeslots: preficsTimeslots,
		PreficsDates:     preficsDates,
		//HeaderAuth:       HeaderAuth,
		HeaderAuth: os.Getenv("HEADER_AUTH"),
	}
}
