package ceye

import (
	"encoding/json"
	"fmt"
	"net/http"
	"xxl-job/config"
	newclient "xxl-job/newClient"
)

type Ceye struct {
	Timeout  int
	Url      string
	Headers  map[string]string
	Proxystr string
	Result   Result
}

type Result struct {
	Meta Meta   `json:"meta"`
	Data []Data `json:"data"`
}

type Meta struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type Data struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Remote_addr string `json:"remote_addr"`
	Created_at  string `json:"created_at"`
}

func NewCeye(token, types, filter, proxy string) *Ceye {
	url := fmt.Sprintf("http://api.ceye.io/v1/records?token=%s&type=%s&filter=%s", token, types, filter)
	return &Ceye{
		Timeout: 20,
		Url:     url,
		Headers: map[string]string{
			"User-Agent": newclient.ReturnUA(),
		},
		Proxystr: proxy,
		Result:   Result{},
	}
}

func GetCeyeResult(proxyStr, filter string) ([]Data, error) {
	config, err := config.ReturnConfig()
	if err != nil {
		return []Data{}, err
	}
	ceye := NewCeye(config.Ceye.Token, config.Ceye.Types, filter, proxyStr)
	response, err := newclient.SendRequest("GET", ceye.Url, ceye.Proxystr, ceye.Timeout, nil, ceye.Headers)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	decoder := json.NewDecoder(response.Body)

	if response.StatusCode == http.StatusOK {
		if err := decoder.Decode(&ceye.Result); err != nil {
			return nil, err
		}
	}
	return ceye.Result.Data, nil
}
