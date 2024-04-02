package S2PoS

import (
	"testing"

	"github.com/FRECNET/core/rawdb"
	"github.com/FRECNET/params"
	"github.com/stretchr/testify/assert"
)

func TestAdaptorShouldShareDbWithV1Engine(t *testing.T) {
	database := rawdb.NewMemoryDatabase()
	config := params.TestS2PoSMockChainConfig.S2PoS
	engine := New(config, database)

	assert := assert.New(t)
	assert.Equal(engine.EngineV1.GetDb(), engine.GetDb())
}
