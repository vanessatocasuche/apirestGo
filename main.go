package main

import (
	"github.com/gin-gonic/gin"
	r "github.com/vanessatocasuche/apirestGo/route"
	_ "strconv"
)

func main() {

	route := gin.Default() //Instance of the server with Framework Gin
	r.EndPointsApp(route)  //Function with endpoints
	route.Run(":9098")     //Run the application

}
