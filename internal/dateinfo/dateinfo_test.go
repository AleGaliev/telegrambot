package dateinfo_test

import (
	"botinfotime/internal/dateinfo"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestTimeList(t *testing.T) {
	// Создаем тестовый сервер
	DateList := []string{
		"2025-03-25",
		"2021-04-05",
		"2021-04-06",
		"2021-04-07",
	}
	responseDate := dateinfo.Response{
		Data: []dateinfo.Data{
			{
				Attributes: dateinfo.Attributes{
					Time: "13:54",
				},
			},
		},
	}
	jsonStr, _ := json.Marshal(responseDate)
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Отправляем тестовый ответ
		w.WriteHeader(http.StatusOK)
		w.Write(jsonStr)
	}))
	defer ts.Close()

	// Заменяем URL на URL тестового сервера
	//oldURL := "https://b611880.yclients.com/api/v1/booking/search/dates"
	//fmt.Println(ts.URL)
	//ts.URL = oldURL

	timeList := dateinfo.TimeList(DateList)

	assert.Equal(t, map[string][]string{
		"2021-04-04": {"13:54"},
		"2021-04-05": {"13:54"},
		"2021-04-06": {"13:54"},
		"2021-04-07": {"13:54"},
	}, timeList)
	// Вызываем тестируемую функцию

}
