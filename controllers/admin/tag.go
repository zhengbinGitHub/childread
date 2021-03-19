package admin

import(
	m "childread/models"
	"strings"
	"github.com/astaxie/beego/validation"
	"time"
)

// TagController operations for Tag
type TagController struct {
	CommonController
}

// URLMapping ...
func (c *TagController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("Index", c.Index)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Create
// @Description create Tag
// @Param	body		body 	models.Tag	true		"body for Tag content"
// @Success 201 {object} models.Tag
// @Failure 403 body is empty
// @router / [post]
func (c *TagController) Post() {
	c.Session()

	t := m.Tag{}
	t.Name = strings.TrimSpace(c.GetString("name"))
	valid := validation.Validation{}
	valid.Required(t.Name, "name")
	// 如果有错误信息，证明验证没通过
	// 打印错误信息
	if valid.HasErrors() {
		for _, err := range valid.Errors {
			c.Rsp(false, err.Key+err.Message, nil)
		}
	}
	t.CreatedAt = time.Now()
	t.UpdatedAt = time.Now()
	if _, err := m.AddTag(&t); err != nil{
		c.Rsp(false, err.Error(), nil)
	}
	c.Rsp(true, "添加成功", nil)
}

// GetAll ...
// @Title GetAll
// @Description get Tag
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Tag
// @Failure 403
// @router / [get]
func (c *TagController) Index() {
	c.Session()

	page, _ := c.GetInt64("page", 0)
	page_size, _ := c.GetInt64("limit", PAGE_SIZE)
	name := c.GetString("name")
	filters := make(map[string]interface{}, 0)
	if len(name) > 0 {
		filters["name__iexact"] = name
	}
	fields := []string{"Id", "Name", "CreatedAt"}
	nodes, count,_ := m.GetAllTag(filters, fields, page, page_size, []string{"Id"}, []string{"desc"})
	c.Data["Lists"] = nodes
	c.SetPaginator(int(page_size), count)
	c.Data["Total"] = count
	c.Data["Page_size"] = page_size
	c.Data["Name"] = name
	c.Data["Current_page"] = page

	c.Layout = "admin/layouts/layout.tpl"
	c.TplName = "admin/tag/index.tpl"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["HtmlHead"] = ""
	c.LayoutSections["Scripts"] = "admin/tag/scripts.tpl"
}

// Put ...
// @Title Put
// @Description update the Tag
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Tag	true		"body for Tag content"
// @Success 200 {object} models.Tag
// @Failure 403 :id is not int
// @router /:id [put]
func (c *TagController) Put() {
	c.Session()

	t := m.Tag{}
	t.Name = strings.TrimSpace(c.GetString("name"))
	id,_ := c.GetInt64("id")
	t.Id = id
	valid := validation.Validation{}
	valid.Required(t.Name, "name")
	// 如果有错误信息，证明验证没通过
	// 打印错误信息
	if valid.HasErrors() {
		for _, err := range valid.Errors {
			c.Rsp(false, err.Key+err.Message, nil)
		}
	}
	t.UpdatedAt = time.Now()
	if err := m.UpdateTagById(&t, []string{"Name", "UpdatedAt"}); err != nil{
		c.Rsp(false, err.Error(), nil)
	}
	c.Rsp(true, "更新成功", nil)
}

// Delete ...
// @Title Delete
// @Description delete the Tag
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *TagController) Delete() {
	c.Session()

	id,err := c.GetInt64(":id")
	if err != nil {
		c.Rsp(false, "参数错误", nil)
	}
	if err := m.DeleteTag(id); err != nil {
		c.Rsp(false, err.Error(), nil)
	}
	c.Rsp(true, "删除成功", nil)
}
