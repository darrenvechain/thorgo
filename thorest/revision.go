package thorest

import (
	"strconv"

	"github.com/ethereum/go-ethereum/common"
)

const (
	revBest      = "best"
	revFinalized = "finalized"
	revJustified = "justified"
	revNext      = "next"
)

// Revision is used to specify the block revision of the chain when making queries.
type Revision struct {
	value string
}

// RevisionID creates a revision based on the block ID.
func RevisionID(id common.Hash) Revision {
	return Revision{value: id.Hex()}
}

// RevisionNumber creates a revision based on the block number.
func RevisionNumber(number int64) Revision {
	return Revision{value: strconv.FormatInt(number, 10)}
}

// RevisionBest creates a revision based on the best/ latest block.
func RevisionBest() Revision {
	return Revision{value: revBest}
}

// RevisionFinalized creates a revision based on the finalized block.
func RevisionFinalized() Revision {
	return Revision{value: revFinalized}
}

// RevisionJustified creates a revision based on the justified block.
func RevisionJustified() Revision {
	return Revision{value: revJustified}
}

// RevisionNext creates a revision based on the next block.
// Only allowed on the /accounts/* endpoint
func RevisionNext() Revision {
	return Revision{value: revNext}
}
