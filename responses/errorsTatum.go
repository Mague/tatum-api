package responses

type ErrorTatum struct {
	StatusCode int    `json:"statusCode" binding:"required"`
	ErrorCode  string `json:"errorCode" binding:"required"`
	Message    string `json:"message" binding:"required"`
}

type Error400 struct {
	ErrorCode  string `json:"errorCode"`
	Message    string `json:"message"`
	StatusCode int    `json:"statusCode"`
	Data       []struct {
		Target struct {
			Property int `json:"property"`
		} `json:"target"`
		Value       int    `json:"value"`
		Property    string `json:"property"`
		Constraints struct {
			Min string `json:"min"`
		} `json:"constraints"`
	} `json:"data"`
}
