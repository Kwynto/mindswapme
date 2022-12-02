package history

type CreateHistoryDTO struct {
	Lid    string `json:"lid"`
	Mid    string `json:"mid"`
	Action string `json:"action"`
}

type UpdateHistoryDTO struct {
	Id     string `json:"id"`
	Lid    string `json:"lid"`
	Mid    string `json:"mid"`
	Action string `json:"action"`
	Desc   string `json:"desc"`
	Date   string `json:"date"`
	Parse  bool   `json:"parse"`
}
