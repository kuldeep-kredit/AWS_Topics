package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type datas struct {
	gorm.Model
	Name string
	Age  uint
}
type User struct {
	gorm.Model
	Name string
	Age  uint
}

func main() {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	// [username[:password]@][protocol[(address)]]/dbname[?param1=value1&...&paramN=valueN]

	dsn := "kanwarpal:kanwar1998@tcp(localhost:3306)/lambda?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	//db.AutoMigrate(&datas{})

	// Create
	db.AutoMigrate(&User{})
	user := User{Name: "Jinzhu", Age: 18}

	result := db.Create(&user) // pass pointer of data to Create

	fmt.Println(user.ID) // returns inserted data's primary key

	fmt.Println(result.Error) // returns error
	fmt.Println(result.RowsAffected)

	// // Read
	// var product datas
	// db.First(&product)                       // find product with integer primary key
	// db.First(&product, "name = ?", "Kanwar") // find product with code D42

	// // // Update - update product's price to 200
	// db.Model(&product).Update("Age", 23)
	// // // Update - update multiple fields
	// db.Model(&product).Updates(datas{Age: 24, name: "kanwarpal"}) // non-zero fields
	// db.Model(&product).Updates(map[string]interface{}{"Age": 25, "name": "kanwarpalsunkariya"})

	// // Delete - delete product
	// db.Delete(&product, 1)

}
