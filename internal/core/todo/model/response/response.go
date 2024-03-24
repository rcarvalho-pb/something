package todo_response

type TodoResponse struct {
	ID          uint32 `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}
