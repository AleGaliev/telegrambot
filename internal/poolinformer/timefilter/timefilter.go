package timefilter

import "time"

const (
	timeMinDefault string = "9:30"
	timeMaxDefault string = "20:40"
)

func genetateMinMAxTime(timeMin, timeMax string) (time.Time, time.Time, error) {
	if timeMin == "" {
		timeMin = timeMinDefault
	}
	if timeMax == "" {
		timeMax = timeMaxDefault
	}
	timeMinForamt, err := time.Parse("15:04", timeMin)
	if err != nil {
		return time.Time{}, time.Time{}, err
	}
	timeMaxForamt, err := time.Parse("15:04", timeMax)
	if err != nil {
		return time.Time{}, time.Time{}, err
	}
	return timeMinForamt, timeMaxForamt, nil
}

func TimeFilter(dateTime map[string][]string, timeMin, timeMax string) (map[string][]string, error) {
	filterDateTime := make(map[string][]string)
	timeMinForamt, timeMaxForamt, err := genetateMinMAxTime(timeMin, timeMax)
	if err != nil {
		return filterDateTime, err
	}

	for date, timeList := range dateTime {
		for _, t := range timeList {
			timeFormat, err := time.Parse("15:04", t)
			if err != nil {
				return filterDateTime, err
			}
			if timeMinForamt.Before(timeFormat) && timeMaxForamt.After(timeFormat) {
				filterDateTime[date] = append(filterDateTime[date], t)
			}
		}
	}
	return filterDateTime, nil
}
