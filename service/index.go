package service

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// IndexHandler 计数器接口
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	data, err := getIndex()
	if err != nil {
		fmt.Fprint(w, "内部错误")
		return
	}
	fmt.Fprint(w, data)
}

// getIndex 获取主页
func getIndex() (string, error) {
	b, err := ioutil.ReadFile("./index.html")
	if err != nil {
		return "", err
	}
	return string(b), nil
}
