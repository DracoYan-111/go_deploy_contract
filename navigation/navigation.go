package navigation

import (
	"database/sql"
	"fmt"
	"github.com/go-ini/ini"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type DB struct {
	SQL *sql.DB
}

// DBConn ...
var dbConn = &DB{}

func CreateData() (*DB, *ini.File) {
	// 读取配置文件
	var dataBase [5]string
	var configIni = []string{"username", "host", "port", "password", "database"}
	cfg, err := ini.Load("config.ini")
	if err != nil {
		log.Println("<==== 配置文件读取异常 ====>", err)
	} else {
		log.Println("<++++ 配置文件读取成功 ++++>")
	}
	for i := 0; i < len(configIni); i++ {
		dataBase[i] = cfg.Section("database").Key(configIni[i]).String()
		if len(dataBase[i]) > 0 {
			log.Println("<++++ " + configIni[i] + "加载成功 ++++>")
		} else {
			log.Println("<==== " + configIni[i] + "加载异常 ====>")
		}
	}

	//连接到数据库
	connection, err := connectSQL(dataBase[0], dataBase[1], dataBase[2], dataBase[3], dataBase[4])
	if err != nil {
		log.Println("<==== 数据库创建异常 ====>", err)
	} else {
		log.Println("<++++ 数据库创建成功 ++++>")
	}
	return connection, cfg
}

// ConnectSQL 连接到数据库
func connectSQL(user, host, port, pass, dbname string) (*DB, error) {
	// 数据库用户名:数据库密码@tcp(127.0.0.1:3306)/数据库名称/?charset=utf-8
	// 生成数据库连接信息
	dbSource := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8",
		user,
		pass,
		host,
		port,
		dbname,
	)

	// 连接到数据库
	dbData, err := sql.Open("mysql", dbSource)
	if err != nil {
		log.Println("<==== 数据库连接异常 ====>", err)
	} else {
		log.Println("<++++ 数据库连接成功 ++++>")
	}

	// 检查数据库连接
	if err = dbData.Ping(); err != nil {
		log.Println("<==== 数据库检查未通过 ====>", err)
	} else {
		log.Println("<++++ 数据库检查通过 ++++>")
	}

	dbConn.SQL = dbData
	return dbConn, err
}
