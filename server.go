package main

import (
	"os"

	"github.com/deyr02/bnzlcrm/auth"
	"github.com/deyr02/bnzlcrm/http"
	"github.com/gin-gonic/gin"
)

//const defaultPort = "8090"
const defaultPort = ":8090"

// func playgroundHandler() gin.HandlerFunc {
// 	//h := playground.Handler("GraphQL playground", "/query")
// 	h := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))
// 	return func(c *gin.Context) {
// 		h.ServeHTTP(c.Writer, c.Request)
// 	}
// }

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

	userRole := server.Group("userrole")
	userRole.Use(auth.AuthMiddleware())
	userRole.GET("/getalluserrole", http.GetAllUserRoleHandler())
	userRole.GET("/getuserrole", http.GetUserRoleByID())
	userRole.POST("/", http.AddNewUserRole())
	userRole.POST("/modifyuserrole", http.ModifyUserRole())
	userRole.POST("/deleteuserrole", http.DeleteUserRole())

	user := server.Group("user")
	user.Use(auth.AuthMiddleware())
	user.GET("/getalluser", http.GetAllUserHandler())
	user.GET("/getuser", http.GetUserByID())
	user.POST("/", http.AddNewUser())
	user.POST("/modifyuser", http.ModifyUser())
	user.POST("/deleteuser", http.DeleteUser())

	metauser := server.Group("userform")
	metauser.Use(auth.AuthMiddleware())
	metauser.GET("/metauser", http.GetAllMetaUserField())
	metauser.GET("/metauserformfield", http.GetMetaUserFieldByID())
	metauser.POST("/addnewfield", http.AddNewMetaUserField())
	metauser.POST("/modifyfield", http.ModifyMetaUserField())
	metauser.POST("/deletefield", http.DeleteMetaUserField())

	activity := server.Group("activity")
	activity.Use(auth.AuthMiddleware())
	activity.GET("/getallactivities", http.GetAllActivityHandler())
	activity.GET("/getactivity", http.GetActivityByID())
	activity.POST("/", http.AddNewActivity())
	activity.POST("/modifyactivity", http.ModifyActivity())
	activity.POST("/deleteactivity", http.DeleteActivity())

	//user.Use(auth.AuthMiddleware())
	server.Run(defaultPort)
}
