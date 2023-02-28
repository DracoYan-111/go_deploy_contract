package navigation

import (
	"GoContractDeployment/utils"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type DB struct {
	SQL *sql.DB
}

var dbConn = &DB{}

// CreateData Read configuration file
func CreateData() *DB {
	dataBase, _ := utils.ConfigurationLoading("database", []string{"username", "host", "port", "password", "database"})
	connection, err := connectSQL(dataBase[0], dataBase[1], dataBase[2], dataBase[3], dataBase[4])
	if err != nil {
		log.Println("<==== navigation:Database creation exception ====>", err)
	} else {
		log.Println("<++++ navigation:Database created successfully ++++>")
	}
	return connection
}

// ConnectSQL Connect to database
func connectSQL(user, host, port, pass, dbname string) (*DB, error) {
	// Generate database connection information
	dbSource := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8",
		user,
		pass,
		host,
		port,
		dbname,
	)

	// connect to database
	dbData, err := sql.Open("mysql", dbSource)
	if err != nil {
		log.Println("<==== navigation:Database connection exception ====>", err)
	} else {
		log.Println("<++++ navigation:Database connection succeeded ++++>")
	}

	// Check database connection
	if err = dbData.Ping(); err != nil {
		log.Println("<==== navigation:Database check failed ====>", err)
	} else {
		log.Println("<++++ navigation:Database check passed ++++>")
	}

	dbConn.SQL = dbData
	return dbConn, err
}
