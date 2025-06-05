// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package testcontract

import (
	"context"
	"errors"
	"math/big"
	"strings"
	"time"

	"github.com/darrenvechain/thorgo/blocks"
	"github.com/darrenvechain/thorgo/contracts"
	"github.com/darrenvechain/thorgo/crypto/tx"
	"github.com/darrenvechain/thorgo/thorest"
	"github.com/darrenvechain/thorgo/transactions"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.Is
	_ = big.NewInt
	_ = strings.ReplaceAll
	_ = abi.ConvertType
	_ = hexutil.Decode
	_ = context.Background
	_ = tx.NewClause
	_ = blocks.New
	_ = time.Sleep
)

// Erc20MetaData contains all meta data concerning the Erc20 contract.
var Erc20MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"allowance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientAllowance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidApprover\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSpender\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"burnFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b5060405162000b1938038062000b19833981016040819052620000349162000123565b818160036200004483826200021c565b5060046200005382826200021c565b5050505050620002e8565b634e487b7160e01b600052604160045260246000fd5b600082601f8301126200008657600080fd5b81516001600160401b0380821115620000a357620000a36200005e565b604051601f8301601f19908116603f01168101908282118183101715620000ce57620000ce6200005e565b81604052838152602092508683858801011115620000eb57600080fd5b600091505b838210156200010f5785820183015181830184015290820190620000f0565b600093810190920192909252949350505050565b600080604083850312156200013757600080fd5b82516001600160401b03808211156200014f57600080fd5b6200015d8683870162000074565b935060208501519150808211156200017457600080fd5b50620001838582860162000074565b9150509250929050565b600181811c90821680620001a257607f821691505b602082108103620001c357634e487b7160e01b600052602260045260246000fd5b50919050565b601f8211156200021757600081815260208120601f850160051c81016020861015620001f25750805b601f850160051c820191505b818110156200021357828155600101620001fe565b5050505b505050565b81516001600160401b038111156200023857620002386200005e565b62000250816200024984546200018d565b84620001c9565b602080601f8311600181146200028857600084156200026f5750858301515b600019600386901b1c1916600185901b17855562000213565b600085815260208120601f198616915b82811015620002b95788860151825594840194600190910190840162000298565b5085821015620002d85787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b61082180620002f86000396000f3fe608060405234801561001057600080fd5b50600436106100a45760003560e01c806306fdde03146100a9578063095ea7b3146100c757806318160ddd146100ea57806323b872dd146100fc578063313ce5671461010f57806340c10f191461011e57806342966c681461013357806370a082311461014657806379cc67901461016f57806395d89b4114610182578063a9059cbb1461018a578063dd62ed3e1461019d575b600080fd5b6100b16101b0565b6040516100be919061061d565b60405180910390f35b6100da6100d5366004610687565b610242565b60405190151581526020016100be565b6002545b6040519081526020016100be565b6100da61010a3660046106b1565b61025c565b604051601281526020016100be565b61013161012c366004610687565b610280565b005b6101316101413660046106ed565b61028e565b6100ee610154366004610706565b6001600160a01b031660009081526020819052604090205490565b61013161017d366004610687565b61029b565b6100b16102b0565b6100da610198366004610687565b6102bf565b6100ee6101ab366004610728565b6102cd565b6060600380546101bf9061075b565b80601f01602080910402602001604051908101604052809291908181526020018280546101eb9061075b565b80156102385780601f1061020d57610100808354040283529160200191610238565b820191906000526020600020905b81548152906001019060200180831161021b57829003601f168201915b5050505050905090565b6000336102508185856102f8565b60019150505b92915050565b60003361026a85828561030a565b610275858585610366565b506001949350505050565b61028a82826103c5565b5050565b61029833826103fb565b50565b6102a682338361030a565b61028a82826103fb565b6060600480546101bf9061075b565b600033610250818585610366565b6001600160a01b03918216600090815260016020908152604080832093909416825291909152205490565b6103058383836001610431565b505050565b600061031684846102cd565b90506000198114610360578181101561035157828183604051637dc7a0d960e11b815260040161034893929190610795565b60405180910390fd5b61036084848484036000610431565b50505050565b6001600160a01b038316610390576000604051634b637e8f60e11b815260040161034891906107b6565b6001600160a01b0382166103ba57600060405163ec442f0560e01b815260040161034891906107b6565b610305838383610506565b6001600160a01b0382166103ef57600060405163ec442f0560e01b815260040161034891906107b6565b61028a60008383610506565b6001600160a01b038216610425576000604051634b637e8f60e11b815260040161034891906107b6565b61028a82600083610506565b6001600160a01b03841661045b57600060405163e602df0560e01b815260040161034891906107b6565b6001600160a01b038316610485576000604051634a1406b160e11b815260040161034891906107b6565b6001600160a01b038085166000908152600160209081526040808320938716835292905220829055801561036057826001600160a01b0316846001600160a01b03167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925846040516104f891815260200190565b60405180910390a350505050565b6001600160a01b03831661053157806002600082825461052691906107ca565b909155506105909050565b6001600160a01b038316600090815260208190526040902054818110156105715783818360405163391434e360e21b815260040161034893929190610795565b6001600160a01b03841660009081526020819052604090209082900390555b6001600160a01b0382166105ac576002805482900390556105cb565b6001600160a01b03821660009081526020819052604090208054820190555b816001600160a01b0316836001600160a01b03167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8360405161061091815260200190565b60405180910390a3505050565b600060208083528351808285015260005b8181101561064a5785810183015185820160400152820161062e565b506000604082860101526040601f19601f8301168501019250505092915050565b80356001600160a01b038116811461068257600080fd5b919050565b6000806040838503121561069a57600080fd5b6106a38361066b565b946020939093013593505050565b6000806000606084860312156106c657600080fd5b6106cf8461066b565b92506106dd6020850161066b565b9150604084013590509250925092565b6000602082840312156106ff57600080fd5b5035919050565b60006020828403121561071857600080fd5b6107218261066b565b9392505050565b6000806040838503121561073b57600080fd5b6107448361066b565b91506107526020840161066b565b90509250929050565b600181811c9082168061076f57607f821691505b60208210810361078f57634e487b7160e01b600052602260045260246000fd5b50919050565b6001600160a01b039390931683526020830191909152604082015260600190565b6001600160a01b0391909116815260200190565b8082018082111561025657634e487b7160e01b600052601160045260246000fdfea2646970667358221220eee086b8f87ffba6a968eeb8f34122f8c6288b914ba69bb1553821ab62d7f32a64736f6c63430008140033",
}

