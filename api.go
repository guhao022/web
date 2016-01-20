package web

type Document struct {
	Version  string 	 `json:"version"`
	ID 		 string		 `json:"id,omitempty"`
	Name 	 string      `json:"name,omitempty"`
	Abstruct string      `json:"abstruct"`
	Link     string      `json:"link"`
	Data     interface{} `json:"data,omitempty"`
	Meta     interface{} `json:"meta"`
}
