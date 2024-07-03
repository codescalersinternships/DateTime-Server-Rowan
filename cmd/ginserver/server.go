package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"time"
	"errors"
)

//	type date struct {
//		year  int
//		month int
//		day   int
//	}
//
//	type timeStruct struct {
//		hour   float64
//		minute float64
//		second float64
//	}

// GetDateTime is the handler function for getting date and time at "/datetime"
func GetDateTime(c *gin.Context) {
	c.String(http.StatusOK, fmt.Sprint(time.Now()))
}

// func GetDate(c *gin.Context) {
// 	currDate := date{int(time.Now().Year()), int(time.Now().Month()), int(time.Now().Day())}
// 	c.IndentedJSON(http.StatusOK, currDate)
// }

// func GetTime(c *gin.Context) {
// 	currTime := timeStruct{float64(time.Now().Hour()), float64(time.Now().Minute()), float64(time.Now().Second())}
// 	c.IndentedJSON(http.StatusOK, currTime)
// }

func main() {
	r := gin.Default()
	r.GET("/datetime", GetDateTime)
	// r.GET("/date", GetDate)
	// r.GET("/time", GetTime)
	err := r.Run(":8089")
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Println("SErver closed")
	} else if err != nil {
		fmt.Printf("Error starting the server: %s\n", err)
		os.Exit(1)
	}
}
