package thorest

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

type debugTraceClause struct {
	Target string `json:"target"`
	Name   string `json:"name"`
	Config any    `json:"config"`
}

// TxRevertResponse represents the response from the output when debugging a reverted clause.
type TxRevertResponse struct {
	From    common.Address `json:"from"`    // From is the contract caller
	Gas     *hexutil.Big   `json:"gas"`     // Gas is the gas limit for the clause
	GasUsed *hexutil.Big   `json:"gasUsed"` // GasUsed is the gas used for the clause
	To      common.Address `json:"to"`      // To is the contract or account address
	Input   hexutil.Bytes  `json:"input"`   // Input is the data provided as part of the clause
	Output  hexutil.Bytes  `json:"output"`  // Output is the data returned from the contract call.
	Error   string         `json:"error"`   // Error is the error message if the EVM experienced an error
	Value   *hexutil.Big   `json:"value"`   // Value is the amount of VET transferred in the clause
	Type    string         `json:"type"`    // Type is the type of clause (e.g., CALL, CREATE, etc.)
}
