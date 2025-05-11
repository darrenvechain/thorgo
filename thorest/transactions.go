package thorest

import (
	"github.com/darrenvechain/thorgo/crypto/tx"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

// SendTransactionResponse represents the response from sending a transaction.
type SendTransactionResponse struct {
	ID common.Hash `json:"id"`
}

// RawTransaction represents a raw transaction in the blockchain.
type RawTransaction struct {
	Raw  hexutil.Bytes `json:"raw"`            // Raw is the raw transaction data
	Meta *TxMeta       `json:"meta,omitempty"` // Meta is null if the transaction is pending
}

// TransactionReceipt represents the receipt of a transaction.
type TransactionReceipt struct {
	GasUsed  int64          `json:"gasUsed"`  // GasUsed is the amount of gas used by the transaction
	GasPayer common.Address `json:"gasPayer"` // GasPayer is the address that paid for the gas
	Paid     *hexutil.Big   `json:"paid"`     // Paid is the amount of VTHO paid for the transaction
	Reward   *hexutil.Big   `json:"reward"`   // Reward is the amount of VTHO given to the block producer
	Reverted bool           `json:"reverted"` // Reverted indicates whether the transaction was reverted
	Meta     *ReceiptMeta   `json:"meta"`     // Meta contains metadata about the transaction receipt
	Outputs  []Output       `json:"outputs"`  // Outputs is the list of outputs for the transaction
}

// Transaction represents a transaction in the blockchain.
type Transaction struct {
	ID                   common.Hash     `json:"id"`                             // ID is the transaction ID
	ChainTag             int64           `json:"chainTag"`                       // ChainTag is the chain tag
	Type                 uint8           `json:"type,omitempty"`                 // Type is the transaction type
	BlockRef             tx.BlockRef     `json:"blockRef"`                       // BlockRef is the block reference
	Expiration           int64           `json:"expiration"`                     // Expiration is the number of blocks on top of BlockRef, after which the transaction is considered expired
	Clauses              []tx.Clause     `json:"clauses"`                        // Clauses is the list of clauses in the transaction
	GasPriceCoef         int64           `json:"gasPriceCoef,omitempty"`         // GasPriceCoef is present if it is a legacy transaction
	MaxFeePerGas         *hexutil.Big    `json:"maxFeePerGas,omitempty"`         // MaxFeePerGas is present if it is a dynamic fee transaction
	MaxPriorityFeePerGas *hexutil.Big    `json:"maxPriorityFeePerGas,omitempty"` // MaxPriorityFeePerGas is present if it is a dynamic fee transaction
	Gas                  int64           `json:"gas"`                            // Gas is the maximum amount of gas that can be used in the transaction
	Origin               common.Address  `json:"origin"`                         // Origin is the address of the sender
	Delegator            *common.Address `json:"delegator,omitempty"`            // Delegator is null if the transaction is not delegated
	Nonce                *hexutil.Big    `json:"nonce"`                          // Nonce is the transaction nonce
	DependsOn            *common.Hash    `json:"dependsOn,omitempty"`            // DependsOn is null if the transaction is not dependent
	Size                 int64           `json:"size"`                           // Size is the size of the transaction in bytes
	Meta                 *TxMeta         `json:"meta,omitempty"`                 // Meta is null if the transaction is pending
}

// Transfer represents a transfer as a result of a transaction.
type Transfer struct {
	Sender    common.Address `json:"sender"`
	Recipient common.Address `json:"recipient"`
	Amount    *hexutil.Big   `json:"amount"`
}

// Output represents the output of a transaction, including VET transfers and smart contract events.
type Output struct {
	ContractAddress *common.Address `json:"contractAddress"`
	Events          []Event         `json:"events"`
	Transfers       []Transfer      `json:"transfers"`
}

// Event represents a smart contract event in the blockchain.
type Event struct {
	Address *common.Address `json:"address"`
	Topics  []common.Hash   `json:"topics"`
	Data    hexutil.Bytes   `json:"data"`
}

// ReceiptMeta contains metadata about a transaction receipt.
type ReceiptMeta struct {
	BlockID        common.Hash    `json:"blockID"`
	BlockNumber    int64          `json:"blockNumber"`
	BlockTimestamp int64          `json:"blockTimestamp"`
	TxID           common.Hash    `json:"txID"`
	TxOrigin       common.Address `json:"txOrigin"`
}

// TxMeta contains metadata about a transaction.
type TxMeta struct {
	BlockID        common.Hash `json:"blockID"`
	BlockNumber    int64       `json:"blockNumber"`
	BlockTimestamp int64       `json:"blockTimestamp"`
}
