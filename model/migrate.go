package model

import "fmt"

func migtation() {
	GlobalDB.Set("gorm:table_options", "charset=utf8mb4").AutoMigrate(&User{}, &Task{})
	fmt.Printf("添加表成功")
	// GlobalDB.Model(&Task{}).Association()
}
