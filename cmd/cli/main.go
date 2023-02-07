package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	rkboot "github.com/rookie-ninja/rk-boot/v2"
	rkgin "github.com/rookie-ninja/rk-gin/v2/boot"
)

//	@title			Swagger service API
//	@version		1.0
//	@description	Service web server.
//	@termsOfService	http://swagger.io/terms/

//	@securityDefinitions.basic	BasicAuth

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

// @license.name	Apache 2.0
// @license.url	http://www.apache.org/licenses/LICENSE-2.0.html
func main() {
	// Create a new boot instance.
	boot := rkboot.NewBoot()

	// Add shutdown hook function
	boot.AddShutdownHookFunc("shutdown-hook", func() {
		fmt.Println("shutting down")
	})

	// Register handler
	entry := rkgin.GetGinEntry("web-service")
	entry.Router.GET("/v1/greeter", Greeter)

	// Bootstrap
	boot.Bootstrap(context.TODO())

	boot.WaitForShutdownSig(context.TODO())
}

// Greeter handler
//
//	@Summary	Greeter
//	@Id			1
//	@Tags		Hello
//	@version	1.0
//	@Param		name	query	string	true	"name"
//	@produce	application/json
//	@Success	200	{object}	GreeterResponse
//	@Router		/v1/greeter [get]
func Greeter(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, &GreeterResponse{
		Message: fmt.Sprintf("Hello %s!", ctx.Query("name")),
	})
}

type GreeterResponse struct {
	Message string
}
