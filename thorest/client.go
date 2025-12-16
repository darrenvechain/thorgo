package thorest

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"sync/atomic"
	"time"

	"github.com/darrenvechain/thorgo/crypto/tx"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

// buildURL constructs a URL path with optional query parameters.
// Parameters with empty values are omitted.
func buildURL(path string, params map[string]string) string {
	if len(params) == 0 {
		return path
	}

	query := url.Values{}
	for k, v := range params {
		if v != "" {
			query.Set(k, v)
		}
	}

	if len(query) == 0 {
		return path
	}
	return path + "?" + query.Encode()
}

// Client is a struct that provides methods to interact with the VeChainThor API.
type Client struct {
	client       *http.Client
	url          string
	ctx          context.Context
	genesisBlock atomic.Pointer[Block]
}

// NewClient creates a new Client instance with the given URL and HTTP client.
func NewClient(url string, client *http.Client) *Client {
	return newClient(url, client)
}

// NewClientFromURL creates a new Client instance with the given URL and a default HTTP client with a 30-second timeout.
func NewClientFromURL(url string) *Client {
	return NewClient(url, &http.Client{Timeout: 30 * time.Second})
}

func newClient(url string, client *http.Client) *Client {
	url = strings.TrimSuffix(url, "/")

	return &Client{
		client: client,
		url:    url,
		ctx:    context.Background(),
	}
}

// WithContext returns a shallow copy of the Client with the given context.
// The context will be used for all HTTP requests made by the returned client.
func (c *Client) WithContext(ctx context.Context) *Client {
	return &Client{
		client:       c.client,
		url:          c.url,
		ctx:          ctx,
		genesisBlock: c.genesisBlock,
	}
}

// Account fetches the account information for the given address.
func (c *Client) Account(addr common.Address) (*Account, error) {
	path := "/accounts/" + addr.Hex()
	return httpGet(c, path, &Account{})
}

// AccountAt fetches the account information for an address at the given revision.
func (c *Client) AccountAt(addr common.Address, revision Revision) (*Account, error) {
	path := buildURL("/accounts/"+addr.Hex(), map[string]string{"revision": revision.value})
	return httpGet(c, path, &Account{})
}

