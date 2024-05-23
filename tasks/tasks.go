package tasks

import (
	"bytes"
	"crypto/rand"
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"sort"
)

var (
	ErrTaskNotFound = errors.New("Task not found")
	counter         = 0
)

// Struct representing a single task
type Task struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

type TaskStorage struct {
	// Tasks stored in memory using a hash map data structure
	// key is id of task, value is task object
	taskMap map[int]Task
}

// Initializes hash map for task storage
func InitTasksStorage() *TaskStorage {
	taskMap := make(map[int]Task)

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

func ParseJson(bytesReader *bytes.Reader) {
	fmt.Println("ParseJson()")
	fmt.Println(bytesReader)
}

func ReadMockData() []Task {
	fmt.Println("\nReadMockData()")

	jsonFile, err := os.Open("./testData/mock_tasks.json")

	if err != nil {
		fmt.Println("Error")
		fmt.Println(err)
	}

	defer jsonFile.Close()

	var testTasks []Task

	byteValue, _ := io.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &testTasks)

	// testRead := bytes.NewReader(jsonFile)
	// fmt.Print(testRead)

	for i := 0; i < len(testTasks); i++ {
		current := testTasks[i]
		fmt.Println(current.Title)
		temp_id := GetId()
		fmt.Println(temp_id)
	}

	return testTasks
}

func getUUID(title string) string {
	hash := sha1.New()
	hash.Write([]byte(title))
	sha1_hash := hex.EncodeToString(hash.Sum(nil))
	return sha1_hash
}

func GetId() string {
	b := make([]byte, 6)
	rand.Read(b)
	s := hex.EncodeToString(b)
	return s
}

func (t TaskStorage) GetNumericId() int {
	counter += 1
	fmt.Println("GetNumericId()")
	fmt.Println(counter)
	return counter
}

// Create new task
func (t TaskStorage) Add(task Task) error {
	// fmt.Println("tasks.go::Add(task Task)\n\n")
	id := task.Id
	// fmt.Println(task)
	t.taskMap[id] = task
	return nil
}

func (table TaskStorage) LoadExamples() {
	var exampleTasks []Task = ReadMockData()
	fmt.Println(exampleTasks)
	for i := 0; i < len(exampleTasks); i++ {
		currentTask := exampleTasks[i]
		currentId := currentTask.Id
		table.taskMap[currentId] = currentTask
		counter = currentId
		fmt.Println(counter)
		fmt.Println(currentId)
	}
}

// Get list of all tasks
func (table TaskStorage) List() ([]Task, error) {
	taskArr := make([]Task, 0, len(table.taskMap))
	for _, taskItem := range table.taskMap {
		taskArr = append(taskArr, taskItem)
	}

	// Ensure listed results are sorted by ascending Id int val
	sort.Slice(taskArr, func(i, j int) bool {
		return taskArr[i].Id < taskArr[j].Id
	})

	return taskArr, nil
}

func (table TaskStorage) Get(id int) (Task, error) {

	if val, ok := table.taskMap[id]; ok {
		return val, nil
	}

	return Task{}, ErrTaskNotFound
}

func (table TaskStorage) Update(id int, task Task) (Task, error) {
	fmt.Println("tasks.go::Update()")
	fmt.Println(task)

	if val, ok := table.taskMap[id]; ok {
		table.taskMap[id] = task
		fmt.Println(val)

		return val, nil
	}

	return Task{}, ErrTaskNotFound
}

func (table TaskStorage) Remove(id int) error {
	if _, ok := table.taskMap[id]; ok {
		delete(table.taskMap, id)
		return nil
	}

	return ErrTaskNotFound
}
