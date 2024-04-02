package vm

import (
	"github.com/FRECNET/FREx/tradingstate"
	"github.com/FRECNET/common"
	"github.com/FRECNET/log"
	"github.com/FRECNET/params"
)

const FREXPriceNumberOfBytesReturn = 32

// FRExPrice implements a pre-compile contract to get token price in FREx

type FRExLastPrice struct {
	tradingStateDB *tradingstate.TradingStateDB
}
type FRExEpochPrice struct {
	tradingStateDB *tradingstate.TradingStateDB
}

func (t *FRExLastPrice) RequiredGas(input []byte) uint64 {
	return params.FREXPriceGas
}

func (t *FRExLastPrice) Run(input []byte) ([]byte, error) {
	// input includes baseTokenAddress, quoteTokenAddress
	if t.tradingStateDB != nil && len(input) == 64 {
		base := common.BytesToAddress(input[12:32]) // 20 bytes from 13-32
		quote := common.BytesToAddress(input[44:])  // 20 bytes from 45-64
		price := t.tradingStateDB.GetLastPrice(tradingstate.GetTradingOrderBookHash(base, quote))
		if price != nil {
			log.Debug("Run GetLastPrice", "base", base.Hex(), "quote", quote.Hex(), "price", price)
			return common.LeftPadBytes(price.Bytes(), FREXPriceNumberOfBytesReturn), nil
		}
	}
	return common.LeftPadBytes([]byte{}, FREXPriceNumberOfBytesReturn), nil
}

func (t *FRExLastPrice) SetTradingState(tradingStateDB *tradingstate.TradingStateDB) {
	if tradingStateDB != nil {
		t.tradingStateDB = tradingStateDB.Copy()
	} else {
		t.tradingStateDB = nil
	}
}

func (t *FRExEpochPrice) RequiredGas(input []byte) uint64 {
	return params.FREXPriceGas
}

func (t *FRExEpochPrice) Run(input []byte) ([]byte, error) {
	// input includes baseTokenAddress, quoteTokenAddress
	if t.tradingStateDB != nil && len(input) == 64 {
		base := common.BytesToAddress(input[12:32]) // 20 bytes from 13-32
		quote := common.BytesToAddress(input[44:])  // 20 bytes from 45-64
		price := t.tradingStateDB.GetMediumPriceBeforeEpoch(tradingstate.GetTradingOrderBookHash(base, quote))
		if price != nil {
			log.Debug("Run GetEpochPrice", "base", base.Hex(), "quote", quote.Hex(), "price", price)
			return common.LeftPadBytes(price.Bytes(), FREXPriceNumberOfBytesReturn), nil
		}
	}
	return common.LeftPadBytes([]byte{}, FREXPriceNumberOfBytesReturn), nil
}

func (t *FRExEpochPrice) SetTradingState(tradingStateDB *tradingstate.TradingStateDB) {
	if tradingStateDB != nil {
		t.tradingStateDB = tradingStateDB.Copy()
	} else {
		t.tradingStateDB = nil
	}
}
