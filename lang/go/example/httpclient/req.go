package client

type Object struct {
	Id int `json:"id,omitempty"`
}

type Req struct {
	Id   int    `json:"id,omitempty"`
	Name string `json:"name"`
}

type Signup struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type Page struct {
	PageNo   int `json:"pageNo"`
	PageSize int `json:"pageSize"`
}

type BachUpdate struct {
	Users []*User `json:"users,omitempty"`
}
