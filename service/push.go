package service

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func PushHandler(w http.ResponseWriter, r *http.Request) {
	res := &JsonResult{}
	content, err := getContent(r)
	if err != nil {
		res.Code = -1
		res.ErrorMsg = err.Error()
	}
	res.Data = content
	msg, err := json.Marshal(res)
	if err != nil {
		fmt.Fprint(w, "内部错误")
		return
	}
	w.Header().Set("content-type", "application/json")
	w.Write(msg)
}

func getContent(r *http.Request) (string, error) {
	decoder := json.NewDecoder(r.Body)
	body := new(WechatMsgRequest)
	if err := decoder.Decode(body); err != nil {
		return "", err
	}
	defer r.Body.Close()
	return body.Content, nil
}
