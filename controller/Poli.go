package controller

import (
	"encoding/json"
	"klinik/basemodel"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PoliController struct {
	Service PoliControllerInterface
}

// Poli godoc
// @Summary Data Master Poli
// @Description Data Master Poli
// @Tags Management Poli
// @Accept json
// @Produce json
// @Param page query int false "Page number" default(1)
// @Param size query int false "Number of items per page" default(10)
// @security api_key
// @Success 200 {object} basemodel.ResponseRequestPagination
// @Failure 400 {object} basemodel.ResponseError
// @Failure 401 {object} basemodel.ResponseError
// @Failure 403 {object} basemodel.ResponseError
// @Failure 500 {object} basemodel.ResponseError
// @Router /poli [get]
func (cc *PoliController) DataPoli(c *gin.Context) {
	c.Writer.Header().Set("Content-Type", "application/json")
	var paginationPoli basemodel.PoliPaginationBaseModel
	var responseError basemodel.ResponseError
	var responseSuccess basemodel.ResponseRequestPagination
	err := c.ShouldBindQuery(&paginationPoli)
	if err != nil {
		responseError.Message = err.Error()
		c.Writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(c.Writer).Encode(responseError)
		return
	}
	code, data, rows, message := cc.Service.DataPoli(paginationPoli)
	if code != 200 {
		responseError.Message = message
		c.Writer.WriteHeader(code)
		json.NewEncoder(c.Writer).Encode(responseError)
		return
	}
	responseSuccess.Data = data
	responseSuccess.Message = message
	responseSuccess.Size = paginationPoli.Size
	responseSuccess.Page = paginationPoli.Page
	responseSuccess.Rows = int(rows)
	c.Writer.WriteHeader(code)
	json.NewEncoder(c.Writer).Encode(responseSuccess)

}

// Poli godoc
// @Summary Data Tambah Poli
// @Description Data Tambah Poli
// @Tags Management Poli
// @Accept json
// @Produce json
// @Param body body basemodel.RequestInsertBaseModel true "Create user body"
// @security api_key
// @Success 201 {object} basemodel.ResponseError
// @Failure 400 {object} basemodel.ResponseError
// @Failure 401 {object} basemodel.ResponseError
// @Failure 403 {object} basemodel.ResponseError
// @Failure 500 {object} basemodel.ResponseError
// @Router /poli [post]
func (cc *PoliController) TambahPoli(c *gin.Context) {
	c.Writer.Header().Set("Content-Type", "application/json")
	var poliData basemodel.RequestInsertBaseModel
	var responseError basemodel.ResponseError
	var responseSuccess basemodel.ResponseError
	err := c.ShouldBindJSON(&poliData)
	if err != nil {
		responseError.Message = err.Error()
		c.Writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(c.Writer).Encode(responseError)
		return
	}
	code, message := cc.Service.TambahPoli(poliData)
	if code != 200 {
		responseError.Message = message
		c.Writer.WriteHeader(code)
		json.NewEncoder(c.Writer).Encode(responseError)
		return
	}

	responseSuccess.Message = message
	c.Writer.WriteHeader(code)
	json.NewEncoder(c.Writer).Encode(responseSuccess)

}

// Poli godoc
// @Summary Data Edit Poli
// @Description Data Edit Poli
// @Tags Management Poli
// @Accept json
// @Produce json
// @Param id path int true "Poli ID"
// @Param body body basemodel.RequestUpdateBaseModel true "Create user body"
// @security api_key
// @Success 201 {object} basemodel.ResponseError
// @Failure 400 {object} basemodel.ResponseError
// @Failure 401 {object} basemodel.ResponseError
// @Failure 403 {object} basemodel.ResponseError
// @Failure 500 {object} basemodel.ResponseError
// @Router /poli/{id} [put]
func (cc *PoliController) EditPoli(c *gin.Context) {
	c.Writer.Header().Set("Content-Type", "application/json")
	var poliData basemodel.RequestUpdateBaseModel
	var responseError basemodel.ResponseError
	var responseSuccess basemodel.ResponseError
	err := c.ShouldBindJSON(&poliData)
	if err != nil {
		responseError.Message = err.Error()
		c.Writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(c.Writer).Encode(responseError)
		return
	}

	idParam, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		responseError.Message = err.Error()
		c.Writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(c.Writer).Encode(responseError)
		return
	}

	code, message := cc.Service.EditPoli(poliData, idParam)
	if code != 200 {
		responseError.Message = message
		c.Writer.WriteHeader(code)
		json.NewEncoder(c.Writer).Encode(responseError)
		return
	}

	responseSuccess.Message = message
	c.Writer.WriteHeader(code)
	json.NewEncoder(c.Writer).Encode(responseSuccess)

}

// Poli godoc
// @Summary Data Delete Poli
// @Description Data Delete Poli
// @Tags Management Poli
// @Accept json
// @Produce json
// @Param id path int true "Poli ID"
// @security api_key
// @Success 201 {object} basemodel.ResponseError
// @Failure 400 {object} basemodel.ResponseError
// @Failure 401 {object} basemodel.ResponseError
// @Failure 403 {object} basemodel.ResponseError
// @Failure 500 {object} basemodel.ResponseError
// @Router /poli/{id} [delete]
func (cc *PoliController) DeletePoli(c *gin.Context) {
	c.Writer.Header().Set("Content-Type", "application/json")
	var responseError basemodel.ResponseError
	var responseSuccess basemodel.ResponseError
	idParam, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		responseError.Message = err.Error()
		c.Writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(c.Writer).Encode(responseError)
		return
	}

	code, message := cc.Service.DeletePoli(idParam)
	if code != 200 {
		responseError.Message = message
		c.Writer.WriteHeader(code)
		json.NewEncoder(c.Writer).Encode(responseError)
		return
	}

	responseSuccess.Message = message
	c.Writer.WriteHeader(code)
	json.NewEncoder(c.Writer).Encode(responseSuccess)

}
