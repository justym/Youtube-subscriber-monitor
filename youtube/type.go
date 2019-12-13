package youtube

//Response struct
type Response struct {
	Kind  string  `json:"kind"`
	Items []Items `json:"items"`
}

//Items struct
type Items struct {
	Kind  string `json:"kind"`
	Id    string `json:"id"`
	Stats Stats  `json:"statistics"`
}

//Stats struct
type Stats struct {
	Views       string `json:"views"`
	Subscribers string `json:"subscriberCount"`
	Videos      string `json:"videoCount"`
}
