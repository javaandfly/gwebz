package router

import (
	"sync"

	"github.com/gin-gonic/gin"
)

type RouterEngine struct {
	R     *gin.Engine
	mutex *sync.Mutex
	port  string
}

func (router *RouterEngine) Run() error {
	err := router.R.Run(router.port)
	return err
}
func (router *RouterEngine) AddGlobalObj(middleware ...gin.HandlerFunc) {
	router.mutex.Lock()
	defer router.mutex.Unlock()
	router.R.Use(middleware...)
}

func NewGinEngine(port string) *RouterEngine {
	return &RouterEngine{
		R:     gin.Default(),
		port:  port,
		mutex: &sync.Mutex{},
	}
}
