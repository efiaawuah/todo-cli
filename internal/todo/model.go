package todo

type Todo struct {
	ID          int    `json:"id"`
	DESCRIPTION string `json:"description"`
	DONE        bool   `json:"done"`
}
