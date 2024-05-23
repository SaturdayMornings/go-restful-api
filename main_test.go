package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/SaturdayMornings/go-restful-api/tasks"

	"github.com/stretchr/testify/assert"
)

func TestPingRoute(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ping", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "pong", w.Body.String())
}

func TestListTasks(t *testing.T) {
	store := tasks.InitTasksStorage()
	store.LoadExamples()

	tasksHandler := NewTasksHandler(store)
	router := setupRouter()

	// Registering routes
	router.GET("/tasks", tasksHandler.ListTask)
	router.GET("/tasks/:id", tasksHandler.GetTask)
	router.POST("/tasks", tasksHandler.CreateTask)
	router.PUT("/tasks/:id", tasksHandler.UpdateTask)
	router.DELETE("/tasks/:id", tasksHandler.RemoveTask)

	recorder := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", "/tasks", strings.NewReader(""))
	router.ServeHTTP(recorder, req)

	// Assert 200 for response code
	assert.Equal(t, 200, recorder.Code)

	var exampleListResponseEq = `[{"id":1,"title":"Task Title 1","description":"description 1","status":"Pending"},{"id":2,"title":"Task Title 2","description":"description 2","status":"In Progress"},{"id":3,"title":"Task Title 3","description":"description 3","status":"Completed"}]`

	// Compare response body with json data. Test assert equal for body
	assert.Equal(t, exampleListResponseEq, recorder.Body.String())

	var exampleListResponseNotEq = `[{"id":0,"title":"Task Title 1","description":"description 1","status":"Pending"},{"id":2,"title":"Task Title 2","description":"description 2","status":"In Progress"},{"id":3,"title":"Task Title 3","description":"description 3","status":"Completed"}]`

	// Test assert not equal for body
	assert.NotEqual(t, exampleListResponseNotEq, recorder.Body.String())

}
