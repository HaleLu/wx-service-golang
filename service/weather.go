package service

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type DailyWeather struct {
	FxDate         string `json:"fxDate"`
	Sunrise        string `json:"sunrise"`
	Sunset         string `json:"sunset"`
	Moonrise       string `json:"moonrise"`
	Moonset        string `json:"moonset"`
	MoonPhase      string `json:"moonPhase"`
	MoonPhaseIcon  string `json:"moonPhaseIcon"`
	TempMax        string `json:"tempMax"`
	TempMin        string `json:"tempMin"`
	IconDay        string `json:"iconDay"`
	TextDay        string `json:"textDay"`
	IconNight      string `json:"iconNight"`
	TextNight      string `json:"textNight"`
	Wind360Day     string `json:"wind360Day"`
	WindDirDay     string `json:"windDirDay"`
	WindScaleDay   string `json:"windScaleDay"`
	WindSpeedDay   string `json:"windSpeedDay"`
	Wind360Night   string `json:"wind360Night"`
	WindDirNight   string `json:"windDirNight"`
	WindScaleNight string `json:"windScaleNight"`
	WindSpeedNight string `json:"windSpeedNight"`
	Humidity       string `json:"humidity"`
	Precip         string `json:"precip"`
	Pressure       string `json:"pressure"`
	Vis            string `json:"vis"`
	Cloud          string `json:"cloud"`
	UvIndex        string `json:"uvIndex"`
}

type WeatherResponse struct {
	Code       string         `json:"code"`
	UpdateTime string         `json:"updateTime"`
	FxLink     string         `json:"fxLink"`
	Daily      []DailyWeather `json:"daily"`
}

func GetWeather() *WeatherResponse {
	url := "https://devapi.qweather.com/v7/weather/3d?key=f74a6a1bbfb94191b8bfbb0ba5900e0c&location=101021500"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return nil
	}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer res.Body.Close()

	decoder := json.NewDecoder(res.Body)
	body := new(WeatherResponse)
	if err := decoder.Decode(body); err != nil {
		fmt.Printf("[Error] err:%v\n", err)
		return nil
	}
	fmt.Printf("weather resp:%+v\n", body)
	return body
}
