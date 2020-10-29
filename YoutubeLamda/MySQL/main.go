package main

import (
    "fmt"
    "github.com/astaxie/beego/orm"
    _ "github.com/go-sql-driver/mysql"
)

func init() {
    orm.RegisterDriver("mysql", orm.DRMySQL)

    orm.RegisterDataBase("default", "mysql", "root:root@/orm_test?charset=utf8")
}

func main() {
	dsn := "kanwarpal:kanwar1998@tcp(localhost:3306)/lambda?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	var empl = []employee{{Name: "Kanwar",Age:}, {Name: "Pal"}}
	db.Create(&empl)
	db.Model(&employee{}).create([]map[string]interface{}{
		{"Name":"kanwarpal","Age":20},
		{"Name":"kanwarpal","Age":20}
	})
	for _, emp := range empl {
		fmt.Println(emp.ID)
	}
}
