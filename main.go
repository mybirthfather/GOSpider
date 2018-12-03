package main

import (
	"fmt"

	"github.com/astaxie/beego"
	_ "spider/config"
	_ "spider/models"
	_ "spider/routers"
	"spider/spider"
)

func main() {
	//ilog.Info("start")
	//s, err := spider.NewSpider("booktxt")
	//if err != nil{
	//	ilog.Fatal("new Spider error: ", err.Error())
	//}
	//err = s.SpiderUrl("http://www.booktxt.net")
	//if err != nil{
	//	ilog.Fatal("new Document error: ", err.Error())
	//}
	go spider.Start()
	s := beego.VERSION
	fmt.Println(s)
	beego.Run(":8089")
}
