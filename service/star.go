package service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type StarResponse struct {
	Code     int        `json:"code"`
	Msg      string     `json:"msg"`
	NewsList []StarNews `json:"newslist"`
}

type StarNews struct {
	Type    string `json:"type"`
	Content string `json:"content"`
}

func GetStar() *StarResponse {
	postValue := url.Values{"key": {"f7303c223b941757d22520d3a73fa52f"}, "astro": {"双子座"}}
	res, _ := http.PostForm("http://api.tianapi.com/star/index", postValue)
	defer res.Body.Close()

	decoder := json.NewDecoder(res.Body)
	body := new(StarResponse)
	if err := decoder.Decode(body); err != nil {
		fmt.Printf("[Error] err:%v\n", err)
		return nil
	}
	fmt.Printf("star resp:%+v\n", body)
	return body
}
