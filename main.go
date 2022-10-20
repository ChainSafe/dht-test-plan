package main

import (
	"errors"

	"github.com/testground/sdk-go/run"
	"github.com/testground/sdk-go/runtime"
)

func main() {
	run.Invoke(testcase)
}

func testcase(runenv *runtime.RunEnv) error {
	switch c := runenv.TestCase; c {
	case "find-providers":
		return runFindProviders(runenv)
	default:
		return errors.New("unknown testcase")
	}
}
