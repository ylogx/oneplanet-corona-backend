package utils

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Message(status bool, message string) map[string]interface{} {
	return map[string]interface{} {"status" : status, "message" : message}
}

func Respond(c *gin.Context, data map[string] interface{})  {
	c.JSON(http.StatusOK, data)
}

func RespondOld(w http.ResponseWriter, data map[string] interface{})  {
	w.Header().Add("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(data)
}
