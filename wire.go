//+build wireinject

package main

import (
	"github.com/google/wire"
)

func InitializeJabba() {
	wire.Build()
}
