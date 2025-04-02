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


## CLIs

- [thorgen](./cmd/thorgen/README.md): A command line tool that generates Go smart contract wrappers for VeChainThor blockchain.

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
- To create your own transaction manager or signer, you can implement the `accounts.TxManager` interface:

```golang
// github.com/darrenvechain/thorgo/accounts
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

### 1: Creating a New Client

```golang
package main

import (
  "context"
  "fmt"

  "github.com/darrenvechain/thorgo"
  "github.com/darrenvechain/thorgo/solo"
  "github.com/ethereum/go-ethereum/common"
)

func main() {
  thor := thorgo.New(context.Background(), solo.URL)

  // Get an accounts balance
  acc, _ := thor.Account(common.HexToAddress("0x0000000000000000000000000000456e6570")).Get()
  fmt.Println(acc.Balance)
}

```

### 2: Interacting with a contract + Delegated Transaction

- It is recommended to create your smart contract wrapper using the `thorgen` CLI. This provides a more idiomatic way to
  interact with the contract.

<details>
  <summary>Expand</summary>

```golang
package main

import (
  "context"
  "log/slog"
  "math/big"

  "github.com/darrenvechain/thorgo"
  "github.com/darrenvechain/thorgo/builtins"
  "github.com/darrenvechain/thorgo/solo"
  "github.com/darrenvechain/thorgo/thorest"
  "github.com/darrenvechain/thorgo/transactions"
  "github.com/darrenvechain/thorgo/txmanager"
)

func main() {
  thor := thorgo.New(context.Background(), "http://localhost:8669")

  // Create a delegated transaction manager
  origin := txmanager.FromPK(solo.Keys()[0], thor.Client())
  gasPayer := txmanager.FromPK(solo.Keys()[1], thor.Client())
  txSender := txmanager.NewDelegated(thor.Client(), origin, gasPayer)

  // Use the `thorgen` CLI to build your own smart contract wrapper
  vtho, _ := builtins.NewVTHOTransactor(thor.Client(), txSender)

  // Create a new account to receive the tokens
  recipient, _ := txmanager.GeneratePK(thor.Client())

  // Call the balanceOf function
  balance, err := vtho.BalanceOf(recipient.Address(), thorest.RevisionBest())
  slog.Info("recipient balance before", "balance", balance, "error", err)

  tx, err := vtho.Transfer(recipient.Address(), big.NewInt(1000000000000000000), &transactions.Options{})
  if err != nil {
    slog.Error("transfer error", "error", err)
    return
  }
  receipt, _ := tx.Wait(context.Background())
  slog.Info("transfer receipt", "error", receipt.Reverted)

  balance, err = vtho.BalanceOf(recipient.Address(), thorest.RevisionBest())
  slog.Info("recipient balance after", "balance", balance, "error", err)
}

```

</details>

### 3: Multi Clause Transaction

<details>
  <summary>Expand</summary>

```golang
package main

import (
  "context"
  "log/slog"
  "math/big"

  "github.com/darrenvechain/thorgo"
  "github.com/darrenvechain/thorgo/builtins"
  "github.com/darrenvechain/thorgo/crypto/tx"
  "github.com/darrenvechain/thorgo/solo"
  "github.com/darrenvechain/thorgo/thorest"
  "github.com/darrenvechain/thorgo/transactions"
  "github.com/darrenvechain/thorgo/txmanager"
)

func main() {
  thor := thorgo.New(context.Background(), "http://localhost:8669")

  // Create a delegated transaction manager
  origin := txmanager.FromPK(solo.Keys()[0], thor.Client())
  recipient1, _ := txmanager.GeneratePK(thor.Client())
  recipient2, _ := txmanager.GeneratePK(thor.Client())

  vtho, _ := builtins.NewVTHOTransactor(thor.Client(), origin)

  clause1, _ := vtho.TransferAsClause(recipient1.Address(), big.NewInt(1000))
  clause2, _ := vtho.TransferAsClause(recipient2.Address(), big.NewInt(9999))

  tx, _ := origin.SendClauses([]*tx.Clause{clause1, clause2}, &transactions.Options{})
  slog.Info("transaction sent", "id", tx.ID())
  trx, _ := tx.Wait(context.Background())
  slog.Info("transaction mined", "reverted", trx.Reverted)

  balance1, _ := vtho.BalanceOf(recipient1.Address(), thorest.RevisionBest())
  balance2, _ := vtho.BalanceOf(recipient2.Address(), thorest.RevisionBest())

  slog.Info("recipient1", "balance", balance1)
  slog.Info("recipient2", "balance", balance2)
}

```
</details>
