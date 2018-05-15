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

	// TODO: Define the method to serve static files
	router.GET(apiPrefix, h.GetSwaggerSpecs)

	router.GET(apiPrefix+"/instance", h.GetAllInstances)

	router.POST(apiPrefix+"/instance/", h.CreateInstance)
	router.GET(apiPrefix+"/instance/:id", h.GetInstancesById)
	router.DELETE(apiPrefix+"/instance/:id", h.DeleteInstanceById)

	log.Println("service is started")

	log.Fatal(fasthttp.ListenAndServe(":8080", router.Handler))
}
