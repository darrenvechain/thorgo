package client

import (
	"strconv"

	"github.com/ethereum/go-ethereum/common"
)

type Revision struct {
	value string
}

func RevisionID(hash common.Hash) Revision {
	return Revision{value: hash.Hex()}
}

func RevisionNumber(number int) Revision {
	return Revision{value: strconv.Itoa(number)}
}

func RevisionBest() Revision {
	return Revision{value: "best"}
}

func RevisionFinalized() Revision {
	return Revision{value: "finalized"}
}

func RevisionJustified() Revision {
	return Revision{value: "justified"}
}
