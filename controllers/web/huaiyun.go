package web

import (
	m "childread/models"
	"github.com/astaxie/beego"
)

// HuaiyunController operations for Huaiyun
type HuaiyunController struct {
	BaseController
}

//子类
func (c *HuaiyunController) NestPrepare() {
	c.Data["InfoPath"] = "huaiyun"
}

// URLMapping ...
func (c *HuaiyunController) URLMapping() {
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetList", c.GetList)
	c.Mapping("GetAll", c.GetAll)
}

// Post ...
// @Title Create
// @Description create Huaiyun
// @Param	body		body 	models.Huaiyun	true		"body for Huaiyun content"
// @Success 201 {object} models.Huaiyun
// @Failure 403 body is empty
// @router / [post]
func (c *HuaiyunController) GetList() {
	page, err := c.GetInt64("p")
	if err != nil {
		page = 1
	}
	alias := c.GetString(":list")
	c.getArticleList(alias, page)

	c.Layout = "layouts/layout.tpl"
	c.TplName = "channel/index.tpl"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["HtmlHead"] = "article/style.tpl"
	c.LayoutSections["Scripts"] = "article/scripts.tpl"
}

// GetOne ...
// @Title GetOne
// @Description get Huaiyun by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Huaiyun
// @Failure 403 :id is empty
// @router /:id [get]
func (c *HuaiyunController) GetOne() {
	id, _ :=c.GetInt64(":id")
	info, err := m.GetArticleById(id)
	if err != nil{
		c.Redirect(beego.AppConfig.String("appdomain"))
		return
	}
	c.Data["Info"] = info

	//文章详情
	c.getArticleInfo(id, info.Group.Id, info.SourceId, info.AuthorId, c.Data["InfoPath"].(string))
	c.setPageTitle(info.Title)

	c.Layout = "layouts/layout.tpl"
	c.TplName = "article/info.tpl"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["HtmlHead"] = "article/style.tpl"
	c.LayoutSections["Scripts"] = "article/scripts.tpl"
}

// GetAll ...
// @Title GetAll
// @Description get Huaiyun
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Huaiyun
// @Failure 403
// @router / [get]
func (c *HuaiyunController) GetAll() {
	page, err := c.GetInt64("p")
	if err != nil {
		page = 1
	}
	c.getArticleList(c.Data["InfoPath"].(string), page)

	c.Layout = "layouts/layout.tpl"
	c.TplName = "channel/index.tpl"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["HtmlHead"] = "channel/style.tpl"
	c.LayoutSections["Scripts"] = "channel/scripts.tpl"
}
