package admin

import(
	m "childread/models"
	"path/filepath"
	"time"
	"os"
	"math/rand"
	"fmt"
	"crypto/md5"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
	"childread/helperfunc"
	"strings"
	"strconv"
)

// ArticleController operations for Article
type ArticleController struct {
	CommonController
}

// URLMapping ...
func (c *ArticleController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("Index", c.Index)
	c.Mapping("Edit", c.Edit)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
	c.Mapping("Upload", c.Upload)
}

// GetAll ...
// @Title GetAll
// @Description get Article
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Article
// @Failure 403
// @router / [get]
func (c *ArticleController) Index() {
	c.Session()

	page, _ := c.GetInt64("page", 0)
	page_size, _ := c.GetInt64("limit", PAGE_SIZE)
	title := c.GetString("title", "")
	filter := make(map[string]interface{}, 0)
	fields := []string{"Id", "Title", "CreatedAt"}
	if len(title) > 0{
		filter["Title__iexact"] = title
	}
	lists, count, err := m.GetAllArticle(filter, fields, page, page_size, "Id")
	if err != nil{
		c.Rsp(false, "数据为空", nil)
		return
	}
	c.SetPaginator(int(page_size), count)
	c.Data["Total"] = count
	c.Data["Title"] = title
	c.Data["Lists"] = lists
	c.Data["Current_page"] = page
	c.Data["Page_size"] = page_size

	c.Layout = "admin/layouts/layout.tpl"
	c.TplName = "admin/article/index.tpl"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["HtmlHead"] = ""
	c.LayoutSections["Scripts"] = "admin/article/scripts.tpl"
}

// Post ...
// @Title Create
// @Description create Article
// @Param	body		body 	models.Article	true		"body for Article content"
// @Success 201 {object} models.Article
// @Failure 403 body is empty
// @router / [post]
func (c *ArticleController) Post() {
	c.Session()
	a := m.Article{}
	isHot, isCommand, isWonderful, status, authorId, sourceId, sort, isBanner, isToutiao, isToday := c.getParams()
	cateId, level := c.getCate()
	a.Title = strings.TrimSpace(c.GetString("title"))
	a.Cover = c.GetString("img")
	a.Memo = strings.TrimSpace(c.GetString("memo"))
	//分类
	group := new(m.Group)
	group.Id = cateId
	a.Group = group

	a.Level = level
	a.IsHot = helperfunc.If(isHot == "on", 1, 0)
	a.IsCommand = helperfunc.If(isCommand == "on", 1, 0)
	a.IsToutiao = helperfunc.If(isToutiao == "on", 1, 0)
	a.IsWonderful = helperfunc.If(isWonderful == "on", 1, 0)
	a.IsBanner = helperfunc.If(isBanner == "on", 1, 0)
	a.IsToday = helperfunc.If(isToday == "on", 1, 0)
	a.Status = int8(status)
	a.AuthorId = authorId
	a.Path = c.GetString("ocate_path", "")
	a.Sort = sort
	a.SourceId = sourceId
	a.CreatedAt = time.Now()
	a.UpdatedAt = time.Now()

	content := strings.TrimSpace(c.GetString("content"))
	if len(content) == 0 {
		c.Rsp(false, "内容为空", nil)
	}
	valid := validation.Validation{}
	valid.Required(a.Title, "title")
	valid.Required(a.Cover, "cover")
	valid.Required(a.Memo, "memo")
	valid.Required(a.Group, "cate")
	valid.Required(a.AuthorId, "author")
	valid.Required(a.SourceId, "source")
	valid.Required(a.Path, "ocate_path")

	// 如果有错误信息，证明验证没通过
	// 打印错误信息
	if valid.HasErrors() {
		for _, err := range valid.Errors {
			c.Rsp(false, err.Key+err.Message, nil)
		}
	}

	id, err := m.AddArticle(&a)
	if err != nil {
		c.Rsp(false, err.Error(), nil)
	}
	d := m.ArticleDetail{}
	d.ArticleId = id
	d.Content = content
	d.CreatedAt = time.Now()
	d.UpdatedAt = time.Now()
	if _, err := m.AddArticleDetail(&d); err != nil {
		c.Rsp(false, err.Error(), nil)
	}
	tags := make([]int64, 0)
	c.Ctx.Input.Bind(&tags, "tag")
	if len(tags) > 0 {
		c.saveArticleTag(tags, id, false)
	}
	c.Rsp(true, "提交成功", map[string]interface{}{"url": ""})
}

