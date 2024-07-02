package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func Test_GetDateTime(t *testing.T) {
	router := gin.Default() //gives empty router --> never test against an empty router
	request, _ := http.NewRequest(http.MethodGet, "/datetime", nil)
	response := httptest.NewRecorder()
	router.GET("/datetime", GetDateTime)
	router.ServeHTTP(response, request)
	fmt.Println(response)
	assert.Equal(t, 200, response.Code)
	assert.Equal(t, strings.Split(time.Now().String(), " ")[0], strings.Split(response.Body.String(), " ")[0])
	assert.Equal(t, strings.Split(strings.Split(response.Body.String(), " ")[1], ":")[:2], strings.Split(strings.Split(time.Now().String(), " ")[1], ":")[:2])
}
