package config

import (
	"os"
	"fmt"
	"time"
)

type service struct {
	Port				string
}

type timeouts struct {
	Read       time.Duration
	Write      time.Duration
	ReadHeader time.Duration
	Idle       time.Duration
}

// Service - entries for all config values
var Service service

// Timeouts contains all important timeouts for HTTP server
var Timeouts timeouts

func init() {
	Service = service{
		// Port:	getEnv("BASIC_SERVICE_PORT"), // 7990
		Port:	"7990", // 7990
	}

	Timeouts = timeouts{
		Read: 10,
		Write: 10,
		ReadHeader: 10,
		Idle: 10,
	}
}

// getEnv extracts the string value from environment variable
func getEnv(env string) string {
	locEnv := os.Getenv(env)
	if locEnv == "" {
		panic(fmt.Errorf("missing env parameter %s", env))
	}
	return locEnv
}