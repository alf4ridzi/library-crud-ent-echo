// using : https://github.com/omniti-labs/jsend for standart json response

package response

type Response struct {
	Status string `json:"status"` // success, fail, error
}

type FailResponse struct {
	Response
	Data any `json:"data"`
}

type ErrorResponse struct {
	Response
	Message string `json:"message"`
}

type SuccessResponse struct {
	Response
	Data any `json:"data"`
}
