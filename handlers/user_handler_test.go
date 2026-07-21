package handlers

import (
	"encoding/json"
	"myapi/db"
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestGetUser_Success(t *testing.T) {

	mockDB, mock, _ := sqlmock.New()

	gormDB, _ := gorm.Open(postgres.New(postgres.Config{
		Conn: mockDB,
	}), &gorm.Config{})

	originalDB := db.DB
	db.DB = gormDB
	defer func() { db.DB = originalDB }()

	rows := sqlmock.NewRows([]string{"id", "username", "password", "email"}).AddRow(1, "maulanairfanf", "password", "maul@gmail.com")

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "users"`)).WillReturnRows(rows)

	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request, _ = http.NewRequest("GET", "/users", nil)

	GetUsers(c)

	assert.Equal(t, 200, w.Code)

	var response map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &response)
	assert.Equal(t, float64(200), response["status"])

	data := response["data"].([]interface{})
	assert.Equal(t, 1, len(data))
}
