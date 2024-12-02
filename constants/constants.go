package constants

const (
	MockTimeString = "2022-01-01 10:00:00"

	ExpectNoErrorPhrase     = "Expected no error, but got %v"
	RedisClosedPhrase       = "redis: client is closed"
	ExpectErrorPhrase       = "Expected error, but got nil"
	MissExpectationPhrase   = "Not all expectations were met: %v"
	FailedJsonBindingPhrase = "failed JSON binding"
	EmptyInputPhrase        = "empty input"
	ErrorPhrase             = "some error"
)
