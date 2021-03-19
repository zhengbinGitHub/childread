package admin

import(
	m "childread/models"
	"time"
	"github.com/astaxie/beego/validation"
	"strconv"
)

// TopicController operations for Topic
type TopicController struct {
	CommonController
}

// URLMapping ...
func (c *TopicController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("Edit", c.Edit)
	c.Mapping("Index", c.Index)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Create
// @Description create Topic
// @Param	body		body 	models.Topic	true		"body for Topic content"
// @Success 201 {object} models.Topic
// @Failure 403 body is empty
// @router / [post]
func (c *TopicController) Post() {
	c.Session()

	status,_ := c.GetInt8("status", 0)
	t := m.Topic{}
	t.Cover = c.GetString("img")
	t.Url = c.GetString("url")

	valid := validation.Validation{}
	valid.Required(t.Cover, "cover")
	valid.Required(t.Url, "url")
	// 如果有错误信息，证明验证没通过
	// 打印错误信息
	if valid.HasErrors() {
		for _, err := range valid.Errors {
			c.Rsp(false, err.Key+err.Message, nil)
		}
	}

	t.Status = status
	t.CreatedAt = time.Now()
	t.UpdatedAt = time.Now()
	if _, err:=m.AddTopic(&t); err != nil{
		c.Rsp(false, "添加失败", nil)
	}
	c.Rsp(true, "添加成功", nil)
}

// GetOne ...
// @Title GetOne
// @Description get Topic by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Topic
// @Failure 403 :id is empty
// @router /:id [get]
func (c *TopicController) GetOne() {
	c.Session()

	id,_ := c.GetInt64(":id")
	info,_ := m.GetTopicById(id)
	c.Data["Info"] = info
	c.Layout = "admin/layouts/layout.tpl"
	c.TplName = "admin/topic/info.tpl"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["HtmlHead"] = ""
	c.LayoutSections["Scripts"] = "admin/topic/edit_scripts.tpl"
}

//编辑
func (c *TopicController) Edit() {
	c.Layout = "admin/layouts/layout.tpl"
	c.TplName = "admin/topic/edit.tpl"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["HtmlHead"] = ""
	c.LayoutSections["Scripts"] = "admin/topic/edit_scripts.tpl"
}

// GetAll ...
// @Title GetAll
// @Description get Topic
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Topic
// @Failure 403
// @router / [get]
func (c *TopicController) Index() {
	c.Session()

	page, _ := c.GetInt64("page", 0)
	page_size, _ := c.GetInt64("limit", PAGE_SIZE)
	filter := make(map[string]interface{}, 0)
	fields := []string{"Id", "Cover", "Url", "CreatedAt"}
	lists, count, err := m.GetAllTopic(filter, fields, page, page_size, "Id")
	if err != nil{
		c.Rsp(false, "数据为空", nil)
		return
	}
	c.SetPaginator(int(page_size), count)
	c.Data["Total"] = count
	c.Data["Lists"] = lists
	c.Data["Current_page"] = page
	c.Data["Page_size"] = page_size

	c.Layout = "admin/layouts/layout.tpl"
	c.TplName = "admin/topic/index.tpl"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["HtmlHead"] = ""
	c.LayoutSections["Scripts"] = "admin/topic/scripts.tpl"
}

// Put ...
// @Title Put
// @Description update the Topic
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Topic	true		"body for Topic content"
// @Success 200 {object} models.Topic
// @Failure 403 :id is not int
// @router /:id [put]
func (c *TopicController) Put() {
	c.Session()

	id, err := c.GetInt64(":id")
	if err != nil{
		c.Rsp(false, "参数缺失", nil)
	}

	status,_ := c.GetInt8("status", 0)
	t := m.Topic{}
	t.Cover = c.GetString("img")
	t.Url = c.GetString("url")

	valid := validation.Validation{}
	valid.Required(t.Cover, "cover")
	valid.Required(t.Url, "url")
	// 如果有错误信息，证明验证没通过
	// 打印错误信息
	if valid.HasErrors() {
		for _, err := range valid.Errors {
			c.Rsp(false, err.Key+err.Message, nil)
		}
	}
	t.Id = id
	t.Status = status
	t.UpdatedAt = time.Now()
	fields := []string{
		"Cover", "Url", "Status", "UpdatedAt"}
	if err:=m.UpdateTopicById(&t, fields); err != nil{
		c.Rsp(false, "更新失败", nil)
	}
	c.Rsp(true, "更新成功", nil)
}

// Delete ...
// @Title Delete
// @Description delete the Topic
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *TopicController) Delete() {
	c.Session()

	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	if err := m.DeleteTopic(id); err != nil {
		c.Rsp(false, err.Error(), nil)
	}
	c.Rsp(true, "删除成功", nil)
}
