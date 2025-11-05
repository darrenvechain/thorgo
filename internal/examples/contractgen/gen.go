package contractgen

//go:generate docker run -v ./:/sources ethereum/solc:0.8.20 --evm-version paris --via-ir --optimize --optimize-runs 200 --overwrite -o /sources/compiled --abi --bin /sources/Echo.sol
//go:generate go run ../../../cmd/thorgen --abi ./compiled/Echo.abi --bin ./compiled/Echo.bin --pkg contractgen --out echo.go --type EchoContract
