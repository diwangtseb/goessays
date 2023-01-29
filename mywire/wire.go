//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"context"

	"github.com/google/wire"
)

func initialize(ctx context.Context) Baz {
	wire.Build(SuperSet)
	return Baz{}
}
