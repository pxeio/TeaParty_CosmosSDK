package main

import (
	pkgadapter "knative.dev/eventing/pkg/adapter/v2"

	adams "github.com/TeaPartyCrypto/partychain/adams"
)

func main() {
	pkgadapter.Main("warren-adapter", adams.EnvAccessorCtor, adams.NewAdapter)
}
