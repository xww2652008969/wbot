package utils

import (
	"os"
	"os/signal"
)

func Wait() {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, os.Kill)
	<-sig
}
