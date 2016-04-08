package main

import (
	"fmt"
	"github.com/deevatech/manager/runner"
	. "github.com/deevatech/manager/types"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
)

func init() {
	log.Println("Deeva Manager!")
}

func main() {
	router := gin.Default()
	router.POST("/run", handleRunRequest)

	port := os.Getenv("DEEVA_MANAGER_PORT")
	if len(port) == 0 {
		port = "8080"
	}

	log.Printf("Starting in %s mode on port %s\n", gin.Mode(), port)
	host := fmt.Sprintf(":%s", port)
	router.Run(host)
}

func handleRunRequest(c *gin.Context) {
	var run RunParams
	if errParams := c.BindJSON(&run); errParams == nil {
		if result, errRun := runner.Run(run); errRun == nil {
			c.JSON(http.StatusOK, result)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": errRun,
			})
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": errParams,
		})
	}
}
