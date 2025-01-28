package tests

import (
	"../services"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestCreateUser(t *testing.T) {
	router := gin.Default()
	router.POST("/users", services.CreateUser)

	body := strings.NewReader(`{"firstname":"John","lastname":"Doe","email":"john.doe@example.com","age":30}`)
	req, _ := http.NewRequest(http.MethodPost, "/users", body)
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestGetUser(t *testing.T) {
	router := gin.Default()
	router.GET("/user/:id", services.GetUser)

	id := "some-uuid"
	req, _ := http.NewRequest(http.MethodGet, "/user/"+id, nil)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestUpdateUser(t *testing.T) {
	router := gin.Default()
	router.PATCH("/user/:id", services.UpdateUser)

	id := "some-uuid"
	body := strings.NewReader(`{"firstname":"Jane","lastname":"Doe","email":"jane.doe@example.com","age":25}`)
	req, _ := http.NewRequest(http.MethodPatch, "/user/"+id, body)
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
}
