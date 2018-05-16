package main

import (
	"log"

	h "github.com/pavlov-tony/xproject/internal/handlers"

	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
)

func main() {
	router := fasthttprouter.New()
	apiPrefix := "/api/v1"

	router.GET(apiPrefix, h.GetSwaggerSpec)

	router.GET(apiPrefix+"/instance", h.GetAllInstances)

	router.GET(apiPrefix+"/instance/:id", h.GetInstancesByID)
	router.DELETE(apiPrefix+"/instance/:id", h.DeleteInstanceByID)
	router.GET(apiPrefix+"/instance/:id/history", h.GetInstanceHistoryByID)

	log.Println("service is started")

	log.Fatal(fasthttp.ListenAndServe(":8080", router.Handler))
}
