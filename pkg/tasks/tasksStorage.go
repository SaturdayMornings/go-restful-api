package tasks

type TasksStore struct {
	// Tasks stored in memory using a hash map data structure
	// key is id of task, value is task object
	tasksTable map[string]Task
}

// Initializes hash map for task storage
func initTaskStore() *TasksStore {
	tasksTable := make(map[string]Task)
	return &TasksStore{
		tasksTable,
	}
}
