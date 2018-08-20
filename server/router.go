package server

import (
	"net/http"

	"github.com/1975210542/wechat_s/controller"
	"github.com/gin-gonic/gin"
)

func init() {
	gin.SetMode(gin.DebugMode)
}

func InitRouter() {
	router := gin.Default()
	router.LoadHTMLGlob("./templates/*")
	initRouters(router)

	println("server listening port on", "8081", "...")
	http.ListenAndServe("0.0.0.0:8081", router)
}

func initRouters(router *gin.Engine) {

	router.StaticFS("/static", http.Dir("static"))

	router.GET("/register", controller.Register)

}
