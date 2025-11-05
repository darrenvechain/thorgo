# Thor GO SDK

`thorgo` is a Golang library designed to provide an easy and intuitive way to interact with the VeChainThor
blockchain. It simplifies blockchain interactions, making it straightforward for developers to build and manage
applications on VeChainThor.

## Key Features

- **Easy-to-Use Interface**: Provides a simple and accessible API for VeChainThor interactions.
- **Blockchain Interaction**: Facilitates transactions, smart contract interactions, and more.
- **Golang Support**: Leverages the power and efficiency of Go for blockchain development.

## Note on Geth

The Thor GO SDK is built on top of the latest version of [geth](https://github.com/ethereum/go-ethereum). Familiarity
with the Geth repository is encouraged, particularly when working with Application Binary Interfaces (ABIs),
cryptographic operations (hashing, signing, and managing private keys), and other low-level blockchain functions.
Understanding these elements can help in effectively utilizing the SDK and troubleshooting any related issues.

## Installation

To install the Thor GO SDK, run the following command:

```bash
go get github.com/darrenvechain/thorgo
``` 

## Quick Start

```go
package main

import (
	"context"

	"github.com/darrenvechain/thorgo"
	"github.com/darrenvechain/thorgo/thorest"
)

func main() {
	thor := thorgo.New(context.Background(), "https://mainnet.vechain.org")

	blockChan := make(chan *thorest.ExpandedBlock)
	sub := thor.Blocks().Subscribe(blockChan)
	defer sub.Unsubscribe()

	for block := range blockChan {
		println("new block: ", block.Number)
	}
}
```

## CLIs

- [thorgen](cmd/thorgen/README.md): A command line tool that generates Go smart contract wrappers for VeChainThor blockchain.

## Packages

### thorgo

- `github.com/darrenvechain/thorgo`
- `thorgo` is the primary package in the Thor GO SDK. It provides a high-level interface for interacting with the
  VeChainThor blockchain. This package includes functions for querying account balances, transactions, blocks, and smart
  contracts. It also supports simulating, building, and sending transactions, as well as interacting with smart
  contracts for reading and transacting.

### thorest

- `github.com/darrenvechain/thorgo/thorest`
- The `thorest` package provides raw REST access to the VeChainThor blockchain. It allows developers to query the
  blockchain directly without the need for higher-level abstractions provided by `thorgo`.

### txmanager

- `github.com/darrenvechain/thorgo/txmanager`
- The `txmanager` package provides a way to sign, send, and delegate transactions.
- The delegation managers can be used to easily delegate transaction gas fees.
- **Note**: The private key implementations in this package are not secure. It is recommended to use a secure key
  management solution in a production environment.
- To create your own transaction manager or signer, you can implement the `contracts.TxManager` interface:

```golang
// github.com/darrenvechain/thorgo/contracts
type TxManager interface {
	SendClauses(clauses []*tx.Clause, opts *transactions.Options) (*transactions.Visitor, error)
}
```

### tx

- `github.com/darrenvechain/thorgo/crypto/tx`
- The `tx` package is a copy of the [vechain/thor/tx](https://github.com/vechain/thor/tree/master/tx) package and can be
  used to build transactions where `thorgo` does not provide the necessary functionality.

### solo

- `github.com/darrenvechain/thorgo/solo`
- The `solo` package provides quick access to Thor solo values for testing and development purposes.

### certificate

- `github.com/darrenvechain/thorgo/crypto/certificate`
- The `certificate` package provides a way to encode, sign, and verify certificates in accordance
  with [VIP-192](https://github.com/vechain/VIPs/blob/master/vips/VIP-192.md)

### hdwallet

- `github.com/darrenvechain/thorgo/crypto/hdwallet`
- The `hdwallet` package provides a way to generate HD wallets and derive keys from them.

## Examples

### 1) Contract Generation with `thorgen` CLI

- See [gen.go](./internal/examples/contractgen/gen.go) for an example of generating a smart contract wrapper using the `thorgen` CLI.
- Run `go generate` in the `internal/examples/contractgen` directory to generate the contract wrapper.
- See the usage at [echo_test.go](./internal/examples/contractgen/echo_test.go).

### 2) Delegated Transaction

- See [delegated_tx.go](./internal/examples/delegatedtx/delegated_tx_test.go) for an example of sending a delegated transaction using the `txmanager` package.

### 3) Multi Clause Transaction

- See [multi_clause_tx.go](./internal/examples/multiclause/multi_clause_tx_test.go) for an example of sending a multi-clause transaction.

### 4) Hardhat Integration

- `thorgen` can natively generate contract bindings for smart contract artifacts produced by [Hardhat](https://hardhat.org/).
- See [gen.go](./internal/examples/hardhat/gen.go) and [counter.go](./internal/examples/hardhat/counter.go) for an example of generating a smart contract wrapper using Hardhat artifacts.

### 5) Custom Transaction Building

- See [custom_tx.go](./internal/examples/txbuilding/tx_building_test.go) to see how to build, simulate and send transactions with custom options.