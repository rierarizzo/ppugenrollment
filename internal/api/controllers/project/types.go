package project

type Response struct {
	ID          int    `json:"id"`
	Company     int    `json:"company"`
	Description string `json:"description"`
	Schedule    string `json:"schedule"`
}
