package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type testData struct {
	ID string `json:"id"`
	Name string `json:"name"`
	Status string `json:"status"`
	
}