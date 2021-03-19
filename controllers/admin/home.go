package admin

// HomeController operations for Home
type HomeController struct {
	CommonController
}

// URLMapping ...
func (c *HomeController) URLMapping() {
	c.Mapping("Index", c.Index)
}

// GetOne ...
// @Title GetOne
// @Description get Home by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Home
// @Failure 403 :id is empty
// @router /:id [get]
func (c *HomeController) Index() {
	c.Session()

	navigations := make([]interface{}, 0)
	groupController := GroupController{}
	articleController := ArticleController{}
	sourceController := SourceController{}
	authorController := AuthorController{}
	topicController := TopicController{}
	tagController := TagController{}
	// 创建三一维数组，各数组长度不同
	row1 := map[string]interface{}{
		"icon": "layui-icon-chart-screen",
		"name":"分类管理",
		"url":"",
		"children":[]map[string]interface{}{
			{"icon":"", "url": groupController.URLFor("GroupController.Index"), "name":"分类列表"},
		},
	}
	navigations = append(navigations, row1)
	row2 := map[string]interface{}{
		"icon": "layui-icon-file-b",
		"name":"文章管理",
		"url":"",
		"children":[]map[string]interface{}{{"icon":"", "url": articleController.URLFor("ArticleController.Index"), "name":"文章列表"}},
	}
	navigations = append(navigations, row2)
	row3 := map[string]interface{}{
		"icon": "layui-icon-username",
		"name":"作者管理",
		"url":"",
		"children":[]map[string]interface{}{{"icon":"", "url": authorController.URLFor("AuthorController.Index"), "name":"作者列表"}},
	}
	navigations = append(navigations, row3)
	row4 := map[string]interface{}{
		"icon": "layui-icon-read",
		"name":"来源管理",
		"url":"",
		"children":[]map[string]interface{}{{"icon":"", "url": sourceController.URLFor("SourceController.Index"), "name":"来源列表"}},
	}
	navigations = append(navigations, row4)
	row5 := map[string]interface{}{
		"icon": "layui-icon-theme",
		"name":"专题管理",
		"url":"",
		"children":[]map[string]interface{}{{"icon":"", "url": topicController.URLFor("TopicController.Index"), "name":"专题列表"}},
	}
	navigations = append(navigations, row5)
	row6 := map[string]interface{}{
		"icon": "layui-icon-note",
		"name":"标签管理",
		"url":"",
		"children":[]map[string]interface{}{{"icon":"", "url": tagController.URLFor("TagController.Index"), "name":"标签列表"}},
	}
	navigations = append(navigations, row6)

	c.Data["Navigations"] = navigations
	c.TplName = "admin/layouts/iframe.tpl"
}