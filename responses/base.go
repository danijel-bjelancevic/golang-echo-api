package responses

// ResponseStatus is used in Response
type ResponseStatus string

const (
	OK    ResponseStatus = "OK"
	ERROR ResponseStatus = "ERROR"
)

// Response represents standard response format
type Response struct {
	Meta Meta        `json:"meta" binding:"required"`
	Data interface{} `json:"data,omitempty"`
}

type Meta struct {
	Status        ResponseStatus `json:"status" enums:"OK,ERROR" binding:"required"`
	StatusCode    int            `json:"status_code" binding:"required"`
	Message       string         `json:"message,omitempty" binding:"required"`
	TransactionID string         `json:"transaction_id" binding:"required"`
}

// BuildResponse will build the Response object and return it
func BuildResponse(data interface{}, statusCode int, msg string, tid string) Response {
	status := OK

	if statusCode >= 400 {
		status = ERROR
	}

	response := Response{
		Meta: Meta{
			Status:        status,
			StatusCode:    statusCode,
			Message:       msg,
			TransactionID: tid,
		},
		Data: data,
	}

	return response
}
