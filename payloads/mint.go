package payloads

type Mint struct {
	Chain           string    `json:"chain"`
	TokenID         string    `json:"tokenId"`
	To              string    `json:"to"`
	ContractAddress string    `json:"contractAddress"`
	URL             string    `json:"url"`
	AuthorAddresses *[]string `json:"authorAddresses,omitempty"`
	CashbackValues  *[]string `json:"cashbackValues,omitempty"`
	Index           *int      `json:"index,omitempty"`
	SignatureID     *string   `json:"signatureId,omitempty"`
	Nonce           *int      `json:"nonce,omitempty"`
	Fee             *struct {
		GasLimit *string `json:"gasLimit,omitempty"`
		GasPrice *string `json:"gasPrice,omitempty"`
	} `json:"fee,omitempty"`
}
