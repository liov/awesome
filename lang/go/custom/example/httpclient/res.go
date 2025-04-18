package client

type User struct {
	Id     int    `json:"id,omitempty"`
	Name   string `json:"name,omitempty"`
	Gender int    `json:"gender,omitempty"`
}

type UserListRes struct {
	Total int64   `json:"total"`
	Users []*User `json:"users,omitempty"`
}
