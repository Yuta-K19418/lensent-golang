package controllers

import (
	"bytes"
	"lensent/server"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var Obj struct {
	Sub  string `json:"sub"`
	Name string `json:"Name"`
}

func TestUserPost(t *testing.T) {
	r := server.Router()
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	body := bytes.NewBufferString("{\"sub\":\"testSub\", \"name\":\"testName\"}")
	c.Request, _ = http.NewRequest("POST", "/api/users", body)
	c.Request.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, c.Request)

	assert.NoError(t, c.BindJSON(&Obj))
	assert.Equal(t, "testSub", Obj.Sub)
	assert.Equal(t, "testName", Obj.Name)
	assert.Empty(t, c.Errors)
	assert.Equal(t, "application/json", c.ContentType())
}

func TestUserGet(t *testing.T) {
	router := server.Router()
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request, _ = http.NewRequest("GET", "/api/users", nil)
	router.ServeHTTP(w, c.Request)
	assert.Equal(t, 200, w.Code)
}

func TestUserPut(t *testing.T) {
	TestUserPost(t)
	r := server.Router()
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	body := bytes.NewBufferString("{\"name\":\"testName2\"}")
	c.Request, _ = http.NewRequest("PUT", "/api/users/testSub", body)
	c.Request.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, c.Request)

	assert.NoError(t, c.BindJSON(&Obj))
	assert.Equal(t, "testSub", Obj.Sub)
	assert.Equal(t, "testName2", Obj.Name)
	assert.Empty(t, c.Errors)
	assert.Equal(t, "application/json", c.ContentType())
}

func TestUserDelete(t *testing.T) {
	TestUserPost(t)
	r := server.Router()
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request, _ = http.NewRequest("DELETE", "/api/users/testSub", nil)
	c.Request.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, c.Request)

	assert.Equal(t, "ユーザー(testSub)を削除しました", w.Body.String())

}
