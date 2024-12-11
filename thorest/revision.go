package thorest

import (
	"strconv"

	"github.com/ethereum/go-ethereum/common"
)

// Revision is used to specify the block revision of the chain when making queries.
type Revision struct {
	value string
}

// RevisionID creates a revision based on the block hash.
func RevisionID(hash common.Hash) Revision {
	return Revision{value: hash.Hex()}
}

// RevisionNumber creates a revision based on the block number.
func RevisionNumber(number int64) Revision {
	return Revision{value: strconv.FormatInt(number, 10)}
}

// RevisionBest creates a revision based on the best/ latest block.
func RevisionBest() Revision {
	return Revision{value: "best"}
}

// RevisionFinalized creates a revision based on the finalized block.
func RevisionFinalized() Revision {
	return Revision{value: "finalized"}
}

// RevisionJustified creates a revision based on the justified block.
func RevisionJustified() Revision {
	return Revision{value: "justified"}
}

// RevisionNext creates a revision based on the next block.
// Only allowed on the /accounts/* endpoint
func RevisionNext() Revision {
	return Revision{value: "next"}
}
