package service

import (
	"bytes"
	"encoding/json"
	"fmt"
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
	jsonData, _ := json.Marshal(body)
	_, err = http.Post("http://api.weixin.qq.com/cgi-bin/message/custom/send", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Printf("http.Post err:%v", err)
		return
	}
	w.Header().Set("content-type", "application/json")
	w.Write(msg)
}
