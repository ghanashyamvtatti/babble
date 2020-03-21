package dtos

import "github.com/gin-gonic/gin/render"

type Response struct {
	Status  bool
	Message string
	Data render.JSON
}
