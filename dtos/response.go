package dtos

import (
	"github.com/gin-gonic/gin"
)

type Response struct {
	Status  bool
	Message string
	Data gin.H
}
