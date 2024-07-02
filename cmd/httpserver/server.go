package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"time"
)

type date struct {
	year  int
	month int
	day   int
}
type timeStruct struct {
	hour   float64
	minute float64
	second float64
}

// GetDateTime is the handler function for getting date and time at "/datetime"
func GetDateTime(w http.ResponseWriter, req *http.Request) {

	fmt.Fprintln(w, time.Now())
}

func getDate(w http.ResponseWriter, req *http.Request) {

	currDate := date{int(time.Now().Year()), int(time.Now().Month()), int(time.Now().Day())}
	fmt.Fprintf(w, "Date is %d/%d/%d\n", currDate.year, currDate.month, currDate.day)
}

func getTime(w http.ResponseWriter, req *http.Request) {

	currTime := timeStruct{float64(time.Now().Hour()), float64(time.Now().Minute()), float64(time.Now().Second())}
	fmt.Fprintf(w, "Time is %v:%v:%v\n", currTime.hour, currTime.minute, currTime.second)
}

func main() {

	http.HandleFunc("/datetime", GetDateTime)
	http.HandleFunc("/date", getDate)
	http.HandleFunc("/time", getTime)

	err := http.ListenAndServe(":8080", nil)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Println("SErver closed")
	} else if err != nil {
		fmt.Printf("Error starting the server: %s\n", err)
		os.Exit(1)
	}
}