// DeployErc20 deploys a new Ethereum contract, binding an instance of Erc20 to it.
func DeployErc20(ctx context.Context, thor *thorest.Client, sender contracts.TxManager, opts *transactions.Options, name string, symbol string) (common.Hash, *Erc20, error) {
	parsed, err := Erc20MetaData.GetAbi()
	if err != nil {
		return common.Hash{}, nil, err
	}

	bytes, err := hexutil.Decode(Erc20MetaData.Bin)
	if err != nil {
		return common.Hash{}, nil, err
	}
	contract, txID, err := contracts.NewDeployer(thor, bytes, parsed).Deploy(ctx, sender, opts, name, symbol)
	if err != nil {
		return common.Hash{}, nil, err
	}
	return txID, &Erc20{thor: thor, contract: contract}, nil
}

// Erc20 is an auto generated Go binding around an Ethereum contract, allowing you to query and create clauses.
type Erc20 struct {
	thor     *thorest.Client     // Thor client connection to use
	contract *contracts.Contract // Generic contract wrapper for the low level calls
}

// NewErc20 creates a new instance of Erc20, bound to a specific deployed contract.
func NewErc20(address common.Address, thor *thorest.Client) (*Erc20, error) {
	parsed, err := Erc20MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	contract := contracts.New(thor, address, parsed)
	return &Erc20{thor: thor, contract: contract}, nil
}

