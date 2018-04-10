package common

var TodoList []TodoDetail

type (
	// SuccessResponse struct of response when success case
	SuccessResponse struct {
		Status  bool        `json:"status"`
		Message string      `json:"message"`
		Data    interface{} `json:"data"`
	}

	//TodoDetail is struct for
	TodoDetail struct {
		ID     int    `json:"id"`
		Title  string `json:"title"`
		Detail string `json:"detail"`
		Status bool   `json:"status"`
	}
)
