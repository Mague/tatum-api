package responses

type TxId struct {
	TxId string `json:"txId" binding:"required"`
}
