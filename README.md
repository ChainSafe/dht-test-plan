# dht-test-plan

This repo contains a (wip) test plan for use with [testground](https://github.com/testground/testground) that tests the content routing functionality of go-libp2p-kad-dht.

## Requirements

- go 1.19
- [testground](https://github.com/testground/testground) - follow install instructions in readme

## Usage

Once testground is installed, run `testground daemon` in one terminal.

Then:

```
git clone https://github.com/ChainSafe/dht-test-plan.git
testground plan import --from ./dht-test-plan
cd dht-test-plan
testground run single --plan=dht-test-plan --testcase=find-providers --builder=docker:go --runner=local:docker --instances=1 --wait
```