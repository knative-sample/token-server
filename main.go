package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

var (
	tokenbucket *TokenBucket
)

func logInfo(msg string) {
	log.Printf("%s %s", time.Now().Format("2006-01-02 15:04:05"), msg)
}

func handler(w http.ResponseWriter, r *http.Request) {
	logInfo("Simple app running...")
	sleepTime := 0
	sleepTimeStr := r.URL.Query().Get("sleep")
	if sleepTimeStr != "" {
		if s, err := strconv.Atoi(sleepTimeStr); err == nil {
			sleepTime = s
		}
	}
	logInfo(fmt.Sprintf("sleepTime: %d", sleepTime))
	if sleepTime > 0 {
		time.Sleep(time.Duration(sleepTime) * time.Millisecond)
	}

	uuid := <-tokenbucket.Tokens

	msg := os.Getenv("MESSAGE")
	if msg == "" {
		msg = ":( MESSAGE EVN not defined"
	}
	logInfo(fmt.Sprintf("%s %s", uuid, msg))

	fmt.Fprintf(w, "uuid: %s", uuid)
}

func main() {
	flag.Parse()
	logInfo("Simple app server started...")

	rate := 10
	rateStr := os.Getenv("RATE")
	if rateStr != "" {
		if r, err := strconv.Atoi(rateStr); err == nil {
			rate = r
		}

	}
	tokenbucket = NewTokenBucket(rate)

	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
