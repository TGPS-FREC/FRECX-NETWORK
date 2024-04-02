package util

import (
	"fmt"
	"math/big"

	"github.com/FRECNET/consensus"
)

func RewardInflation(chain consensus.ChainReader, chainReward *big.Int, number uint64, blockPerYear uint64) *big.Int {
	// if chain != nil && chain.Config().IsTIPNoHalvingMNReward(new(big.Int).SetUint64(number)) {
	// 	return chainReward
	// }

	// if blockPerYear*2 <= number && number < blockPerYear*5 {
	// 	chainReward.Div(chainReward, new(big.Int).SetUint64(2))
	// }
	// if blockPerYear*5 <= number {
	// 	chainReward.Div(chainReward, new(big.Int).SetUint64(4))
	// }

	// return chainReward
	blockPer5Years := blockPerYear * 5

	fmt.Println("blockPer5Years", blockPer5Years)
	// Ensure we do not divide by zero or perform the operation if it's not needed.
	if blockPer5Years == 0 {
		fmt.Println("Error: blockPer5Years is zero, division by zero is not allowed.")
		return chainReward
	}

	if number <= blockPer5Years {
		fmt.Println("blockPer5Years: number::", number)
		// If the number of blocks is less than or equal to blockPer5Years,
		// no halving should occur, return the original reward.
		return chainReward
	}

	// Calculate the number of 5-year periods that have elapsed since reaching the 80,000,000 block mark.
	// Since we're dealing with uint64, the division operation is safe from overflow by nature,
	// but we've already checked for division by zero above.
	periodsElapsed := (number - blockPer5Years) / blockPer5Years
	fmt.Println("periodsElapsed", periodsElapsed)
	// For each elapsed period, halve the chainReward.
	for i := uint64(0); i <= periodsElapsed; i++ {
		chainReward.Div(chainReward, big.NewInt(2))
	}

	return chainReward

}
