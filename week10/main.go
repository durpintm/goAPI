package main

import (
	"encoding/json"
	"io"
	"net/http"
)

const worldTimeAPI = "http://worldtimeapi.org/api/timezone/America/Toronto"

type TimeInfo struct {
	Datetime string `json:"datetime"`
}

func getTorontoTime() (string, error) {
	resp, err := http.Get(worldTimeAPI)
	if(err != nil){
		return "Error retriving date", err
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

if(err != nil){
	return "Error reading data", err
}

var timeInfo TimeInfo 
err = json.Unmarshal(body, &timeInfo)

if(err != nil){
	return "Error parsing data", err
}

	return timeInfo.Datetime, nil
}

func main() {
	println("Server started")
 data, err:=	getTorontoTime()

if err != nil{
	println(data)
}

 println(data)
}