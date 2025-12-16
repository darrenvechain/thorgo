package builtins

//go:generate mkdir -p contracts
//go:generate curl -o ./contracts/staker.sol https://raw.githubusercontent.com/vechain/thor/refs/heads/master/builtin/gen/staker.sol
//go:generate rm -rf ./compiled/
//go:generate docker run -v ./:/sources ethereum/solc:0.8.20 --evm-version paris --via-ir --optimize --optimize-runs 200 --overwrite -o /sources/compiled --abi --bin /sources/contracts/staker.sol
//go:generate go run ../cmd/thorgen --abi https://raw.githubusercontent.com/vechain/b32/refs/heads/master/ABIs/authority.json --pkg builtins --out authority.go --type Authority
//go:generate go run ../cmd/thorgen --abi https://raw.githubusercontent.com/vechain/b32/refs/heads/master/ABIs/energy.json --pkg builtins --out vtho.go --type VTHO
//go:generate go run ../cmd/thorgen --abi https://raw.githubusercontent.com/vechain/b32/refs/heads/master/ABIs/executor.json --pkg builtins --out executor.go --type Executor
//go:generate go run ../cmd/thorgen --abi https://raw.githubusercontent.com/vechain/b32/refs/heads/master/ABIs/params.json --pkg builtins --out params.go --type Params
//go:generate go run ../cmd/thorgen --abi https://raw.githubusercontent.com/vechain/b32/refs/heads/master/ABIs/prototype.json --pkg builtins --out prototype.go --type Prototype
//go:generate go run ../cmd/thorgen --abi ./compiled/Staker.abi --pkg builtins --out staker.go --type Staker
//go:generate sh -c "find . -name '*.go' ! -name 'gen.go' -exec sed -i.bak 's/(address common.Address, /(/' {} +"
//go:generate sh -c "find . -name '*.go' ! -name 'gen.go' -exec sed -i.bak 's/(address, thor)/(thor)/g' {} +"
//go:generate sh -c "sed -i.bak 's/contracts.New(thor, address/contracts.New(thor, common.HexToAddress(\"0x0000000000000000000000417574686f72697479\")/g' authority.go"
//go:generate sh -c "sed -i.bak 's/contracts.New(thor, address/contracts.New(thor, common.HexToAddress(\"0x0000000000000000000000000000456e65726779\")/g' vtho.go"
//go:generate sh -c "sed -i.bak 's/contracts.New(thor, address/contracts.New(thor, common.HexToAddress(\"0x0000000000000000000000004578656375746f72\")/g' executor.go"
//go:generate sh -c "sed -i.bak 's/contracts.New(thor, address/contracts.New(thor, common.HexToAddress(\"0x0000000000000000000000000000506172616d73\")/g' params.go"
//go:generate sh -c "sed -i.bak 's/contracts.New(thor, address/contracts.New(thor, common.HexToAddress(\"0x000000000000000000000050726f746f74797065\")/g' prototype.go"
//go:generate sh -c "sed -i.bak 's/contracts.New(thor, address/contracts.New(thor, common.HexToAddress(\"0x00000000000000000000000000005374616b6572\")/g' staker.go"
//go:generate sh -c "find . -name '*.bak' -exec rm -f {} +"
