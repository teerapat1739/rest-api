package model

// Posts - represent
type Posts struct {
	Userid int    `json:"userId ,omitempty"`
	ID     int    `json:"id, omitempty"`
	Title  string `json:"title, omitempty"`
	Body   string `json:"body, omitempty"`
}
