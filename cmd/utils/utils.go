package utils

import (
	"github.com/FRECNET/FREx"
	"github.com/FRECNET/FRExlending"
	"github.com/FRECNET/eth"
	"github.com/FRECNET/eth/downloader"
	"github.com/FRECNET/ethstats"
	"github.com/FRECNET/les"
	"github.com/FRECNET/node"
	whisper "github.com/FRECNET/whisper/whisperv6"
)

// RegisterEthService adds an Ethereum client to the stack.
func RegisterEthService(stack *node.Node, cfg *eth.Config) {
	var err error
	if cfg.SyncMode == downloader.LightSync {
		err = stack.Register(func(ctx *node.ServiceContext) (node.Service, error) {
			return les.New(ctx, cfg)
		})
	} else {
		err = stack.Register(func(ctx *node.ServiceContext) (node.Service, error) {
			var FREXServ *FREx.FREX
			ctx.Service(&FREXServ)
			var lendingServ *FRExlending.Lending
			ctx.Service(&lendingServ)
			fullNode, err := eth.New(ctx, cfg, FREXServ, lendingServ)
			if fullNode != nil && cfg.LightServ > 0 {
				ls, _ := les.NewLesServer(fullNode, cfg)
				fullNode.AddLesServer(ls)
			}
			return fullNode, err
		})
	}
	if err != nil {
		Fatalf("Failed to register the Ethereum service: %v", err)
	}
}

// RegisterShhService configures Whisper and adds it to the given node.
func RegisterShhService(stack *node.Node, cfg *whisper.Config) {
	if err := stack.Register(func(n *node.ServiceContext) (node.Service, error) {
		return whisper.New(cfg), nil
	}); err != nil {
		Fatalf("Failed to register the Whisper service: %v", err)
	}
}

// RegisterEthStatsService configures the Ethereum Stats daemon and adds it to
// th egiven node.
func RegisterEthStatsService(stack *node.Node, url string) {
	if err := stack.Register(func(ctx *node.ServiceContext) (node.Service, error) {
		// Retrieve both eth and les services
		var ethServ *eth.Ethereum
		ctx.Service(&ethServ)

		var lesServ *les.LightEthereum
		ctx.Service(&lesServ)

		return ethstats.New(url, ethServ, lesServ)
	}); err != nil {
		Fatalf("Failed to register the Ethereum Stats service: %v", err)
	}
}

func RegisterFREXService(stack *node.Node, cfg *FREx.Config) {
	FREX := FREx.New(cfg)
	if err := stack.Register(func(n *node.ServiceContext) (node.Service, error) {
		return FREX, nil
	}); err != nil {
		Fatalf("Failed to register the FREX service: %v", err)
	}

	// register FRExlending service
	if err := stack.Register(func(n *node.ServiceContext) (node.Service, error) {
		return FRExlending.New(FREX), nil
	}); err != nil {
		Fatalf("Failed to register the FREXLending service: %v", err)
	}
}
