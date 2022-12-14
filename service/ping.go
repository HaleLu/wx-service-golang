package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type Text struct {
	Content string `json:"content"`
}

type CustomSendBody struct {
	ToUser  string `json:"touser"`
	MsgType string `json:"msgtype"`
	Text    Text   `json:"text"`
}

func PingHandler(w http.ResponseWriter, r *http.Request) {
	res := &JsonResult{}

	res.Data = "Pong! now:" + time.Now().String()
	msg, err := json.Marshal(res)
	if err != nil {
		fmt.Fprint(w, "内部错误")
		return
	}

	//body := &CustomSendBody{
	//	ToUser:  "o3dzp575sUNoB3_KyJj_Aq7ZOJKw",
	//	MsgType: "text",
	//	Text: Text{
	//		Content: time.Now().String(),
	//	},
	//}
	//SendMessage(body)
	w.Header().Set("content-type", "application/json")
	w.Write(msg)

	Push()
}

func SendMessage(body *CustomSendBody) {
	fmt.Printf("SendMessage body:%+v\n", body)
	jsonData, _ := json.Marshal(body)
	resp, err := http.Post("http://api.weixin.qq.com/cgi-bin/message/custom/send?from_appid=wx5ea4de9d9451abaf", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Printf("http.Post err:%v", err)
		return
	}
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("ioutil.ReadAll err:%v", err)
		return
	}
	fmt.Printf("SendMessage resp:%+v\n", string(bodyBytes))
}
