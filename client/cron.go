package client

import (
	"fmt"
	"time"
)

func HeartBeat() {
	fmt.Println("start heartbeat")
	c := time.NewTicker(40 * time.Second)

	for _ = range c.C {
		fmt.Println("time:", time.Now().Unix())
		Report()
	}
}
