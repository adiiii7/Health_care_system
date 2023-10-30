package app

import (
	"healthcareapp/controllers" // Import your controller package
	"healthcareapp/middleware"  // Import your middleware package (if applicable)

	"github.com/gin-gonic/gin"
)

// SetupRoutes defines and configures the application routes.
func SetupRoutes(router *gin.Engine) {
	// Unauthenticated routes (for login and signup)
	unauthenticated := router.Group("/api")
	{
		unauthenticated.POST("/login", controllers.Login)
		unauthenticated.POST("/signup", controllers.Signup)
	}

	// Authenticated routes (protected by middleware, e.g., user authentication)
	authenticated := router.Group("/api")
	authenticated.Use(middleware.AuthMiddleware) // Use your authentication middleware here
	{
		// Appointment management
		authenticated.POST("/appointment/create", controllers.CreateAppointment)
		authenticated.GET("/appointment/:id", controllers.GetAppointment)
		authenticated.PUT("/appointment/approve/:id", controllers.ApproveAppointment)
		authenticated.PUT("/appointment/reject/:id", controllers.RejectAppointment)
		authenticated.PUT("/appointment/update/:id", controllers.UpdateAppointment)
		authenticated.DELETE("/appointment/remove/:id", controllers.RemoveAppointment)
		authenticated.GET("/appointment/status", controllers.GetAppointmentStatus)

		// Admin operations
		authenticated.POST("/admin/diagnostic-center/add", controllers.AddDiagnosticCenter)
		authenticated.PUT("/admin/diagnostic-center/modify/:id", controllers.ModifyDiagnosticCenter)
		authenticated.GET("/admin/diagnostic-center/view/:id", controllers.ViewDiagnosticCenter)
		authenticated.DELETE("/admin/diagnostic-center/remove/:id", controllers.RemoveDiagnosticCenter)
		authenticated.POST("/admin/diagnostic-test/add", controllers.AddDiagnosticTest)
		authenticated.PUT("/admin/diagnostic-test/modify/:id", controllers.ModifyDiagnosticTest)
		authenticated.GET("/admin/diagnostic-test/view/:id", controllers.ViewDiagnosticTest)
		authenticated.DELETE("/admin/diagnostic-test/remove/:id", controllers.RemoveDiagnosticTest)
	}
}
