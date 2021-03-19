package web

import (
	m "childread/models"
)

//  TagController operations for Tag
type TagController struct {
	BaseController
}

// URLMapping ...
func (c *TagController) URLMapping() {
	c.Mapping("GetAll", c.GetAll)
}

// GetAll ...
// @Title Get All
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
func (c *TagController) GetAll() {
	tagId,_ := c.GetInt64(":id")
	page, err := c.GetInt64("p")
	if err != nil {
		page = 1
	}
	filters := make(map[string]interface{}, 0)
	tagInfo,_ := m.GetTagById(tagId)
	c.Data["TagInfo"] = tagInfo
	c.setPageTitle(tagInfo.Name)

	filters["ArticleHasTag__TagId"] = tagId
	fields := []string{"Id", "Title", "Path", "Memo", "Cover", "CreatedAt", "Group", "View"}
	lists,count,_ := m.GetAllArticle(filters, fields, page, PAGE_SIZE, "-Id")
	articleDatas := []*ArticleData{}
	for _, item := range lists {
		articleId := item.Id
		groupInfo := item.Group
		node := &ArticleData{
			Id: articleId,
			Title: item.Title,
			Path: item.Path,
			Cover: item.Cover,
			Memo: item.Memo,
			View: item.View,
			Cid: groupInfo.Id,
			GroupName: groupInfo.Name,
			CreatedAt: item.CreatedAt,
			Tags: c.getArticleTags(articleId),
		}
		articleDatas = append(articleDatas, node)
	}
	c.Data["Lists"] = articleDatas
	c.SetPaginator(PAGE_SIZE, count)
	c.Data["Total"] = count
	//本类文章
	viewArticles,_ := m.GetLimtArticle(nil, []string{"Id", "Title", "Group", "Path"}, 10, "-View")
	c.Data["ViewArticles"] = viewArticles
	//热门标签
	tagList,_, _ := m.GetAllTag(nil, []string{"Id", "Name"}, 1, 18, nil, nil)
	c.Data["HotTags"] = tagList

	c.Layout = "layouts/layout.tpl"
	c.TplName = "tag/index.tpl"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["HtmlHead"] = "tag/style.tpl"
	c.LayoutSections["Scripts"] = ""
}