//保存文章标签
func (c *ArticleController) saveArticleTag(tags []int64, articleId int64, isUp bool) {
	tagParams := []m.ArticleHasTag{}
	for _, tagId := range tags {
		if isUp == true {
			c.delArticleTag(tagId, articleId)
		}
		article := new(m.Article)
		article.Id = articleId
		node := m.ArticleHasTag{
			Article: article,
			TagId: tagId,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}
		tagParams = append(tagParams, node)
	}
	m.InsertArticleHasTagMulti(tagParams)
}

//删除旧的文章标签
func (c *ArticleController) delArticleTag(tagId int64, articleId int64) {
	m.DeleteArticleHasTag(map[string]interface{}{"ArticleId": articleId, "TagId": tagId})
}

//分类ID、分类等级
func (c *ArticleController) getCate() (cateId int64, level int) {
	level = 1
	cate, _ := c.GetInt64("cate", 0)
	cateId = cate
	if childCateId,err := c.GetInt64("cate_child", 0); err == nil{
		if childCateId > 0 {
			level = 2
			cateId = childCateId
		}
	}
	if sunChildCateId,err := c.GetInt64("cate_sun", 0); err == nil{
		if sunChildCateId > 0 {
			level = 3
			cateId = sunChildCateId
		}
	}
	return
}

//属性
func (c *ArticleController) getParams() (isHot string, isCommand string, isWonderful string, status int8, authorId int64, sourceId int64, sort int, isBanner string, isToutiao string, isToday string){
	hot := c.GetString("is_hot", "")
	isHot = hot
	command := c.GetString("is_command", "")
	isCommand = command
	wonderful := c.GetString("is_wonderful", "")
	isWonderful = wonderful
	state,_ := c.GetInt8("status", 0)
	status = state
	author, _ := c.GetInt64("author", 0)
	authorId = author
	source, _ := c.GetInt64("source", 0)
	sourceId = source
	snum,_ := c.GetInt("sort", 0)
	sort = snum
	banner := c.GetString("is_banner", "")
	isBanner = banner
	toutiao := c.GetString("is_toutiao", "")
	isToutiao = toutiao
	today := c.GetString("is_today", "")
	isToday = today
	return
}

type TreeList struct {
	Id int64		`json:"id"`
	Name string		`json:"name"`
}

// GetOne ...
// @Title GetOne
// @Description get Article by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Article
// @Failure 403 :id is empty
// @router /:id [get]
// beego tpl 中在range 中使用controller传递过来的变量会报错
// 原因是 range 相当于一个闭包  需要使用$. 来引用上下文
func (c *ArticleController) GetOne() {
	c.Session()

	id, _ := c.GetInt64(":id", 0)
	info, err := m.GetArticleById(id)
	if err != nil {
		c.Rsp(false, err.Error(), nil)
	}
	c.param()
	c.tree(info.Group.Id, info.Level)

	c.Data["Info"] = info
	detail,_ := m.GetArticleDetail(map[string]interface{}{"ArticleId": info.Id})
	c.Data["Detail"] = detail

	c.getTag()
	c.getCheckTagIds(id)

	c.Layout = "admin/layouts/layout.tpl"
	c.TplName = "admin/article/info.tpl"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["HtmlHead"] = "admin/article/style.tpl"
	c.LayoutSections["Scripts"] = "admin/article/edit_scripts.tpl"
}

//标签
func (c *ArticleController) getTag() {
	tags := m.GetTagList(nil, []string{"Id", "Name"})
	c.Data["Tags"] = tags
}

//获取选中标签
func (c *ArticleController) getCheckTagIds(articleId int64) {
	tags, err := m.GetAllArticleHasTag(map[string]interface{}{"ArticleId": articleId}, []string{"TagId"}, nil, nil,0, 0)
	var tagIds []int64
	if err == nil {
		for _, tag := range tags {
			TagId := tag.(map[string]interface{})["TagId"].(int64)
			tagIds = append(tagIds, TagId)
		}
	}
	c.Data["TagIds"] = tagIds
}