// Inspect will send an array of clauses to the node to simulate the execution of the clauses.
// This can be used to:
// - Read contract(s) state
// - Simulate the execution of a transaction
func (c *Client) Inspect(body InspectRequest) ([]InspectResponse, error) {
	path := "/accounts/*"
	response := make([]InspectResponse, 0)
	_, err := httpPost(c, path, body, &response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

// InspectAt will send an array of clauses to the node to simulate the execution of the clauses at the given revision.
func (c *Client) InspectAt(body InspectRequest, revision Revision) ([]InspectResponse, error) {
	path := buildURL("/accounts/*", map[string]string{"revision": revision.value})
	response := make([]InspectResponse, 0)
	_, err := httpPost(c, path, body, &response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

// AccountCode fetches the code for the account at the given address.
func (c *Client) AccountCode(addr common.Address) (*AccountCode, error) {
	path := "/accounts/" + addr.Hex() + "/code"
	return httpGet(c, path, &AccountCode{})
}

// AccountCodeAt fetches the code for the account at the given address and revision.
func (c *Client) AccountCodeAt(addr common.Address, revision Revision) (*AccountCode, error) {
	path := buildURL("/accounts/"+addr.Hex()+"/code", map[string]string{"revision": revision.value})
	return httpGet(c, path, &AccountCode{})
}

// AccountStorage fetches the storage value for the account at the given address and key.
func (c *Client) AccountStorage(addr common.Address, key common.Hash) (*AccountStorage, error) {
	path := "/accounts/" + addr.Hex() + "/storage/" + key.Hex()
	return httpGet(c, path, &AccountStorage{})
}

// AccountStorageAt fetches the storage value for the account at the given address and key at the given revision.
func (c *Client) AccountStorageAt(
	addr common.Address,
	key common.Hash,
	revision Revision,
) (*AccountStorage, error) {
	path := buildURL("/accounts/"+addr.Hex()+"/storage/"+key.Hex(), map[string]string{"revision": revision.value})
	return httpGet(c, path, &AccountStorage{})
}

// Block fetches the block for the given revision.
func (c *Client) Block(revision Revision) (*Block, error) {
	path := "/blocks/" + revision.value
	return httpGet(c, path, &Block{})
}

// BestBlock returns the best block.
func (c *Client) BestBlock() (*Block, error) {
	return c.Block(RevisionBest())
}

// GenesisBlock returns the genesis block.
func (c *Client) GenesisBlock() (*Block, error) {
	if c.genesisBlock.Load() == nil {
		block, err := c.Block(RevisionNumber(0))
		if err != nil {
			return nil, err
		}
		c.genesisBlock.Store(block)
	}
	return c.genesisBlock.Load(), nil
}

// ExpandedBlock fetches the block at the given revision with all the transactions expanded.
func (c *Client) ExpandedBlock(revision Revision) (*ExpandedBlock, error) {
	path := "/blocks/" + revision.value + "?expanded=true"
	return httpGet(c, path, &ExpandedBlock{})
}

// ChainTag returns the chain tag of the genesis block.
func (c *Client) ChainTag() (byte, error) {
	gen, err := c.GenesisBlock()
	if err != nil {
		return 0, err
	}
	return gen.ID[:][31], nil
}

// SendTransaction sends a transaction to the node.
func (c *Client) SendTransaction(transaction *tx.Transaction) (*SendTransactionResponse, error) {
	body := make(map[string]string)
	rlpTx, err := transaction.MarshalBinary()
	if err != nil {
		return nil, err
	}
	body["raw"] = hexutil.Encode(rlpTx)
	return httpPost(c, "/transactions", body, &SendTransactionResponse{})
}

// SendRawTransaction sends a raw transaction to the node.
func (c *Client) SendRawTransaction(bytes []byte) (*SendTransactionResponse, error) {
	body := make(map[string]string)
	body["raw"] = hexutil.Encode(bytes)
	return httpPost(c, "/transactions", body, &SendTransactionResponse{})
}

// Transaction fetches a transaction by its ID.
func (c *Client) Transaction(id common.Hash) (*Transaction, error) {
	path := "/transactions/" + id.Hex()
	return httpGet(c, path, &Transaction{})
}

// TransactionAt fetches a transaction by its ID for the given head block ID.
func (c *Client) TransactionAt(id common.Hash, head common.Hash) (*Transaction, error) {
	path := "/transactions/" + id.Hex() + "?head=" + head.Hex()
	return httpGet(c, path, &Transaction{})
}

// RawTransaction fetches a transaction by its ID and returns the raw transaction.
func (c *Client) RawTransaction(id common.Hash) (*RawTransaction, error) {
	path := "/transactions/" + id.Hex() + "?raw=true"
	return httpGet(c, path, &RawTransaction{})
}

// RawTransactionAt fetches a transaction by its ID for the given head block ID and returns the raw transaction.
func (c *Client) RawTransactionAt(id common.Hash, head common.Hash) (*RawTransaction, error) {
	path := "/transactions/" + id.Hex() + "?head=" + head.Hex() + "&raw=true"
	return httpGet(c, path, &RawTransaction{})
}

// PendingTransaction includes the pending block when fetching a transaction.
func (c *Client) PendingTransaction(id common.Hash) (*Transaction, error) {
	path := "/transactions/" + id.Hex() + "?pending=true"
	return httpGet(c, path, &Transaction{})
}

// TransactionReceipt fetches a transaction receipt by its ID.
func (c *Client) TransactionReceipt(id common.Hash) (*TransactionReceipt, error) {
	path := "/transactions/" + id.Hex() + "/receipt"
	return httpGet(c, path, &TransactionReceipt{})
}

// TransactionReceiptAt fetches a transaction receipt by its ID for the given head block ID.
func (c *Client) TransactionReceiptAt(id common.Hash, head common.Hash) (*TransactionReceipt, error) {
	path := "/transactions/" + id.Hex() + "/receipt?revision=" + head.Hex()
	return httpGet(c, path, &TransactionReceipt{})
}

// FilterEvents fetches the event logs that match the given filter.
func (c *Client) FilterEvents(filter *EventFilter) ([]*EventLog, error) {
	path := "/logs/event"
	events := make([]*EventLog, 0)
	_, err := httpPost(c, path, filter, &events)
	if err != nil {
		return nil, err
	}
	return events, nil
}

// FilterTransfers fetches the transfer logs that match the given filter.
func (c *Client) FilterTransfers(filter *TransferFilter) ([]*TransferLog, error) {
	path := "/logs/transfer"
	transfers := make([]*TransferLog, 0)
	_, err := httpPost(c, path, filter, &transfers)
	if err != nil {
		return nil, err
	}
	return transfers, nil
}

// Peers fetches the list of peers connected to the node.
func (c *Client) Peers() ([]Peer, error) {
	path := "/node/network/peers"
	peers := make([]Peer, 0)
	_, err := httpGet(c, path, &peers)
	if err != nil {
		return nil, err
	}
	return peers, nil
}

// DebugRevertReason fetches the revert reason for the transaction.
func (c *Client) DebugRevertReason(receipt *TransactionReceipt) (*TxRevertResponse, error) {
	path := "/debug/tracers"
	config := make(map[string]any)
	config["OnlyTopCall"] = true
	body := debugTraceClause{
		Config: config,
		Name:   "call",
		Target: fmt.Sprintf("%s/%s/%d", receipt.Meta.BlockID.Hex(), receipt.Meta.TxID.Hex(), len(receipt.Outputs)),
	}
	return httpPost(c, path, body, &TxRevertResponse{})
}

// FeesHistory fetches the fee history for the given block range.
func (c *Client) FeesHistory(revision Revision, blockCount int64, rewardPercentiles []float64) (*FeesHistory, error) {
	var url strings.Builder
	url.WriteString("/fees/history?blockCount=" + fmt.Sprint(blockCount) + "&newestBlock=" + revision.value)
	if len(rewardPercentiles) > 0 {
		var values []string
		for _, v := range rewardPercentiles {
			values = append(values, strconv.FormatFloat(v, 'f', -1, 64))
		}
		url.WriteString("&rewardPercentiles=" + strings.Join(values, ","))
	}
	return httpGet(c, url.String(), &FeesHistory{})
}

// FeesPriority fetches the suggested priority fee for the next block.
func (c *Client) FeesPriority() (*FeesPriority, error) {
	return httpGet(c, "/fees/priority", &FeesPriority{})
}

func httpGet[T any](c *Client, endpoint string, v *T) (*T, error) {
	req, err := http.NewRequestWithContext(c.ctx, http.MethodGet, c.url+endpoint, nil)
	if err != nil {
		return nil, err
	}
	return httpDo(c, req, v)
}

func httpPost[T any](c *Client, path string, body any, v *T) (*T, error) {
	reqBody, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	request, err := http.NewRequestWithContext(c.ctx, http.MethodPost, c.url+path, bytes.NewReader(reqBody))
	if err != nil {
		return nil, err
	}
	request.Header.Set("Content-Type", "application/json")
	return httpDo(c, request, v)
}

func httpDo[T any](c *Client, req *http.Request, v *T) (*T, error) {
	response, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if statusOK := response.StatusCode >= 200 && response.StatusCode < 300; !statusOK {
		return nil, newHttpError(response)
	}

	// Read the entire body into a buffer
	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	// Check if the body is "null"
	if strings.TrimSpace(string(responseBody)) == "null" {
		return nil, ErrNotFound
	}

	// Decode the JSON response
	err = json.Unmarshal(responseBody, v)
	if err != nil {
		return nil, err
	}

	return v, nil
}
