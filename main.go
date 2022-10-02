package main

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"log"
	"net/http"
	"wxcloudrun-golang/db"
	"wxcloudrun-golang/service"
)

func main() {
	if err := db.Init(); err != nil {
		panic(fmt.Sprintf("mysql init failed with %+v", err))
	}

	c := cron.New() //精确到秒

	//定时任务
	spec := "* */2 * * * ?" //cron表达式，每秒一次
	c.AddFunc(spec, service.Push)

	http.HandleFunc("/", service.IndexHandler)
	http.HandleFunc("/api/count", service.CounterHandler)
	http.HandleFunc("/api/ping", service.PingHandler)
	http.HandleFunc("/api/push", service.PushHandler)

	log.Fatal(http.ListenAndServe(":80", nil))
}
