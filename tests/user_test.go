package tests

import (
	"bytes"
	"lensent/server"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestPost(t *testing.T) {
	router := server.Router()
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	body := bytes.NewBufferString("{\"sub\":\"12345\",\"name\":\"test\"}")
	c.Request, _ = http.NewRequest("POST", "/users", body)
	router.ServeHTTP(w, c.Request)

	assert.JSONEq(t, w.Body.String(), "{\"sub\":\"12345\",\"name\":\"test\"}")
	assert.Equal(t, w.Code, 200)
}

func TestUsersRouter(t *testing.T) {
	router := server.Router()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/users", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
}
