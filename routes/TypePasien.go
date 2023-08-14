package routes

import (
	"klinik/controller"

	"github.com/gin-gonic/gin"
)

func TypePasienRoutes(routesPoint *gin.RouterGroup, typePasienController *controller.TypePasienController) {
	mainRouter := routesPoint.Group("/type-pasien")
	mainRouter.GET("", typePasienController.DataTypePasien)
	mainRouter.POST("", typePasienController.TambahTypePasien)
	mainRouter.PUT("/:id", typePasienController.EditTypePasien)
	mainRouter.DELETE("/:id", typePasienController.DeleteTypePasien)
}
