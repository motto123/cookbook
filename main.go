package main

import (
	"net/http"
	"log"
	"cookbook/menu"
	"cookbook/db"
	"github.com/gin-gonic/gin"
	"time"
	"math/rand"
	"cookbook/public/middleware"
)

func main() {
	rand.Seed(time.Now().Unix())
	db.InitDB()
	engine := gin.New()
	engine.Use(middleware.Cors())
	menuGroup := engine.Group("menu")
	menuGroup.Handle(http.MethodGet, "get", menu.HandleGetAllMenu)
	menuGroup.Handle(http.MethodGet, "/user/look", menu.HandleLookMenu)
	menuGroup.Handle(http.MethodPost, "/inert", menu.HandleInsertMenu)
	menuGroup.Handle(http.MethodPost, "/delete", menu.HandleDeleteMenu)
	menuGroup.Handle(http.MethodPost, "/alter", menu.HandleAlterMenu)
	menuGroup.Handle(http.MethodGet, "/random", menu.HandleRandomMenu)

	engine.LoadHTMLFiles("./website/page/index.html")
	engine.LoadHTMLFiles("./website/page/menu/home.html")
	engine.Static("/static", "./website/static")
	engine.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	engine.GET("/home", func(c *gin.Context) {
		c.HTML(http.StatusOK, "home.html", nil)
	})



	err := engine.Run(":7111")
	log.Fatal(err)
}
