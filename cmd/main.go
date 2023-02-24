package main

import (
	"GoContractDeployment/cron"
	phMysql "GoContractDeployment/handler/http"
	"GoContractDeployment/navigation"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-ini/ini"
	"log"
	"net"
	"net/http"
)

func main() {
	connection, cfg := navigation.CreateData()
	phHandler := phMysql.NewJobHandler(connection)

	//one, _ := phHandler.Repo.GetOne()
	//
	//log.Println(one.GasUsed, "++++++++++++++++++++")
	//
	//a := internal.GetBnbToUsdt(big.NewInt(one.GasUsed))
	//
	//log.Panicln(a, "++++++++++++++++++++")

	cron.UpdateLibrary(cfg, phHandler)
	cron.ReturnStatus(cfg, phHandler)
	basicConfiguration(phHandler, cfg)
	//// 定义传入 HTTP 请求的路由规则
	//router := chi.NewRouter()
	//// 防止在http请求时报错崩溃
	//router.Use(middleware.Recoverer)
	//// 记录每次请求的详细信息
	//router.Use(middleware.Logger)
	//
	////定义了一条新路线Router路由它有两个参数：路由的路径和一个以Router实例作为参数的函数。
	////路径设置为“/”，这意味着该路由将匹配任何传入的 HTTP 请求，其 URL 以正斜杠开头。
	////传递给该方法的函数Route然后将Router实例作为参数并使用该方法定义一个新的子路由器Mount。
	//router.Route("/", func(rt chi.Router) {
	//	rt.Mount("/tianyun", postRouter(phHandler))
	//})
	//fmt.Println("Server listen at :8005")
	//
	//// 启动定时任务
	//log.Println("++++定时任务启动成功++++")
	//
	//// 将请求路由到正确的处理程序
	//if err := http.ListenAndServe(":"+cfg.Section("server").Key("port").String(), router); err != nil {
	//	log.Println("====服务启动异常====", err)
	//} else {
	//	log.Println("++++服务启动成功++++")
	//}
}

// postRouter 一个完全独立的路由器
func postRouter(phHandler *phMysql.CreateTask) http.Handler {
	router := chi.NewRouter()
	//router.Get("/{id:[0-9]+}", phHandler.Operate)
	router.Post("/", phHandler.CreateJob)
	return router
}

// basicConfiguration 加载基础配置
func basicConfiguration(phHandler *phMysql.CreateTask, cfg *ini.File) {
	// 服务相关配置
	port := cfg.Section("server").Key("port").String()
	local := getLocal()
	fmt.Println("<service address: " + local + ":" + port + ">")

	// 定义传入 HTTP 请求的路由规则
	router := chi.NewRouter()
	// 防止在http请求时报错崩溃
	router.Use(middleware.Recoverer)
	// 记录每次请求的详细信息
	router.Use(middleware.Logger)

	//定义了一条新路线Router路由它有两个参数：路由的路径和一个以Router实例作为参数的函数。
	//路径设置为“/”，这意味着该路由将匹配任何传入的 HTTP 请求，其 URL 以正斜杠开头。
	//传递给该方法的函数Route然后将Router实例作为参数并使用该方法定义一个新的子路由器Mount。
	router.Route("/", func(rt chi.Router) {
		rt.Mount("/tianyun", postRouter(phHandler))
	})

	// 将请求路由到正确的处理程序
	if err := http.ListenAndServe(":"+port, router); err != nil {
		log.Println("<==== 服务启动异常 ====>", err)
	} else {
		log.Println("<++++ 服务启动成功 ++++>")
	}
}

// getLocal 获取本机ip地址
func getLocal() string {
	localIP := ""
	ifAces, err := net.Interfaces()
	if err != nil {
		log.Println("<==== IP地址获取异常 ====>")
	} else {
		log.Println("<++++ IP地址获取成功 ++++>")
	}
	for _, face := range ifAces {
		address, err := face.Addrs()
		if err != nil {
			panic(err)
		}
		for _, addr := range address {
			aspnet, ok := addr.(*net.IPNet)
			if ok && !aspnet.IP.IsLoopback() && aspnet.IP.To4() != nil {
				localIP = aspnet.IP.String()
			}
		}
	}
	return localIP
}
