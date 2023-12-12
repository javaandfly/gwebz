package router

import (
	"sync"

	"github.com/gin-gonic/gin"
)

type RouterEngine struct {
	R          *gin.Engine
	mutex      *sync.Mutex
	port       string
	middleware []gin.HandlerFunc
}

func (router *RouterEngine) Router() error {

	router.R.RedirectTrailingSlash = false

	for _, middleware := range router.middleware {
		router.R.Use(middleware)
	}

	return nil
}

func (router *RouterEngine) Run() {
	router.R.Run(router.port)
}
func (router *RouterEngine) AddGlobalObj(middleware ...gin.HandlerFunc) {
	router.mutex.Lock()
	defer router.mutex.Unlock()
	router.middleware = append(router.middleware, middleware...)
}

func NewGinEngine(port string) *RouterEngine {
	return &RouterEngine{
		R:     gin.Default(),
		port:  port,
		mutex: &sync.Mutex{},
	}
}
