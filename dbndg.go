package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

func setupRouter(destBase string) *gin.Engine {
	r := gin.Default()

	r.GET("/:user/:repo/*req", func(c *gin.Context) {
		repo := strings.TrimSuffix(c.Param("repo"), ".git")
		var params url.Values = c.Request.URL.Query()

		// switch out base URL (domain), remove ".git" from repo, and pass on everything else
		dest := destBase + "/" + c.Param("user") + "/" + repo + c.Param("req") + "?" + params.Encode()
		log.Printf("redirecting to %s\n", dest)
		c.Redirect(http.StatusMovedPermanently, dest)
	})

	return r
}

var destEnvName string = "DESTINATION_BASE"

func main() {
	destBase := os.Getenv(destEnvName)
	_, err := url.ParseRequestURI(destBase)
	if err != nil {
		log.Fatal(fmt.Errorf("couldn't parse %s as URL: %w", destEnvName, err))
	}

	r := setupRouter(destBase)

	// Listen and Server in 0.0.0.0:8080
	err = r.Run(":8080")
	if err != nil {
		log.Fatal(fmt.Errorf("couldn't start gin: %w", err))
	}
}
