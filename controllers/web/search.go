package web

import (
	m "childread/models"
)

// SearchController operations for Search
type SearchController struct {
	BaseController
}

// URLMapping ...
func (c *SearchController) URLMapping() {
	c.Mapping("GetAll", c.GetAll)
}

// GetAll ...
// @Title GetAll
// @Description get Search
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Search
// @Failure 403
// @router / [get]
func (c *SearchController) GetAll() {
	page, err := c.GetInt64("p")
	if err != nil {
		page = 1
	}
	keyword := c.GetString("q", "")
	c.setPageTitle(keyword)
	c.Data["Keyword"] = keyword
	filters := make(map[string]interface{}, 0)
	if len(keyword) > 0 {
		filters["Title__icontains"] = keyword
	}
	lists,count,_ := m.GetAllArticle(filters, []string{"Id", "Title", "Cover", "Memo", "CreatedAt", "Path", "Group"}, page, PAGE_SIZE, "-Id")
	c.Data["Lists"] = lists
	c.SetPaginator(PAGE_SIZE, count)
	c.Data["Total"] = count

	c.Layout = "layouts/layout.tpl"
	c.TplName = "search/index.tpl"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["HtmlHead"] = "search/style.tpl"
	c.LayoutSections["Scripts"] = ""
}
