package main

import (
	"fmt"
	"time"
)

type TokenBucket struct {
	Rate   int
	Tokens chan string
}

func NewTokenBucket(rate int) *TokenBucket {
	tb := &TokenBucket{
		Rate:   rate,
		Tokens: make(chan string, rate),
	}

	go tb.start()

	return tb
}

func (tb *TokenBucket) start() {
	//ticker := time.NewTicker(time.Second)
	//for _ = range ticker.C {
	for {
		for i := 0; i < tb.Rate-len(tb.Tokens); i++ {
			logInfo(fmt.Sprintf("time:%s create token: %d", time.Now(), i))
			tb.Tokens <- Uuid()
		}
		time.Sleep(time.Second)
	}
}
