package controllers

import (
	"lensent/models/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct{}

// Get /users
func (uc UserController) Get(c *gin.Context) {
	var u repository.UserRepository
	users, err := u.GetAll()
	if err != nil {
		c.AbortWithStatus(404)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(200, users)
	}
}

// Post /users
func (uc UserController) Post(c *gin.Context) {
	var u repository.UserRepository
	user, err := u.CreateModel(c)

	if err != nil {
		c.AbortWithStatus(400)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(201, user)
	}
}

// Get /users/:sub
func (uc UserController) GetBySub(c *gin.Context) {
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
func (uc UserController) Put(c *gin.Context) {
	sub := c.Params.ByName("sub")
	var u repository.UserRepository
	user, err := u.UpdateBySub(sub, c)

	if err != nil {
		c.AbortWithStatus(404)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(200, user)
	}
}

// Delete /users/:sub
func (uc UserController) Delete(c *gin.Context) {
	sub := c.Params.ByName("sub")
	var u repository.UserRepository
	if err := u.DeleteBySub(sub); err != nil {
		c.AbortWithStatus(403)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"success": "sub" + sub + "のユーザーを削除しました"})
}
