package responses

type ErrorTatum struct {
	StatusCode int    `json:"statusCode" binding:"required"`
	ErrorCode  string `json:"errorCode" binding:"required"`
	Message    string `json:"message" binding:"required"`
}
