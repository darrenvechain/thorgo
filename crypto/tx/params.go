// Copyright (c) 2018 The VeChainThor developers

// Distributed under the GNU Lesser General Public License v3.0 software license, see the accompanying
// file LICENSE or <https://www.gnu.org/licenses/lgpl-3.0.html>

package tx

import (
	"github.com/ethereum/go-ethereum/params"
)

// Constants of vechain thor blockchain.
const (
	blockInterval             uint64 = 10 // time interval between two consecutive blocks.
	txGas                     uint64 = 5000
	clauseGas                        = params.TxGas - txGas
	clauseGasContractCreation        = params.TxGasContractCreation - txGas
	maxTxWorkDelay            uint32 = 30 // (unit: block) if tx delay exceeds this value, no energy can be exchanged.
	txDataNonZeroGas          uint64 = 68
)
