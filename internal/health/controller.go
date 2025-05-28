package health

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type HealthController struct {
	// Add any necessary fields here
	healthSvc IHealthService
}

func NewHealthController(svc IHealthService) *HealthController {
	return &HealthController{
		// Initialize any necessary fields here
		healthSvc: svc,
	}
}

func (hc *HealthController) RegisterRoutes(rg *gin.RouterGroup) {
	rg.GET("/health", hc.getHealthCheck)
}

// getHealthCheck godoc
// @Summary ping example
// @Schemes
// @Description do ping
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {object} HealthCheckResponse
// @Failure 500
// @Router /v1/helpers/health [get]
func (hc *HealthController) getHealthCheck(c *gin.Context) {
	var log = logrus.New()

	// Call the health service to perform the health check
	status, err := hc.healthSvc.HealthCheck()
	if err != nil {
		// Handle error (e.g., log it, return an error response)
		log.Errorf(fmt.Sprintf("Health check failed: %+v", err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	// Return the health status as a JSON response
	log.Infof(fmt.Sprintf("Health check status: %+v", status))
	//c.JSON(http.StatusOK, gin.H{"status": status})
	c.JSON(http.StatusOK, status)
}
