package service

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func PushHandler(w http.ResponseWriter, r *http.Request) {
	req, err := getRequest(r)
	if err != nil {
		fmt.Fprint(w, "内部错误")
		return
	}
	fmt.Printf("push req:%+v\n", r)
	defer fmt.Printf("push resp:%+v\n", w)
	resp := &WechatMsgBody{
		ToUserName:   req.ToUserName,
		FromUserName: req.FromUserName,
		CreateTime:   req.CreateTime + 1,
		MsgType:      "text",
		Content:      "response to " + req.Content,
	}
	msg, err := json.Marshal(resp)
	if err != nil {
		fmt.Fprint(w, "内部错误")
		return
	}
	w.Header().Set("content-type", "application/json")
	w.Write(msg)
}

func getRequest(r *http.Request) (*WechatMsgBody, error) {
	decoder := json.NewDecoder(r.Body)
	body := new(WechatMsgBody)
	if err := decoder.Decode(body); err != nil {
		return nil, err
	}
	defer r.Body.Close()
	return body, nil
}
