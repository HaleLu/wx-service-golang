package service

import (
	"encoding/json"
	"fmt"
	"github.com/Admingyu/go-workingday"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
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

var WeekDayMap = map[string]string{
	"Monday":    "星期一",
	"Tuesday":   "星期二",
	"Wednesday": "星期三",
	"Thursday":  "星期四",
	"Friday":    "星期五",
	"Saturday":  "星期六",
	"Sunday":    "星期日",
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

	weather := GetWeather()
	if weather == nil {
		fmt.Println("weather == nil")
		return
	}
	daily := weather.Daily[0]

	star := GetStar()
	if star == nil {
		fmt.Println("star == nil")
		return
	}
	var newsMap = make(map[string]string)
	for _, n := range star.NewsList {
		newsMap[n.Type] = n.Content
	}

	one := GetOne()
	if one == nil {
		fmt.Println("one == nil")
		return
	}

	url := "https://qyapi.weixin.qq.com/cgi-bin/message/send?access_token=" + token
	method := "POST"

	fallInLoveDay := time.Date(2021, 7, 18, 0, 0, 0, 0, time.Local)

	now := time.Now()

	toNextRestDay := GetNextRestDay(now)
	toNewYearDay := int(time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local).Sub(now).Hours())/24 + 1
	toSpringFestivalDay := int(time.Date(2023, 1, 22, 0, 0, 0, 0, time.Local).Sub(now).Hours())/24 + 1
	var dayCountDesc string
	if toNextRestDay == 0 {
		dayCountDesc = fmt.Sprintf("今天不工作！距离元旦还有%d天，距离春节还有%d天", toNewYearDay, toSpringFestivalDay)
	} else {
		dayCountDesc = fmt.Sprintf("工作日要加油哦！距离周末还有%d天，距离元旦还有%d天，距离春节还有%d天", toNextRestDay, toNewYearDay, toSpringFestivalDay)
	}

	desc := "今天是" + now.Format("2006-01-02") + " " + WeekDayMap[now.Weekday().String()] + "\n" +
		"今日气温：" + daily.TempMin + "至" + daily.TempMax + "℃\n" +
		"白天天气：" + daily.TextDay + "\n" +
		"晚间天气：" + daily.TextNight + "\n" +
		"穿衣推荐：" + GetClothes(convInt64Default(daily.TempMin, 0), convInt64Default(daily.TempMax, 0)) + "\n" +
		"带伞提醒：" + GetUmbrella(daily.TextDay, daily.TextNight) + "\n" +
		"\n" +
		"今天是我们在一起的第" + strconv.FormatInt(int64(now.Sub(fallInLoveDay).Hours()/24), 10) + "天，也是我的宝贝最可爱的一天~" + "\n" +
		dayCountDesc + "\n" +
		"\n" +
		//"双子今日运势：\n" +
		//"综合：" + newsMap["综合指数"] + "\n" +
		//"爱情：" + newsMap["爱情指数"] + "\n" +
		//"工作：" + newsMap["工作指数"] + "\n" +
		//"财运：" + newsMap["财运指数"] + "\n" +
		//"健康：" + newsMap["健康指数"] + "\n" +
		//"\n" +
		one.NewsList[0].Word

	payload := strings.NewReader(`{
	   "touser" : "HaiErYouZhiXingXingKouDai|CaoCao",
	   "agentid" : 1000002,
		"msgtype": "news",
        "news": {
	        "articles": [
	            {
	                "title": "我亲爱的小充电宝，早上好(*´▽｀)ノノ",
	                "description": "` + desc + `",
	                "picurl": "` + one.NewsList[0].Imgurl + `"
	            }
	        ]
	    },
	   "safe":0,
	   "enable_id_trans": 0,
	   "enable_duplicate_check": 1,
	   "duplicate_check_interval": 10
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

func GetNextRestDay(now time.Time) int {
	for i := 0; i < 10; i++ {
		dt := now.AddDate(0, 0, i)
		isWork, _ := workingday.IsWorkDay(dt, "CN")
		if !isWork {
			return i
		}
	}
	return 0
}

func convInt64Default(s string, d int64) int64 {
	res, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return d
	}
	return res
}
