package routers

import (
	"github.com/astaxie/beego"
	"childread/controllers/admin"
    "childread/controllers/web"
)

func init() {
    beego.Router("/", &web.MainController{}, "get:Index")
    beego.Router("/search", &web.SearchController{}, "get:GetAll")

    //备孕
    beego.Router("/beiyun", &web.BeiyunController{}, "get:GetAll")
    beego.Router("/beiyun/:list", &web.BeiyunController{}, "get:GetList")
    beego.Router("/beiyun/info/cms_:id([0-9]+).html", &web.BeiyunController{}, "get:GetOne")

    //怀孕
    beego.Router("/huaiyun", &web.HuaiyunController{}, "get:GetAll")
    beego.Router("/huaiyun/:list", &web.HuaiyunController{}, "get:GetList")
    beego.Router("/huaiyun/info/cms_:id([0-9]+).html", &web.HuaiyunController{}, "get:GetOne")

    //产后
    beego.Router("/chanhou", &web.ChanhouController{}, "get:GetAll")
    beego.Router("/chanhou/:list", &web.ChanhouController{}, "get:GetList")
    beego.Router("/chanhou/info/cms_:id([0-9]+).html", &web.ChanhouController{}, "get:GetOne")

    beego.Router("/life", &web.LifeController{}, "get:GetAll")
    beego.Router("/life/:list", &web.LifeController{}, "get:GetList")
    beego.Router("/life/info/cms_:id([0-9]+).html", &web.LifeController{}, "get:GetOne")

    beego.Router("/yuer", &web.YuerController{}, "get:GetAll")
    beego.Router("/yuer/:list", &web.YuerController{}, "get:GetList")
    beego.Router("/yuer/info/cms_:id([0-9]+).html", &web.YuerController{}, "get:GetOne")

    beego.Router("/zaojiao", &web.ZaojiaoController{}, "get:GetAll")
    beego.Router("/zaojiao/:list", &web.ZaojiaoController{}, "get:GetList")
    beego.Router("/zaojiao/info/cms_:id([0-9]+).html", &web.ZaojiaoController{}, "get:GetOne")
    //tag标签
    beego.Router("/tag/:id([0-9]+)", &web.TagController{}, "get:GetAll")

    beego.Router("/admin/login", &admin.LoginController{}, "get:Index")
    beego.Router("/admin/signon", &admin.LoginController{}, "post:Post")
    beego.Router("/admin/password", &admin.LoginController{}, "get:Password")
    beego.Router("/admin/logout", &admin.LoginController{}, "get:Logout")
    beego.Router("/admin/home", &admin.HomeController{}, "get:Index")

    beego.Router("admin/group/?:level([0-9]+/?:id([0-9]+)", &admin.GroupController{}, "get:Index")
    beego.Router("admin/group/switch/:id", &admin.GroupController{}, "post:Switch")
    beego.Router("admin/group/store", &admin.GroupController{}, "post:Post")
    beego.Router("admin/group/update", &admin.GroupController{}, "put:Put")
    beego.Router("admin/group/ajax", &admin.GroupController{}, "get:GetAjaxGroup")

    beego.Router("admin/article", &admin.ArticleController{}, "get:Index")
    beego.Router("admin/article/edit/?:id", &admin.ArticleController{}, "get:Edit")
    beego.Router("admin/article/upload", &admin.ArticleController{}, "post:Upload")
    beego.Router("admin/article/post", &admin.ArticleController{}, "post:Post")
    beego.Router("admin/article/delete/:id", &admin.ArticleController{}, "delete:Delete")
    beego.Router("admin/article/show/:id", &admin.ArticleController{}, "get:GetOne")
    beego.Router("admin/article/update/:id", &admin.ArticleController{}, "put:Put")

    beego.Router("admin/author", &admin.AuthorController{}, "get:Index")
    beego.Router("admin/author/store", &admin.AuthorController{}, "post:Post")
	beego.Router("admin/author/update", &admin.AuthorController{}, "put:Put")
    beego.Router("admin/author/delete/:id", &admin.AuthorController{}, "delete:Delete")

    beego.Router("admin/source", &admin.SourceController{}, "get:Index")
    beego.Router("admin/source/store", &admin.SourceController{}, "post:Post")
    beego.Router("admin/source/update", &admin.SourceController{}, "put:Put")
    beego.Router("admin/source/delete/:id", &admin.SourceController{}, "delete:Delete")

    beego.Router("admin/topic", &admin.TopicController{}, "get:Index")
    beego.Router("admin/topic/edit/?:id", &admin.TopicController{}, "get:Edit")
    beego.Router("admin/topic/show/:id", &admin.TopicController{}, "get:GetOne")
    beego.Router("admin/topic/store", &admin.TopicController{}, "post:Post")
    beego.Router("admin/topic/update/:id", &admin.TopicController{}, "put:Put")
    beego.Router("admin/topic/delete/:id", &admin.TopicController{}, "delete:Delete")

    beego.Router("admin/tag", &admin.TagController{}, "get:Index")
    beego.Router("admin/tag/store", &admin.TagController{}, "post:Post")
    beego.Router("admin/tag/update", &admin.TagController{}, "put:Put")
    beego.Router("admin/tag/delete/:id", &admin.TagController{}, "delete:Delete")
}
