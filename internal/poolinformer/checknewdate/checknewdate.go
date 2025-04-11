package checknewdate

func CheckNewDate(newDateTimeList, oldDateTimeList map[string][]string) map[string][]string {
	changeDataTime := make(map[string][]string)
	for newDate, newTimes := range newDateTimeList {
		if _, exists := oldDateTimeList[newDate]; exists {
			for _, newTime := range newTimes {
				check := false
				for _, oldTime := range oldDateTimeList[newDate] {
					if newTime == oldTime {
						check = true
					}
				}
				if !check {
					changeDataTime[newDate] = append(changeDataTime[newDate], newTime)
				}
			}
		} else {
			changeDataTime[newDate] = newDateTimeList[newDate]
			continue
		}
	}
	return changeDataTime
}
