package controllers

import (
	"lensent/models/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

type WordController struct{}

// Get /words
func (wc WordController) Get(c *gin.Context) {
	var w repository.WordRepository
	words, err := w.GetAll()
	if err != nil {
		c.AbortWithStatus(404)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(200, words)
	}
}

// Post /words
func (wc WordController) Post(c *gin.Context) {
	var w repository.WordRepository
	word, err := w.CreateModel(c)

	if err != nil {
		c.AbortWithStatus(400)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(201, word)
	}
}

// Get /words/:sentense_id
func (wc WordController) GetBySentenseId(c *gin.Context) {
	sentense_id := c.Params.ByName("sentense_id")
	var w repository.WordRepository
	words, err := w.GetAllBySentense(sentense_id)
	if err != nil {
		c.AbortWithStatus(404)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(200, words)
	}
}

// Get /words/:word_id
func (wc WordController) GetByWordId(c *gin.Context) {
	word_id := c.Params.ByName("word_id")
	var w repository.WordRepository
	word, err := w.GetByWordId(word_id)

	if err != nil {
		c.AbortWithStatus(400)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(200, word)
	}
}

// Put /words/:word_id
func (wc WordController) Put(c *gin.Context) {
	word_id := c.Params.ByName("word_id")
	var w repository.WordRepository
	word, err := w.UpdateByWordId(word_id, c)

	if err != nil {
		c.AbortWithStatus(404)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(200, word)
	}
}

// Delete /sentense/:word_id
func (wc WordController) Delete(c *gin.Context) {
	word_id := c.Params.ByName("word_id")
	var w repository.WordRepository
	if err := w.DeleteByWordId(word_id); err != nil {
		c.AbortWithStatus(403)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"success": "word_id" + word_id + "の語句を削除しました"})
}
