package service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type OneResponse struct {
	Code     int       `json:"code"`
	Msg      string    `json:"msg"`
	NewsList []OneNews `json:"newslist"`
}

type OneNews struct {
	Oneid     int    `json:"oneid"`
	Word      string `json:"word"`
	Wordfrom  string `json:"wordfrom"`
	Imgurl    string `json:"imgurl"`
	Imgauthor string `json:"imgauthor"`
	Date      string `json:"date"`
}

func GetOne() *OneResponse {
	postValue := url.Values{"key": {"f7303c223b941757d22520d3a73fa52f"}}
	res, _ := http.PostForm("http://api.tianapi.com/one/index", postValue)
	defer res.Body.Close()

	decoder := json.NewDecoder(res.Body)
	body := new(OneResponse)
	if err := decoder.Decode(body); err != nil {
		fmt.Printf("[Error] err:%v\n", err)
		return nil
	}
	fmt.Printf("one resp:%+v\n", body)
	return body
}
