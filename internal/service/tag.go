package service

import (
	"github.com/gin-gonic/gin"
	"log"
)

type Tag struct {

}

func NewTag() Tag {
	return Tag{}
}

func (t Tag) Get (c *gin.Context) {
	log.Println("訪問了Tag的Get方法")
	c.JSON(200, gin.H{
		"msg": "訪問了Tag的Get方法",
	})
}
