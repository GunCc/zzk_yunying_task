package initialize

import (
	"database/sql"
	"fmt"
)

// 创建初始化数据库
func InitDatabase(dsn string, createSql string) error {
	// 原生sql
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return err
	}

	defer func(db *sql.DB) {
		err = db.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(db)

	if err = db.Ping(); err != nil {
		return err
	}
	_, err = db.Exec(createSql)
	return err
}
