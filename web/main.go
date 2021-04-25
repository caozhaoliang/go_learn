package main

import (
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"go_learn/web/router"
)

func main() {
	rz := gin.Default()
	r := router.NewRouter()
	pprof.Register(rz)
	rz.GET("/test",r.Test)
	rz.Run("127.0.0.1:8081")
}
