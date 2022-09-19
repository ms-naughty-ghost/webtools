package service

import (
	crand "crypto/rand"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func SecureRandomStr(c *gin.Context) {
	length := c.Query("length")
	if length == "" {
		c.String(http.StatusBadRequest, "expected length parameter")
		return
	}
	length_int, err := strconv.ParseInt(length, 0, 32)
	if err != nil {
		c.String(http.StatusBadRequest, "Enter a number")
		return
	}
	k := make([]byte, length_int)
	if _, err := crand.Read(k); err != nil {
		c.String(http.StatusBadRequest, "Failed to generate")
		return
	}
	c.String(http.StatusOK, fmt.Sprintf("%x", k))
}
