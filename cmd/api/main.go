package main

import (
	"fmt"
	"leboncointest/endpoint/fizzbuzz"
	"leboncointest/endpoint/statistics"
	"leboncointest/storage"
	"leboncointest/storage/endpointstore/sql"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("unable to read env file")
	}
}

func main() {
	db, err := storage.NewPostgresql()
	if err != nil {
		log.Fatalf("error initializing sql database configuration: %v", err)
	}

	gin.SetMode(gin.ReleaseMode)

	r := gin.Default()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.Use(timeoutMiddleware(time.Second * 30))
	r.Use(statisticsMiddleware(db))

	fizzbuzzGroup := r.Group("/fizzbuzz")
	{
		fizzbuzzGroup.GET("/", fizzbuzz.GetEndpoint())
	}
	r.GET("/statistics", statistics.GetStatistics(db))

	r.Run(":8080")
}

func timeoutMiddleware(timeout time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {
		cCp := c.Copy()
		ch := make(chan struct{})
		go func() {
			cCp.Next()
			ch <- struct{}{}
		}()
		select {
		case <-ch:
			return
		case <-time.After(timeout):
			c.JSON(http.StatusGatewayTimeout, gin.H{"error": "Request timed out"})
			c.Abort()
		}
	}
}

func statisticsMiddleware(db *storage.Repositories) gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := db.EndpointStatistics.CountUp(sql.EndpointStatistics{
			Route:  c.Request.URL.Path,
			Method: c.Request.Method,
		}); err != nil {
			log.New(os.Stderr, fmt.Sprintf("statistics database met an error : %s", err), log.LstdFlags)
		}

		c.Next()
	}
}
