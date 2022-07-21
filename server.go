package main

import (
	"os"

	"github.com/deyr02/bnzlcrm/auth"
	"github.com/deyr02/bnzlcrm/http"
	"github.com/gin-gonic/gin"
)

//const defaultPort = "8090"
const defaultPort = ":8090"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	// srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	// http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	// http.Handle("/query", srv)

	// log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	// log.Fatal(http.ListenAndServe(":"+port, nil))

	server := gin.Default()
	server.Use(auth.CORSMiddleware())

	account := server.Group("account")
	account.POST("/login", http.LoginHandler())

	server.Run(defaultPort)
}
