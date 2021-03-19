package admin

import (
	"github.com/astaxie/beego"
	"html/template"
	"childread/modules/utils"
)

const (
	PAGE_SIZE = 20
)

// CommonController operations for Common
type CommonController struct {
	beego.Controller
}

func (c *CommonController) Rsp(status bool, str string, data map[string]interface{})  {
	c.Data["json"] = &map[string]interface{}{"status": status, "message":str, "info": str, "data": data}
	c.ServeJSON()
	c.StopRun()
}

//用户扩展
func (c *CommonController) Prepare()  {
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
	c.Data["token"] = c.XSRFToken()
}

//session 值
func (c *CommonController) Session() {
	userinfo := c.GetSession("userinfo")
	if userinfo == nil {
		if c.IsAjax() {
			c.Rsp(false, "请登录",nil)
		} else {
			c.Ctx.Redirect(302, "/admin/login")
		}
		return
	}
	c.Data["userinfo"] = userinfo
}

//分页
func (c *CommonController) SetPaginator(per int, nums int64) *utils.Paginator {
	p := utils.NewPaginator(c.Ctx.Request, per, nums)
	c.Data["paginator"] = p
	return p
}

