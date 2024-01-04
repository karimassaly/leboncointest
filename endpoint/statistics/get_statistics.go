package statistics

import (
	"net/http"

	"leboncointest/storage"

	"github.com/gin-gonic/gin"
)

func GetStatistics(db *storage.Repositories) gin.HandlerFunc {
	return func(c *gin.Context) {
		stats, err := db.EndpointStatistics.GetAll()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "database met an error"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"statistics": stats})
	}
}
