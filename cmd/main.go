package main

import (
	models "GoContractDeployment/models"
	navigation "GoContractDeployment/navigation"
	"fmt"
	"log"
	"time"
)

func main() {
	dbHost := "rm-uf61g6789v5v1fc19vo.mysql.rds.aliyuncs.com"
	dbPort := "3306"
	dbPass := "NMTTYps!2021"
	dbName := "go_test"

	connection, err := navigation.ConnectSQL(dbHost, dbPort, dbPass, dbName)
	if err != nil {
		log.Fatal("====数据库创建异常====", err)
	}

	//r := chi.NewRouter()
	//r.Use(middleware.Recoverer)
	//r.Use(middleware.Logger)
	query, err := connection.SQL.Query("SELECT * FROM go_test_db")
	if err != nil {
		log.Fatal("====数据库查询异常====", err)

	}

	//defer func(query *sql.Rows) {
	//	err := query.Close()
	//	if err != nil {
	//
	//	}
	//}(query)
	//defer func(SQL *sql.DB) {
	//	err := SQL.Close()
	//	if err != nil {
	//
	//	}
	//}(connection.SQL)

	var post models.Post
	var created []byte

	//遍历查询结果
	for query.Next() {
		err = query.Scan(
			&post.ID,
			&post.Opcode,
			&post.ContractName,
			&post.ContractAddr,
			&post.ContractHash,
			&post.GasUsed,
			&post.GasUST,
			&post.ChainId,
			&created,
			&post.CurrentStatus,
		)
		if err != nil {
			panic(err.Error())
		}
		post.CreatedAt, err = time.Parse("2006-01-02 15:04:05", string(created))
		if err != nil {
			// handle error
		}
		fmt.Println(post)
	}

	//pHandler := ph.NewPostHandler(connection)
	//r.Route("/", func(rt chi.Router) {
	//	rt.Mount("/posts", postRouter(pHandler))
	//})
	//
	//fmt.Println("Server listen at :8005")
	//http.ListenAndServe(":8005", r)
}

// A completely separate router for posts routes
//func postRouter(pHandler *models.Post) http.Handler {
//	r := chi.NewRouter()
//	r.Get("/", pHandler.Fetch)
//	r.Get("/{id:[0-9]+}", pHandler.GetByID)
//	r.Post("/", pHandler.Create)
//	r.Put("/{id:[0-9]+}", pHandler.Update)
//	r.Delete("/{id:[0-9]+}", pHandler.Delete)
//
//	return r
//}
