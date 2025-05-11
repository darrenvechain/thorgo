package thorest

import "github.com/ethereum/go-ethereum/common"

// Peer represents a connected peer to the current node.
type Peer struct {
	Name        string      `json:"name"`
	BestBlockID common.Hash `json:"bestBlockID"`
	TotalScore  int64       `json:"totalScore"`
	PeerID      string      `json:"peerID"`
	NetAddr     string      `json:"netAddr"`
	Inbound     bool        `json:"inbound"`
	Duration    int64       `json:"duration"`
}
