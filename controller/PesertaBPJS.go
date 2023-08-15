package controller

import (
	"encoding/json"
	"klinik/basemodel"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PesertaBPJSController struct {
	Service PesertaBPJSControllerInterface
}

// Peserta BPJS godoc
// @Summary Data Peserta BPJS WITH Medical Record
// @Description Data Peserta BPJS WITH Medical Record
// @Tags Management Peserta BPJS
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
// @Router /peserta-bpjs [get]
func (cc *PesertaBPJSController) DataPesertaBPJS(c *gin.Context) {
	c.Writer.Header().Set("Content-Type", "application/json")
	var paginationPesertaBPJS basemodel.PesertaBPJSPaginationBaseModel
	var responseError basemodel.ResponseError
	var responseSuccess basemodel.ResponseRequestPagination
	err := c.ShouldBindQuery(&paginationPesertaBPJS)
	if err != nil {
		responseError.Message = err.Error()
		c.Writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(c.Writer).Encode(responseError)
		return
	}
	code, data, rows, message := cc.Service.DataPesertaBPJS(paginationPesertaBPJS)
	if code != 200 {
		responseError.Message = message
		c.Writer.WriteHeader(code)
		json.NewEncoder(c.Writer).Encode(responseError)
		return
	}
	responseSuccess.Data = data
	responseSuccess.Message = message
	responseSuccess.Size = paginationPesertaBPJS.Size
	responseSuccess.Page = paginationPesertaBPJS.Page
	responseSuccess.Rows = int(rows)
	c.Writer.WriteHeader(code)
	json.NewEncoder(c.Writer).Encode(responseSuccess)

}

// Peserta BPJS godoc
// @Summary Data Tambah Peserta BPJS
// @Description Data Tambah Peserta BPJS
// @Tags Management Peserta BPJS
// @Accept json
// @Produce json
// @Param body body basemodel.PesertaBPJSRequestInsertBaseModel true "Create Peserta BPJS body"
// @security api_key
// @Success 201 {object} basemodel.ResponseError
// @Failure 400 {object} basemodel.ResponseError
// @Failure 401 {object} basemodel.ResponseError
// @Failure 403 {object} basemodel.ResponseError
// @Failure 500 {object} basemodel.ResponseError
// @Router /peserta-bpjs [post]
func (cc *PesertaBPJSController) TambahPesertaBPJS(c *gin.Context) {
	c.Writer.Header().Set("Content-Type", "application/json")
	var pesertaBPJSData basemodel.PesertaBPJSRequestInsertBaseModel
	var responseError basemodel.ResponseError
	var responseSuccess basemodel.ResponseError
	err := c.ShouldBindJSON(&pesertaBPJSData)
	if err != nil {
		responseError.Message = err.Error()
		c.Writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(c.Writer).Encode(responseError)
		return
	}

	code, message := cc.Service.TambahPesertaBPJS(pesertaBPJSData)
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

// Peserta BPJS godoc
// @Summary Data Edit Peserta BPJS
// @Description Data Edit Peserta BPJS
// @Tags Management Peserta BPJS
// @Accept json
// @Produce json
// @Param id path int true "Peserta BPJS ID"
// @Param body body basemodel.PesertaBPJSRequestUpdateBaseModel true "Edit Peserta BPJS body"
// @security api_key
// @Success 201 {object} basemodel.ResponseError
// @Failure 400 {object} basemodel.ResponseError
// @Failure 401 {object} basemodel.ResponseError
// @Failure 403 {object} basemodel.ResponseError
// @Failure 500 {object} basemodel.ResponseError
// @Router /peserta-bpjs/{id} [put]
func (cc *PesertaBPJSController) EditPesertaBPJS(c *gin.Context) {
	c.Writer.Header().Set("Content-Type", "application/json")
	var pesertaBPJSData basemodel.PesertaBPJSRequestUpdateBaseModel
	var responseError basemodel.ResponseError
	var responseSuccess basemodel.ResponseError
	err := c.ShouldBindJSON(&pesertaBPJSData)
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

	code, message := cc.Service.EditPesertaBPJS(pesertaBPJSData, idParam)
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

// Peserta BPJS godoc
// @Summary Data Delete Peserta BPJS
// @Description Data Delete Peserta BPJS
// @Tags Management Peserta BPJS
// @Accept json
// @Produce json
// @Param id path int true "Peserta BPJS ID"
// @security api_key
// @Success 201 {object} basemodel.ResponseError
// @Failure 400 {object} basemodel.ResponseError
// @Failure 401 {object} basemodel.ResponseError
// @Failure 403 {object} basemodel.ResponseError
// @Failure 500 {object} basemodel.ResponseError
// @Router /peserta-bpjs/{id} [delete]
func (cc *PesertaBPJSController) DeletePesertaBPJS(c *gin.Context) {
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

	code, message := cc.Service.DeletePesertaBPJS(idParam)
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
