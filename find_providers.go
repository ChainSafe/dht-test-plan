package main

import (
	"context"
	"time"

	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p/core/crypto"
	"github.com/libp2p/go-libp2p/core/host"

	"github.com/testground/sdk-go/runtime"
	"github.com/testground/sdk-go/sync"
)

var (
	defaultDuration  = time.Minute * 10
	defaultNodeCount = 40
)

func runFindProviders(runenv *runtime.RunEnv) error {
	ctx, cancel := context.WithTimeout(context.Background(), defaultDuration)
	defer cancel()

	client := sync.MustBoundClient(ctx, runenv)
	defer client.Close()

	hosts := make([]host.Host, defaultNodeCount)
	for i := range hosts {
		h, err := newHost(ctx)
		if err != nil {
			return err
		}
		hosts[i] = h
	}

	return nil
}

func newHost(ctx context.Context) (host.Host, error) {
	priv, _, err := crypto.GenerateKeyPair(crypto.Ed25519, 256)
	if err != nil {
		return nil, err
	}

	return libp2p.New(
		libp2p.Identity(priv),
		libp2p.NoListenAddrs,
	)
}
