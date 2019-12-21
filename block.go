package navcoind

type Block struct {
	Hash                       string   `json:"hash"`
	Confirmations              uint64   `json:"confirmations"`
	StrippedSize               uint64   `json:"strippedsize"`
	Size                       uint64   `json:"size"`
	Weight                     uint64   `json:"weight"`
	Height                     uint64   `json:"height"`
	Version                    uint32   `json:"version"`
	VersionHex                 string   `json:"versionHex"`
	Merkleroot                 string   `json:"merkleroot"`
	Tx                         []string `json:"tx"`
	TxCount                    uint     `json:"tx_count"`
	ProposalCount              uint     `json:"proposal_count"`
	PaymentRequestCount        uint     `json:"payment_request_count"`
	ProposalVotesCount         uint     `json:"proposal_votes_count"`
	PaymentRequestVotesCount   uint     `json:"payment_request_votes_count"`
	PaymentRequestPayoutsCount uint     `json:"payment_request_payouts_count"`
	Time                       int64    `json:"time"`
	MedianTime                 int64    `json:"mediantime"`
	Nonce                      uint64   `json:"nonce"`
	Bits                       string   `json:"bits"`
	Difficulty                 float64  `json:"difficulty"`
	Chainwork                  string   `json:"chainwork,omitempty"`
	Previousblockhash          string   `json:"previousblockhash"`
	Nextblockhash              string   `json:"nextblockhash"`
}
