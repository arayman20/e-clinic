package main

import (
	"fmt"
	"klinik/controller"
	"klinik/database"
	_ "klinik/docs"
	"klinik/domain"
	"klinik/migration"
	"klinik/repository"
	"klinik/routes"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title E-Clinic
// @description Documentation API E-Clinic
// @version 1.0
// @BasePath /api/e-clinic/v1
// @securityDefinitions.apikey api_key
// @in header
// @name Authorization
func main() {
	if os.Getenv("DOCKER") == "" {
		godotenv.Load()
	}
	db, err := database.ConnectionORM()
	if os.Getenv("MIGRATION") == "ACTIVE" {
		migration.Migrate(db)
	}

	poliRepo := repository.PoliConn{DB: db, Err: err}
	typePasienRepo := repository.TypePasienConn{DB: db, Err: err}
	pesertaBPJSRepo := repository.PesertaBPJSConn{DB: db, Err: err}
	medicalRecordRepo := repository.MedicalRecordConn{DB: db, Err: err}

	poliService := domain.PoliDomain{PoliRepo: &poliRepo}
	typePasienService := domain.TypePasienDomain{TypePasienRepo: &typePasienRepo}
	pesertaBPJSService := domain.PesertaBPJSDomain{PesertaBPJSRepo: &pesertaBPJSRepo}
	medicalRecordService := domain.MedicalRecordDomain{MedicalRecordRepo: &medicalRecordRepo}

	poliController := controller.PoliController{Service: &poliService}
	typePasienController := controller.TypePasienController{Service: &typePasienService}
	pesertaBPJSController := controller.PesertaBPJSController{Service: &pesertaBPJSService}
	medicalRecordController := controller.MedicalRecordController{Service: &medicalRecordService}

	router := gin.Default()
	router.Use(gin.Recovery())
	if os.Getenv("SWAGGER") == "1" {
		swaggerURL := ginSwagger.URL(fmt.Sprintf("http://%s:%s/swagger/doc.json", os.Getenv("IP_SWAGGER"), os.Getenv("PORT_SWAGGER")))
		router.GET("/swagger/*any", controller.CorsOriginsMiddleware(), ginSwagger.WrapHandler(swaggerFiles.Handler, swaggerURL))
	}
	routes.PoliRoutes(router.Group("/api/e-clinic/v1", controller.CorsOriginsMiddleware()), &poliController)
	routes.TypePasienRoutes(router.Group("/api/e-clinic/v1", controller.CorsOriginsMiddleware()), &typePasienController)
	routes.PesertaBPJSRoutes(router.Group("/api/e-clinic/v1", controller.CorsOriginsMiddleware()), &pesertaBPJSController)
	routes.MedicalRecordRoutes(router.Group("/api/e-clinic/v1", controller.CorsOriginsMiddleware()), &medicalRecordController)
	router.Run(":" + os.Getenv("PORT_SWAGGER"))
}
