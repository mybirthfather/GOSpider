# 简单爬虫实现

使用golang实现简单的爬虫

使用cron定时爬取数据

## api接口

`/v1/book` : 获取小说列表

`/v1/book/getchapters?bookid=1&page=1` : 获取指定小说的章节，分页查询，每页10章

`/v1/book/chapter?id=950` : 获取指定章节详细内容