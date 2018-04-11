package common

var (
	TodoList []TodoDetail
	UID      int
	TodoMap  map[int]TodoDetail
)

type (
	// SuccessResponse struct of response when success case
	SuccessResponse struct {
		Status  bool        `json:"status"`
		Message string      `json:"message"`
		Data    interface{} `json:"data"`
	}

	//TodoDetail is struct for
	TodoDetail struct {
		ID     int    `json:"id"`
		Title  string `json:"title"`
		Detail string `json:"detail"`
		Status string `json:"status"`
	}
)

//AddTodo is function for add Todo
func AddTodo(title string, detail string) TodoDetail {
	todoDetail := TodoDetail{
		ID:     UID,
		Title:  title,
		Detail: detail,
		Status: "pending",
	}
	TodoMap[UID] = todoDetail
	UID = UID + 1

	return todoDetail
}

//GetAllToDo is function for return TodoList
func GetAllToDo() (*[]TodoDetail, error) {
	var todoList = []TodoDetail{}
	for _, value := range TodoMap {
		todoList = append(todoList, value)
	}

	return &todoList, nil
}

//GetToDo is function for return Todo
func GetToDo(key int) (*TodoDetail, error) {
	value := TodoMap[key]
	if value.ID == 0 {
		return nil, nil
	}
	return &value, nil
}

//ModifyTodo is function for add Todo
func ModifyTodo(key int, title string, detail string, status string) (*TodoDetail, error) {
	if TodoMap[key].ID == 0 {
		//return error
		return nil, nil
	}
	todoDetail := TodoDetail{
		ID:     key,
		Title:  title,
		Detail: detail,
		Status: status,
	}
	TodoMap[key] = todoDetail

	return &todoDetail, nil
}

//RemoveTodo is function for remove Todo
func RemoveTodo(key int) error {
	delete(TodoMap, key)

	return nil
}
