package constants

import (
	"github.com/smartcontractkit/chainlink/v2/core/assets"
	"math/big"
)

var (
	SubscriptionBalanceJuels = assets.Ether(10).ToInt()

	// optional flags
	FallbackWeiPerUnitLink = big.NewInt(6e16)

	MinConfs                        = 3
	NodeSendingKeyFundingAmountGwei = assets.GWei(0).Int64() //100000000 = 0.1 ETH
	MaxGasLimit                     = int64(2.5e6)
	StalenessSeconds                = int64(86400)
	GasAfterPayment                 = int64(33285)
	FlatFeeTier1                    = int64(500)
	FlatFeeTier2                    = int64(500)
	FlatFeeTier3                    = int64(500)
	FlatFeeTier4                    = int64(500)
	FlatFeeTier5                    = int64(500)
	ReqsForTier2                    = int64(0)
	ReqsForTier3                    = int64(0)
	ReqsForTier4                    = int64(0)
	ReqsForTier5                    = int64(0)
)