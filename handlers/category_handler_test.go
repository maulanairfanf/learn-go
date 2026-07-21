package handlers

import (
	"bytes"
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

func TestGetCategories_Success(t *testing.T) {

	mockDB, mock, _ := sqlmock.New()

	gormDB, _ := gorm.Open(postgres.New(postgres.Config{
		Conn: mockDB,
	}), &gorm.Config{})

	originalDB := db.DB
	db.DB = gormDB
	defer func() { db.DB = originalDB }()

	rows := sqlmock.NewRows([]string{"id", "name"}).AddRow(1, "Category A").AddRow(2, "Category B")

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "categories"`)).WillReturnRows(rows)

	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request, _ = http.NewRequest("GET", "/category", nil)

	GetCategories(c)

	assert.Equal(t, 200, w.Code)

	var response map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &response)
	assert.Equal(t, float64(200), response["status"])

	data := response["data"].([]interface{})
	assert.Equal(t, 2, len(data))
}

func TestGetCategory_Success(t *testing.T) {

	mockDB, mock, _ := sqlmock.New()

	gormDB, _ := gorm.Open(postgres.New(postgres.Config{
		Conn: mockDB,
	}), &gorm.Config{})

	originalDB := db.DB
	db.DB = gormDB
	defer func() { db.DB = originalDB }()

	rows := sqlmock.NewRows([]string{"id", "name"}).AddRow(1, "Category A")

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "categories"`)).WillReturnRows(rows)

	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request, _ = http.NewRequest("GET", "/category", nil)
	c.Params = []gin.Param{{Key: "id", Value: "1"}}

	GetCategory(c)

	assert.Equal(t, 200, w.Code)

	var response map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &response)
	assert.Equal(t, float64(200), response["status"])

	data := response["data"].(map[string]interface{})
	assert.Equal(t, float64(1), data["id"])
	assert.Equal(t, "Category A", data["name"])
}

func TestCreateCategory_Success(t *testing.T) {
	mockDB, mock, _ := sqlmock.New()
	gormDB, _ := gorm.Open(postgres.New(postgres.Config{Conn: mockDB}), &gorm.Config{})

	originalDB := db.DB
	db.DB = gormDB
	defer func() { db.DB = originalDB }()

	// Mock INSERT — GORM pake RETURNING id
	mock.ExpectBegin()
	mock.ExpectQuery(`INSERT INTO "categories"`).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
	mock.ExpectCommit()

	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	// Kirim JSON body
	body := `{"name":"Elektronik"}`
	c.Request, _ = http.NewRequest("POST", "/category",
		bytes.NewBufferString(body))
	c.Request.Header.Set("Content-Type", "application/json")

	CreateCategory(c)

	assert.Equal(t, 200, w.Code)

	var response map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &response)
	data := response["data"].(map[string]interface{})
	assert.Equal(t, "Elektronik", data["name"])
}

func TestUpdateCategory_Success(t *testing.T) {
	mockDB, mock, _ := sqlmock.New()
	gormDB, _ := gorm.Open(postgres.New(postgres.Config{Conn: mockDB}), &gorm.Config{})

	originalDB := db.DB
	db.DB = gormDB
	defer func() { db.DB = originalDB }()

	// 1. SELECT existing — QuoteMeta cocok sampe WHERE aja
	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "categories" WHERE`)).
		WillReturnRows(sqlmock.NewRows(
			[]string{"id", "name"},
		).AddRow(1, "Category A"))

	// 2. UPDATE — QuoteMeta cocok sampe SET
	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(`UPDATE "categories" SET`)).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	body := `{"name":"Elektronik"}`
	c.Request, _ = http.NewRequest("PUT", "/category/1",
		bytes.NewBufferString(body))
	c.Request.Header.Set("Content-Type", "application/json")
	c.Params = []gin.Param{{Key: "id", Value: "1"}}

	UpdateCategory(c)

	assert.Equal(t, 200, w.Code)

	var response map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &response)
	assert.Equal(t, float64(200), response["status"])

	data := response["data"].(map[string]interface{})
	assert.Equal(t, "Elektronik", data["name"])
}
