package routes

import (
	"klinik/controller"

	"github.com/gin-gonic/gin"
)

func MedicalRecordRoutes(routesPoint *gin.RouterGroup, medicalRecordController *controller.MedicalRecordController) {
	mainRouter := routesPoint.Group("/medical-record")
	mainRouter.GET("", medicalRecordController.DataMedicalRecord)
	mainRouter.POST("", medicalRecordController.TambahMedicalRecord)
	mainRouter.PUT("/:id", medicalRecordController.EditMedicalRecord)
	mainRouter.DELETE("/:id", medicalRecordController.DeleteMedicalRecord)
}
