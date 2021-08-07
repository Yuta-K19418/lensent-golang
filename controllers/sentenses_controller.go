package controllers

import (
	"lensent/models/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SentensesController struct{}

// Get /sentenses
func (sc SentensesController) Get(c *gin.Context) {
	var s repository.SentensesRepository
	p, err := s.GetAll()
	if err != nil {
		c.AbortWithStatus(404)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(200, p)
	}
}

// Post /sentenses
func (sc SentensesController) Post(c *gin.Context) {
	var s repository.SentensesRepository
	p, err := s.CreateModel(c)

	if err != nil {
		c.AbortWithStatus(400)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(201, p)
	}
}

// Get /sentenses/:sub
func (sc SentensesController) GetBySub(c *gin.Context) {
	sub := c.Params.ByName("sub")
	var s repository.SentensesRepository
	p, err := s.GetAllByUser(sub)
	if err != nil {
		c.AbortWithStatus(404)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(200, p)
	}
}

// Get /sentenses/:sentense_id
func (sc SentensesController) GetBySentenseId(c *gin.Context) {
	sentense_id := c.Params.ByName("sentense_id")
	var s repository.SentensesRepository
	user, err := s.GetBySentenseId(sentense_id)

	if err != nil {
		c.AbortWithStatus(400)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(200, user)
	}
}

// Put /sentenses/:sentense_id
func (sc SentensesController) Put(c *gin.Context) {
	sentense_id := c.Params.ByName("sentense_id")
	var s repository.SentensesRepository
	p, err := s.UpdateBySentenseId(sentense_id, c)

	if err != nil {
		c.AbortWithStatus(404)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(200, p)
	}
}

// Delete /sentense/:sentense_id
func (sc SentensesController) Delete(c *gin.Context) {
	sentense_id := c.Params.ByName("sentense_id")
	var s repository.SentensesRepository
	if err := s.DeleteBySentenseId(sentense_id); err != nil {
		c.AbortWithStatus(403)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"success": "sentense_id" + sentense_id + "の英文を削除しました"})
}
