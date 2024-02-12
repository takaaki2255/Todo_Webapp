package db

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

const DBName = "model.db"

type Todo struct {
	gorm.Model
	Text   string
	Status string
}

// エラーチェック
func CheckError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func init() {
	db := openDB()
	db.AutoMigrate(&Todo{})
}

// データベースを開く
func openDB() (db *gorm.DB) {
	db, err := gorm.Open(sqlite.Open(DBName), &gorm.Config{})
	CheckError(err)
	return db
}

// Todo追加
func AddTodo(text string, status string) {
	db := openDB()
	todo := Todo{Text: text, Status: status}
	db.Create(&todo)
}

// 全Todo取得
func SelectAllTodo() []Todo {
	var todos []Todo
	db := openDB()
	db.Find(&todos)
	return todos
}
