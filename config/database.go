package config

import "github.com/jinzhu/gorm"

/**
* @created by Reja Nur Rochmat
* @email rezanurrochmat3@gmail.com
**/

// DatabaseConn function does initialize database connection
func DatabaseConn() *gorm.DB {
	db, err := gorm.Open("mysql", "reza:reza@tcp(128.0.0.1:3306)/codemi?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		panic(err)
	}

	return db
}
