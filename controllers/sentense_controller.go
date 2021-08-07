package controllers

import (
	"lensent/models/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SentenseController struct{}

// Get /sentenses
func (sc SentenseController) Get(c *gin.Context) {
	var s repository.SentenseRepository
	sentenses, err := s.GetAll()
	if err != nil {
		c.AbortWithStatus(404)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(200, sentenses)
	}
}

// Post /sentenses
func (sc SentenseController) Post(c *gin.Context) {
	var s repository.SentenseRepository
	sentense, err := s.CreateModel(c)

	if err != nil {
		c.AbortWithStatus(400)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(201, sentense)
	}
}

// Get /sentenses/:sub
func (sc SentenseController) GetBySub(c *gin.Context) {
	sub := c.Params.ByName("sub")
	var s repository.SentenseRepository
	sentenses, err := s.GetAllByUser(sub)
	if err != nil {
		c.AbortWithStatus(404)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(200, sentenses)
	}
}

// Get /sentenses/:sentense_id
func (sc SentenseController) GetBySentenseId(c *gin.Context) {
	sentense_id := c.Params.ByName("sentense_id")
	var s repository.SentenseRepository
	sentense, err := s.GetBySentenseId(sentense_id)

	if err != nil {
		c.AbortWithStatus(400)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(200, sentense)
	}
}

// Put /sentenses/:sentense_id
func (sc SentenseController) Put(c *gin.Context) {
	sentense_id := c.Params.ByName("sentense_id")
	var s repository.SentenseRepository
	sentense, err := s.UpdateBySentenseId(sentense_id, c)

	if err != nil {
		c.AbortWithStatus(404)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(200, sentense)
	}
}

// Delete /sentense/:sentense_id
func (sc SentenseController) Delete(c *gin.Context) {
	sentense_id := c.Params.ByName("sentense_id")
	var s repository.SentenseRepository
	if err := s.DeleteBySentenseId(sentense_id); err != nil {
		c.AbortWithStatus(403)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"success": "sentense_id" + sentense_id + "の英文を削除しました"})
}
