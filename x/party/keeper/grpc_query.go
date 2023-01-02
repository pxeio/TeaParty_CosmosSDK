package keeper

import (
	"github.com/TeaPartyCrypto/partychain/x/party/types"
)

var _ types.QueryServer = Keeper{}
