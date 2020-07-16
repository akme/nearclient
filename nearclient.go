package nearclient

import (
	"context"
	"encoding/json"

	"github.com/akme/nearclient/types"
	"github.com/ethereum/go-ethereum/rpc"
)

// Client defines typed wrappers for the Ethereum RPC API.
type Client struct {
	c *rpc.Client
}

// Dial connects a client to the given URL.
func Dial(rawurl string) (*Client, error) {
	return DialContext(context.Background(), rawurl)
}

// DialContext with context
func DialContext(ctx context.Context, rawurl string) (*Client, error) {
	c, err := rpc.DialContext(ctx, rawurl)
	if err != nil {
		return nil, err
	}
	return NewClient(c), nil
}

// NewClient creates a client that uses the given RPC client.
func NewClient(c *rpc.Client) *Client {
	return &Client{c}
}

// Close client
func (nc *Client) Close() {
	nc.c.Close()
}

// Validators gets list of validators
func (nc *Client) Validators(ctx context.Context, params interface{}) (*types.ValidatorsResponse, error) {
	var raw json.RawMessage

	if err := nc.c.CallContext(ctx, &raw, "validators", params); err != nil {
		return nil, err
	}
	var resp *types.ValidatorsResponse
	if err := json.Unmarshal(raw, &resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// GasPrice get the price of Gas
func (nc *Client) GasPrice(ctx context.Context, params interface{}) (*types.GasPrice, error) {
	var raw json.RawMessage
	if err := nc.c.CallContext(ctx, &raw, "gas_price", params); err != nil {
		return nil, err
	}
	var gasprice types.GasPrice
	if err := json.Unmarshal(raw, &gasprice); err != nil {
		return nil, err
	}

	return &gasprice, nil
}

// Status gets status of the network
func (nc *Client) Status(ctx context.Context) (*types.Status, error) {
	var raw json.RawMessage
	var params interface{}
	if err := nc.c.CallContext(ctx, &raw, "status", params); err != nil {
		return nil, err
	}
	var status types.Status
	if err := json.Unmarshal(raw, &status); err != nil {
		return nil, err
	}
	return &status, nil
}

// TransactionStatus gets status of transaction
func (nc *Client) TransactionStatus(ctx context.Context, txHash, AccountID string) (*types.TransactionStatus, error) {
	var raw json.RawMessage
	if err := nc.c.CallContext(ctx, &raw, "tx", txHash, AccountID); err != nil {
		return nil, err
	}
	var txstatus types.TransactionStatus
	if err := json.Unmarshal(raw, &txstatus); err != nil {
		return nil, err
	}
	return &txstatus, nil
}

// Block get block information
func (nc *Client) Block(ctx context.Context, params interface{}) (*types.Block, error) {
	var raw json.RawMessage

	if err := nc.c.CallContext(ctx, &raw, "block", params); err != nil {
		return nil, err
	}
	var block types.Block
	if err := json.Unmarshal(raw, &block); err != nil {
		return nil, err
	}
	return &block, nil
}

// Chunk gets chunk info
func (nc *Client) Chunk(ctx context.Context, params interface{}) (*types.ChunkResponse, error) {
	var raw json.RawMessage
	if err := nc.c.CallContext(ctx, &raw, "chunk", params); err != nil {
		return nil, err
	}
	var chunk types.ChunkResponse
	if err := json.Unmarshal(raw, &chunk); err != nil {
		return nil, err
	}
	return &chunk, nil
}