//分类
func (c *ArticleController) tree(cateId int64, level int) {
	tree := m.MenuList()
	ptrees := []*TreeList{}
	ctrees := make(map[int64][]interface{}, 0)
	strees := make(map[int64][]interface{}, 0)
	var(
		pid int64=0
		cid int64=0
		sid int64=0
	)
	for _, pitem := range tree{
		ptrees = append(ptrees, &TreeList{
			Id:pitem.Id,
			Name:pitem.Name,
		})
		if level == 1 && cateId == pitem.Id {
			pid = pitem.Id
		}
		if len(pitem.Children) > 0{
			for _,citem := range pitem.Children{
				ctrees[pitem.Id] = append(ctrees[pitem.Id], &TreeList{
					Id:citem.Id,
					Name:citem.Name,
				})
				if level == 2 && cateId == citem.Id {
					pid = pitem.Id
					cid = citem.Id
				}
				if len(citem.Children) > 0 {
					for _, sitem := range citem.Children {
						strees[citem.Id] = append(strees[citem.Id], &TreeList{
							Id:   sitem.Id,
							Name: sitem.Name,
						})
						if level == 3 && cateId == sitem.Id {
							pid = pitem.Id
							cid = citem.Id
							sid = sitem.Id
						}
					}
				}
			}
		}
	}
	c.Data["Pid"] = pid
	c.Data["Cid"] = cid
	c.Data["Sid"] = sid
	c.Data["Pclass"] = ptrees
	c.Data["Cclass"] = ctrees
	c.Data["Sclass"] = strees
}

//属性
func (c *ArticleController) param() {
	sources, _, _ := m.GetAllSource(nil, []string{"Id", "Name"}, 0, 0, "Id")
	c.Data["Sources"] = sources
	authors, _, _ := m.GetAllAuthor(nil, []string{"Id", "Name"}, 0, 0, "Id")
	c.Data["Authors"] = authors
}

//创建
func (c *ArticleController) Edit() {
	c.Session()

	var(
		pid int64=0
		cid int64=0
		sid int64=0
	)
	c.Data["Pid"] = pid
	c.Data["Cid"] = cid
	c.Data["Sid"] = sid

	c.param()
	c.getTag()

	c.Layout = "admin/layouts/layout.tpl"
	c.TplName = "admin/article/edit.tpl"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["HtmlHead"] = "admin/article/style.tpl"
	c.LayoutSections["Scripts"] = "admin/article/edit_scripts.tpl"
}