// Address returns the address of the contract.
func (_Erc20 *Erc20) Address() common.Address {
	return _Erc20.contract.Address
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Erc20 *Erc20) Allowance(owner common.Address, spender common.Address) *contracts.Caller[*big.Int] {
	return contracts.NewCaller[*big.Int](_Erc20.contract, "allowance", owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Erc20 *Erc20) BalanceOf(account common.Address) *contracts.Caller[*big.Int] {
	return contracts.NewCaller[*big.Int](_Erc20.contract, "balanceOf", account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Erc20 *Erc20) Decimals() *contracts.Caller[uint8] {
	return contracts.NewCaller[uint8](_Erc20.contract, "decimals")
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Erc20 *Erc20) Name() *contracts.Caller[string] {
	return contracts.NewCaller[string](_Erc20.contract, "name")
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Erc20 *Erc20) Symbol() *contracts.Caller[string] {
	return contracts.NewCaller[string](_Erc20.contract, "symbol")
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Erc20 *Erc20) TotalSupply() *contracts.Caller[*big.Int] {
	return contracts.NewCaller[*big.Int](_Erc20.contract, "totalSupply")
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_Erc20 *Erc20) Approve(spender common.Address, value *big.Int) *contracts.Sender {
	return contracts.NewSender(_Erc20.contract, "approve", spender, value)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 value) returns()
func (_Erc20 *Erc20) Burn(value *big.Int) *contracts.Sender {
	return contracts.NewSender(_Erc20.contract, "burn", value)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 value) returns()
func (_Erc20 *Erc20) BurnFrom(account common.Address, value *big.Int) *contracts.Sender {
	return contracts.NewSender(_Erc20.contract, "burnFrom", account, value)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns()
func (_Erc20 *Erc20) Mint(to common.Address, amount *big.Int) *contracts.Sender {
	return contracts.NewSender(_Erc20.contract, "mint", to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_Erc20 *Erc20) Transfer(to common.Address, value *big.Int) *contracts.Sender {
	return contracts.NewSender(_Erc20.contract, "transfer", to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_Erc20 *Erc20) TransferFrom(from common.Address, to common.Address, value *big.Int) *contracts.Sender {
	return contracts.NewSender(_Erc20.contract, "transferFrom", from, to, value)
}

// Erc20Approval represents a Approval event raised by the Erc20 contract.
type Erc20Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Log     *thorest.EventLog
}

type Erc20ApprovalCriteria struct {
	Owner   *common.Address `abi:"owner"`
	Spender *common.Address `abi:"spender"`
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Erc20 *Erc20) FilterApproval(criteria []Erc20ApprovalCriteria, filters *thorest.LogFilters) ([]Erc20Approval, error) {
	topicHash := _Erc20.contract.ABI.Events["Approval"].ID

	criteriaSet := make([]thorest.EventCriteria, len(criteria))
	for i, c := range criteria {
		crteria := thorest.EventCriteria{
			Address: &_Erc20.contract.Address,
			Topic0:  &topicHash,
		}
		if c.Owner != nil {
			matcher := *c.Owner
			topics, err := abi.MakeTopics([]interface{}{matcher})
			if err != nil {
				return nil, err
			}
			crteria.Topic1 = &topics[0][0]
		}
		if c.Spender != nil {
			matcher := *c.Spender
			topics, err := abi.MakeTopics([]interface{}{matcher})
			if err != nil {
				return nil, err
			}
			crteria.Topic2 = &topics[0][0]
		}

		criteriaSet[i] = crteria
	}

	if len(criteriaSet) == 0 {
		criteriaSet = append(criteriaSet, thorest.EventCriteria{
			Address: &_Erc20.contract.Address,
			Topic0:  &topicHash,
		})
	}

	logs, err := _Erc20.thor.FilterEvents(criteriaSet, filters)
	if err != nil {
		return nil, err
	}

	events := make([]Erc20Approval, len(logs))
	for i, log := range logs {
		event := new(Erc20Approval)
		if err := _Erc20.contract.UnpackLog(event, "Approval", log); err != nil {
			return nil, err
		}
		event.Log = log
		events[i] = *event
	}

	return events, nil
}

// WatchApproval listens for on chain events binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Erc20 *Erc20) WatchApproval(criteria []Erc20ApprovalCriteria, ctx context.Context, bufferSize int64) (chan *Erc20Approval, error) {
	topicHash := _Erc20.contract.ABI.Events["Approval"].ID
	criteriaSet := make([]thorest.EventCriteria, len(criteria))

	for i, c := range criteria {
		crteria := thorest.EventCriteria{
			Address: &_Erc20.contract.Address,
			Topic0:  &topicHash,
		}
		if c.Owner != nil {
			matcher := *c.Owner
			topics, err := abi.MakeTopics([]interface{}{matcher})
			if err != nil {
				return nil, err
			}
			crteria.Topic1 = &topics[0][0]
		}
		if c.Spender != nil {
			matcher := *c.Spender
			topics, err := abi.MakeTopics([]interface{}{matcher})
			if err != nil {
				return nil, err
			}
			crteria.Topic2 = &topics[0][0]
		}

		criteriaSet[i] = crteria
	}

	eventChan := make(chan *Erc20Approval, bufferSize)
	blocks := blocks.New(ctx, _Erc20.thor)
	ticker := blocks.Ticker()
	best, err := blocks.Best()
	if err != nil {
		return nil, err
	}

	go func(current int64) {
		defer close(eventChan)

		for {
			select {
			case <-ticker.C():
				for { // loop until the current block is not found
					block, err := blocks.Expanded(thorest.RevisionNumber(current))
					if errors.Is(thorest.ErrNotFound, err) {
						break
					}
					if err != nil {
						time.Sleep(250 * time.Millisecond)
						continue
					}
					current++

					for _, log := range block.FilteredEvents(criteriaSet) {
						ev := new(Erc20Approval)
						if err := _Erc20.contract.UnpackLog(ev, "Approval", log); err != nil {
							continue
						}
						ev.Log = log
						eventChan <- ev
					}
				}
			case <-ctx.Done():
				return
			}
		}
	}(best.Number + 1)

	return eventChan, nil
}

// Erc20Transfer represents a Transfer event raised by the Erc20 contract.
type Erc20Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Log   *thorest.EventLog
}

type Erc20TransferCriteria struct {
	From *common.Address `abi:"from"`
	To   *common.Address `abi:"to"`
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Erc20 *Erc20) FilterTransfer(criteria []Erc20TransferCriteria, filters *thorest.LogFilters) ([]Erc20Transfer, error) {
	topicHash := _Erc20.contract.ABI.Events["Transfer"].ID

	criteriaSet := make([]thorest.EventCriteria, len(criteria))
	for i, c := range criteria {
		crteria := thorest.EventCriteria{
			Address: &_Erc20.contract.Address,
			Topic0:  &topicHash,
		}
		if c.From != nil {
			matcher := *c.From
			topics, err := abi.MakeTopics([]interface{}{matcher})
			if err != nil {
				return nil, err
			}
			crteria.Topic1 = &topics[0][0]
		}
		if c.To != nil {
			matcher := *c.To
			topics, err := abi.MakeTopics([]interface{}{matcher})
			if err != nil {
				return nil, err
			}
			crteria.Topic2 = &topics[0][0]
		}

		criteriaSet[i] = crteria
	}

	if len(criteriaSet) == 0 {
		criteriaSet = append(criteriaSet, thorest.EventCriteria{
			Address: &_Erc20.contract.Address,
			Topic0:  &topicHash,
		})
	}

	logs, err := _Erc20.thor.FilterEvents(criteriaSet, filters)
	if err != nil {
		return nil, err
	}

	events := make([]Erc20Transfer, len(logs))
	for i, log := range logs {
		event := new(Erc20Transfer)
		if err := _Erc20.contract.UnpackLog(event, "Transfer", log); err != nil {
			return nil, err
		}
		event.Log = log
		events[i] = *event
	}

	return events, nil
}

// WatchTransfer listens for on chain events binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Erc20 *Erc20) WatchTransfer(criteria []Erc20TransferCriteria, ctx context.Context, bufferSize int64) (chan *Erc20Transfer, error) {
	topicHash := _Erc20.contract.ABI.Events["Transfer"].ID
	criteriaSet := make([]thorest.EventCriteria, len(criteria))

	for i, c := range criteria {
		crteria := thorest.EventCriteria{
			Address: &_Erc20.contract.Address,
			Topic0:  &topicHash,
		}
		if c.From != nil {
			matcher := *c.From
			topics, err := abi.MakeTopics([]interface{}{matcher})
			if err != nil {
				return nil, err
			}
			crteria.Topic1 = &topics[0][0]
		}
		if c.To != nil {
			matcher := *c.To
			topics, err := abi.MakeTopics([]interface{}{matcher})
			if err != nil {
				return nil, err
			}
			crteria.Topic2 = &topics[0][0]
		}

		criteriaSet[i] = crteria
	}

	eventChan := make(chan *Erc20Transfer, bufferSize)
	blocks := blocks.New(ctx, _Erc20.thor)
	ticker := blocks.Ticker()
	best, err := blocks.Best()
	if err != nil {
		return nil, err
	}

	go func(current int64) {
		defer close(eventChan)

		for {
			select {
			case <-ticker.C():
				for { // loop until the current block is not found
					block, err := blocks.Expanded(thorest.RevisionNumber(current))
					if errors.Is(thorest.ErrNotFound, err) {
						break
					}
					if err != nil {
						time.Sleep(250 * time.Millisecond)
						continue
					}
					current++

					for _, log := range block.FilteredEvents(criteriaSet) {
						ev := new(Erc20Transfer)
						if err := _Erc20.contract.UnpackLog(ev, "Transfer", log); err != nil {
							continue
						}
						ev.Log = log
						eventChan <- ev
					}
				}
			case <-ctx.Done():
				return
			}
		}
	}(best.Number + 1)

	return eventChan, nil
}
