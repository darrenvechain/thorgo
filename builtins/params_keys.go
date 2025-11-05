package builtins

import "github.com/ethereum/go-ethereum/common"

var (
	KeyExecutorAddress           = common.BytesToHash([]byte("executor"))
	KeyRewardRatio               = common.BytesToHash([]byte("reward-ratio"))
	KeyValidatorRewardPercentage = common.BytesToHash([]byte("validator-reward-percentage"))
	KeyLegacyTxBaseGasPrice      = common.BytesToHash([]byte("base-gas-price")) // the legacy tx default gas price
	KeyProposerEndorsement       = common.BytesToHash([]byte("proposer-endorsement"))
	KeyMaxBlockProposers         = common.BytesToHash([]byte("max-block-proposers"))
	KeyCurveFactor               = common.BytesToHash([]byte("curve-factor")) // curve factor to define VTHO issuance after PoS
	KeyDelegatorContractAddress  = common.BytesToHash([]byte("delegator-contract-address"))
	KeyStakerSwitches            = common.BytesToHash([]byte("staker-switches"))
)
