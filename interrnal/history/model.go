package history

type History struct {
	Id     string `json:"id,omitempty"`
	Lid    string `json:"lid,omitempty"`
	Mid    string `json:"mid,omitempty"`
	Action string `json:"action,omitempty"`
	Desc   string `json:"desc,omitempty"`
	Date   string `json:"date,omitempty"`
	Parse  bool   `json:"parse,omitempty"`
}
