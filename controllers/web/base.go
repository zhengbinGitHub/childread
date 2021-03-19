package web

import (
	"github.com/astaxie/beego"
	"github.com/beego/i18n"
	"html/template"
	"childread/modules/utils"
	m "childread/models"
	"time"
	"strings"
)

// BaseController operations for Base
type BaseController struct {
	beego.Controller
	i18n.Locale
	controllerName string
	actionName     string
	menuParents []interface{}
	menuPaths []string
}

const (
	PAGE_SIZE = 20
)

//基类
type NestPreparer interface {
	NestPrepare()
}

//用户扩展
func (c *BaseController) Prepare()  {
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
	c.Data["xsrf_token"] = c.XSRFToken()
	navs := m.MenuList()
	c.Data["Navs"] = navs
	c.Data["Domain"] = beego.AppConfig.String("appdomain")
	c.hotNews()
	controlerName, actionName := c.GetControllerAndAction()
	c.controllerName = strings.ToLower(controlerName[0 : len(controlerName)-10])
	c.actionName = strings.ToLower(actionName)

	//c.AppController当前controller,看他有没有实现NestPreparer
	if app, ok := c.AppController.(NestPreparer); ok {
		app.NestPrepare()
	}
}

//热点新闻
func (c *BaseController) hotNews() {
	hots, _ := m.GetLimtArticle(map[string]interface{}{"IsHot":1}, []string{"Id", "Title", "Path"}, 4, "Id")
	c.Data["Hots"] = hots
}

// redirect to url
func (c *BaseController) Redirect(url string) {
	c.Controller.Redirect(url, 302)
	c.StopRun()
}

//json结构
type JSONStruct struct {
	status bool
	message string
	info string
	data map[string]interface{}
}

//返回json信息
func (c *BaseController) Rsp(status bool, str string, data map[string]interface{})  {
	res := JSONStruct{
		status,
		str,
		str,
		data,
	}
	c.Data["json"] = res
	c.ServeJSON()
	c.StopRun()
}

//分页
func (c *BaseController) SetPaginator(per int, nums int64) *utils.Paginator {
	p := utils.NewPaginator(c.Ctx.Request, per, nums)
	c.Data["paginator"] = p
	return p
}

type ArticleData struct {
	Id int64
	Title string
	Memo string
	Cover string
	View int
	Cid int64
	Group string
	Path string
	GroupName string
	CreatedAt time.Time
	Tags []interface{}
}

//返回文章大类下子类ID
func (c *BaseController) getArticleGroupIds(Path string) (groupIds []int64) {
	group,_ := m.GetGroupByFields(map[string]interface{}{"Path": Path}, []string{"Id", "Level"})
	groupIds = append(groupIds, group.Id)
	if 1 == group.Level {
		groups := m.GroupList(map[string]interface{}{"Parent_id": group.Id})
		for _, item := range groups {
			groupIds = append(groupIds, item.Id)
		}
	}
	return
}

