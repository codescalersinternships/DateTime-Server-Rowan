package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
	"slices"
	"strconv"
)

func Test_GetDateTime(t *testing.T) {
	request,_ := http.NewRequest(http.MethodGet, "/datetime", nil)
	response := httptest.NewRecorder()
	GetDateTime(response, request)
	t.Run("testing correct date", func (t *testing.T) {
		got := strings.Split(response.Body.String(), " ")[0]
		want := strings.Split(time.Now().String(), " ")[0]
		fmt.Println(got)
		if want != got {
			t.Errorf("got: %v, wanted: %v", got, want)
		}
	})

	t.Run("testing correct time", func (t* testing.T) {
		got := strings.Split(strings.Split(response.Body.String(), " ")[1], ":")
		want := strings.Split(strings.Split(time.Now().String(), " ")[1], ":")
		secondsGot, err := strconv.ParseFloat(got[2], 64)
		if err != nil {
			t.Error()
		}
		secondsWant, e := strconv.ParseFloat(want[2], 64)
		if e != nil {
			t.Error()
		}
		fmt.Println(got)
		if !slices.Equal(want[:2], got[:2])  || int(secondsGot) !=  int(secondsWant) {
			t.Errorf("got: %v, wanted: %v", got, want)
		}
	})
}