package types

import (
	"fmt"
	"math/big"
)

// GasPrice Gas Price
type GasPrice struct {
	GasPrice string `json:"gas_price"`
}

// Status struct
type Status struct {
	ChainID    string     `json:"chain_id"`
	RPCAddr    string     `json:"rpc_addr"`
	SyncInfo   SyncInfo   `json:"sync_info"`
	Validators Validators `json:"validators"`
	Version    Version    `json:"version"`
}

// SyncInfo struct
type SyncInfo struct {
	LatestBlockHash   string `json:"latest_block_hash"`
	LatestBlockHeight int64  `json:"latest_block_height"`
	LatestBlockTime   string `json:"latest_block_time"`
	LatestStateRoot   string `json:"latest_state_root"`
	Syncing           bool   `json:"syncing"`
}

// Version struct
type Version struct {
	Build   string `json:"build"`
	Version string `json:"version"`
}

// Validators array
type Validators []Validator

// ValidatorsResponse struct
type ValidatorsResponse struct {
	CurrentFishermans []Fisherman    `json:"current_fishermen,omitempty"`
	NextFishermans    []Fisherman    `json:"next_fishermen,omitempty"`
	CurrentValidators Validators     `json:"current_validators,omitempty"`
	NextValidators    Validators     `json:"next_validators,omitempty"`
	CurrentProposal   string         `json:"current_proposal,omitempty"`
	EpochStartHeight  int64          `json:"epoch_start_height"`
	PrevEpochKickout  []EpochKickout `json:"prev_epoch_kickout"`
}

// EpochKickout struct
type EpochKickout struct {
	AccountID string                 `json:"account_id"`
	Reason    map[string]interface{} `json:"reason"`
}

// Fisherman struct
type Fisherman struct {
	AccountID string `json:"account_id"`
	PublicKey string `json:"public_key"`
	Stake     string `json:"stake"`
}

// Validator struct
type Validator struct {
	AccountID         string  `json:"account_id"`
	IsSlashed         bool    `json:"is_slashed"`
	ExpectedBlocksNum int64   `json:"num_expected_blocks,omitempty"`
	ProducedBlocksNum int64   `json:"num_produced_blocks,omitempty"`
	PublicKey         string  `json:"public_key,omitempty"`
	Shards            []int64 `json:"shards,omitempty"`
	Stake             string  `json:"stake,omitempty"`
}

// TransactionStatus struct
type TransactionStatus struct {
	ReceiptsOutcome    []RootOutcome `json:"receipts_outcome"`
	Transaction        Transaction   `json:"transaction"`
	Status             Status        `json:"status"`
	TransactionOutcome RootOutcome   `json:"transaction_outcome"`
}

// Transaction struct
type Transaction struct {
	SignerID   string   `json:"signer_id"`
	PublicKey  string   `json:"public_key"`
	Nonce      int64    `json:"nonce"`
	Hash       string   `json:"hash"`
	ReceiverID string   `json:"receiver_id"`
	BlockHash  string   `json:"block_hash"`
	Signature  string   `json:"signature"`
	Actions    []Action `json:"actions"`
}

// Action struct
type Action struct {
}

// RootOutcome struct
type RootOutcome struct {
	BlockHash string  `json:"block_hash"`
	ID        string  `json:"id"`
	Outcome   Outcome `json:"outcome"`
	Proof     []Proof `json:"proof,omitempty"`
}

// Outcome struct
type Outcome struct {
	GasBurnt   int64       `json:"gas_burnt"`
	Logs       []string    `json:"logs"`
	ReceiptIDs []string    `json:"receipt_ids"`
	Status     interface{} `json:"status"`
}

// Proof struct
type Proof struct {
	Direction string `json:"direction,omitempty"`
	Hash      string `json:"hash,omitempty"`
}

// Block struct
type Block struct {
	Author string        `json:"author"`
	Chunks []ChunkHeader `json:"chunks"`
	Header BlockHeader   `json:"header"`
}