//文章分页
func (c *BaseController) getArticlePaginator(groupIds []int64, page int64) {
	filters := make(map[string]interface{}, 0)
	filters["Group__in"] = groupIds
	fields := []string{"Id", "Title", "Memo", "Cover", "View", "Group", "CreatedAt", "Path"}
	lists, count, err := m.GetAllArticle(filters, fields, page, PAGE_SIZE, "-Id")
	if err != nil {
		beego.Info(err)
	}
	articleDatas := []*ArticleData{}
	for _, item := range lists {
		articleId := item.Id
		groupInfo := item.Group
		node := &ArticleData{
			Id: articleId,
			Title: item.Title,
			Path: item.Path,
			Group: groupInfo.Path,
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
	c.SetPaginator(PAGE_SIZE, count)
	c.Data["Total"] = count
	c.Data["Lists"] = articleDatas
}

//推荐文章
func (c *BaseController) getArticleCommand(groupIds []int64, page_size int64) {
	filters := make(map[string]interface{}, 0)
	filters["Group__in"] = groupIds
	filters["IsCommand"] = 1
	fields := []string{"Id", "Title", "View", "Group", "Cover", "CreatedAt", "Path"}
	lists,_ := m.GetLimtArticle(filters, fields, page_size, "-Id")
	c.Data["CommandArticles"] = lists
}

//热门文章
func (c *BaseController) getArticleHot(groupIds []int64, page_size int64) {
	filters := make(map[string]interface{}, 0)
	filters["Group__in"] = groupIds
	filters["IsHot"] = 1
	fields := []string{"Id", "Title", "View", "Group", "Cover", "CreatedAt", "Path"}
	lists,_ := m.GetLimtArticle(filters, fields, page_size, "-Id")
	c.Data["HotArticles"] = lists
}

//本类排行
func (c *BaseController) getArticleSort(groupIds []int64, page_size int64) {
	filters := make(map[string]interface{}, 0)
	filters["Group__in"] = groupIds
	fields := []string{"Id", "Title", "Group", "Path"}
	lists,_ := m.GetLimtArticle(filters, fields, page_size, "-View")
	c.Data["ViewArticles"] = lists
}

type TagParams struct {
	Id int64
	Name string
}

//文章标签
func (c *BaseController) getArticleTags(articleId int64) (tags []interface{}) {
	tagList := m.GetArticleTags(articleId)
	for _, item := range tagList{
		node := &TagParams{
			Id: item.Id,
			Name: item.Name,
		}
		tags = append(tags, node)
	}
	return
}

//分类标签
func (c *BaseController) getGroupTags(groupIds []int64, page_size int64) {
	filters := make(map[string]interface{}, 0)
	filters["GroupHasTag__GroupId__in"] = groupIds
	tagList,_, _ := m.GetAllTag(filters, []string{"Id", "Name"}, 1, page_size, nil, nil)
	c.Data["HotTags"] = tagList
}

//设置标题
func (c *BaseController) setPageTitle(pageTitle string) {
	c.Data["pageTitle"] = pageTitle
}

// 是否POST提交
func (c *BaseController) IsPost() bool {
	return c.Ctx.Request.Method == "POST"
}

//面包屑
func (c *BaseController) getBreadcrumb (groupId int64) {
	c.getCategoryByChild(groupId,nil)
	c.Data["MenuParents"] = c.menuParents
	c.Data["MenuPaths"] = c.menuPaths
}

//通过子类获取所有的父类
func (c *BaseController) getCategoryByChild(childId int64, category []*m.Group) {
	if len(category) == 0 {
		category = m.GroupList(nil)
	}
	arr := make(map[string]interface{}, 0)
	for _, item := range category {
		if item.Id == childId {
			arr["Id"] = item.Id
			arr["Title"] = item.Title
			arr["Path"] = item.Path
			if item.Parent_id != 0 {
				c.getCategoryByChild(item.Parent_id, category)
			}
			c.menuPaths = append(c.menuPaths, item.Path)
			c.menuParents = append(c.menuParents, arr)
		}
	}
	return
}

//最新文章
func (c *BaseController) getBestNewArticle(groupIds []int64, pageSize int64) {
	filters := make(map[string]interface{}, 0)
	filters["IsHot"] = 0
	filters["IsCommand"] = 0
	filters["IsWonderful"] = 0
	filters["IsToutiao"] = 0
	filters["IsToday"] = 0
	filters["Group__in"] = groupIds
	fields := []string{"Id", "Title", "Cover", "Group"}
	lists,_ := m.GetLimtArticle(filters, fields, pageSize, "-Id")
	c.Data["BestNewArticles"] = lists
}

//上下一篇
func (c *BaseController) getRelationArticle(groupIds []int64, ArticleId int64)  {
	filters := make(map[string]interface{}, 0)
	filters["Group__in"] = groupIds
	filters["Id__lt"] = ArticleId
	prevArticle, _ := m.GetArticleByField(filters)
	c.Data["Prev"] = prevArticle

	delete(filters, "Id__lt")
	filters["Id__gt"] = ArticleId
	nextArticle, _ := m.GetArticleByField(filters)
	c.Data["Next"] = nextArticle
}

//子类文章列表
func (c *BaseController) getArticleList(alias string, page int64)  {
	groupIds := c.getArticleGroupIds(alias)
	c.getArticlePaginator(groupIds, page)
	c.getArticleCommand(groupIds, 4)
	c.getArticleSort(groupIds, 8)
	c.getGroupTags(groupIds, 25)

	//Breadcrumb
	c.getBreadcrumb(groupIds[0])
}

//文章详情
func (c *BaseController) getArticleInfo(articleId int64, groupId int64, sourceId int64, authorId int64, alias string) {
	//面包屑
	c.getBreadcrumb(groupId)
	//文章详情
	filter := make(map[string]interface{}, 0)
	filter["ArticleId"] = articleId
	detail, _ := m.GetArticleDetail(filter)
	c.Data["Detail"] = detail

	//来源
	source, _ := m.GetSourceById(sourceId)
	c.Data["Source"] = source
	//作者
	author, _ := m.GetAuthorById(authorId)
	c.Data["Author"] = author
	//分类标签
	c.getGroupTags([]int64{groupId}, 25)
	//热门文章
	c.getArticleHot([]int64{groupId}, 6)
	//最新文章
	groupIds := c.getArticleGroupIds(alias)
	c.getBestNewArticle(groupIds, 8)
	//上下一篇
	c.getRelationArticle(groupIds, articleId)
}