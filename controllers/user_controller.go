package controllers

import (
	"lensent/models/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct{}

// Get /users
func (pc UserController) Get(c *gin.Context) {
	var u repository.UserRepository
	p, err := u.GetAll()
	if err != nil {
		c.AbortWithStatus(404)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(200, p)
	}
}

// Post /users
func (pc UserController) Post(c *gin.Context) {
	var u repository.UserRepository
	p, err := u.CreateModel(c)

	if err != nil {
		c.AbortWithStatus(400)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(201, p)
	}
}

// Get /users/:sub
func (pc UserController) GetBySub(c *gin.Context) {
	sub := c.Params.ByName("sub")
	var u repository.UserRepository
	user, err := u.GetBySub(sub)

	if err != nil {
		c.AbortWithStatus(400)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(200, user)
	}
}

// Put /users/:sub
func (pc UserController) Put(c *gin.Context) {
	sub := c.Params.ByName("sub")
	var u repository.UserRepository
	p, err := u.UpdateBySub(sub, c)

	if err != nil {
		c.AbortWithStatus(404)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(200, p)
	}
}

// Delete /users/:sub
func (pc UserController) Delete(c *gin.Context) {
	sub := c.Params.ByName("sub")
	var u repository.UserRepository
	if err := u.DeleteBySub(sub); err != nil {
		c.AbortWithStatus(403)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"success": "ID" + sub + "のユーザーを削除しました"})
}
