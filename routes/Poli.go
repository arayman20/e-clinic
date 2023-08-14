package routes

import (
	"klinik/controller"

	"github.com/gin-gonic/gin"
)

func PoliRoutes(routesPoint *gin.RouterGroup, poliController *controller.PoliController) {
	mainRouter := routesPoint.Group("/poli")
	mainRouter.GET("", poliController.DataPoli)
	mainRouter.POST("", poliController.TambahPoli)
	mainRouter.PUT("/:id", poliController.EditPoli)
	mainRouter.DELETE("/:id", poliController.DeletePoli)
}
