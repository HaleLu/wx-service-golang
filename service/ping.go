package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func PingHandler(w http.ResponseWriter, r *http.Request) {
	res := &JsonResult{}

	res.Data = "Pong! now:" + time.Now().String()
	msg, err := json.Marshal(res)
	if err != nil {
		fmt.Fprint(w, "内部错误")
		return
	}

	body := &WechatMsgBody{
		ToUserName:   "o3dzp575sUNoB3_KyJj_Aq7ZOJKw",
		FromUserName: "gh_f61f83ed88cf",
		CreateTime:   int(time.Now().Unix()),
		MsgType:      "text",
		Content:      time.Now().String(),
	}
	SendMessage(body)
	w.Header().Set("content-type", "application/json")
	w.Write(msg)
}

func SendMessage(body *WechatMsgBody) {
	fmt.Printf("SendMessage body:%+v", body)
	jsonData, _ := json.Marshal(body)
	resp, err := http.Post("http://api.weixin.qq.com/cgi-bin/message/custom/send", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Printf("http.Post err:%v", err)
		return
	}
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("ioutil.ReadAll err:%v", err)
		return
	}
	fmt.Printf("SendMessage resp:%+v", string(bodyBytes))
}