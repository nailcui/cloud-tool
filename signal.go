package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func processSignal() {
	c := make(chan os.Signal)
	// 监听所有信号
	signal.Notify(c, syscall.SIGTERM)
	for s := range c {
		fmt.Printf("receive signal: %s\n", s)
	}
}
