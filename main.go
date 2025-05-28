package main

import (
	"gin-api-test/internal/health"

	docs "gin-api-test/infrastructure/api"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @contact.name   API Support
// @contact.url    http://www.teste.com.br
// @contact.email  support@teste.com.br
func main() {
	docs.SwaggerInfo.Title = "Swagger Example API"
	docs.SwaggerInfo.Description = "This is a sample server Petstore server."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.BasePath = "/api"

	// Initialize the Gin router
	r := gin.Default()
	r.ForwardedByClientIP = false
	//r.SetTrustedProxies([]string{"127.0.0.1", "10.160.0.0/24"})

	// Initialize the health service and controller
	healthSvc := health.NewHealthService()
	healthCtrl := health.NewHealthController(healthSvc)

	api := r.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			//healthCtrl.Routes(v1)
			helpers := v1.Group("/helpers")
			{
				healthCtrl.RegisterRoutes(helpers) // acessa:  http://localhost:8080/api/v1/helpers/health
			}
		}

	}

	// Set up the routes
	// apiv1 := r.Group("/v1")
	// healthCtrl.Routes(apiv1)

	// apiv2 := r.Group("/v2")
	// healthCtrl.Routes(apiv2)

	// Start the server
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.Run(":8080")
}
