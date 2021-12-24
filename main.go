package main

import (
	"fmt"
	"library/libraryDemo/cs"
	"library/libraryDemo/middleware"
	"library/libraryDemo/model"
	"library/libraryDemo/rest"
	"library/libraryDemo/rest/wechat"
	"library/libraryDemo/router"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/xormplus/xorm"
	"golang.org/x/sync/errgroup"

	"net/http"
	"time"
)

var (
	g errgroup.Group
)

func main() {
	//将数据库拉起
	params := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=true", "welfare", "welfare2021", "39.100.19.104:3306", "welfare")
	var err error
	//连接数据库
	cs.Sql, err = xorm.NewEngine("mysql", params)
	if err != nil {
		panic(err)
	}
	//首次运行时加载
	model.NewBD()
	//启动基础的Http服务
	app := gin.Default()
	root := app.Group("/api")
	root.Use(middleware.CorsHandler())
	router.Register(root, wechat.User)
	router.Register(root, rest.Book)
	router.Register(root, rest.Table)
	router.Register(root, rest.Library)
	router.Register(root, rest.Message)
	router.Register(root, rest.Notes)
	router.Register(root, rest.Cabinet)
	router.Register(root, rest.Mybooks)
	server := &http.Server{
		Addr:         ":3001",
		Handler:      app,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	g.Go(func() error {
		return server.ListenAndServe()
	})
	fmt.Print("listen:3001")
	if err := g.Wait(); err != nil {
		fmt.Print(err)
	}

}
