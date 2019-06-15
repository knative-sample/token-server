package main

import (
	"crypto/rand"
	"fmt"
	"time"
)

func Uuid() string {
	// generate 32 bits timestamp
	unix32bits := uint32(time.Now().UTC().UnixNano())

	buff := make([]byte, 12)

	numRead, err := rand.Read(buff)

	if numRead != len(buff) || err != nil {
		panic(err)
	}

	return fmt.Sprintf("%x-%x-%x-%x-%x-%x\n", unix32bits, buff[0:2], buff[2:4], buff[4:6], buff[6:8], buff[8:])
}
