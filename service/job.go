package service

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func Push() {
	url := "https://qyapi.weixin.qq.com/cgi-bin/message/send?access_token=z7ING515ZPc2jkrk6vI7_0wBDCuviAc6J9rEPgN0tQBOZ35YLi1VajhS7Wd-7fnc1IjfnP5Zg3asLxJ5uZo9U1SsA5EEjlPvoac1a_b6swB9m4vguuHV_7TPL-WBkBt_1SugpkGutf1vMDbpF9TmCsNCtRLaH9gd1tiom3arzOqpSvv91ERZ5MAmrPkI7tmDkTrLNdph2iMxytTHxOhbJA"
	method := "POST"

	payload := strings.NewReader(`{
	   "touser" : "HaiErYouZhiXingXingKouDai",
	   "agentid" : 1000002,
		"msgtype": "news",
        "news": {
	        "articles": [
	            {
	                "title": "中秋节礼品领取",
	                "description": "今年中秋节公司有豪礼相送",
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
