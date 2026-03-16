package main

import "go.uber.org/zap"

func main() {
	l, _ := NewLogger("log.json")
	l.Info("Hello", zap.String("name", "Arkady"))
	l.Info("Hello", zap.String("name", "rkady"))

}
