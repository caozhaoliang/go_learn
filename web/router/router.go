package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
)
func NewRouter() *Router {
	return &Router{}
}
type Router struct {

}
func (r *Router) Test(c *gin.Context) {
	res := make([]string,10,100)
	fmt.Println(res)
	return
}