// Put ...
// @Title Put
// @Description update the Article
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Article	true		"body for Article content"
// @Success 200 {object} models.Article
// @Failure 403 :id is not int
// @router /:id [put]
func (c *ArticleController) Put() {
	c.Session()

	id, err := c.GetInt64(":id")
	if err != nil{
		c.Rsp(false, "参数缺失", nil)
	}

	a := m.Article{}
	isHot, isCommand, isWonderful, status, authorId, sourceId, Sort, isBanner, isToutiao, isToday:= c.getParams()
	cateId, level := c.getCate()
	a.Title = strings.TrimSpace(c.GetString("title"))
	a.Cover = c.GetString("img")
	a.Memo = strings.TrimSpace(c.GetString("memo"))
	//分类
	group := new(m.Group)
	group.Id = cateId
	a.Group = group

	a.Level = level
	a.Sort = Sort
	a.IsHot = helperfunc.If(isHot == "on", 1, 0)
	a.IsCommand = helperfunc.If(isCommand == "on", 1, 0)
	a.IsToutiao = helperfunc.If(isToutiao == "on", 1, 0)
	a.IsWonderful = helperfunc.If(isWonderful == "on", 1, 0)
	a.IsBanner = helperfunc.If(isBanner == "on", 1, 0)
	a.IsToday = helperfunc.If(isToday == "on", 1, 0)
	a.Status = int8(status)
	a.Path = c.GetString("ocate_path", "")
	a.AuthorId = authorId
	a.SourceId = sourceId
	a.UpdatedAt = time.Now()

	content := strings.TrimSpace(c.GetString("content"))
	if len(content) == 0 {
		c.Rsp(false, "内容为空", nil)
	}
	valid := validation.Validation{}
	valid.Required(a.Title, "title")
	valid.Required(a.Cover, "cover")
	valid.Required(a.Memo, "memo")
	valid.Required(a.Group, "cate")
	valid.Required(a.AuthorId, "author")
	valid.Required(a.SourceId, "source")
	valid.Required(a.Path, "ocate_path")

	// 如果有错误信息，证明验证没通过
	// 打印错误信息
	if valid.HasErrors() {
		for _, err := range valid.Errors {
			c.Rsp(false, err.Key+err.Message, nil)
		}
	}

	fields := []string{
		"Title", "Sort", "Group", "Memo",
		"Cover", "Status", "IsHot", "IsCommand", "IsBanner", "IsToday",
		"IsWonderful", "AuthorId", "SourceId", "UpdatedAt"}
	a.Id = id
	if err := m.UpdateArticleById(&a, fields); err != nil{
		c.Rsp(false, err.Error(), nil)
	}
	d := m.ArticleDetail{}
	detail, err :=m.GetArticleDetail(map[string]interface{}{"ArticleId": id})
	if err != nil {
		d.ArticleId = id
		d.Content = content
		d.CreatedAt = time.Now()
		d.UpdatedAt = time.Now()
		m.AddArticleDetail(&d)
	} else {
		d.Id = detail.Id
		d.Content = content
		d.UpdatedAt = time.Now()
		detFields := []string{"content", "updated_at"}
		if err := m.UpdateArticleDetailByArticleId(&d, detFields); err != nil {
			c.Rsp(false, err.Error(), nil)
		}
	}

	tags := make([]int64, 0)
	c.Ctx.Input.Bind(&tags, "tag")
	if len(tags) > 0 {
		c.saveArticleTag(tags, id, true)
	}

	data := map[string]interface{}{"url": ""}
	c.Rsp(true, "更新成功", data)

}

//上传文件
func (c *ArticleController) Upload() {
	c.Session()

	var fileInput string
	fileInput = c.GetString("filename", "files")
	file, header, err := c.GetFile(fileInput)
	if err != nil {
		c.Rsp(false, err.Error(), nil)
	}

	//验证后缀名是否符合要求
	var AllowExtMap = map[string]bool{
		".jpg":  true,
		".jpeg": true,
		".png":  true,
	}

	filename := header.Filename
	ext := filepath.Ext(filename)
	if _, ok := AllowExtMap[ext]; !ok {
		c.Rsp(false, "后缀名不符合上传要求", nil)
	}

	uploadDir := "static/uploads/" + time.Now().Format("2006-01-02")
	err = os.MkdirAll(uploadDir, os.ModePerm)
	if err != nil {
		c.Rsp(false, err.Error(), nil)
	}
	rand.Seed(time.Now().UnixNano())
	randNum := fmt.Sprintf("%d", rand.Intn(9999)+1000)
	hashName := md5.Sum([]byte(time.Now().Format("2006_01_02_15_04_05_") + randNum))
	fileName := fmt.Sprintf("%x", hashName) + ext

	defer file.Close()
	fpath := uploadDir + "/" + fileName
	err = c.SaveToFile(fileInput, fpath)
	if err != nil {
		c.Rsp(false, fmt.Sprintf("%v", err), nil)
	}
	data := make(map[string]interface{}, 0)
	data["url"] = "/" + fpath
	data["name"] = fileName
	data["path"] = beego.AppConfig.String("appdomain") + "/" + fpath
	data["size"] = header.Size

	c.Rsp(true, "upload success", data)
}

// Delete ...
// @Title Delete
// @Description delete the Article
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *ArticleController) Delete() {
	c.Session()
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	if err := m.DeleteArticle(id); err != nil {
		c.Rsp(false, err.Error(), nil)
	}
	if info, err:= m.GetArticleDetail(map[string]interface{}{"ArticleId": id}); err==nil {
		if err := m.DeleteArticleDetail(info.Id); err!=nil{
			c.Rsp(false, err.Error(), nil)
		}
	}
	c.Rsp(true, "删除成功", nil)
}