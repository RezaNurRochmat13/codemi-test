package config

import "github.com/jinzhu/gorm"
import _ "github.com/go-sql-driver/mysql"

/**
* @created by Reja Nur Rochmat
* @email rezanurrochmat3@gmail.com
**/

// DatabaseConn function does initialize database connection
func DatabaseConn() *gorm.DB {
	db, err := gorm.Open("mysql", "reza:reza@tcp(127.0.0.1:3306)/codemi?charset=utf8&parseTime=true&loc=Local")

	if err != nil {
		panic(err)
	}

	return db
}
