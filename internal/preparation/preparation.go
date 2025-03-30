package preparation

import (
	"net/http"
)

type ConfigPostPayload struct {
	Req    *http.Request
	Client *http.Client
}

func InitConfigPayload(auth, baseUrl string) ConfigPostPayload {
	client := &http.Client{}
	var req, err = http.NewRequest(http.MethodPost, baseUrl, nil)
	if err != nil {
		return ConfigPostPayload{}
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", auth)
	//req.URL.Path = path.Join(req.URL.Path, "config")
	//fmt.Println(req.URL)
	return ConfigPostPayload{req, client}
}
