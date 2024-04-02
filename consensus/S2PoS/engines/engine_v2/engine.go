package engine_v2

import (
	"github.com/FRECNET/common"
	"github.com/FRECNET/consensus"
	"github.com/FRECNET/core/types"
	"github.com/FRECNET/ethdb"
	"github.com/FRECNET/params"
)

type S2PoS_v2 struct {
	config *params.S2PoSConfig // Consensus engine configuration parameters
	db     ethdb.Database      // Database to store and retrieve snapshot checkpoints
}

func New(config *params.S2PoSConfig, db ethdb.Database) *S2PoS_v2 {
	return &S2PoS_v2{
		config: config,
		db:     db,
	}
}

func NewFaker(db ethdb.Database, config *params.S2PoSConfig) *S2PoS_v2 {
	var fakeEngine *S2PoS_v2
	// Set any missing consensus parameters to their defaults
	conf := config

	// Allocate the snapshot caches and create the engine
	fakeEngine = &S2PoS_v2{
		config: conf,
		db:     db,
	}
	return fakeEngine
}

func (consensus *S2PoS_v2) Author(header *types.Header) (common.Address, error) {
	return common.Address{}, nil
}

func (consensus *S2PoS_v2) VerifyHeader(chain consensus.ChainReader, header *types.Header, fullVerify bool) error {
	return nil
}
