package http

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/deyr02/bnzlcrm/graph"
	"github.com/deyr02/bnzlcrm/graph/generated"
	"github.com/gin-gonic/gin"
)

//checking for valid graphql command
func CheckRequest(c *gin.Context, operationName string) bool {
	buf := make([]byte, 1024)
	num, _ := c.Request.Body.Read(buf)
	reqBody := string(buf[0:num])

	splitedString := strings.Split(reqBody, "{")
	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer([]byte(reqBody)))
	return strings.Contains(splitedString[2], operationName)
}

//Login Handler

func LoginHandler() gin.HandlerFunc {
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	return func(c *gin.Context) {
		if !CheckRequest(c, "Login") {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": "Invalid Request",
			})
			return
		}
		srv.ServeHTTP(c.Writer, c.Request)

	}
}
