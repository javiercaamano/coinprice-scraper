package config

import "time"

type Category string

const (
	CATEGORY_DEFI           = "DEFI"
	CATEGORY_STABLECOIN     = "STABLE_COIN"
	CATEGORY_PROOF_OF_WORK  = "PROOF_OF_WORK"
	CATEGORY_PROOF_OF_STAKE = "PROOF_OF_STAKE"
)

type Blockchain string

const (
	BLOCKCHAIN_DEFAULT = "" // Blockchain is same as symbol name
	BLOCKCHAIN_ETH     = "ETH"
)

type SupportedAsset struct {
	Symbol string
	// Maybe Blockchain and Category should be pointers since they are enums
	Blockchain   Blockchain
	Categories   []Category
	LastSyncTime time.Time
}

func GetSupportedAssets() map[string]SupportedAsset {
	return map[string]SupportedAsset{
		"BTC": {
			Symbol:     "BTC",
			Categories: []Category{CATEGORY_PROOF_OF_WORK},
		},
		"ETH": {
			Symbol:     "ETH",
			Categories: []Category{CATEGORY_PROOF_OF_STAKE},
		},
		"USDT": {
			Symbol:     "USDT",
			Categories: []Category{CATEGORY_STABLECOIN},
			Blockchain: BLOCKCHAIN_ETH,
		},
		"LTC": {
			Symbol:     "LTC",
			Categories: []Category{CATEGORY_PROOF_OF_WORK},
		},
		"XRP": {
			Symbol:     "XRP",
			Categories: []Category{CATEGORY_DEFI},
		},
	}
}
