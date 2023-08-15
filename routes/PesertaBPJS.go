package routes

import (
	"klinik/controller"

	"github.com/gin-gonic/gin"
)

func PesertaBPJSRoutes(routesPoint *gin.RouterGroup, pesertaBPJSController *controller.PesertaBPJSController) {
	mainRouter := routesPoint.Group("/peserta-bpjs")
	mainRouter.GET("", pesertaBPJSController.DataPesertaBPJS)
	mainRouter.POST("", pesertaBPJSController.TambahPesertaBPJS)
	mainRouter.PUT("/:id", pesertaBPJSController.EditPesertaBPJS)
	mainRouter.DELETE("/:id", pesertaBPJSController.DeletePesertaBPJS)
}
