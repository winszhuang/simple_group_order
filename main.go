package main

import (
	"github.com/gin-gonic/gin"
	"simple_group_order/db"
	"simple_group_order/routes"
)

func main() {
	dsn := "root:root@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := db.NewDatabase(dsn)
	if err != nil {
		panic(err)
	} else {
		println("------------")
		println("db connected")
		println("------------")
	}
	if err = db.Migrate(); err != nil {
		panic(err)
	}

	r := gin.Default()
	v1 := r.Group("/v1")
	routes.AddAuthRoutes(v1, db.GetDB())
	routes.AddUserRoutes(v1, db.GetDB())
	routes.AddProductRoutes(v1, db.GetDB())
	routes.AddCartRoutes(v1, db.GetDB())

	err = r.Run()
	if err != nil {
		return
	} // listen and serve on 0.0.0.0:8080
}
