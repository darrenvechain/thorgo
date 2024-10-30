# `thorgen`

`thorgen` is a command line tool that generates Go smart contract wrappers for VeChainThor blockchain. This CLI tool has been copied from [go-ethereum](https://github.com/ethereum/go-ethereum) and modified to work with VeChainThor blockchain.

## Features

`thorgen` provides typed contract wrappers to:
- Query contract state
- Create transaction clauses
- Execute single clause transactions
- Query on chain events

## Installation

```bash
go install github.com/darrenvechain/thorgo/cmd/thorgen
```

## Usage

### 1. Print help

Prints the command usage and options.

```bash
thorgen help
```

### 2. Contract Wrapper - ABI and BIN

Generates a Go smart contract wrapper for the given contract address.

```bash
thorgen \
    --abi /path/to/your/contract.abi \
    --bin /path/to/your/contract.bin \
    --pkg main \
    --out your-contract-wrapper.go \
    --type YourContract
```

### 3. Contract Wrapper - Hardhat Artifact

Generates a Go smart contract wrapper for the given hardhat artifact.

```bash
thorgen \
    --artifact /path/to/your/artifact.json \
    --pkg main \
    --out your-contract-wrapper.go \
    --type YourContract
```

### 4. Contract Wrapper - ABI URL

```bash
thorgen \
    --abi https://raw.githubusercontent.com/vechain/b32/refs/heads/master/ABIs/VeBetterDAO-b3tr.json \
    --pkg main \
    --out b3tr-erc20.go \
    --type B3tr
```
