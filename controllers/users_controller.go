package controllers

import (
	"lensent/models/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UsersController struct{}

// Get /users
func (uc UsersController) Get(c *gin.Context) {
	var u repository.UsersRepository
	p, err := u.GetAll()
	if err != nil {
		c.AbortWithStatus(404)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(200, p)
	}
}

// Post /users
func (uc UsersController) Post(c *gin.Context) {
	var u repository.UsersRepository
	p, err := u.CreateModel(c)

	if err != nil {
		c.AbortWithStatus(400)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(201, p)
	}
}

// Get /users/:sub
func (uc UsersController) GetBySub(c *gin.Context) {
	sub := c.Params.ByName("sub")
	var u repository.UsersRepository
	user, err := u.GetBySub(sub)

	if err != nil {
		c.AbortWithStatus(400)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(200, user)
	}
}

// Put /users/:sub
func (uc UsersController) Put(c *gin.Context) {
	sub := c.Params.ByName("sub")
	var u repository.UsersRepository
	p, err := u.UpdateBySub(sub, c)

	if err != nil {
		c.AbortWithStatus(404)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(200, p)
	}
}

// Delete /users/:sub
func (uc UsersController) Delete(c *gin.Context) {
	sub := c.Params.ByName("sub")
	var u repository.UsersRepository
	if err := u.DeleteBySub(sub); err != nil {
		c.AbortWithStatus(403)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"success": "sub" + sub + "のユーザーを削除しました"})
}
