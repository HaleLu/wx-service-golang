package service

import (
	"github.com/Admingyu/go-workingday"
	"log"
	"testing"
	"time"
)

func TestGetStar(t *testing.T) {
	tests := []struct {
		name string
	}{
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			GetStar()
		})
	}
}

func TestWorking(t *testing.T) {
	dt := time.Now().AddDate(0, 0, -1)
	isWork, dayType := workingday.IsWorkDay(dt, "CN")

	log.Print("现在是：", dt)
	log.Print("今天需要上班？", isWork)
	log.Print("原因：", dayType)
}

func TestGetNextRestDay(t *testing.T) {
	now := time.Now()
	now = time.Now().AddDate(0, 0, -1)

}
