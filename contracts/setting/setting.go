// Copyright (c) 2018 FRECNET
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with this program. If not, see <http://www.gnu.org/licenses/>.

package setting

import (
	"github.com/FRECNET/accounts/abi/bind"
	"github.com/FRECNET/common"
	contract "github.com/FRECNET/contracts/setting/contract"
)

type Setting struct {
	*contract.SettingSession
	contractBackend bind.ContractBackend
}

func NewSetting(transactOpts *bind.TransactOpts, contractAddr common.Address, contractBackend bind.ContractBackend) (*Setting, error) {
	setting, err := contract.NewSetting(contractAddr, contractBackend)
	if err != nil {
		return nil, err
	}

	return &Setting{
		&contract.SettingSession{
			Contract:     setting,
			TransactOpts: *transactOpts,
		},
		contractBackend,
	}, nil
}

func DeploySetting(transactOpts *bind.TransactOpts, contractBackend bind.ContractBackend, reward uint64, maxnode uint64, epoch uint64, tipincrease uint64) (common.Address, *Setting, error) {
	settingAddr, _, _, err := contract.DeploySetting(transactOpts, contractBackend, maxnode, epoch, reward, tipincrease)
	if err != nil {
		return settingAddr, nil, err
	}

	setting, err := NewSetting(transactOpts, settingAddr, contractBackend)
	if err != nil {
		return settingAddr, nil, err
	}

	return settingAddr, setting, nil
}
