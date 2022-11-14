package main

import (
	"github.com/testground/sdk-go/run"
)

func main() {
	run.InvokeMap(testcases)
}

var testcases = map[string]interface{}{
	"find-providers": run.InitializedTestCaseFn(RunFindProviders),
}
