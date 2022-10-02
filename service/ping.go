package service

import (
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
	w.Header().Set("content-type", "application/json")
	w.Write(msg)
}
