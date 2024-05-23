package main

import (
	"fmt"
	"net/http"

	"strconv"

	"github.com/SaturdayMornings/go-restful-api/tasks"
	"github.com/gin-gonic/gin"
)

func main() {
	router := setupRouter()

	// Initialize TaskStorage for in memory storage of tasks via map
	store := tasks.InitTasksStorage()

	// Load several mock Task objects from local JSON file
	store.LoadExamples()

	// Instantiate task Handler
	tasksHandler := NewTasksHandler(store)

	// Registering routes
	router.GET("/tasks", tasksHandler.ListTask)
	router.GET("/tasks/:id", tasksHandler.GetTask)
	router.POST("/tasks", tasksHandler.CreateTask)
	router.PUT("/tasks/:id", tasksHandler.UpdateTask)
	router.DELETE("/tasks/:id", tasksHandler.RemoveTask)
	// Listen and Server in 0.0.0.0:8080
	router.Run(":8080")
}

// Struct with a store as an attribute for use with each handler method
type TasksHandler struct {
	store taskStore
}

// Interface defining CRUD operations for use with data storage structure
type taskStore interface {
	// Create Task
	Add(task tasks.Task) error
	// Read single Task by id
	Get(id int) (tasks.Task, error)
	// PUT request for task by id
	Update(id int, task tasks.Task) (tasks.Task, error)
	// Delete task by id
	Remove(id int) error
	// List Tasks
	List() ([]tasks.Task, error)
	GetNumericId() int
}

// POST /tasks handler. Allows for creation of tasks
func (h TasksHandler) CreateTask(c *gin.Context) {
	var task tasks.Task

	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get id that is based off of total item count
	id := h.store.GetNumericId()
	task.Id = id

	// Tasks by default have Status of Pending
	task.Status = "Pending"

	// Add task to the map storage solution
	h.store.Add(task)

	// Return status indicating task was successfully created along with created task
	c.JSON(http.StatusOK, gin.H{"status": "success", "task": task})
}

// GET /tasks/{id}
func (h TasksHandler) GetTask(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	task, err := h.store.Get(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	}

	c.JSON(200, task)
}

// PUT /tasks/{id}
func (h TasksHandler) UpdateTask(c *gin.Context) {
	// Get request body and convert it to recipes.Recipe
	var task tasks.Task

	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	task.Id = id

	updatedTask, err := h.store.Update(id, task)

	if err != nil {
		if err == tasks.ErrTaskNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "task": updatedTask})
}

// handler for DELETE /tasks/{id}
func (h TasksHandler) RemoveTask(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		fmt.Println("strconv.Atoi() error")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	deleteErr := h.store.Remove(id)

	if deleteErr != nil {
		if deleteErr == tasks.ErrTaskNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": deleteErr.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": deleteErr.Error()})
		return
	}

	// return success payload
	c.JSON(http.StatusOK, gin.H{"status": "success"})
}

// GET /tasks
func (h TasksHandler) ListTask(c *gin.Context) {
	res, err := h.store.List()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(200, res)
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func NewTasksHandler(t taskStore) *TasksHandler {
	return &TasksHandler{
		store: t,
	}
}

func setupRouter() *gin.Engine {
	r := gin.New()
	r.Use(CORSMiddleware())
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	// Generic test for a GET route handler with a JSON response
	r.GET("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  "success",
			"message": "test data",
		})
	})

	return r
}
