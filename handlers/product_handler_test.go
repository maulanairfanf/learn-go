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

func TestGetProducts_Success(t *testing.T) {
	// 1. Setup mock DB
	mockDB, mock, _ := sqlmock.New()
	defer mockDB.Close()

	gormDB, _ := gorm.Open(postgres.New(postgres.Config{
		Conn: mockDB,
	}), &gorm.Config{})

	// Simpan DB asli, ganti pake mock
	originalDB := db.DB
	db.DB = gormDB
	defer func() { db.DB = originalDB }()

	// 2. Mock query yang bakal dijalanin GORM
	// Step 1: SELECT products
	rows := sqlmock.NewRows([]string{"id", "name", "quantity", "price", "description"}).
		AddRow(1, "Product A", 10, 25000.50, "Desc A").
		AddRow(2, "Product B", 5, 15000.00, "Desc B")

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "products"`)).
		WillReturnRows(rows)

	// Step 2: SELECT product_categories (join table many2many)
	// GORM butuh ini buat tau product_id 1 dan 2 punya category_id berapa
	joinRows := sqlmock.NewRows([]string{"product_id", "category_id"}).
		AddRow(1, 1). // Product 1 → Category 1
		AddRow(2, 1)  // Product 2 → Category 1
	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "product_categories"`)).
		WillReturnRows(joinRows)

	// Step 3: SELECT categories (data category-nya)
	categoryRows := sqlmock.NewRows([]string{"id", "name"}).
		AddRow(1, "Elektronik")
	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "categories"`)).
		WillReturnRows(categoryRows)

	// 3. Setup Gin
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request, _ = http.NewRequest("GET", "/product", nil)

	// 4. Panggil handler
	GetProducts(c)

	// 5. Assert response
	assert.Equal(t, 200, w.Code)

	var response map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &response)
	assert.Equal(t, float64(200), response["status"])

	data := response["data"].([]interface{})
	assert.Equal(t, 2, len(data))
}
