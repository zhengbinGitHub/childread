package main

import (
	_ "childread/routers"
	"github.com/astaxie/beego"
	"os"
	"childread/models"
	. "childread/helperfunc"
	"github.com/astaxie/beego/session"
	"github.com/astaxie/beego/context"
)
var globalSessions *session.Manager

var FilterMethod = func(ctx *context.Context) {
	if ctx.Input.Query("_method")!="" && ctx.Input.IsPost(){
		ctx.Request.Method = ctx.Input.Query("_method")
	}
}

func main() {
	sessionConfig := &session.ManagerConfig{
		CookieName:      "childsessionid",
		EnableSetCookie: true,
		Gclifetime:      3600,
		Maxlifetime:     3600,
		Secure:          false,
		CookieLifeTime:  3600,
		ProviderConfig:  "./tmp",
	}
	globalSessions, _ := session.NewManager("memory", sessionConfig)
	go globalSessions.GC()

	beego.SetStaticPath("/static", "static")
	//初始化
	initialize()
	//在表单中使用 PUT 方法
	beego.InsertFilter("*", beego.BeforeRouter, FilterMethod)

	beego.Run()
}

func initialize() {
	//判断初始化参数
	initArgs()

	models.InitDb()

	beego.AddFuncMap("stringsToJson", StringsToJson)
	beego.AddFuncMap("Pwdhash", Pwdhash)
	beego.AddFuncMap("Index", Index)
	beego.AddFuncMap("FmtSprintf", FmtSprintf)
	beego.AddFuncMap("InArray", InArray)
	beego.AddFuncMap("Breadcrumb", Breadcrumb)
}

func initArgs() {
	args := os.Args
	for _, v := range args {
		//$ ./demo -syncdb 创建数据库、初始化数据库表。
		if v == "-syncdb" {
			models.Syncdb()
			os.Exit(0)
		}
	}
}

