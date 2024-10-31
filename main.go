package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type testdata struct {
	ID string `json:"id"`
	Name string `json:"name"`
	Status string `json:"status"`
}

var testdatas = []testdata{
	{ID: "1", Name: "Cattu", Status: "Denied" },
	{ID: "2", Name: "Sourcetwo", Status: "Approved" },
}

func getDatas(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, testdatas)
}

func main() {
	router := gin.Default()
	router.GET("/testdatas", getDatas)

	router.Run("localhost:3000")
}