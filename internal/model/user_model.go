package model

type RegsisterRequest struct {
	Username string `json:"username" validate:"required,max=100"`
	FullName string `json:"fullName" validate:"required,max=100"`
	Password string `json:"password" validate:"required,max=100"`
	Email    string `json:"email" validate:"required,email,max=100"`
}


type WebResponse[T any] struct {
	Data   T             `json:"data"`
	// Paging *PageMetadata `json:"paging,omitempty"`
	Errors string        `json:"errors,omitempty"`
}


type RegisterResponse struct {
	FullName      string `json:"name,omitempty"`
	Username  string `json:"username,omitempty"`
	CreatedAt int64  `json:"created_at,omitempty"`
}
