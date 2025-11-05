package hardhat

//go:generate npm install
//go:generate npx hardhat compile
//go:generate go run ../../../cmd/thorgen --artifact ./artifacts/contracts/Counter.sol/Counter.json --pkg hardhat --out counter.go --type Counter
