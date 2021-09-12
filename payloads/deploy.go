package payloads

type Deploy struct {
	Chain          string `json:"chain" binding:"required"`
	Name           string `json:"name" binding:"required"`
	Symbol         string `json:"symbol" binding:"required"`
	FromPrivateKey string `json:"fromPrivateKey" binding:"required"`
	//Nonce          int    `json:"nonce" binding:"required"`
	FeeCurrency string `json:"feeCurrency" binding:"required"`
}
