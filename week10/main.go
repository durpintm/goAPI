package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
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

func TorontoTImeHandler(w http.ResponseWriter, r *http.Request){

torontoTime, err:=	getTorontoTime()

if err != nil{
	http.Error(w, "Error fetching Toronto time", http.StatusInternalServerError)
	return
}

// fmt.Fprintf(w, "Toronto time is %s", torontoTime)

resp := map[string]string{"Current_Time_Toronto": torontoTime}
w.Header().Set("Content-Type", "application/json")
json.NewEncoder(w).Encode(resp)
}


func main() {

	http.HandleFunc("/api/torontotime", TorontoTImeHandler)
	fmt.Println("Server started on port 8015!")
	log.Fatal(http.ListenAndServe(":8015", nil))

}

// docker build -t durpintm/torontotime:v01

// docker run -p 8015:8015 durpintm/torontotime:v01

// docker push durpintm/torontotime:v01