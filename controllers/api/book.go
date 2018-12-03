package api

import (
	"spider/models"
	"time"
)

type BookController struct{
    ApiController
}

//获取小说列表
func (self *BookController) GetAll(){
	books, _ := models.GetBookList()
	self.ToJson(MSG_OK, "成功", books)
}

// 分页获取指定小说的章节列表, 每页10条
// url参数：bookid => 小说id; page => 页码
func (self *BookController) GetChapters(){
	bookid, err := self.GetInt("bookid")
	if err != nil{
		self.ToJson(MSG_ERR, err.Error(), nil)
	}
	page, err := self.GetInt("page")
	if err != nil{
		self.ToJson(MSG_ERR, err.Error(), nil)
	}
	chapters, _ := models.GetChapterPage(page, 10, "book_id",bookid)
	self.ToJson(MSG_OK, "success", chapters)
}

// 获取指定章节详细信息
// url参数: id => 章节id
func (self *BookController) GetChapter(){
	id, err := self.GetInt("id")
    if err != nil{
		self.ToJson(MSG_ERR, err.Error(), nil)
	}
	chapter, err := models.GetChapterById(id)
	if err != nil{
		self.ToJson(MSG_ERR, err.Error(), nil)
	}
	self.ToJson(MSG_OK, "success", chapter)
}
//后面是我加的
func (self *BookController) Register() {
	username := self.GetString("username")
	password := self.GetString("password")
	if len(username) == 0 {
		self.ToJson(MSG_ERR, "用户名不能为空", nil)
	}
	if len(password) == 0 || len(password) < 6 {
		self.ToJson(MSG_ERR, "密码长度不规范", nil)
	}
	user := models.User{Username: username, Password: password, Level: 99, Status: 1, CreatedAt: time.Now()}
	_, e := models.UserAdd(&user)
	if e != nil {
		self.ToJson(MSG_ERR, "用户名已存在", nil)
	} else {
		self.Ctx.SetCookie("auth", "ceshishiyong", 7*86400)//这个返回直接就在postman上就能看到的
		self.ToJson(MSG_OK, "注册成功", nil)
	}
}