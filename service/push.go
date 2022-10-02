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
	resp := &WechatMsgBody{
		ToUserName:   req.FromUserName,
		FromUserName: req.ToUserName,
		CreateTime:   req.CreateTime + 1,
		MsgType:      "text",
		Content:      "response to " + req.Content,
	}
	fmt.Printf("push resp:%+v\n", resp)
	msg, err := json.Marshal(resp)
	if err != nil {
		fmt.Printf("[Error] err:%v\n", err)
		return
	}
	w.Header().Set("content-type", "application/json")
	w.Write(msg)
}

func getRequest(r *http.Request) (*WechatMsgBody, error) {
	decoder := json.NewDecoder(r.Body)
	body := new(WechatMsgBody)
	if err := decoder.Decode(body); err != nil {
		fmt.Printf("[Error] err:%v\n", err)
		return nil, err
	}
	defer r.Body.Close()
	fmt.Printf("push req:%+v\n", body)
	return body, nil
}
