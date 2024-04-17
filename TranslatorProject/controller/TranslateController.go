package controller

import (
	"net/http"

	"github.com/AllenOps08/TranslatorBot/helper"
	"github.com/AllenOps08/TranslatorBot/model"
	"github.com/gin-gonic/gin"
)

func Translate(c *gin.Context){
	var translateObj  model.Translate
	err := c.BindJSON(&translateObj)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message":"request can't fulfill",
		})

		return
	}

 	resp, err := helper.Translate(translateObj)
    if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message":"request can't fulfill by google",
		})
		return
	}

	
	c.JSON(http.StatusOK, gin.H{
		"message":"Success",
		"data": string(resp),
	})
 }