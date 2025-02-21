package thorest

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/darrenvechain/thorgo/crypto/tx"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

type Client struct {
	client       *http.Client
	url          string
	genesisBlock *Block
}

func NewClient(url string, client *http.Client) *Client {
	return newClient(url, client)
}

func NewClientFromURL(url string) *Client {
	return NewClient(url, &http.Client{})
}

func newClient(url string, client *http.Client) *Client {
	url = strings.TrimSuffix(url, "/")

	return &Client{
		client: client,
		url:    url,
	}
}

// Account fetches the account information for the given address.
func (c *Client) Account(addr common.Address) (*Account, error) {
	url := "/accounts/" + addr.Hex()
	return httpGet(c, url, &Account{})
}

// AccountAt fetches the account information for an address at the given revision.
func (c *Client) AccountAt(addr common.Address, revision Revision) (*Account, error) {
	url := "/accounts/" + addr.Hex() + "?revision=" + revision.value
	return httpGet(c, url, &Account{})
}

// Inspect will send an array of clauses to the node to simulate the execution of the clauses.
// This can be used to:
// - Read contract(s) state
// - Simulate the execution of a transaction
func (c *Client) Inspect(body InspectRequest) ([]InspectResponse, error) {
	url := "/accounts/*"
	response := make([]InspectResponse, 0)
	_, err := httpPost(c, url, body, &response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

// InspectAt will send an array of clauses to the node to simulate the execution of the clauses at the given revision.
func (c *Client) InspectAt(body InspectRequest, revision Revision) ([]InspectResponse, error) {
	url := "/accounts/*?revision=" + revision.value
	response := make([]InspectResponse, 0)
	_, err := httpPost(c, url, body, &response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

// AccountCode fetches the code for the account at the given address.
func (c *Client) AccountCode(addr common.Address) (*AccountCode, error) {
	url := "/accounts/" + addr.Hex() + "/code"
	return httpGet(c, url, &AccountCode{})
}

// AccountCodeAt fetches the code for the account at the given address and revision.
func (c *Client) AccountCodeAt(addr common.Address, revision Revision) (*AccountCode, error) {
	url := "/accounts/" + addr.Hex() + "/code?revision=" + revision.value
	return httpGet(c, url, &AccountCode{})
}

// AccountStorage fetches the storage value for the account at the given address and key.
func (c *Client) AccountStorage(addr common.Address, key common.Hash) (*AccountStorage, error) {
	url := "/accounts/" + addr.Hex() + "/storage/" + key.Hex()
	return httpGet(c, url, &AccountStorage{})
}

// AccountStorageAt fetches the storage value for the account at the given address and key at the given revision.
func (c *Client) AccountStorageAt(
	addr common.Address,
	key common.Hash,
	revision Revision,
) (*AccountStorage, error) {
	url := "/accounts/" + addr.Hex() + "/storage/" + key.Hex() + "?revision=" + revision.value
	return httpGet(c, url, &AccountStorage{})
}

// Block fetches the block for the given revision.
func (c *Client) Block(revision Revision) (*Block, error) {
	url := "/blocks/" + revision.value
	return httpGet(c, url, &Block{})
}

// BestBlock returns the best block.
func (c *Client) BestBlock() (*Block, error) {
	return c.Block(RevisionBest())
}

// GenesisBlock returns the genesis block.
func (c *Client) GenesisBlock() (*Block, error) {
	if c.genesisBlock == nil {
		block, err := c.Block(RevisionNumber(0))
		if err != nil {
			return nil, err
		}
		c.genesisBlock = block
	}
	return c.genesisBlock, nil
}

// ExpandedBlock fetches the block at the given revision with all the transactions expanded.
func (c *Client) ExpandedBlock(revision Revision) (*ExpandedBlock, error) {
	url := "/blocks/" + revision.value + "?expanded=true"
	return httpGet(c, url, &ExpandedBlock{})
}

// ChainTag returns the chain tag of the genesis block.
func (c *Client) ChainTag() (byte, error) {
	gen, err := c.GenesisBlock()
	if err != nil {
		return 0, err
	}
	return gen.ChainTag(), nil
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
	url := "/transactions/" + id.Hex()
	return httpGet(c, url, &Transaction{})
}

// TransactionAt fetches a transaction by its ID for the given head block ID.
func (c *Client) TransactionAt(id common.Hash, head common.Hash) (*Transaction, error) {
	url := "/transactions/" + id.Hex() + "?head=" + head.Hex()
	return httpGet(c, url, &Transaction{})
}

// RawTransaction fetches a transaction by its ID and returns the raw transaction.
func (c *Client) RawTransaction(id common.Hash) (*RawTransaction, error) {
	url := "/transactions/" + id.Hex() + "?raw=true"
	return httpGet(c, url, &RawTransaction{})
}

// RawTransactionAt fetches a transaction by its ID for the given head block ID and returns the raw transaction.
func (c *Client) RawTransactionAt(id common.Hash, head common.Hash) (*RawTransaction, error) {
	url := "/transactions/" + id.Hex() + "?head=" + head.Hex() + "&raw=true"
	return httpGet(c, url, &RawTransaction{})
}

// PendingTransaction includes the pending block when fetching a transaction.
func (c *Client) PendingTransaction(id common.Hash) (*Transaction, error) {
	url := "/transactions/" + id.Hex() + "?pending=true"
	return httpGet(c, url, &Transaction{})
}

// TransactionReceipt fetches a transaction receipt by its ID.
func (c *Client) TransactionReceipt(id common.Hash) (*TransactionReceipt, error) {
	url := "/transactions/" + id.Hex() + "/receipt"
	return httpGet(c, url, &TransactionReceipt{})
}

// TransactionReceiptAt fetches a transaction receipt by its ID for the given head block ID.
func (c *Client) TransactionReceiptAt(id common.Hash, head common.Hash) (*TransactionReceipt, error) {
	url := "/transactions/" + id.Hex() + "/receipt?revision=" + head.Hex()
	return httpGet(c, url, &TransactionReceipt{})
}

// FilterEvents fetches the event logs that match the given filter.
func (c *Client) FilterEvents(criteriaSet []EventCriteria, filters *LogFilters) ([]EventLog, error) {
	path := "/logs/event"
	events := make([]EventLog, 0)
	request := eventFilter{
		Criteria: &criteriaSet,
	}
	if filters != nil {
		request.Range = filters.filterRange
		request.Options = filters.options
		request.Order = filters.order
	}
	_, err := httpPost(c, path, request, &events)
	if err != nil {
		return nil, err
	}
	return events, nil
}

// FilterTransfers fetches the transfer logs that match the given filter.
func (c *Client) FilterTransfers(criteriaSet []TransferCriteria, filters *LogFilters) ([]TransferLog, error) {
	path := "/logs/transfer"
	transfers := make([]TransferLog, 0)
	request := transferFilter{
		Criteria: &criteriaSet,
	}
	if filters != nil {
		request.Range = filters.filterRange
		request.Options = filters.options
		request.Order = filters.order
	}
	_, err := httpPost(c, path, request, &transfers)
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
	url := "/debug/tracers"
	config := make(map[string]interface{})
	config["OnlyTopCall"] = true
	body := debugTraceClause{
		Config: config,
		Name:   "call",
		Target: fmt.Sprintf("%s/%s/%d", receipt.Meta.BlockID.Hex(), receipt.Meta.TxID.Hex(), len(receipt.Outputs)),
	}
	return httpPost(c, url, body, &TxRevertResponse{})
}

// FeesHistory fetches the fee history for the given block range.
func (c *Client) FeesHistory(newestBlock int64, blockCount int64) (*FeesHistory, error) {
	url := fmt.Sprintf("/fees/history?newestBlock=%d&blockCount=%d", newestBlock, blockCount)
	return httpGet(c, url, &FeesHistory{})
}

func httpGet[T any](c *Client, endpoint string, v *T) (*T, error) {
	req, err := http.NewRequest(http.MethodGet, c.url+endpoint, nil)
	if err != nil {
		return nil, err
	}
	return httpDo(c, req, v)
}

func httpPost[T any](c *Client, path string, body interface{}, v *T) (*T, error) {
	reqBody, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	request, err := http.NewRequest(http.MethodPost, c.url+path, strings.NewReader(string(reqBody)))
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
	statusOK := response.StatusCode >= 200 && response.StatusCode < 300
	if !statusOK {
		return nil, newHttpError(response)
	}
	defer response.Body.Close()

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
