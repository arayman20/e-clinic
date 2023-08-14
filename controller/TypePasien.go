package controller

import (
	"encoding/json"
	"klinik/basemodel"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TypePasienController struct {
	Service TypePasienControllerInterface
}

// Type Pasien godoc
// @Summary Data Master Type Pasien
// @Description Data Master Type Pasien
// @Tags Management Type Pasien
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
// @Router /type-pasien [get]
func (cc *TypePasienController) DataTypePasien(c *gin.Context) {
	c.Writer.Header().Set("Content-Type", "application/json")
	var paginationTypePasien basemodel.TypePasienPaginationBaseModel
	var responseError basemodel.ResponseError
	var responseSuccess basemodel.ResponseRequestPagination
	err := c.ShouldBindQuery(&paginationTypePasien)
	if err != nil {
		responseError.Message = err.Error()
		c.Writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(c.Writer).Encode(responseError)
		return
	}
	code, data, rows, message := cc.Service.DataTypePasien(paginationTypePasien)
	if code != 200 {
		responseError.Message = message
		c.Writer.WriteHeader(code)
		json.NewEncoder(c.Writer).Encode(responseError)
		return
	}
	responseSuccess.Data = data
	responseSuccess.Message = message
	responseSuccess.Size = paginationTypePasien.Size
	responseSuccess.Page = paginationTypePasien.Page
	responseSuccess.Rows = int(rows)
	c.Writer.WriteHeader(code)
	json.NewEncoder(c.Writer).Encode(responseSuccess)

}

// Type Pasien godoc
// @Summary Data Tambah Type Pasien
// @Description Data Tambah Type Pasien
// @Tags Management Type Pasien
// @Accept json
// @Produce json
// @Param body body basemodel.TypePasienRequestInsertBaseModel true "Create user body"
// @security api_key
// @Success 201 {object} basemodel.ResponseError
// @Failure 400 {object} basemodel.ResponseError
// @Failure 401 {object} basemodel.ResponseError
// @Failure 403 {object} basemodel.ResponseError
// @Failure 500 {object} basemodel.ResponseError
// @Router /type-pasien [post]
func (cc *TypePasienController) TambahTypePasien(c *gin.Context) {
	c.Writer.Header().Set("Content-Type", "application/json")
	var typePasienData basemodel.TypePasienRequestInsertBaseModel
	var responseError basemodel.ResponseError
	var responseSuccess basemodel.ResponseError
	err := c.ShouldBindJSON(&typePasienData)
	if err != nil {
		responseError.Message = err.Error()
		c.Writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(c.Writer).Encode(responseError)
		return
	}
	code, message := cc.Service.TambahTypePasien(typePasienData)
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

// Type Pasien godoc
// @Summary Data Edit Type Pasien
// @Description Data Edit Type Pasien
// @Tags Management Type Pasien
// @Accept json
// @Produce json
// @Param id path int true "Type Pasien ID"
// @Param body body basemodel.TypePasienRequestUpdateBaseModel true "Create user body"
// @security api_key
// @Success 201 {object} basemodel.ResponseError
// @Failure 400 {object} basemodel.ResponseError
// @Failure 401 {object} basemodel.ResponseError
// @Failure 403 {object} basemodel.ResponseError
// @Failure 500 {object} basemodel.ResponseError
// @Router /type-pasien/{id} [put]
func (cc *TypePasienController) EditTypePasien(c *gin.Context) {
	c.Writer.Header().Set("Content-Type", "application/json")
	var typePasienData basemodel.TypePasienRequestUpdateBaseModel
	var responseError basemodel.ResponseError
	var responseSuccess basemodel.ResponseError
	err := c.ShouldBindJSON(&typePasienData)
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

	code, message := cc.Service.EditTypePasien(typePasienData, idParam)
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

// Type Pasien godoc
// @Summary Data Delete Type Pasien
// @Description Data Delete Type Pasien
// @Tags Management Type Pasien
// @Accept json
// @Produce json
// @Param id path int true "Type Pasien ID"
// @security api_key
// @Success 201 {object} basemodel.ResponseError
// @Failure 400 {object} basemodel.ResponseError
// @Failure 401 {object} basemodel.ResponseError
// @Failure 403 {object} basemodel.ResponseError
// @Failure 500 {object} basemodel.ResponseError
// @Router /type-pasien/{id} [delete]
func (cc *TypePasienController) DeleteTypePasien(c *gin.Context) {
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

	code, message := cc.Service.DeleteTypePasien(idParam)
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
