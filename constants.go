package main

const (
	MockTimeString = "2022-01-01 10:00:00"

	ExpectedErrorPhrase   = "Expected no error, but got %v"
	RedisClosedPhrase     = "redis: client is closed"
	ExpectErrorPhrase     = "Expected error, but got nil"
	MissExpectationPhrase = "Not all expectations were met: %v"
)
