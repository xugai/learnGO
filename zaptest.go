package main

import (
	"go.uber.org/zap"
)
func main() {

	log, _ := zap.NewProduction()
	log.Warn("Warning Test")
}
