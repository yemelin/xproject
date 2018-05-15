package handlers

import (
	"fmt"
	"log"

	"github.com/valyala/fasthttp"
)

// GetSwaggerSpecs - returns a swagger specs file
func GetSwaggerSpecs(ctx *fasthttp.RequestCtx) {
	log.Println("request for swagger specs")
	fmt.Fprintf(ctx, "You have requested for swagger specs\n")
}

// GetAllInstances - returns a list of all available instances
func GetAllInstances(ctx *fasthttp.RequestCtx) {
	log.Println("request for all instances")
	fmt.Fprintf(ctx, "You have requested for all instances\n")
}

// GetInstancesById - returns an instance by id
func GetInstancesById(ctx *fasthttp.RequestCtx) {
	id := ctx.UserValue("id").(string)
	log.Println("request for instance with id", id)
	fmt.Fprintf(ctx, "You have requested for instance with id \"%v\"\n", id)
}

// CreateInstance - creates new instance
func CreateInstance(ctx *fasthttp.RequestCtx) {
	log.Println("request for instance creation")
}

// CreateInstance - deletes an instance by id
func DeleteInstanceById(ctx *fasthttp.RequestCtx) {
	id := ctx.UserValue("id").(string)
	log.Println("request to delete instance with id", id)
}
