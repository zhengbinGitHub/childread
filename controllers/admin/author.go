package admin

import (
	m "childread/models"
	"strings"
	"github.com/astaxie/beego/validation"
	"time"
)

//  AuthorController operations for Author
type AuthorController struct {
	CommonController
}

// URLMapping ...
func (c *AuthorController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("Index", c.Index)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// GetAll ...
// @Title Get All
// @Description get Author
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Author
// @Failure 403
// @router / [get]
func (c *AuthorController) Index() {
	c.Session()
	page, _ := c.GetInt64("page", 0)
	page_size, _ := c.GetInt64("limit", PAGE_SIZE)
	name := c.GetString("name", "")
	filter := make(map[string]interface{}, 0)
	fields := []string{"Id", "Name", "CreatedAt"}
	if len(name) > 0{
		filter["Name__iexact"] = name
	}
	lists, count, err := m.GetAllAuthor(filter, fields, page, page_size, "Id")
	if err != nil{
		c.Rsp(false, "数据为空", nil)
		return
	}
	c.SetPaginator(int(page_size), count)
	c.Data["Total"] = count
	c.Data["Name"] = name
	c.Data["Lists"] = lists
	c.Data["Current_page"] = page
	c.Data["Page_size"] = page_size
	c.Layout = "admin/layouts/layout.tpl"
	c.TplName = "admin/author/index.tpl"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["HtmlHead"] = ""
	c.LayoutSections["Scripts"] = "admin/author/scripts.tpl"
}

// Post ...
// @Title Post
// @Description create Author
// @Param	body		body 	models.Author	true		"body for Author content"
// @Success 201 {int} models.Author
// @Failure 403 body is empty
// @router / [post]
func (c *AuthorController) Post() {
	c.Session()
	var v m.Author
	v.Name = strings.TrimSpace(c.GetString("name", ""))
	valid := validation.Validation{}
	valid.Required(v.Name, "name")
	// 如果有错误信息，证明验证没通过，打印错误信息
	if valid.HasErrors() {
		for _, err := range valid.Errors {
			c.Rsp(false, err.Key+err.Message, nil)
		}
	}
	v.CreatedAt = time.Now()
	v.UpdatedAt = time.Now()
	if _, err := m.AddAuthor(&v); err != nil {
		c.Rsp(false, err.Error(), nil)
	}
	c.Rsp(true, "添加成功", nil)
}

// Put ...
// @Title Put
// @Description update the Author
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Author	true		"body for Author content"
// @Success 200 {object} models.Author
// @Failure 403 :id is not int
// @router /:id [put]
func (c *AuthorController) Put() {
	c.Session()
	id, _ := c.GetInt64("id")
	v := m.Author{}
	v.Id = id
	v.Name = strings.TrimSpace(c.GetString("name", ""))
	valid := validation.Validation{}
	valid.Required(v.Name, "name")
	// 如果有错误信息，证明验证没通过，打印错误信息
	if valid.HasErrors() {
		for _, err := range valid.Errors {
			c.Rsp(false, err.Key+err.Message, nil)
		}
	}
	v.UpdatedAt = time.Now()
	if err := m.UpdateAuthorById(&v, []string{"Name", "UpdatedAt"}); err != nil {
		c.Rsp(false, err.Error(), nil)
	}
	c.Rsp(true, "更新成功", nil)
}

// Delete ...
// @Title Delete
// @Description delete the Author
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *AuthorController) Delete() {
	c.Session()
	id,_ := c.GetInt64(":id")
	count, err := m.CountArticle(map[string]interface{}{"AuthorId": id})
	if err==nil && count>0 {
		c.Rsp(false, "关联文章不能删除", nil)
	}
	if err := m.DeleteAuthor(id); err != nil {
		c.Rsp(false, err.Error(), nil)
	}
	c.Rsp(true, "删除成功", nil)
}