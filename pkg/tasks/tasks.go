package tasks

// Struct representing a single task
type Task struct {
	Id          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

type TasksStorage struct {
	// Tasks stored in memory using a hash map data structure
	// key is id of task, value is task object
	tasksTable map[string]Task
}

// Initializes hash map for task storage
func initTasksStorage() *TasksStorage {
	tasksTable := make(map[string]Task)
	return &TasksStorage{
		tasksTable,
	}
}
