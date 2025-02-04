package response

type BaseResponse struct {
	Code    int
	Message string
}

type Ping struct {
	BaseResponse
	ID int
}
