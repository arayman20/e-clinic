package unitest

import (
	"klinik/basemodel"
	"klinik/database"
	"klinik/repository"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetDataPoli(t *testing.T) {
	db, err := database.ConnectionORM()
	conn := repository.PoliConn{DB: db, Err: err}

	pagination := basemodel.PoliPaginationBaseModel{
		Page: 1,
		Size: 10,
	}
	data, count, err := conn.DataPoli(pagination)
	assert.NoError(t, err)
	assert.Equal(t, int64(2), count)
	assert.Len(t, data, 2)
}
