package admin

import (
	m "childread/models"
	"strconv"
	"github.com/astaxie/beego/validation"
)

//  GroupController operations for Group
type GroupController struct {
	CommonController
}

// Post ...
// @Title Post
// @Description create Group
// @Param	body		body 	models.Group	true		"body for Group content"
// @Success 201 {int} models.Group
// @Failure 403 body is empty
// @router / [post]
func (c *GroupController) Post() {
	v := m.Group{}
	v.Name = c.GetString("name")
	v.Title = c.GetString("title")
	valid := validation.Validation{}
	valid.Required(v.Name, "name")
	valid.Required(v.Title, "title")
	// 如果有错误信息，证明验证没通过
	// 打印错误信息
	if valid.HasErrors() {
		for _, err := range valid.Errors {
			c.Rsp(false, err.Key+err.Message, nil)
		}
	}
	ParentId, _ := c.GetInt64("parent_id", 0)
	v.Parent_id = ParentId
	v.Status = 1
	if err := c.ParseForm(&v); err != nil {
		//handle error
		c.Rsp(false, err.Error(), nil)
	}
	if _, err := m.AddGroup(&v); err != nil {
		c.Rsp(false, err.Error(), nil)
	}
	c.Rsp(true, "提交成功", nil)
}

// GetOne ...
// @Title Get One
// @Description get Group by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Group
// @Failure 403 :id is empty
// @router /:id [get]
func (c *GroupController) Index() {
	c.Session()

	page, _ := c.GetInt64("page", 0)
	page_size, _ := c.GetInt64("limit", PAGE_SIZE)
	sort := c.GetString("sort")
	name := c.GetString("name")
	order := c.GetString("order")
	if len(order) > 0 {
		if order == "desc" {
			sort = "-" + sort
		}
	} else {
		sort = "Id"
	}
	parentId,_ := c.GetInt64(":id", 0)
	filters := make(map[string]interface{}, 0)
	filters["parent_id"] = parentId
	if len(name) > 0 {
		filters["title__iexact"] = name
	}

	nodes, count := m.GetGrouplist(filters, page, page_size, sort)
	c.Data["Lists"] = nodes
	c.SetPaginator(int(page_size), count)
	c.Data["Total"] = count
	c.Data["Page_size"] = page_size
	c.Data["Name"] = name
	c.Data["Current_page"] = page
	c.Data["Parent_id"] = parentId

	c.Layout = "admin/layouts/layout.tpl"
	c.TplName = "admin/group/index.tpl"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["HtmlHead"] = ""
	c.LayoutSections["Scripts"] = "admin/group/scripts.tpl"
}

// Delete ...
// @Title Delete
// @Description delete the Group
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *GroupController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	if _, err := m.DelGroupById(id); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

// Switch
func (c *GroupController) Switch() {
	id, err := c.GetInt64(":id")
	value, _ := c.GetInt("value")
	if err != nil{
		c.Rsp(false, err.Error(), nil)
	}
	var g m.Group
	g.Id = id
	g.Status = value
	if _, err := m.UpdateGroup(&g); err == nil {
		c.Rsp(true, "更新成功", nil)
	}
	c.Rsp(false, err.Error(), nil)
}

// Put ...
// @Title Put
// @Description update the Home
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Home	true		"body for Home content"
// @Success 200 {object} models.Home
// @Failure 403 :id is not int
// @router /:id [put]
func (c *GroupController) Put() {
	id, err := c.GetInt64("id")
	if err != nil{
		c.Rsp(false, err.Error(), nil)
	}
	v := m.Group{}
	v.Title = c.GetString("title")
	valid := validation.Validation{}
	valid.Required(v.Title, "title")
	// 如果有错误信息，证明验证没通过
	// 打印错误信息
	if valid.HasErrors() {
		for _, err := range valid.Errors {
			c.Rsp(false, err.Key+err.Message, nil)
		}
	}
	v.Id = id
	v.Status = 1
	if _, err := m.UpdateGroup(&v); err == nil {
		c.Rsp(true, "更新成功", nil)
	}
	c.Rsp(false, err.Error(), nil)
}

//分类
func (c *GroupController) GetAjaxGroup()  {
	id, err := c.GetInt64("id", 0)
	if err != nil{
		c.Rsp(false, err.Error(), nil)
	}
	filters := make(map[string]interface{}, 0)
	filters["Parent_id"] = id

	groupList := m.GroupList(filters)
	c.Rsp(true, "ok", map[string]interface{}{"group": groupList})

}
