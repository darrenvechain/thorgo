package builtins

//go:generate go run ../cmd/thorgen --abi https://raw.githubusercontent.com/vechain/b32/refs/heads/master/ABIs/authority.json --pkg builtins --out authority.go --type Authority
//go:generate go run ../cmd/thorgen --abi https://raw.githubusercontent.com/vechain/b32/refs/heads/master/ABIs/energy.json --pkg builtins --out vtho.go --type VTHO
//go:generate go run ../cmd/thorgen --abi https://raw.githubusercontent.com/vechain/b32/refs/heads/master/ABIs/executor.json --pkg builtins --out executor.go --type Executor
//go:generate go run ../cmd/thorgen --abi https://raw.githubusercontent.com/vechain/b32/refs/heads/master/ABIs/params.json --pkg builtins --out params.go --type Params
//go:generate go run ../cmd/thorgen --abi https://raw.githubusercontent.com/vechain/b32/refs/heads/master/ABIs/prototype.json --pkg builtins --out prototype.go --type Prototype
//go:generate sh -c "find . -name '*.go' ! -name 'gen.go' -exec sed -i.bak 's/(address common.Address, /(/' {} +"
//go:generate sh -c "find . -name '*.go' ! -name 'gen.go' -exec sed -i.bak 's/(address, thor)/(thor)/g' {} +"
//go:generate sh -c "sed -i.bak 's/accounts.New(thor, address)/accounts.New(thor, common.HexToAddress(\"0x0000000000000000000000417574686f72697479\"))/g' authority.go"
//go:generate sh -c "sed -i.bak 's/accounts.New(thor, address)/accounts.New(thor, common.HexToAddress(\"0x0000000000000000000000000000456e65726779\"))/g' vtho.go"
//go:generate sh -c "sed -i.bak 's/accounts.New(thor, address)/accounts.New(thor, common.HexToAddress(\"0x0000000000000000000000004578656375746f72\"))/g' executor.go"
//go:generate sh -c "sed -i.bak 's/accounts.New(thor, address)/accounts.New(thor, common.HexToAddress(\"0x0000000000000000000000000000506172616d73\"))/g' params.go"
//go:generate sh -c "sed -i.bak 's/accounts.New(thor, address)/accounts.New(thor, common.HexToAddress(\"0x000000000000000000000050726f746f74797065\"))/g' prototype.go"
//go:generate sh -c "find . -name '*.bak' -exec rm -f {} +"
