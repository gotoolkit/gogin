package handler

import "github.com/gin-gonic/gin"

func googleNearbySearch(c *gin.Context) {
	c.String(200, "pong")
}