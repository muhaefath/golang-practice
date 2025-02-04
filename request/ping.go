package request

type Ping struct {
	ID int `json:"id" binding:"required"`
}
