package controller

import (
	"encoding/json"
	"klinik/basemodel"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type MedicalRecordController struct {
	Service MedicalRecordControllerInterface
}

// Medical Record godoc
// @Summary Data Medical Record
// @Description Data Medical Record
// @Tags Management Medical Record
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
// @Router /medical-record [get]
func (cc *MedicalRecordController) DataMedicalRecord(c *gin.Context) {
	c.Writer.Header().Set("Content-Type", "application/json")
	var paginationMedicalRecord basemodel.MedicalRecordPaginationBaseModel
	var responseError basemodel.ResponseError
	var responseSuccess basemodel.ResponseRequestPagination
	err := c.ShouldBindQuery(&paginationMedicalRecord)
	if err != nil {
		responseError.Message = err.Error()
		c.Writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(c.Writer).Encode(responseError)
		return
	}
	code, data, rows, message := cc.Service.DataPesertaBPJS(paginationMedicalRecord)
	if code != 200 {
		responseError.Message = message
		c.Writer.WriteHeader(code)
		json.NewEncoder(c.Writer).Encode(responseError)
		return
	}
	responseSuccess.Data = data
	responseSuccess.Message = message
	responseSuccess.Size = paginationMedicalRecord.Size
	responseSuccess.Page = paginationMedicalRecord.Page
	responseSuccess.Rows = int(rows)
	c.Writer.WriteHeader(code)
	json.NewEncoder(c.Writer).Encode(responseSuccess)

}

// Medical Record godoc
// @Summary Data Tambah Medical Record
// @Description Data Tambah Medical Record
// @Tags Management Medical Record
// @Accept json
// @Produce json
// @Param body body basemodel.MedicalRecordRequestInsertBaseModel true "Create Medical Record body"
// @security api_key
// @Success 201 {object} basemodel.ResponseError
// @Failure 400 {object} basemodel.ResponseError
// @Failure 401 {object} basemodel.ResponseError
// @Failure 403 {object} basemodel.ResponseError
// @Failure 500 {object} basemodel.ResponseError
// @Router /medical-record [post]
func (cc *MedicalRecordController) TambahMedicalRecord(c *gin.Context) {
	c.Writer.Header().Set("Content-Type", "application/json")
	var medicalRecordData basemodel.MedicalRecordRequestInsertBaseModel
	var responseError basemodel.ResponseError
	var responseSuccess basemodel.ResponseError
	err := c.ShouldBindJSON(&medicalRecordData)
	if err != nil {
		responseError.Message = err.Error()
		c.Writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(c.Writer).Encode(responseError)
		return
	}

	code, message := cc.Service.TambahMedicalRecord(medicalRecordData)
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

// Medical Record godoc
// @Summary Data Edit Medical Record
// @Description Data Edit Medical Record
// @Tags Management Medical Record
// @Accept json
// @Produce json
// @Param id path int true "Medical Record ID"
// @Param body body basemodel.PesertaBPJSRequestUpdateBaseModel true "Edit Medical Record body"
// @security api_key
// @Success 201 {object} basemodel.ResponseError
// @Failure 400 {object} basemodel.ResponseError
// @Failure 401 {object} basemodel.ResponseError
// @Failure 403 {object} basemodel.ResponseError
// @Failure 500 {object} basemodel.ResponseError
// @Router /medical-record/{id} [put]
func (cc *MedicalRecordController) EditMedicalRecord(c *gin.Context) {
	c.Writer.Header().Set("Content-Type", "application/json")
	var medicalRecordData basemodel.MedicalRecordRequestUpdateBaseModel
	var responseError basemodel.ResponseError
	var responseSuccess basemodel.ResponseError
	err := c.ShouldBindJSON(&medicalRecordData)
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

	code, message := cc.Service.EditMedicalRecord(medicalRecordData, idParam)
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

// Medical Record godoc
// @Summary Data Delete Medical Record
// @Description Data Delete Medical Record
// @Tags Management Medical Record
// @Accept json
// @Produce json
// @Param id path int true "Medical Record ID"
// @security api_key
// @Success 201 {object} basemodel.ResponseError
// @Failure 400 {object} basemodel.ResponseError
// @Failure 401 {object} basemodel.ResponseError
// @Failure 403 {object} basemodel.ResponseError
// @Failure 500 {object} basemodel.ResponseError
// @Router /medical-record/{id} [delete]
func (cc *MedicalRecordController) DeleteMedicalRecord(c *gin.Context) {
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

	code, message := cc.Service.DeleteMedicalRecord(idParam)
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
