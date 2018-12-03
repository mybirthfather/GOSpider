package spider

import (
	"errors"
	"fmt"
	"github.com/robfig/cron"
	"spider/config"
	"spider/models"
)

type SBook struct{
	Name string
	Image string
	Url string
	Chapters []*SChapter
}

type SChapter struct{
	Title string
	Url string
	Order int
	Pre int
	Next int
	Content string
}

type Spider interface{
    SpiderUrl(url string) error
}

func NewSpider(from string) (Spider, error){
	switch from{
	case "booktxt":
		return new(BookTextSpider), nil
	default:
		return nil, errors.New("系统暂未处理该类型的配置文件")
	}
}

func Start(){

    c := cron.New()
	spec := config.AppConfig.GetString("task.spec")
    c.AddFunc(spec,getBook)
	c.Start()
    select{

	}
}

func getBook(){
	fmt.Println("getBook()")
	s, err := NewSpider("booktxt")
	if err != nil{
	}
//http://www.booktxt.net/2_2219/
	err = s.SpiderUrl("http://www.booktxt.net")
	if err != nil{
	}
	var str string
	fmt.Scan(&str)
	books, _ := models.GetBookList("status", 0)
	for _, book := range books{
		go func(book *models.Book){
			s, err := NewSpider(book.From)
			if err != nil{
				return
			}
			err = s.SpiderUrl(book.Url)
			if err != nil{
			}
			fmt.Println("已爬取完毕")
		}(book)
	}
}