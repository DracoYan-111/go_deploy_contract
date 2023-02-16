package navigation

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type DB struct {
	SQL *sql.DB
}

// DBConn ...
var dbConn = &DB{}

// ConnectSQL 连接到数据库
func ConnectSQL(host, port, pass, dbname string) (*DB, error) {
	// 数据库用户名:数据库密码@tcp(127.0.0.1:3306)/数据库名称/?charset=utf-8
	// 生成数据库连接信息
	dbSource := fmt.Sprintf(
		"root:%s@tcp(%s:%s)/%s?charset=utf8",
		pass,
		host,
		port,
		dbname,
	)
	println(dbSource)

	// 连接到数据库
	db, err := sql.Open("mysql", dbSource)
	if err != nil {
		log.Fatal("====数据库连接异常====", err)
	}

	// 检查数据库连接
	if err = db.Ping(); err != nil {
		log.Fatal("====数据库检查未通过====", err)
	}

	dbConn.SQL = db
	return dbConn, err
}
