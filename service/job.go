package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func GetAccessToken() (string, error) {
	url := "https://qyapi.weixin.qq.com/cgi-bin/gettoken?corpid=ww06d5d98feda07249&corpsecret=PUE3HYm_PHgGuvUH0lkut57vUzuhcvhmrzSuRyq9IUE"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return "", err
	}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	defer res.Body.Close()

	decoder := json.NewDecoder(res.Body)
	body := make(map[string]interface{})
	if err := decoder.Decode(&body); err != nil {
		fmt.Printf("[Error] err:%v\n", err)
		return "", err
	}
	fmt.Printf("get token resp:%+v\n", body)
	return body["access_token"].(string), nil
}

func Push() {
	token, err := GetAccessToken()
	if err != nil {
		token, err = GetAccessToken()
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	url := "https://qyapi.weixin.qq.com/cgi-bin/message/send?access_token=" + token
	method := "POST"

	payload := strings.NewReader(`{
	   "touser" : "HaiErYouZhiXingXingKouDai",
	   "agentid" : 1000002,
		"msgtype": "news",
        "news": {
	        "articles": [
	            {
	                "title": "每日提醒",
	                "description": "文案",
	                "picurl": "http://res.mail.qq.com/node/ww/wwopenmng/images/independent/doc/test_pic_msg1.png"
	            }
	        ]
	    },
	   "safe":0,
	   "enable_id_trans": 0,
	   "enable_duplicate_check": 0,
	   "duplicate_check_interval": 1800
	}`)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}
