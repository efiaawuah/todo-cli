package todo

type Task struct {
	ID          int    `json:"id"`
	DESCRIPTION string `json:"description"`
	DONE        bool   `json:"done"`
}

type Todos struct {
	NEXT_ID int          `json:"next_id"`
	TASKS   map[int]Task `json:"tasks"`
}
