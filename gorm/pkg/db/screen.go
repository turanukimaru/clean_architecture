package dummydb

import (
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

// Screen テーブル
type Screen struct {
	gorm.Model
	Name     string
	Capacity uint
	Memo     string
}

//IsFirstTouch(startTime time.Time, seats []Seat) (bool, error)
//LetGo(startTime time.Time, seats []Seat) error
//Book(startTime time.Time, seats []Seat) error
//Rollback(startTime time.Time, seats []Seat) error

// github.com/denisenkom/go-mssqldb
func DbAccess() error {
	dsn := "sqlserver://gorm:pass@127.0.0.1:1433?database=gorm"
	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&Screen{})

	// Create
	db.Create(&Screen{Name: "D42", Capacity: 100, Memo: "First Screen"})

	// Read
	var product Screen
	db.First(&product, 1)                 // find product with integer primary key
	db.First(&product, "code = ?", "D42") // find product with code D42

	// Update - update product's price to 200
	db.Model(&product).Update("Price", 200)
	// Update - update multiple fields
	db.Model(&product).Updates(Screen{Name: "Main Screen", Capacity: 200, Memo: "First Screen"}) // non-zero fields
	db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})              // これはエラーになるんだっけ…？

	// Delete - delete product
	deleted := db.Delete(&product, 1)
	return deleted.Error
}