// ChunkHeader struct
type ChunkHeader struct {
	BalanceBurnt         string   `json:"balance_burnt"`
	ChunkHash            string   `json:"chunk_hash"`
	EncodedLength        int64    `json:"encoded_length"`
	EncodedMerkleRoot    string   `json:"encoded_merkle_root"`
	GasLimit             int64    `json:"gas_limit"`
	GasUsed              int64    `json:"gas_used"`
	HeightCreated        int64    `json:"height_created"`
	HeightIncluded       int64    `json:"height_included"`
	OutcomeRoot          string   `json:"outcome_root"`
	OutgoingReceiptsRoot string   `json:"outgoing_receipts_root"`
	PrevBlockHash        string   `json:"prev_block_hash"`
	PrevStateRoot        string   `json:"prev_state_root"`
	RentPaid             string   `json:"rent_paid"`
	ShardID              int64    `json:"shard_id"`
	Signature            string   `json:"signature"`
	TxRoot               string   `json:"tx_root"`
	ValidatorProposals   []string `json:"validator_proposals"`
	ValidatorReward      string   `json:"validator_reward"`
}

// BlockHeader struct
type BlockHeader struct {
	Approvals          []string    `json:"approvals"`
	BlockMerkleRoot    string      `json:"block_merkle_root"`
	ChallengesResult   []string    `json:"challenges_result"`
	ChallengesRoot     string      `json:"challenges_root"`
	ChunkHeadersRoot   string      `json:"chunk_headers_root"`
	ChunkMask          interface{} `json:"chunk_mask"`
	ChunkReceiptsRoot  string      `json:"chunk_receipts_root"`
	ChunkTxRoot        string      `json:"chunk_tx_root"`
	ChunksIncluded     int64       `json:"chunks_included"`
	EpochID            string      `json:"epoch_id"`
	GasPrice           string      `json:"gas_price"`
	Hash               string      `json:"hash"`
	Height             int64       `json:"height"`
	LastDSFinalBlock   string      `json:"last_ds_final_block"`
	LastFinalBlock     string      `json:"last_final_block"`
	LatestFinalBlock   int64       `json:"latest_protocol_version"`
	NextBPHash         string      `json:"next_bp_hash"`
	NextEpochID        string      `json:"next_epoch_id"`
	OutcomeRoot        string      `json:"outcome_root"`
	PrevHash           string      `json:"prev_hash"`
	PrevStateRoot      string      `json:"prev_state_root"`
	RandomValue        string      `json:"random_value"`
	RentPaid           string      `json:"rent_paid"`
	Signature          string      `json:"signature"`
	Timestamp          string      `json:"timestamp"`
	TotalSupply        string      `json:"total_supply"`
	ValidatorProposals interface{} `json:"validator_proposals"`
	ValidatorReward    string      `json:"validator_reward"`
}

// ChunkResponse struct
type ChunkResponse struct {
	Author       string          `json:"author"`
	Header       ChunkHeader     `json:"header"`
	Receipts     []ReceiptHeader `json:"receipts"`
	Transactions []Transaction   `json:"transactions"`
}

type ReceiptHeader struct {
	PredecessorID string  `json:"predecessor_id"`
	Receipt       Receipt `json:"receipt"`
	ReceiptID     string  `json:"receipt_id"`
	ReceiverID    string  `json:"receiver_id"`
}

type Receipt struct {
	Action ActionRoot `json:"Action"`
}

type ActionRoot struct {
	Actions             []Action `json:"actions"`
	GasPrice            string   `json:"gas_price"`
	InputDataIDs        []string `json:"input_data_ids"`
	OutputDataReceivers []string `json:"output_data_receivers"`
	SignerID            string   `json:"signer_id"`
	SignerPublicKey     string   `json:"signer_public_key"`
}

// TotalStake computes total staked NEARs
func (v *Validators) TotalStake() *big.Int {
	totalStake := new(big.Int)
	for i := range *v {
		n := new(big.Int)
		n, ok := n.SetString((*v)[i].Stake, 10)
		if !ok {
			fmt.Println("SetString: error")
			//return
		}

		totalStake.Add(totalStake, n)
	}
	return totalStake
}
