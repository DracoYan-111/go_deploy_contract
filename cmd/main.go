package main

import (
	"GoContractDeployment/cron"
	phMysql "GoContractDeployment/handler/http"
	"GoContractDeployment/navigation"
	"GoContractDeployment/utils"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"log"
	"net"
	"net/http"
)

func main() {

	// connect to database
	connection := navigation.CreateData()
	phHandler := phMysql.NewJobHandler(connection)

	// Start a scheduled task
	cron.UpdateLibrary(phHandler)
	cron.ReturnStatus(phHandler)

	basicConfiguration(phHandler)
}

// postRouter a completely separate router
func postRouter(phHandler *phMysql.CreateTask) http.Handler {
	router := chi.NewRouter()
	router.Post("/", phHandler.CreateJob)
	return router
}

// basicConfiguration Load the basic configuration
func basicConfiguration(phHandler *phMysql.CreateTask) {
	var datalist = []string{"port"}
	loading, err := utils.ConfigurationLoading("server", datalist)
	if err != nil {
		log.Println(err)
	}

	port := loading[0]
	local := getLocal()
	fmt.Println("<service address: " + local + ":" + port + ">")

	router := chi.NewRouter()
	router.Use(middleware.Recoverer)
	router.Use(middleware.Logger)

	router.Route("/", func(rt chi.Router) {
		rt.Mount("/tianyun", postRouter(phHandler))
	})

	// Route the request to the correct handler
	if err := http.ListenAndServe(":"+port, router); err != nil {
		log.Println("<==== main:Service startup exception ====>", err)
	} else {
		log.Println("<++++ main:Service started successfully ++++>")
	}
}

// getLocal Get local ip address
func getLocal() string {
	localIP := ""
	ifAces, err := net.Interfaces()
	if err != nil {
		log.Println("<==== main:Abnormal IP address acquisition ====>")
	} else {
		log.Println("<++++ main:IP address obtained successfully ++++>")
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
