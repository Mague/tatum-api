package payloads

type Mint struct {
	Chain           string   `json:"chain"`
	TokenID         string   `json:"tokenId"`
	To              string   `json:"to"`
	ContractAddress string   `json:"contractAddress"`
	URL             string   `json:"url"`
	AuthorAddresses []string `json:"authorAddresses"`
	CashbackValues  []string `json:"cashbackValues"`
	Index           int      `json:"index"`
	SignatureID     string   `json:"signatureId"`
	Nonce           int      `json:"nonce"`
	Fee             struct {
		GasLimit string `json:"gasLimit"`
		GasPrice string `json:"gasPrice"`
	} `json:"fee"`
}
