package main

import (
	phMysql "GoContractDeployment/handler/http"
	navigation "GoContractDeployment/navigation"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"log"
	"net/http"
)

func main() {
	dbHost := "rm-uf61g6789v5v1fc19vo.mysql.rds.aliyuncs.com"
	dbPort := "3306"
	dbPass := "NMTTYps!2021"
	dbName := "go_test"

	connection, err := navigation.ConnectSQL(dbHost, dbPort, dbPass, dbName)
	if err != nil {
		log.Fatal("====数据库创建异常====", err)
	} else {
		log.Println("++++数据库创建成功++++")
	}

	// 定义传入 HTTP 请求的路由规则
	router := chi.NewRouter()
	// 防止在http请求时报错崩溃
	router.Use(middleware.Recoverer)
	// 记录每次请求的详细信息
	router.Use(middleware.Logger)

	pHandler := phMysql.NewPostHandler(connection)

	//定义了一条新路线Router路由它有两个参数：路由的路径和一个以Router实例作为参数的函数。
	//路径设置为“/”，这意味着该路由将匹配任何传入的 HTTP 请求，其 URL 以正斜杠开头。
	//传递给该方法的函数Route然后将Router实例作为参数并使用该方法定义一个新的子路由器Mount。
	router.Route("/", func(rt chi.Router) {
		rt.Mount("/tianyun", postRouter(pHandler))
	})

	fmt.Println("Server listen at :8005")

	// 将请求路由到正确的处理程序
	if err = http.ListenAndServe(":8005", router); err != nil {
		log.Fatal("====服务启动异常====", err)
	} else {
		log.Println("++++服务启动成功++++")
	}
}

// postRouter 一个完全独立的路由器
func postRouter(pHandler *phMysql.Post) http.Handler {
	r := chi.NewRouter()
	r.Get("/{id:[0-9]+}", pHandler.Fetch)
	return r
}
