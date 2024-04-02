package FREx

import (
	"github.com/FRECNET/accounts/abi/bind"
	"github.com/FRECNET/common"
	"github.com/FRECNET/contracts/FREx/contract"
)

type FREXListing struct {
	*contract.FREXListingSession
	contractBackend bind.ContractBackend
}

func NewMyFREXListing(transactOpts *bind.TransactOpts, contractAddr common.Address, contractBackend bind.ContractBackend) (*FREXListing, error) {
	smartContract, err := contract.NewFREXListing(contractAddr, contractBackend)
	if err != nil {
		return nil, err
	}

	return &FREXListing{
		&contract.FREXListingSession{
			Contract:     smartContract,
			TransactOpts: *transactOpts,
		},
		contractBackend,
	}, nil
}

func DeployFREXListing(transactOpts *bind.TransactOpts, contractBackend bind.ContractBackend) (common.Address, *FREXListing, error) {
	contractAddr, _, _, err := contract.DeployFREXListing(transactOpts, contractBackend)
	if err != nil {
		return contractAddr, nil, err
	}
	smartContract, err := NewMyFREXListing(transactOpts, contractAddr, contractBackend)
	if err != nil {
		return contractAddr, nil, err
	}

	return contractAddr, smartContract, nil
}
