package main

import (
	"agnos/config"
	"agnos/internal/auth"
	"agnos/internal/delivery/http"
	"agnos/internal/domain"
	"agnos/internal/middleware"
	"agnos/internal/repository"
	"agnos/internal/usecase"

	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDB()
	config.DB.AutoMigrate(&domain.Staff{}, &domain.Hospital{}, &domain.Patient{})

	staffRepo := repository.NewPostgresStaffRepository(config.DB)
	authService := auth.NewAuthService()
	staffUsecase := usecase.NewStaffUsecase(staffRepo, authService)
	staffHandler := http.NewStaffHandler(staffUsecase)

	patientRepo := repository.NewPostgresPatientRepository(config.DB)
	patientUsecase := usecase.NewPatientUsecase(patientRepo, staffRepo)
	patientHandler := http.NewPatientHandler(patientUsecase)

	middleware := middleware.NewMiddleware(authService)

	router := gin.Default()
	router.POST("/staff/create", staffHandler.CreateNewStaff)
	router.POST("/staff/login", staffHandler.LoginStaff)

	authorized := router.Group("/")
	authorized.Use(middleware.ValidateToken)
	{
		authorized.POST("/patient/search", patientHandler.FindPatient)
		authorized.POST("/patient/create", patientHandler.CreateNewPatient)
		authorized.GET("/patient/search/:id", patientHandler.FindPatientByID)

		authorized.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})
	}

	router.Run() // listens on 0.0.0.0:8080 by default
}
