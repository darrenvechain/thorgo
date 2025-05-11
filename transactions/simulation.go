package transactions

import (
	"github.com/darrenvechain/thorgo/thorest"
)

// Simulation represents the result of a transaction simulation.
type Simulation struct {
	consumedGas  uint64
	reverted     bool
	outputs      []thorest.InspectResponse
	vmError      string
	intrinsicGas uint64
}

// TotalGas returns the total gas used in the simulation, which is the sum of consumed gas and intrinsic gas.
func (s *Simulation) TotalGas() uint64 {
	return s.consumedGas + s.intrinsicGas
}

// IsSuccess checks if the simulation was successful.
func (s *Simulation) IsSuccess() bool {
	return !s.reverted && s.vmError == ""
}

// ConsumedGas returns the amount of gas consumed during the simulation.
func (s *Simulation) ConsumedGas() uint64 {
	return s.consumedGas
}

// Reverted checks if the simulation was reverted.
func (s *Simulation) Reverted() bool {
	return s.reverted
}

// Outputs returns the outputs of the simulation.
func (s *Simulation) Outputs() []thorest.InspectResponse {
	return s.outputs
}

// VMError returns the VM error message if the simulation failed.
func (s *Simulation) VMError() string {
	return s.vmError
}

// IntrinsicGas returns the intrinsic gas used in the simulation.
func (s *Simulation) IntrinsicGas() uint64 {
	return s.intrinsicGas
}
