package main

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
	"time"
)

func Test_GetDateTime(t *testing.T) {
	for i := 0; i < 1000; i++ {
		request, _ := http.NewRequest(http.MethodGet, "/datetime", nil)
		response := httptest.NewRecorder()
		GetDateTime(response, request)
		t.Run("testing successs of response", func(t *testing.T) {
			assert.Equal(t, 200, response.Code)
		})
		t.Run("testing correct date", func(t *testing.T) {
			got := strings.Split(response.Body.String(), " ")[0]
			want := strings.Split(time.Now().String(), " ")[0]
			assert.Equal(t, want, got)
		})

		t.Run("testing correct time", func(t *testing.T) {
			got := strings.Split(strings.Split(response.Body.String(), " ")[1], ":")
			want := strings.Split(strings.Split(time.Now().String(), " ")[1], ":")
			assert.Equal(t, want[:2], got[:2])

			secondsGot, err := strconv.ParseFloat(got[2], 64)
			if err != nil {
				t.Error()
			}
			secondsWant, e := strconv.ParseFloat(want[2], 64)
			if e != nil {
				t.Error()
			}
			if int(secondsWant) != int(secondsGot) && int(secondsWant) != int(secondsGot)+1 {
				t.Error()
			}
		})
	}

}
