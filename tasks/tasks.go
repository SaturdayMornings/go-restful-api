package tasks

import (
	"bytes"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"os"
)

// Struct representing a single task
type Task struct {
	Id          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

type TaskStorage struct {
	// Tasks stored in memory using a hash map data structure
	// key is id of task, value is task object
	taskMap map[string]Task
}

// Initializes hash map for task storage
func InitTasksStorage() *TaskStorage {
	taskMap := make(map[string]Task)
	return &TaskStorage{
		taskMap,
	}
}

func GetCurrentDir() {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	fmt.Println(dir)
}

func ReadMockData() {
	fmt.Println("\nReadMockData()")

	GetCurrentDir()

	content, err := os.ReadFile("./testData/mock_tasks.json")

	if err != nil {
		fmt.Println("Error")
		fmt.Println(err)
	}
	testRead := bytes.NewReader(content)
	fmt.Print(testRead)

	// return content
}

func (table TaskStorage) LoadExamples() {
	fmt.Println("LoadExamples()")
	s := "test"
	hash := sha1.New()
	hash.Write([]byte(s))
	sha1_hash := hex.EncodeToString(hash.Sum(nil))
	fmt.Println(s, sha1_hash)

	ReadMockData()

}

// Create new task
func (table TaskStorage) Add(title string, description string, status string) error {
	return nil
}

// Get list of all tasks
func (table TaskStorage) List() (map[string]Task, error) {
	return table.taskMap, nil
}
